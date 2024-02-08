package main

import (
	"github.com/emaforlin/banking-api/config"
	"github.com/emaforlin/banking-api/database"
	"github.com/emaforlin/banking-api/server"
)

func main() {
	cfg := config.GetConfig()
	database := database.NewPosgresDatabase(&cfg)
	server.NewGinServer(&cfg, database.GetDb()).Start()

}
