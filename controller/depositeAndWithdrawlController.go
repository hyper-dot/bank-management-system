package controller

import (
	"bank/model"
	"bank/utils"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func Deposit() {
	// Account Number identification
	var accountNumber int
	fmt.Print("Please Enter Account Number:")
	fmt.Scan(&accountNumber)

	// Amount Identification
	var amount float64
	fmt.Print("Please Enter Amount to Deposite:")
	fmt.Scan(&amount)

	// Find account with account number
	var account model.Account
	ctx := context.Background()
	filter := bson.D{{Key: "accountNo", Value: accountNumber}}
	err := utils.Collection.FindOne(ctx, filter).Decode(&account)
	account.Balance += amount

	// Update the account balance in the database
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "balance", Value: account.Balance}}}}
	updateResult, err := utils.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"Found %v Account(s) and modified %v Account(s)\n",
		updateResult.MatchedCount,
		updateResult.ModifiedCount,
	)
	fmt.Println("Amount Added Successfully !!")
	utils.Return()
}
