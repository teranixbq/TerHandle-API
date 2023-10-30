package service

import (
	"context"
	"os"
	"terhandle/internal/features/chathandle/dto"
	"github.com/sashabaranov/go-openai"
)

type ChatHandleService interface {
	ChatSolution(request dto.RequestChat, key string) (string, error)
}

type ServiceHandle struct{}

func NewServiceHandle() ChatHandleService {
	return &ServiceHandle{}
}

func (uc *ServiceHandle) getCompletionMessages(ctx context.Context, client *openai.Client, messages []openai.ChatCompletionMessage, model string) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}

func (us *ServiceHandle) ChatSolution(request dto.RequestChat, key string) (string, error) {
	filePath := "internal/utils/prompt/ChatHandlePrompt.txt" 
	
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "",err
	}

	ctx := context.Background()
	client := openai.NewClient(key)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: string(content),
		},
		{
			Role:    "user",
			Content: request.Pertanyaan,
		},
	}

	response, err := us.getCompletionMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := response.Choices[0].Message.Content
	return answer, nil
}
