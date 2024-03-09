package main

import (
	// "github.com/Yespolovaz/golangFinal/cmd",
	"github.com/Yespolovaz/golangFinal/pkg/bank/database"
	// "github.com/Yespolovaz/golangFinal/pkg/bank/migrations"
)

func main() {
	// migrations.Migrate()
	StartApi()
	// migrations.MigrateTransactions()
	database.InitDatabase()

}
