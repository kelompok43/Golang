package handlerAPI

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/kelompok43/Golang/payment_method/domain"
	"github.com/labstack/echo/v4"
)

type PaymentMethodHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewPaymentMethodHandler(service domain.Service) PaymentMethodHandler {
	return PaymentMethodHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (pmh PaymentMethodHandler) AddData(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := pmh.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := pmh.service.InsertData(toDomain(req))

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

func (pmh PaymentMethodHandler) GetAllData(ctx echo.Context) error {
	trainerRes, err := pmh.service.GetAllData()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	trainerObj := []ResponseJSON{}

	for _, value := range trainerRes {
		trainerObj = append(trainerObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    trainerObj,
	})
}

func (pmh PaymentMethodHandler) GetByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	trainerRes, err := pmh.service.GetByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	trainerObj := fromDomain(trainerRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    trainerObj,
	})
}

func (pmh PaymentMethodHandler) UpdateData(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := pmh.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	trainerRes, err := pmh.service.UpdateData(id, toDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	trainerObj := fromDomain(trainerRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    trainerObj,
	})
}

func (pmh PaymentMethodHandler) DeleteData(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := pmh.service.DeleteData(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
	})
}
