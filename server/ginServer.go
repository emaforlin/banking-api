package server

import (
	"fmt"

	"github.com/emaforlin/banking-api/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ginServer struct {
	app *gin.Engine
	db  *gorm.DB
	cfg *config.Config
}

func NewGinServer(cfg *config.Config, db *gorm.DB) Server {
	return &ginServer{
		app: gin.New(),
		db:  db,
		cfg: cfg,
	}
}

func (gs *ginServer) Start() {
	// init routers here ...

	gs.app.Use(gin.Logger())

	serverAddr := fmt.Sprintf(":%d", gs.cfg.App.Port)
	gs.app.Run(serverAddr)
}
