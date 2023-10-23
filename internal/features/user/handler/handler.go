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
		
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(http.StatusBadRequest,err.Error()))

	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Succes Create User"))
}

func (uc *userHandler) Login(e echo.Context) error {
	input := dto.RequestLogin{}
	inputmain := dto.RequestLoginToCore(input)

	if err := e.Bind(&inputmain); err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(400,err.Error()))
	}

	fmt.Println(inputmain)

	user, token, err := uc.userService.Login(inputmain.Email, inputmain.Password)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(400,err.Error()))
	}

	response := dto.ResponseLogin{
		Name : user.Name,
		Token: token,
	}

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Succes Login", response))
}

func (uc *userHandler) CreateDetail(e echo.Context) error {
	input := dto.RequestCreateDetail{}

	user_id, _ := jwt.ExtractToken(e)
	if user_id == 0 {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse(401,"Unauthorized"))
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest,helper.FailedResponse(400,"Failed Convert"))
	}

	

	if id != user_id {
        return e.JSON(http.StatusForbidden,helper.FailedResponse(401,"Access denied"))
    }

	if err := e.Bind(&input); err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(400,"Invalid Request"))
	}

	err = uc.userService.CreateUpdateDetail(user_id, dto.RequestCreateDetailToCore(input))
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("success deactivate account "))

}


func (uc *userHandler) CreateTeknisiRole(e echo.Context) error {
	user_id, _ := jwt.ExtractToken(e)
	// if user_id == 0 {
	// 	return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	// }
	fmt.Println(user_id)
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest,helper.FailedResponse(400,"Failed Convert"))
	}

	if id != user_id {
        return e.JSON(http.StatusForbidden,helper.FailedResponse(401,"Access denied"))
    }

	err = uc.userService.RequestTeknisiRole(user_id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.FailedResponse(500,err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("success request teknisi"))

}
