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
	var req RequestCategoryJSON
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
	categoryRes, err := mh.service.GetAllCategory()

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

func (mh MembershipHandler) GetCategoryByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	categoryRes, err := mh.service.GetCategoryByID(id)

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

func (mh MembershipHandler) UpdateCategory(ctx echo.Context) error {
	var req RequestCategoryJSON
	ctx.Bind(&req)
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := mh.validation.Struct(req)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	categoryRes, err := mh.service.UpdateCategory(id, toCategoryDomain(req))

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

func (mh MembershipHandler) GetAllData(ctx echo.Context) error {
	membershipRes, err := mh.service.GetAllData()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	membershipObj := []ResponseJSON{}

	for _, value := range membershipRes {
		membershipObj = append(membershipObj, fromDomain(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    membershipObj,
	})
}

func (mh MembershipHandler) GetByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	membershipRes, err := mh.service.GetByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	membershipObj := fromDomain(membershipRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    membershipObj,
	})

}

func (mh MembershipHandler) GetByUserID(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}

	membershipRes, err := mh.service.GetByUserID(userId)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}

	membershipObj := fromDomain(membershipRes)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"data":    membershipObj,
	})
}
