package menu

import (
	"bank/model"
	"bank/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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

	// Returns to Parent Menu
	utils.ReturnTo(ShowDepositeAndWithdrawlMenu)
}

func TransferMoney() {
	//Get account numbers
	var receiverAccNo int
	var senderAccNo int
	var amount float64
	fmt.Printf("Receiver Account No : ")
	fmt.Scan(&receiverAccNo)
	fmt.Printf("Sender Account No :")
	fmt.Scan(&senderAccNo)
	fmt.Printf("Amount :")
	fmt.Scan(&amount)

	//Get details of the account number
	var receiver model.Account
	var sender model.Account
	ctx := context.Background()
	receiverFilter := bson.D{{Key: "accountNo", Value: receiverAccNo}}
	senderFilter := bson.D{{Key: "accountNo", Value: senderAccNo}}
	if err := utils.Collection.FindOne(ctx, receiverFilter).Decode(&receiver); err != nil {
		fmt.Println("receiver's account not found !!")
		return
	}
	if err := utils.Collection.FindOne(ctx, senderFilter).Decode(&sender); err != nil {
		fmt.Println("sender's account not found !!")
		return
	}

	// Checks if sender has enough balance
	if sender.Balance < amount {
		fmt.Println("sender don't have enough balance !!")
		return
	}

	// Add Subtract the amounts
	senderNewAmount := sender.Balance - amount
	receiverNewAmount := receiver.Balance + amount

	// Update the values in database
	updateReceiver := bson.D{
		{Key: "$set", Value: bson.D{{Key: "balance", Value: receiverNewAmount}}},
	}
	updateSender := bson.D{{Key: "$set", Value: bson.D{{Key: "balance", Value: senderNewAmount}}}}

	if _, err := utils.Collection.UpdateOne(ctx, receiverFilter, updateReceiver); err != nil {
		fmt.Println("receiver's account updation failed!!")
		return
	}
	if _, err := utils.Collection.UpdateOne(ctx, senderFilter, updateSender); err != nil {
		fmt.Println("sender's account updation failed!!")
		return
	}
	// Print success message and return to menu
	fmt.Println("Account Updated Succesfully!!!")

	// Returns to Parent Menu
	utils.ReturnTo(ShowDepositeAndWithdrawlMenu)
}
