package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func Success[T any](c *fiber.Ctx, message string, data T) error {
	return c.Status(http.StatusOK).JSON(Response[T]{
		Message: message,
		Data:    data,
	})
}

func UnprocessableEntity[T any](c *fiber.Ctx, message string, data T) error {
	return c.Status(http.StatusUnprocessableEntity).JSON(Response[T]{
		Message: message,
	})
}

func BadRequest[T any](c *fiber.Ctx, message string, data T) error {
	return c.Status(http.StatusBadRequest).JSON(Response[T]{
		Message: message,
	})
}

func InternalServerError[T any](c *fiber.Ctx, message string, data T) error {
	return c.Status(http.StatusInternalServerError).JSON(Response[T]{
		Message: message,
	})
}
