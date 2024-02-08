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
		{FullName: "Alejandro Forlin", Funds: 10549, Username: "alef11", Password: "11012001"},
		{FullName: "Emanuel Forlin", Funds: 12403, Username: "emaf21", Password: "21082003"},
		{FullName: "Alejandro Forlin", Funds: 8096, Username: "maurif27", Password: "27042006"},
	}, 10)
}
