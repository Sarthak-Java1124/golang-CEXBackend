package routes

import (
	"golang-CEX/internal/handlers"

	"github.com/gin-gonic/gin"
)

func incomingAuthRoutes() {
	r := gin.Default()

	r.GET("api/auth/:id", handlers.GetProfileInfo)
	r.POST("/api/signup", handlers.CreateUserProfile)
	r.POST("/api/login", handlers.LoginUser)
	r.DELETE("/api/delete/:id", handlers.DeleteUser)
	r.PATCH("/api/update-user/:id", handlers.UpdateUserInfo)

}
