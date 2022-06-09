package handlerAPI

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/kelompok43/Golang/user/domain"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewUserHandler(service domain.Service) UserHandler {
	return UserHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (uh UserHandler) Register(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := uh.validation.Struct(req)
	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := uh.service.InsertData(toDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
	})
}

func (uh UserHandler) GetAllData(ctx echo.Context) error {
	userRes, _ := uh.service.GetAllData()

	// if err != nil {
	// 	return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"message": err.Error(),
	// 		"rescode": http.StatusInternalServerError,
	// 	})
	// }

	userObj := []ResponseJSON{}

	for _, value := range userRes {
		userObj = append(userObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    userObj,
	})
}

func (uh UserHandler) GetByID(ctx echo.Context) error {

	id, _ := strconv.Atoi(ctx.Param("id"))

	userRes, err := uh.service.GetByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userObj := fromDomain(userRes)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    userObj,
	})
}
