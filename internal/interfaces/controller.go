package interfaces

import "github.com/gofiber/fiber/v2"

//go:generate mockgen -source=controller.go -destination=./mock/controller_mock.go -package=mocks
type IBookController interface {
	Create(ctx *fiber.Ctx) error
}
