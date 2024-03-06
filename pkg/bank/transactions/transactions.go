package transactions

import (
	"github.com/Yespolovaz/golangFinal/pkg/bank/helpers"
	"github.com/Yespolovaz/golangFinal/pkg/bank/interfaces"
)

func CreateTransaction(From uint, To uint, Amount int) {
	db := helpers.ConnectDB()
	transaction := &interfaces.Transaction{From: From, To: To, Amount: Amount}
	db.Create(&transaction)

	defer db.Close()
}
