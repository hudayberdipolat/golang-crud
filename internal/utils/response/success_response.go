package response

import "github.com/gofiber/fiber/v3"

type successResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

func SuccessResponse(ctx fiber.Ctx, statusCode int, message string, data any) error {
	return ctx.Status(statusCode).JSON(successResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}
