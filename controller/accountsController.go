package controller

import (
	"bank/model"
	"bank/utils"
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateAccount() {
	var account model.Account
	var firstName string
	var lastName string

	// Title
	fmt.Println("---------------------------------------")
	fmt.Println("Create Account")
	fmt.Println("---------------------------------------")

	// Scans firstname and lastname
	fmt.Printf("FullName(first and last): ")
	fmt.Scan(&firstName, &lastName)

	// Sets Accounts Name
	account.Name = firstName + " " + lastName

	// Scans account number and set it in account
	fmt.Printf("Account Number :")
	fmt.Scan(&account.AccountNo)

	// Scans account number and set it in account
	fmt.Printf("Phone Number :")
	fmt.Scan(&account.Phone)

	//Scans balance and set it in balance
	fmt.Printf("Balance :")
	fmt.Scan(&account.Balance)

	//Mongodb specifics
	_, err := utils.Collection.InsertOne(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}

	// Message
	fmt.Println("Account Created Succesfully !!")
	utils.Return()
}

func ShowAllAccount() {
	// Retrieve all accounts from MongoDB
	cursor, err := utils.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Create tab writer with custom formatting
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, " Name\t Account Number\t Phone\t Balance")
	fmt.Fprintln(
		w,
		"--------------------\t--------------------\t--------------------\t--------------------",
	)

	for cursor.Next(context.Background()) {
		var account model.Account

		if err := cursor.Decode(&account); err != nil {
			log.Fatal(err)
		}

		// Format and write account details to the tab writer
		fmt.Fprintf(
			w,
			"%s \t %d \t %d \t $ %.2f\n",
			account.Name,
			account.AccountNo,
			account.Phone,
			account.Balance,
		)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Flush the tab writer to output the formatted table
	w.Flush()

	utils.Return()
}

func DeleteAccount() {
	// Take Account Number
	var accountNo int
	fmt.Print("Enter Account Number: ")
	fmt.Scan(&accountNo)

	ctx := context.TODO()
	filter := bson.D{{Key: "accountNo", Value: accountNo}}
	deleteResult, err := utils.Collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	if deleteResult.DeletedCount > 1 {
		fmt.Printf("Found %v Account(s) and Deleted Succesfully!! \n", deleteResult.DeletedCount)
	} else {
		fmt.Println("No accounts found !!")
	}
	utils.Return()
}
