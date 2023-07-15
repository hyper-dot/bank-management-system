package menu

import (
	"bank/model"
	"bank/utils"
	"context"
	"fmt"
	"log"

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
	// Returns to parent function
	utils.ReturnTo(SearchAccount)
}

func SearchByName() {
	// Retrieve all accounts from MongoDB
	var firstName string
	var lastName string

	// Asks User for name
	fmt.Printf("Enter the Account Name: ")
	fmt.Scan(&firstName, &lastName)

	fullName := firstName + " " + lastName
	// To track whether account is found or not
	var found bool

	// Search Filter
	filter := bson.D{{Key: "name", Value: fullName}}
	cursor, err := utils.Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var account model.Account

		if err := cursor.Decode(&account); err != nil {
			log.Fatal(err)
		}

		// Output
		fmt.Println("--------------------------")
		fmt.Println("Name      :", account.Name)
		fmt.Println("Accoun No :", account.AccountNo)
		fmt.Println("Phone No  :", account.Phone)
		fmt.Println("Balance   :", account.Balance)

		found = true
	}
	if found {
		fmt.Println("--------------------------")
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	if !found {
		fmt.Println("No account found !!")
	}

	// Returns to parent function
	utils.ReturnTo(SearchAccount)
}

func SearchAccountbyAccNo() {
	var account model.Account
	fmt.Print("Enter Account Number: ")
	var acNo int
	fmt.Scan(&acNo)
	ctx := context.Background()
	filter := bson.D{{Key: "accountNo", Value: acNo}}
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

	// Returns to parent function
	utils.ReturnTo(SearchAccount)
}
