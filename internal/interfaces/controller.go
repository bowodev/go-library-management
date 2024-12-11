package interfaces

import "github.com/gofiber/fiber/v2"

//go:generate mockgen -source=controller.go -destination=./mock/controller_mock.go -package=mocks
type IBookController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	GetByAuthor(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type IAuthorController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
