package handlers

import (
	"context"
	"golang-CEX/internal/database/sqlc"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) MatchingEngine(order sqlc.Order, userId pgtype.UUID) {

	orders_from_orderbook, err := h.FindMatchingOrders(&order)
	if err != nil {
		log.Fatal("There was an error finding a matching order in the matching engine")
	}
	h.ExecuteTrades(&order, &orders_from_orderbook, userId)

}

func (h *Handler) ExecuteTrades(orders *sqlc.Order, orders_from_orderbook *[]sqlc.Order) (bool, error) {
	//code for case 1
	var ok bool
	var err error
	err = nil
	ok = true
	for _, i := range *orders_from_orderbook {
		qty, _ := orders.Quantity.Int64Value()
		orderbook_qty, _ := i.Quantity.Int64Value()

		if orderbook_qty.Int64 > qty.Int64 {
			orderbook_qty.Int64 = orderbook_qty.Int64 - qty.Int64
			ok, err = SettleBalances(orders.UserID, orders.Price, orders.Quantity)

			remaining_quantity, _ := i.RemainingQuantity.Int64Value()
			remaining_quantity.Int64 = orderbook_qty.Int64
			i.Status = "filled"

		}
		if orderbook_qty.Int64 == qty.Int64 {
			remaining_quantity, _ := i.RemainingQuantity.Int64Value()
			remaining_quantity.Int64 = 0

			ok, err = SettleBalances(orders.UserID, orders.Price, orders.Quantity)
			i.Status = "filled"
		}
		if orderbook_qty.Int64 < qty.Int64 {
			ok, err = SettleBalances(orders.UserID, orders.Price, orders.Quantity)

			remaining_quantity, _ := i.RemainingQuantity.Int64Value()
			remaining_quantity.Int64 = 0
			i.Status = "partially_filled"

		}
		if orderbook_qty.Int64 == 0 {
			log.Fatal("No orderbook quantity available")

		}
	}
	return ok, err
}

func (h *Handler) FindMatchingOrders(order *sqlc.Order) ([]sqlc.Order, error) {
	var final_orders []sqlc.Order
	if order.Status == "filled" {
		log.Fatal("The order is already filled")
		return final_orders, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var n pgtype.Numeric
	price := order.Price
	err := n.Scan(price)
	if err != nil {
		log.Fatal("There was an error in converting the int into pg numeric")
	}
	side := order.Side

	if side == "buy" {
		order_match_param_sell := sqlc.GetMatchingSellOrdersParams{
			Market: order.Market,
			Price:  n,
		}
		sell_orders, err := h.Queries.Queries.GetMatchingSellOrders(ctx, order_match_param_sell)
		if err != nil {
			log.Fatal("An error occured while getting the orders from the db")
			return nil, err
		}
		for _, orders := range sell_orders {
			qty, _ := orders.Quantity.Int64Value()
			if int64(qty.Int64) == 0 {
				log.Fatal("There is no available quantity for the given order")
			}
			final_orders = append(final_orders, orders)

		}
		return final_orders, nil
	} else if side == "sell" {
		order_match_param_buy := sqlc.GetMatchingBuyOrdersParams{
			Market: order.Market,
			Price:  n,
		}
		sell_orders, err := h.Queries.Queries.GetMatchingBuyOrders(ctx, order_match_param_buy)
		if err != nil {
			log.Fatal("An error occured while getting the orders from the db")
			return nil, err
		}
		for _, orders := range sell_orders {
			qty, _ := orders.Quantity.Int64Value()
			if int64(qty.Int64) == 0 {
				log.Fatal("There is no available quantity for the given order")
			}
			final_orders = append(final_orders, orders)

		}
		return final_orders, nil
	}

	return final_orders, nil

}

func (h *Handler) SettleBalances(userId pgtype.UUID, price pgtype.Numeric, qty pgtype.Numeric) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user_balance, err := h.Queries.Queries.GetBalanceByID(ctx, userId)
	if err != nil {
		log.Fatal("There was an error fetching the user balances")
		return false, err
	}
	use_price, _ := price.Int64Value()
	use_qty, _ := qty.Int64Value()
	var err1 error

	if user_balance.Balance < (int32(use_price.Int64) * int32(use_qty.Int64)) {
		log.Fatal("User balance is less")
		return false, err1
	}
	spend_amt := int32(use_price.Int64) * int32(use_qty.Int64)

	userBalanceParams := sqlc.UpdateBalanceAmountParams{
		ID:      userId,
		Balance: spend_amt,
	}
	updated_balance, err := h.Queries.Queries.UpdateBalanceAmount(ctx, userBalanceParams)
	if err != nil {
		log.Fatal("There was an error in updating user balance")
		return false, err
	}
	log.Println("The updated user balance is  : ", updated_balance)
	return true, nil

}

/*
case 1 : quantity more
 order_from_user {
    price : 200
	market : sol-btc
	side : buy
	quantity : 100
 }

 order_from_orderbooks : {
   price : 200
	market : sol-btc
	side : buy
	quantity : 400
 }
 order gets eaten up directly
 <======================================>

 case 2 : same quantity
 order_from_user {
    price : 200
	market : sol-btc
	side : buy
	quantity : 100
 }

 order_from_orderbooks : {
   price : 200
	market : sol-btc
	side : buy
	quantity : 100
 }
  order gets eaten up directly

 <======================================>
case 3 : less quantity
 order_from_user {
    price : 200
	market : sol-btc
	side : buy
	quantity : 100
 }

 order_from_orderbooks : {
   price : 200
	market : sol-btc
	side : buy
	quantity : 50
 }

 the available quantity gets picked up
 remaining_quantity gets updated from order_book
 order_quantity gets updated
 the remaining order sits in the order book
 order status is partially_filled
 <======================================>
 case 4 : zero quantity
 order_from_user {
    price : 200
	market : sol-btc
	side : buy
	quantity : 100
 }

 order_from_orderbooks : {
   price : 200
	market : sol-btc
	side : buy
	quantity : 0
 }

order sits idle in the orderbook waiting for the right match
 <======================================>






*/
