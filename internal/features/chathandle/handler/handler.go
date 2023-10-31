package handler

import (
	"log"
	"net/http"

	"os"
	"terhandle/internal/features/chathandle/dto"
	"terhandle/internal/features/chathandle/service"
	"terhandle/internal/utils/helper"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type ChatHandleHandler struct {
	service service.ChatHandleService
}

func NewChatHandleHandler(uc service.ChatHandleService) *ChatHandleHandler {
	return &ChatHandleHandler{service: uc}
}

func (uc *ChatHandleHandler) GetChatSolution(c echo.Context) error {
	var request dto.RequestChat

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	answer, err := uc.service.ChatSolution(request, os.Getenv("OPEN_API_KEY"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	responseDTO := dto.ResponseRecomended{
		Status:     "Success",
		ChatHandle: answer,
	}

	return c.JSON(http.StatusOK, responseDTO)
}
