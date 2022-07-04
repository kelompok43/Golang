package handlerAPI

import (
	"fmt"
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

func (uh UserHandler) Login(ctx echo.Context) error {
	var req RequestLoginJSON
	ctx.Bind(&req)
	err := uh.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	email := req.Email
	password := req.Password
	token, userRes, err := uh.service.CreateToken(email, password)

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
		"token":   token,
		"data":    userObj,
	})
}

func (uh UserHandler) GetAllData(ctx echo.Context) error {
	userRes, err := uh.service.GetAllData()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

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

func (uh UserHandler) AddDetail(ctx echo.Context) error {
	var req RequestDetailJSON
	ctx.Bind(&req)
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	req.UserID = id
	userRes, err := uh.service.InsertDetailData(detailToDomain(req))

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

func (uh UserHandler) GetByEmail(ctx echo.Context) error {
	email := ctx.QueryParam("email")
	fmt.Println(email)
	userRes, err := uh.service.GetByEmail(email)

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

func (uh UserHandler) ChangePassword(ctx echo.Context) error {
	var req RequestJSON
	ctx.Bind(&req)
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}
	userRes, err := uh.service.ChangePassword(id, toDomain(req))

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

func (uh UserHandler) UpdateStatus(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	userRes, err := uh.service.UpdateStatus(id)

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

func (uh UserHandler) UserStatus(id int) (status string, err error) {
	userObj, err := uh.service.GetByID(id)

	if err != nil {
		return "", err
	}

	status = userObj.Status
	return status, err
}
