package handlerAPI

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/kelompok43/Golang/class/domain"
	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewClassHandler(service domain.Service) ClassHandler {
	return ClassHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (th ClassHandler) AddCategory(ctx echo.Context) error {
	var req RequestCategoryJSON
	ctx.Bind(&req)
	errVal := th.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	picture, _ := ctx.FormFile("picture")
	src, _ := picture.Open()
	defer src.Close()
	req.Picture = src
	_, err := th.service.InsertCategory(toCategoryDomain(req))

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

func (th ClassHandler) GetAllCategory(ctx echo.Context) error {
	categoryRes, err := th.service.GetAllCategory()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	categoryObj := []ResponseCategoryJSON{}

	for _, value := range categoryRes {
		categoryObj = append(categoryObj, fromCategoryDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    categoryObj,
	})
}

func (th ClassHandler) GetCategoryByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	categoryRes, err := th.service.GetCategoryByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	categoryObj := fromCategoryDomain(categoryRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    categoryObj,
	})
}

func (th ClassHandler) UpdateCategory(ctx echo.Context) error {
	var req RequestCategoryJSON
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	categoryRes, err := th.service.UpdateCategory(id, toCategoryDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	categoryObj := fromCategoryDomain(categoryRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    categoryObj,
	})
}

func (th ClassHandler) DeleteCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.service.DeleteCategory(id)

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

func (th ClassHandler) AddOnline(ctx echo.Context) error {
	var req RequestOnlineJSON
	ctx.Bind(&req)
	errVal := th.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := th.service.InsertOnline(toOnlineDomain(req))

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

func (th ClassHandler) GetAllOnline(ctx echo.Context) error {
	onlineRes, err := th.service.GetAllOnline()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	onlineObj := []ResponseOnlineJSON{}

	for _, value := range onlineRes {
		onlineObj = append(onlineObj, fromOnlineDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    onlineObj,
	})
}

func (th ClassHandler) GetOnlineByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	onlineRes, err := th.service.GetOnlineByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	onlineObj := fromOnlineDomain(onlineRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    onlineObj,
	})
}

func (th ClassHandler) UpdateOnline(ctx echo.Context) error {
	var req RequestOnlineJSON
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	onlineRes, err := th.service.UpdateOnline(id, toOnlineDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	onlineObj := fromOnlineDomain(onlineRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    onlineObj,
	})
}

func (th ClassHandler) DeleteOnline(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.service.DeleteOnline(id)

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

func (th ClassHandler) AddOffline(ctx echo.Context) error {
	var req RequestOfflineJSON
	ctx.Bind(&req)
	errVal := th.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := th.service.InsertOffline(toOfflineDomain(req))

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

func (th ClassHandler) GetAllOffline(ctx echo.Context) error {
	offlineRes, err := th.service.GetAllOffline()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	offlineObj := []ResponseOfflineJSON{}

	for _, value := range offlineRes {
		offlineObj = append(offlineObj, fromOfflineDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    offlineObj,
	})
}

func (th ClassHandler) GetOfflineByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	offlineRes, err := th.service.GetOfflineByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	offlineObj := fromOfflineDomain(offlineRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    offlineObj,
	})
}

func (th ClassHandler) UpdateOffline(ctx echo.Context) error {
	var req RequestOfflineJSON
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	offlineRes, err := th.service.UpdateOffline(id, toOfflineDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	offlineObj := fromOfflineDomain(offlineRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    offlineObj,
	})
}

func (th ClassHandler) DeleteOffline(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := th.service.DeleteOffline(id)

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
