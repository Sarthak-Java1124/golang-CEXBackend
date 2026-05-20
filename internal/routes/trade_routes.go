package routes

import (
	"golang-CEX/internal/handlers"

	"github.com/gin-gonic/gin"
)

func incomingTradeRoutes() {
	r := gin.Default()
	r.GET("/api/trades/:market", handlers.MarketTradeHandler)
}
