package routes

import (
	"golang-CEX/internal/handlers"

	"github.com/gin-gonic/gin"
)

func incomingWalletRoutes() {
	r := gin.Default()
	r.POST("/api/wallet/deposit", handlers.WalletDepositHandler)
	r.GET("/api/wallet/balances", handlers.WalletBalanceHandler)
	
}
