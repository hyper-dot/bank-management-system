package controller

import (
	"bank/model"
	"bank/utils"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// Account Management Handler

func SearchByPhone() {
	var account model.Account
	fmt.Print("Enter Phone Number: ")
	var phone int
	fmt.Scan(&phone)
	ctx := context.Background()
	filter := bson.D{{Key: "phone", Value: phone}}
	err := utils.Collection.FindOne(ctx, filter).Decode(&account)
	if err != nil {
		fmt.Println("Not found !!")
	}
	if err == nil {
		fmt.Println("---------------------------------------")
		fmt.Printf("Name        : %v\n", account.Name)
		fmt.Printf("Account No. : %v\n", account.AccountNo)
		fmt.Printf("Phone       : %v\n", account.Phone)
		fmt.Printf("Balance     : %v\n", account.Balance)
		fmt.Println("---------------------------------------")
	}

	utils.Return()
}

func SearchByName() {
	fmt.Print("Enter Full Name: ")
	return
}

func SearchAccountbyAccNo() {
	fmt.Print("Enter Account Number: ")
	return
}
