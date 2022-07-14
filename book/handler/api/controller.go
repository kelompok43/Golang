package handlerAPI

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/kelompok43/Golang/book/domain"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewBookHandler(service domain.Service) BookHandler {
	return BookHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (th BookHandler) AddOnlineClass(ctx echo.Context) error {
	var req RequestOnlineClassJSON
	ctx.Bind(&req)
	errVal := th.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := th.service.InsertOnlineClass(toOnlineClassDomain(req))

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

func (th BookHandler) GetAllOnlineClass(ctx echo.Context) error {
	onlineRes, err := th.service.GetAllOnlineClass()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	onlineObj := []ResponseOnlineClassJSON{}

	for _, value := range onlineRes {
		onlineObj = append(onlineObj, fromOnlineClassDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    onlineObj,
	})
}

func (th BookHandler) GetOnlineClassByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	onlineRes, err := th.service.GetOnlineClassByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	onlineObj := fromOnlineClassDomain(onlineRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    onlineObj,
	})
}

func (th BookHandler) UpdateOnlineClass(ctx echo.Context) error {
	var req RequestOnlineClassJSON
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	onlineRes, err := th.service.UpdateOnlineClass(id, toOnlineClassDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	onlineObj := fromOnlineClassDomain(onlineRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    onlineObj,
	})
}

func (th BookHandler) DeleteOnlineClass(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.service.DeleteOnlineClass(id)

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

func (th BookHandler) AddOfflineClass(ctx echo.Context) error {
	var req RequestOfflineClassJSON
	ctx.Bind(&req)
	errVal := th.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := th.service.InsertOfflineClass(toOfflineClassDomain(req))

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

func (th BookHandler) GetAllOfflineClass(ctx echo.Context) error {
	offlineRes, err := th.service.GetAllOfflineClass()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	offlineObj := []ResponseOfflineClassJSON{}

	for _, value := range offlineRes {
		offlineObj = append(offlineObj, fromOfflineClassDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    offlineObj,
	})
}

func (th BookHandler) GetOfflineClassByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	offlineRes, err := th.service.GetOfflineClassByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	offlineObj := fromOfflineClassDomain(offlineRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    offlineObj,
	})
}

func (th BookHandler) UpdateOfflineClass(ctx echo.Context) error {
	var req RequestOfflineClassJSON
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	offlineRes, err := th.service.UpdateOfflineClass(id, toOfflineClassDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	offlineObj := fromOfflineClassDomain(offlineRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    offlineObj,
	})
}

func (th BookHandler) DeleteOfflineClass(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.service.DeleteOfflineClass(id)

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
