package migrations

import (
	"github.com/Yespolovaz/golangFinal/pkg/bank/database"
	"github.com/Yespolovaz/golangFinal/pkg/bank/helpers"
	"github.com/Yespolovaz/golangFinal/pkg/bank/interfaces"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createAccounts() {
	users := &[3]interfaces.User{
		{Username: "zharkynay", Email: "zh_yespolova@kbtu.kz", Name: "Zharkynay", Surname: "Surname"},
		{Username: "baleygr", Email: "d_mombekov@kbtu.kz", Name: "Dias", Surname: "Mombekov"},
		{Username: "nk", Email: "n_kidirmaganbetov@kbtu.kz", Name: "Nurken", Surname: "Surname"},
	}

	for i := 0; i < len(users); i++ {
		// MARK: - change subject of salting
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))

		user := &interfaces.User{
			Username: users[i].Username,
			Email:    users[i].Email,
			Name:     users[i].Username,
			Surname:  users[i].Surname,
			Password: generatedPassword}
		database.DB.Create(&user)

		account := &interfaces.Account{
			Type:    "Account",
			Name:    string(users[i].Username + "'s" + " account"),
			Balance: int(10000 * int(i+1)),
			UserId:  user.ID}
		database.DB.Create(&account)
	}
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	Transactions := &interfaces.Transaction{}
	database.DB.AutoMigrate(User, Account, Transactions)

	createAccounts()
}

func MigrateTransactions() {
	Transactions := &interfaces.Transaction{}

	db := helpers.ConnectDB()
	db.AutoMigrate(&Transactions)
	defer db.Close()
}
