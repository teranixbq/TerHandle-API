package handler

import (
	"fmt"
	"net/http"
	"strconv"

	//"strings"
	"terhandle/internal/features/user/dto"
	"terhandle/internal/features/user/entity"
	"terhandle/internal/utils/helper"
	"terhandle/internal/utils/jwt"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService entity.UserServiceInterface
}

func NewUserHandler(userService entity.UserServiceInterface) *userHandler {
	return &userHandler{userService: userService}
}

func (uc *userHandler) Create(e echo.Context) error {
	input := dto.RequestCreate{}

	if err := e.Bind(&input); err != nil {
		return err
	}
	inputmain := dto.RequestToCore(input)
	err := uc.userService.Create(inputmain)
	if err != nil {

		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))

	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Succes Create User"))
}

func (uc *userHandler) Login(e echo.Context) error {
	input := dto.RequestLogin{}
	inputmain := dto.RequestLoginToCore(input)

	if err := e.Bind(&inputmain); err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	fmt.Println(inputmain)

	user, token, err := uc.userService.Login(inputmain.Email, inputmain.Password)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}
	jwt.SetTokenCookie(e, token)
	response := dto.ResponseLogin{
		Name: user.Nama,
	}
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Succes Login", response))
}

func (uc *userHandler) CreateDetail(e echo.Context) error {
	input := dto.RequestCreateDetail{}

	user_id, role := jwt.ExtractToken(e)
	if user_id == 0 {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}

	if id != user_id {
		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
	}

	if role == "admin" {
		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access Denied"))
	}

	if err := e.Bind(&input); err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Request"))
	}

	err = uc.userService.CreateUpdateDetail(uint(user_id), dto.RequestCreateDetailToCore(input))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("success memperbarui data"))

}

func (uc *userHandler) CreateTeknisiRole(e echo.Context) error {
	user_id, role := jwt.ExtractToken(e)
	// if user_id == 0 {
	// 	return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	// }

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}

	if id != user_id {
		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
	}
	if role == "admin" {
		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access Denied"))
	}
	err = uc.userService.RequestTeknisiRole(uint(user_id))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("success request teknisi"))

}
func (uc *userHandler) GetAll(e echo.Context) error {
	users, err := uc.userService.GetAll()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.FailedResponse("Error"))
	}

	respons := dto.CoreToResponseByIdList(users)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("success", respons))
}

func (uc *userHandler) GetById(e echo.Context) error {
	user_id, role := jwt.ExtractToken(e)
	if user_id == 0 {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}

	// id, err := strconv.Atoi(e.Param("id"))
	// if err != nil {
	// 	return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	// }

	// if id != user_id {
	// 	return e.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
	// }

	if role == "admin" {
		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access Denied"))
	}
	response, err := uc.userService.GetUserById(uint(user_id), role)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	responseProfile := dto.ResponsesProfileList(response, role)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Data Profile "+role, responseProfile))

}

func (uc *userHandler) GetByIdWithFeedback(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}

	if id == 0 {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Teknisi tidak ada"))
	}

	users, err := uc.userService.GetById(uint(id))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.FailedResponse("Error"))
	}

	respons := dto.CoreToResponseList(users)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("success", respons))
}

func (uc *userHandler) DeleteById(e echo.Context) error {
	user_id, _ := jwt.ExtractToken(e)

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}

	if id == 0 {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Teknisi tidak ada"))
	}

	if id != user_id {
		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
	}

	err = uc.userService.DeleteById(uint(user_id))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.FailedResponse("Error"))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Success Delete Account"))
}
