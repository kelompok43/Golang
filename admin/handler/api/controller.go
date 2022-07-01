package handlerAPI

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/kelompok43/Golang/admin/domain"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewAdminHandler(service domain.Service) AdminHandler {
	return AdminHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (ah AdminHandler) Register(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	errVal := ah.validation.Struct(req)

	if errVal != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errVal.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	_, err := ah.service.InsertData(toDomain(req))

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

func (ah AdminHandler) Login(ctx echo.Context) error {
	var req RequestLoginJSON
	ctx.Bind(&req)
	err := ah.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	email := req.Email
	password := req.Password
	token, adminRes, err := ah.service.CreateToken(email, password)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	adminObj := fromDomain(adminRes)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"token":   token,
		"data":    adminObj,
	})
}

func (ah AdminHandler) GetAllData(ctx echo.Context) error {
	adminRes, err := ah.service.GetAllData()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	adminObj := []ResponseJSON{}

	for _, value := range adminRes {
		adminObj = append(adminObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    adminObj,
	})
}

func (ah AdminHandler) GetByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	adminRes, err := ah.service.GetByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	adminObj := fromDomain(adminRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    adminObj,
	})
}

func (ah AdminHandler) GetByEmail(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	email := req.Email
	adminRes, err := ah.service.GetByEmail(email)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	adminObj := fromDomain(adminRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    adminObj,
	})
}

func (ah AdminHandler) ChangePassword(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}
	adminRes, err := ah.service.ChangePassword(id, toDomain(req))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	adminObj := fromDomain(adminRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    adminObj,
	})
}

func (ah AdminHandler) AdminRole(id int) (role string, err error) {
	adminObj, err := ah.service.GetByID(id)

	if err != nil {
		return "", err
	}

	role = adminObj.Role
	return role, err
}
