package controller

import (
	"github.com/bowodev/go-library-management/internal/controller/request"
	"github.com/bowodev/go-library-management/internal/controller/response"
	"github.com/bowodev/go-library-management/internal/dto"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type bookController struct {
	creator interfaces.ICreateBook
}

var _ interfaces.IBookController = (*bookController)(nil)

func NewBookController(
	creator interfaces.ICreateBook,
) *bookController {
	return &bookController{
		creator: creator,
	}
}

func (b *bookController) Create(c *fiber.Ctx) error {
	var request request.CreateBook

	if err := c.BodyParser(&request); err != nil {
		return response.BadRequest[any](c, err.Error(), nil)
	}

	bookDto, err := request.ToDTO()
	if err != nil {
		return response.BadRequest[any](c, err.Error(), nil)
	}

	book, err := b.creator.Do(c.Context(), bookDto)
	if err != nil {
		return response.UnprocessableEntity[any](c, err.Error(), nil)
	}

	return response.Success[dto.Book](c, "Ok", book)
}
