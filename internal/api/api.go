package api

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/todalist/app/internal/globals"
	"go.uber.org/zap"
)

type ApiResult struct {
	fiber.Ctx
}

func (c ApiResult) Or(data any, err error) error {
	if err != nil {
		if errors.Is(err, &ApiResponse{}) {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err400(err.Error()))
		}
		if errors.Is(err, &fiber.Error{}) {
			return c.Status(err.(*fiber.Error).Code).JSON(err400(err.Error()))
		}
		globals.LOG.Error("uncatched error", zap.String("error", err.Error()))
		// mask error message for security
		return c.Status(fiber.StatusInternalServerError).JSON(err500())
	}
	return c.JSON(ok(data))
}

func (c ApiResult) Ok(data any) error {
	return c.JSON(ok(data))
}

func (c ApiResult) Err500() error {
	return c.Status(fiber.StatusInternalServerError).JSON(err500())
}

func (c ApiResult) Err400(message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(err400(message))
}

func (c ApiResult) Err(code int, message string) error {
	return c.Status(code).JSON(err(code, message))
}

type ApiResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func (e *ApiResponse) Error() string {
	return e.Message
}

func err400(message string) error {
	return &ApiResponse{Code: fiber.ErrBadRequest.Code, Message: message}
}

func err500() error {
	return &ApiResponse{Code: fiber.ErrInternalServerError.Code, Message: fiber.ErrInternalServerError.Message}
}

func ok(data interface{}) ApiResponse {
	return ApiResponse{Code: fiber.StatusOK, Data: data}
}

func err(code int, message string) error {
	return &ApiResponse{Code: code, Message: message}
}

func Result(c fiber.Ctx) ApiResult {
	return ApiResult{c}
}
