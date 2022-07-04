package handlerAPI

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/kelompok43/Golang/membership/domain"
	"github.com/labstack/echo/v4"
)

type MembershipHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewMembershipHandler(service domain.Service) MembershipHandler {
	return MembershipHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (mh MembershipHandler) AddCategory(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := mh.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := mh.service.InsertCategory(toCategoryDomain(req))

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

func (mh MembershipHandler) GetAllCategory(ctx echo.Context) error {
	trainerRes, err := mh.service.GetAllCategory()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	trainerObj := []ResponseJSON{}

	for _, value := range trainerRes {
		trainerObj = append(trainerObj, fromCategoryDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    trainerObj,
	})
}

func (mh MembershipHandler) GetCategoryByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	trainerRes, err := mh.service.GetCategoryByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	trainerObj := fromCategoryDomain(trainerRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    trainerObj,
	})
}

func (mh MembershipHandler) UpdateCategory(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := mh.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	trainerRes, err := mh.service.UpdateCategory(id, toCategoryDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	trainerObj := fromCategoryDomain(trainerRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    trainerObj,
	})
}

func (mh MembershipHandler) DeleteCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := mh.service.DeleteCategory(id)

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
