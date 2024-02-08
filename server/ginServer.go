package server

import (
	"fmt"

	userHandlers "github.com/emaforlin/banking-api/banking/handlers"
	userRepository "github.com/emaforlin/banking-api/banking/repositories"
	userUsecases "github.com/emaforlin/banking-api/banking/usecases"
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
	// init routers here
	gs.initUsersHttpHandler()
	// end of routers

	gs.app.Use(gin.Logger(), gin.Recovery())

	serverAddr := fmt.Sprintf(":%d", gs.cfg.App.Port)
	gs.app.Run(serverAddr)
}

func (gs *ginServer) initUsersHttpHandler() {
	usersRepository := userRepository.NewUserPostgresRepository(gs.db)
	usersUsecase := userUsecases.NewUserUsecaseImpl(usersRepository)
	usersHttpHandler := userHandlers.NewUserHttpHandler(usersUsecase)

	// Routers
	usersRouters := gs.app.Group("v1/users")
	{
		usersRouters.POST("/register", usersHttpHandler.RegisterUser)
	}
}
