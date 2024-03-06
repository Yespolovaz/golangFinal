package accounts

import (
	"github.com/Yespolovaz/golangFinal/pkg/bank/helpers"
	"github.com/Yespolovaz/golangFinal/pkg/bank/interfaces"
)

func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB()

	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)

	defer db.Close()
}
