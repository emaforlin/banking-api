package main

import (
	"github.com/emaforlin/banking-api/banking/entities"
	"github.com/emaforlin/banking-api/config"
	"github.com/emaforlin/banking-api/database"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPosgresDatabase(&cfg)
	db.GetDb().Migrator().CreateTable(&entities.User{})
	db.GetDb().CreateInBatches([]entities.User{
		{FullName: "Alejandro Forlin", Funds: 10549},
		{FullName: "Emanuel Forlin", Funds: 12403},
		{FullName: "Alejandro Forlin", Funds: 8096},
	}, 10)
}
