package handler

import (
	"net/http"
	"strconv"

	//"strings"
	"terhandle/internal/features/request-teknisi/dto"
	"terhandle/internal/features/request-teknisi/entity"
	"terhandle/internal/utils/helper"
	"terhandle/internal/utils/jwt"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService entity.UserRequestServiceInterface
}

func NewUserHandler(userService entity.UserRequestServiceInterface) *userHandler {
	return &userHandler{userService: userService}
}

func (uc *userHandler) Create(e echo.Context) error {
	input := dto.RequestCreate{}

	user_id, role := jwt.ExtractToken(e)

	if user_id == 0 {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}

	if role == "teknisi" {
		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
	}

	if err := e.Bind(&input); err != nil {
		return err
	}

	inputmain := dto.RequestCreateToCore(input)

	err := uc.userService.Create(inputmain)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Succes Request Teknisi"))
}

func (uc *userHandler) GetHistoryRequest(e echo.Context) error {
	role_id, role := jwt.ExtractToken(e)
	if role_id == 0 {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}

	if id != role_id {
		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
	}

	history,err := uc.userService.GetById(role_id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}
	
	respon := dto.ResponsesHistory(history,role)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("succes",respon))

}

// func (uc *userHandler) CreateDetail(e echo.Context) error {
// 	input := dto.RequestCreateDetail{}

// 	user_id, _ := jwt.ExtractToken(e)
// 	if user_id == 0 {
// 		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
// 	}

// 	id, err := strconv.Atoi(e.Param("id"))
// 	if err != nil {
// 		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
// 	}

// 	if id != user_id {
// 		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
// 	}

// 	if err := e.Bind(&input); err != nil {
// 		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Request"))
// 	}

// 	err = uc.userService.CreateUpdateDetail(user_id, dto.RequestCreateDetailToCore(input))
// 	if err != nil {
// 		return err
// 	}

// 	return e.JSON(http.StatusOK, helper.SuccessResponse("success memperbarui data"))

// }

// func (uc *userHandler) SelectAll(e echo.Context) error {
// 	users, err := uc.userService.GetAll()
// 	if err != nil {
// 		return e.JSON(http.StatusInternalServerError, helper.FailedResponse("Error"))
// 	}

// 	// Menggunakan CoreToResponseList untuk mengubah data ke format yang diinginkan.
// 	respons := dto.CoreToResponseList(users)

// 	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("success", respons))
// }
