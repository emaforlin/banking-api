package handlers

import (
	"net/http"

	"github.com/emaforlin/banking-api/banking/models"
	"github.com/emaforlin/banking-api/banking/usecases"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type userHttpHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHttpHandler(useCase usecases.UserUsecase) UserHandler {
	return &userHttpHandler{
		userUsecase: useCase,
	}
}

func (h *userHttpHandler) RegisterUser(c *gin.Context) {
	reqBody := new(models.AddUserData)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("error binding request body: %v", err)
		response(c, http.StatusBadRequest, "Bad request")
	}
	response(c, http.StatusOK, "Created")
}
