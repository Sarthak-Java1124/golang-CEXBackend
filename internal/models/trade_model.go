package models

type TradeModel struct {
	ID           string `json:"id"`
	Market       string `json:"market"`
	BuyOrderId   string `json:"buy_order_id"`
	SellOrderId  string `json:"sell_order_id"`
	BuyerUserId  string `json:"buyer_user_id"`
	SellerUserId string `json:"seller_user_id"`
	Price        int    `json:"price"`
	Quantity     int    `json:"quantity"`
	CreatedAt    string `json:"created_at"`
}
