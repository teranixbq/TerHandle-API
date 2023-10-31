package handler

import (
	"net/http"
	"strconv"

	"terhandle/internal/features/feedback/dto"
	"terhandle/internal/features/feedback/entity"
	"terhandle/internal/utils/helper"
	"terhandle/internal/utils/jwt"

	"github.com/labstack/echo/v4"
)

type feedbackHandler struct {
	feedbackService entity.FeedbackServiceInterface
}

func NewUserHandler(feedbackService entity.FeedbackServiceInterface) *feedbackHandler {
	return &feedbackHandler{feedbackService: feedbackService}
}

func (uc *feedbackHandler) Create(e echo.Context) error {
	input := dto.RequestCreate{}

	user_id, role := jwt.ExtractToken(e)

	if user_id == 0 || role == "teknisi" {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}

	if user_id != id {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Acces Denied"))
	}

	if err := e.Bind(&input); err != nil {
		return err
	}

	if user_id != int(input.UsersID) {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Acces Denied"))

	}
	inputModel := dto.RequestCreateToCore(input)

	err = uc.feedbackService.Create(inputModel)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Success Create Feedback"))
}

func (uc *feedbackHandler) UpdateDataFeedback(e echo.Context) error {
	input := dto.RequestCreate{}

	user_id, role := jwt.ExtractToken(e)

	if user_id == 0 || role == "teknisi" {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}

	id_feedback, err := strconv.Atoi(e.Param("id_feedback"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse("Failed Convert"))
	}

	if user_id != id {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Acces Denied"))
	}

	if err := e.Bind(&input); err != nil {
		return err
	}

	if user_id != int(input.UsersID) {
		return e.JSON(http.StatusUnauthorized, helper.FailedResponse("Acces Denied"))

	}
	inputModel := dto.RequestCreateToCore(input)

	err = uc.feedbackService.UpdateFeedback(uint(user_id),uint(id_feedback),input.RequestTeknisiID, inputModel)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Success Create Feedback"))
}

