package helper

import "net/http"

type ErrorResponseJson struct {
    Status_code int    `json:"status_code"`
    Status     string `json:"status"`
    Message    string `json:"message"`
}

type SuccessResponseJson struct {
    Status_code int    `json:"status_code"`
    Status     string `json:"status"`
    Message    string `json:"message"`
    Data       interface{} `json:"data,omitempty"`
}

func FailedResponse(status int, message string) ErrorResponseJson {
    return ErrorResponseJson{
        Status_code: status,
        Status:     "failed",
        Message:    message,
    }
}

func SuccessResponse(message string) SuccessResponseJson {
    return SuccessResponseJson{
        Status_code: http.StatusOK,
        Status:     "success",
        Message:    message,
    }
}

func SuccessWithDataResponse(message string, data interface{}) SuccessResponseJson {
    return SuccessResponseJson{
        Status_code: http.StatusOK,
        Status:     "success",
        Message:    message,
        Data:       data,
    }
}
