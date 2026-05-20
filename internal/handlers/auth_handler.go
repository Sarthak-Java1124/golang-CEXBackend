package handlers

import (
	"context"
	"golang-CEX/internal/database"
	"golang-CEX/internal/database/sqlc"
	"golang-CEX/internal/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

var userCount int = 0

type Handler struct {
	Queries *database.Store
}
type UserUpdateRequest struct {
	Email string 
	Name string
	Password string 
}

func (h *Handler) GetProfileInfo(c *gin.Context) {

}

func (h *Handler) CreateUserProfile(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user sqlc.CreateUserParams
	if err := c.BindJSON(&user); err != nil {
		log.Fatal("The error in binding the json request is : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "There is an error binding the request body to json"})
	}
	isUserPresent, err := h.Queries.Queries.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Fatal("The error in finding user is : ", err)
	}
	if isUserPresent.Email != "" {
		log.Fatal("The user is already present")
		c.JSON(http.StatusNotFound, gin.H{"message": "The User With this email already exists"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	userCount++
	if err != nil {
		log.Fatal("The error in hashing the user password is : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "There was an error in hashing the users password"})
		return
	}
	user.Password = hashedPassword
	user.ID = strconv.Itoa(userCount)
	user.CreatedAt = pgtype.Timestamp{Time: time.Now()}
	createdUser, err := h.Queries.Queries.CreateUser(ctx, user)
	if err != nil {
		log.Fatal("An error occured while signing up", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "An error occurred while signing up the user "})
	}
	access_token, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		log.Fatal("The error in generating token is : ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred in generating token"})
		return
	}
	c.SetCookie(
		"access_token",
		access_token,
		3600,
		"/",
		"",
		true,
		true,
	)
	c.JSON(http.StatusAccepted, gin.H{"message": "Signed up successfully", "data": createdUser})

}
func LoginUser(c *gin.Context) {

}
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	err := h.Queries.Queries.DeleteUser(ctx, id)
	if err != nil {
		log.Fatal("Failed to delete the user")
		c.JSON(http.StatusBadGateway, gin.H{"message": "An error occurred in deleting the user"})
		return
	}
	c.JSON(http.StatusAccepted , gin.H{"message" : "User deleted successfully"})
}

func UpdateUserInfo(c *gin.Context) {

}
