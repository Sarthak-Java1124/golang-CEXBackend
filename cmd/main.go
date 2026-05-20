package main

import (
	"context"
	"fmt"
	"golang-CEX/internal/database"
	"log"
	"time"
)

func main() {
	store, err := database.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatal("The error in db connection is : ", err)
	}
	defer cancel()
	users, err := store.Queries.ListUsers(ctx)
	if err != nil {
		log.Fatal("The error in getting users is :", err)

	}
	fmt.Println("The users we got is ", users)
}
