package routes

import (
	"golang-CEX/internal/handlers"

	"github.com/gin-gonic/gin"
)

func incomingOrderRoutes() {
	r := *gin.Default()
	r.POST("/api/order", handlers.OrderPostHandler)
	r.GET("/api/order/me", handlers.OrderGetHandler)
	r.DELETE("/api/order/:id", handlers.DeleteOrder)
	r.GET("/api/orderbooks/:market", handlers.GetOrderBookByMarket)
	
}
