package handler

import (
	"net/http"
	"strconv"

	"terhandle/internal/features/request-teknisi/dto"
	"terhandle/internal/features/request-teknisi/entity"
	"terhandle/internal/utils/cloudflare"
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

	if user_id == 0 || role == "teknisi" {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}

	if err := e.Bind(&input); err != nil {
		return err
	}

	fileForm, err := e.MultipartForm()
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Failed to receive files")
	}
	files := fileForm.File["foto"]

	fotos := []dto.RequestFoto{}
	for _, file := range files {
		url, err := cloudflare.UploadFile(file)
		if err != nil {
			return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		}

		fotos = append(fotos, dto.RequestFoto{Foto: url})
	}
	
	input.Foto = fotos
	inputmain := dto.RequestCreateToCore(input)

	err = uc.userService.Create(inputmain)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Success Request Teknisi"))
}

func (uc *userHandler) GetAllHistoryRequest(e echo.Context) error {
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

	history, err := uc.userService.GetAllById(role_id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	respon := dto.ResponsesHistoryList(history, role)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("succes", respon))

}

func (uc *userHandler) GetHistoryRequestById(e echo.Context) error {
	user_id, role := jwt.ExtractToken(e)
	if user_id == 0 {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}

	id, idErr := strconv.Atoi(e.Param("id"))
	id_request, idRequestErr := strconv.Atoi(e.Param("id_request"))

	if idErr != nil || idRequestErr != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}

	if id != user_id {
		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
	}

	result, err := uc.userService.GetById(id, id_request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	respon := dto.ResponsesHistoryList(result, role)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("succes", respon))

}

func (uc *userHandler) UpdateStatusRequest(e echo.Context) error {

	user_id, role := jwt.ExtractToken(e)
	if user_id == 0 {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}

	id, idErr := strconv.Atoi(e.Param("id"))
	id_request, idRequestErr := strconv.Atoi(e.Param("id_request"))

	if idErr != nil || idRequestErr != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}

	if id != user_id {
		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
	}

	if role == "teknisi" {
		input := dto.RequestClaims{}
		if err := e.Bind(&input); err != nil {
			return err
		}

		inputmodel := dto.RequestClaimsToCore(input)
		err := uc.userService.ClaimRequest(id_request, inputmodel)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
		}

		return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil"))
	} else if role == "user" {
		input := dto.RequestUpdateStatus{}
		if err := e.Bind(&input); err != nil {
			return err
		}

		inputmodel := dto.RequestUpdateStatusToCore(input)
		err := uc.userService.KonfirmasiBiaya(id, id_request, inputmodel)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
		}

		return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil"))
	} else {
		return e.JSON(http.StatusForbidden, helper.FailedResponse("Access denied"))
	}
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
