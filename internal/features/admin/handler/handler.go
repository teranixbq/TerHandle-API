package handler

import (
	"net/http"
	"strconv"
	"terhandle/internal/features/admin/dto"
	"terhandle/internal/features/admin/entity"
	"terhandle/internal/utils/helper"
	"terhandle/internal/utils/jwt"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService entity.AdminServiceInterface
}

func NewUserHandler(AdminService entity.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{AdminService: AdminService}
}

func (uc *AdminHandler) Create(e echo.Context) error {
	input := dto.RequestCreate{}

	_, role := jwt.ExtractToken(e)

	if role != "admin" {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Acces Denied"))
	}

	if err := e.Bind(&input); err != nil {
		return err
	}

	inputModel := dto.RequestCreateToCore(input)
	err := uc.AdminService.Create(inputModel)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Success Create Profesi"))
}

func (uc *AdminHandler) GetAllProfesi(e echo.Context) error {
	profesi, err := uc.AdminService.GetAll()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.FailedResponse("Error"))
	}

	respons := dto.CoreToResponseProfesiList(profesi)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Data Profesi", respons))
}

func (uc *AdminHandler) UpdateDataProfesi(e echo.Context) error {
	input := dto.RequestCreate{}

	_, role := jwt.ExtractToken(e)
	if role != "admin" {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Acces Denied"))
	}

	id_profesi,err := strconv.Atoi(e.Param("id_profesi"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}


	if err := e.Bind(&input); err != nil {
		return err
	}

	inputModel := dto.RequestCreateToCore(input)

	err = uc.AdminService.Update(uint(id_profesi),inputModel)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Success Update Profesi"))
}

func (uc *AdminHandler) DeleteDataProfesi(e echo.Context) error {
	_, role := jwt.ExtractToken(e)
	if role != "admin" {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Acces Denied"))
	}

	id_profesi,err := strconv.Atoi(e.Param("id_profesi"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}

	err = uc.AdminService.Delete(uint(id_profesi))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Success Delete Profesi"))
}



func (uc *AdminHandler) CreateBiaya(e echo.Context) error {
	input := dto.RequestCreateBiaya{}
	_, role := jwt.ExtractToken(e)

	if role != "admin" {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Acces Denied"))
	}

	if err := e.Bind(&input); err != nil {
		return err
	}

	inputModel := dto.RequestCreateBiayaToCore(input)
	err := uc.AdminService.InsertBiaya(inputModel)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Success Create Biaya Transport"))
}