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
	if err := c.ShouldBindJSON(reqBody); err != nil {
		log.Errorf("error binding request body: %v", err)
		response(c, http.StatusBadRequest, "Bad request", reqBody)
		return
	}

	if err := h.userUsecase.RegisterUser(reqBody); err != nil {
		response(c, http.StatusInternalServerError, "Processing data failed", nil)
		return
	}
	response(c, http.StatusOK, "Created", reqBody)
}

func (h *userHttpHandler) QueryUser(c *gin.Context) {
	name := c.Param("fullname")
	obtainedUserData, err := h.userUsecase.GetUserByName(name)
	if err != nil {
		response(c, http.StatusInternalServerError, "error obtaining data", "")
	}
	response(c, http.StatusOK, "ok", *obtainedUserData)

}
