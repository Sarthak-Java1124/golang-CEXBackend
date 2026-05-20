package handlers

import (
	"context"
	"golang-CEX/internal/database/sqlc"
	"golang-CEX/internal/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Req struct {
	Market   string
	Side     string
	Type     string
	Price    int
	Quantity int
}

func (h *Handler) OrderPostHandler(c *gin.Context) {
	userId, ok := utils.GetUserId(c)
	if !ok {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	var order Req
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "There was an error in binding json"})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	params := sqlc.GetUserAssetBalanceParams{
		UserID: userId,
		Asset:  order.Type,
	}
	
	userBalance, err := h.Queries.Queries.GetUserAssetBalance(ctx, params)
	if err != nil {
		log.Println("The error in getting the user balance is : " , err)
		c.JSON(http.StatusBadRequest , gin.{"error" : "An error occured in fetching the user balances"})
	}
	if userBalance < (order.Price * order.Quantity) {
		c.JSON(http.StatusBadRequest , gin.H{"error" : "There is not enough balance in the account to carry out this transaction"})
	}
	order_id := pgtype.UUID{
    Bytes: [16]byte(uuid.New()),  
    Valid: true,
}
	order_params := sqlc.CreateOrderParams {
		ID:  order_id,
Market: order.Market,
Side: order.Side,
Type : order.Type,
UserID: userId,
Price: order.Price,
Quantity: order.Quantity,
RemainingQuantity: order.Quantity,
Status: "open",
CreatedAt:time.Now().UTC()



	}
	_, err := h.Queries.Queries.CreateOrder(ctx , order_params )
	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"message" : "An error occured while saving the order to the database"})
	}

	c.JSON(http.StatusAccepted , gin.H{"message"  : "Order saved success"  , "data" : {
		"orderId" : order_id,
		"status" : "open"
	} })
}

func (h *Handler) OrderGetHandler(c *gin.Context) {
userId, ok := utils.GetUserId(c)
	if !ok {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

}

func DeleteOrder(c *gin.Context) {

}

func GetOrderBookByMarket(c *gin.Context) {

}
