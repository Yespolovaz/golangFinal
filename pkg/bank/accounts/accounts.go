package accounts

import (
	"fmt"

	"github.com/Yespolovaz/golangFinal/pkg/bank/database"
	"github.com/Yespolovaz/golangFinal/pkg/bank/helpers"
	"github.com/Yespolovaz/golangFinal/pkg/bank/interfaces"
	"github.com/Yespolovaz/golangFinal/pkg/bank/transactions"
)

func updateAccount(id uint, amount int) interfaces.ResponseAccount {
	account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}

	database.DB.Where("id = ? ", id).First(&account)
	account.Balance = amount
	database.DB.Save(&account)

	responseAcc.Id = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = int(account.Balance)
	return responseAcc
}

func getAccount(id uint) *interfaces.Account {
	account := &interfaces.Account{}
	if database.DB.Where("id = ? ", id).First(&account).RecordNotFound() {
		return nil
	}
	return account
}

func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
	userIdString := fmt.Sprint(userId)
	isValid := helpers.ValidateToken(userIdString, jwt)
	if isValid {

	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
	fromAccount := getAccount(from)
	toAccount := getAccount(to)

	if fromAccount == nil || toAccount == nil {
		return map[string]interface{}{"message": "Account not found"}
	} else if fromAccount.UserId != userId {
		return map[string]interface{}{"message": "You are not owner of the account"}
	} else if int(fromAccount.Balance) < amount {
		return map[string]interface{}{"message": "Account balance is too small"}
	}

	updatedAccount := updateAccount(from, int(fromAccount.Balance)-amount)
	updateAccount(to, int(toAccount.Balance)+amount)

	transactions.CreateTransaction(from, to, amount)
	var response = map[string]interface{}{"message": "all is fine"}
	response["data"] = updatedAccount
	return response
}
