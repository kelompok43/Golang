package handlerAPI

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/kelompok43/Golang/transaction/domain"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	service domain.Service
	// userService userDomain.Service
	// pmService   pmDomain.Service
	validation *validator.Validate
}

func NewTransactionHandler(service domain.Service) TransactionHandler {
	return TransactionHandler{
		service: service,
		// userService: userService,
		// pmService:   pmService,
		validation: validator.New(),
	}
}

func (th TransactionHandler) AddData(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := th.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	picture, _ := ctx.FormFile("payment_receipt")
	src, _ := picture.Open()
	defer src.Close()
	req.PaymentReceipt = src
	_, err := th.service.InsertData(toDomain(req))

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

func (th TransactionHandler) GetAllData(ctx echo.Context) error {
	transactionRes, err := th.service.GetAllData()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	transactionObj := []ResponseJSON{}

	for _, value := range transactionRes {
		transactionObj = append(transactionObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    transactionObj,
	})
}

func (th TransactionHandler) GetByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	transactionRes, err := th.service.GetByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	transactionObj := fromDomain(transactionRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    transactionObj,
	})
}

func (th TransactionHandler) UpdateData(ctx echo.Context) error {
	var req RequestStatus
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	transactionRes, err := th.service.UpdateData(id, statusToDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	//add user to membership
	// if transactionRes.Status == "Diterima"{

	// }

	transactionObj := fromDomain(transactionRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    transactionObj,
	})
}

func (th TransactionHandler) DeleteData(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.service.DeleteData(id)

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
