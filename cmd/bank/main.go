package main

import (
	"github.com/Yespolovaz/golangFinal/cmd/bank/api"
	"github.com/Yespolovaz/golangFinal/pkg/bank/database"
	// "github.com/Yespolovaz/golangFinal/pkg/bank/migrations"
)

func main() {
	database.InitDatabase()
	// migrations.Migrate()
	// migrations.MigrateTransactions()
	api.StartApi()
}
