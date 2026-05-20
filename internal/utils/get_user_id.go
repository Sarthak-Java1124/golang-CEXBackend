package utils

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetUserId(c *gin.Context) (pgtype.UUID, bool) {
	var uuid pgtype.UUID
	userId, exists := c.Get("userId")
	if !exists {
		log.Fatal("The user with this id is not found")
		return pgtype.UUID{}, false
	}
	err := uuid.Scan(userId)

	if err != nil {
		return pgtype.UUID{}, false
	}

	return uuid, true
}
