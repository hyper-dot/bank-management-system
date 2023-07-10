package controller

import (
	"bank/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"text/tabwriter"
)

func createConnection() *mongo.Client {
	ctx := context.TODO()
	URI := "mongodb://127.0.0.1:27017"
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func createAccountCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("Bank").Collection("accounts")
	return collection
}

var client = createConnection()
var collection = createAccountCollection(client)

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
	_, err := collection.InsertOne(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}

	// Message
	fmt.Println("Account Created Succesfully !!")
	fmt.Println("Insert 0 and Hit Enter to continue !")
	var choice int
	fmt.Scan(&choice)
	if choice == 0 {
		return
	}
	for choice != 0 {
		fmt.Println("Invalid Option!!")
		fmt.Scan(&choice)
	}
}

// func ShowAllAccount() {
// 	fmt.Println("---------------------------------------")
// 	fmt.Println("All Accounts")
// 	fmt.Println("---------------------------------------")
//
// 	// Retrieve all accounts from MongoDB
// 	cursor, err := collection.Find(context.Background(), bson.D{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cursor.Close(context.Background())
//
// 	fmt.Println("--------------------------------------------------------------")
// 	fmt.Println("Name           | Account Number      |     Phone     | Balance|")
// 	fmt.Println("--------------------------------------------------------------")
// 	for cursor.Next(context.Background()) {
// 		var account model.Account
//
// 		if err := cursor.Decode(&account); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf(
// 			"%v              | %v                   | %v            | %v    |\n",
// 			account.Name,
// 			account.AccountNo,
// 			account.Phone,
// 			account.Balance,
// 		)
// 		// fmt.Println("Name           : ", account.Name)
// 		// fmt.Println("Account Number : ", account.AccountNo)
// 		// fmt.Println("Phone          : ", account.Phone)
// 		// fmt.Println("Balance        : $", account.Balance)
// 		// fmt.Println("---------------------------------------")
// 		// fmt.Println()
// 	}
//
// 	if err := cursor.Err(); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Println("Press 0 And Hit Enter to return")
// 	var choice int
// 	fmt.Scan(&choice)
// 	if choice != 0 {
// 		fmt.Println("Invalid Options")
// 	}
// }

// func ShowAllAccount() {
//
//		// Retrieve all accounts from MongoDB
//		cursor, err := collection.Find(context.Background(), bson.D{})
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer cursor.Close(context.Background())
//
//		// Create tab writer with custom formatting
//		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)
//		fmt.Fprintln(w, "--------------------------------------------------------------")
//		fmt.Fprintln(w, "Name\tAccount Number\tPhone\tBalance")
//		fmt.Fprintln(w, "--------------------------------------------------------------")
//
//		for cursor.Next(context.Background()) {
//			var account model.Account
//
//			if err := cursor.Decode(&account); err != nil {
//				log.Fatal(err)
//			}
//
//			// Format and write account details to the tab writer
//			fmt.Fprintf(
//				w,
//				"%s\t%d\t%d\t%.2f\n",
//				account.Name,
//				account.AccountNo,
//				account.Phone,
//				account.Balance,
//			)
//		}
//
//		if err := cursor.Err(); err != nil {
//			log.Fatal(err)
//		}
//
//		// Flush the tab writer to output the formatted table
//		w.Flush()
//
//		fmt.Println("Press 0 And Hit Enter to return")
//		var choice int
//		fmt.Scan(&choice)
//		if choice != 0 {
//			fmt.Println("Invalid Options")
//		}
//	}
func ShowAllAccount() {
	// Retrieve all accounts from MongoDB
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Create tab writer with custom formatting
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, " Name\t Account Number\t Phone\t Balance")

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

	fmt.Println("Press 0 And Hit Enter to return")
	var choice int
	fmt.Scan(&choice)
	if choice != 0 {
		fmt.Println("Invalid Options")
	}
}

func SearchAccount() {

}
