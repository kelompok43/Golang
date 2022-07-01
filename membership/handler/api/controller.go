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

func (mh MembershipHandler) AddData(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := mh.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := mh.service.InsertData(toDomain(req))

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

func (mh MembershipHandler) GetAllData(ctx echo.Context) error {
	trainerRes, err := mh.service.GetAllData()

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

func (mh MembershipHandler) GetByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	trainerRes, err := mh.service.GetByID(id)

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

func (mh MembershipHandler) UpdateData(ctx echo.Context) error {
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

	trainerRes, err := mh.service.UpdateData(id, toDomain(req))

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

func (mh MembershipHandler) DeleteData(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := mh.service.DeleteData(id)

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
