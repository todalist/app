package common

import "github.com/gofiber/fiber/v3"

type ApiResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func (e *ApiResponse) Error() string {
	return e.Message
}

func Err400(message string) ApiResponse {
	return ApiResponse{Code: fiber.ErrBadRequest.Code, Message: message}
}

func Ok(data interface{}) ApiResponse {
	return ApiResponse{Code: fiber.StatusOK, Data: data}
}

func Or(data interface{}, err error) ApiResponse {
	if err != nil {
		return Err400(err.Error())
	}
	return Ok(data)
}
