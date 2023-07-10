package model

// Account Model
type Account struct {
	Name      string  `bson:"name"`
	AccountNo int     `bson:"accountNo"`
	Balance   float64 `bson:"balance"`
	Phone     int     `bson:"phone"`
}
