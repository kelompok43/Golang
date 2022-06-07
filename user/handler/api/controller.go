package handlerAPI

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/kelompok43/Golang/user/domain"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service    domain.Service
	Validation *validator.Validate
}

func NewUserHandler(service domain.Service) UserHandler {
	return UserHandler{
		Service:    service,
		Validation: validator.New(),
	}
}

func (uh UserHandler) Register(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
	})
}
