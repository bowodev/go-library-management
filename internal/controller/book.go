package controller

import (
	"strconv"

	"github.com/bowodev/go-library-management/internal/controller/request"
	"github.com/bowodev/go-library-management/internal/controller/response"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/bowodev/go-library-management/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type bookController struct {
	creator        interfaces.ICreateBook
	updater        interfaces.IUpdateBook
	deleter        interfaces.IDeleteBook
	getter         interfaces.IGetBookById
	allGetter      interfaces.IGetBookAll
	byAuthorGetter interfaces.IGetBookByAuthor
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
		if utils.IsUnprocessableEntityErrors(err) {
			return response.UnprocessableEntity[any](c, err.Error(), nil)
		}
		return response.InternalServerError[any](c, err.Error(), nil)
	}

	return response.Success(c, "Ok", book)
}

// Delete implements interfaces.IBookController.
func (b *bookController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	if id < 1 {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	if err := b.deleter.Do(c.Context(), id); err != nil {
		return response.InternalServerError[any](c, err.Error(), nil)
	}

	return response.Success[any](c, "Ok", nil)
}

// GetAll implements interfaces.IBookController.
func (b *bookController) GetAll(c *fiber.Ctx) error {
	books, err := b.allGetter.Do(c.Context())
	if err != nil {
		return response.InternalServerError[any](c, err.Error(), nil)
	}

	return response.Success(c, "Ok", books)
}

// GetByAuthor implements interfaces.IBookController.
func (b *bookController) GetByAuthor(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("authorId"), 10, 64)
	if err != nil {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	if id < 1 {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	books, err := b.byAuthorGetter.Do(c.Context(), id)
	if err != nil {
		return response.InternalServerError[any](c, err.Error(), nil)
	}

	return response.Success(c, "Ok", books)
}

// GetById implements interfaces.IBookController.
func (b *bookController) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	if id < 1 {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	book, err := b.getter.Do(c.Context(), id)
	if err != nil {
		return response.InternalServerError[any](c, err.Error(), nil)
	}

	return response.Success(c, "Ok", book)
}

// Update implements interfaces.IBookController.
func (b *bookController) Update(c *fiber.Ctx) error {
	ctx := c.Context()
	req := request.UpdateBook{}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	if id < 1 {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest[any](c, err.Error(), nil)
	}

	bookDto, err := req.ToDTO()
	if err != nil {
		return response.BadRequest[any](c, err.Error(), nil)
	}

	book, err := b.updater.Do(ctx, bookDto)
	if err != nil {
		if utils.IsUnprocessableEntityErrors(err) {
			return response.UnprocessableEntity[any](c, err.Error(), nil)
		}
		return response.InternalServerError[any](c, err.Error(), nil)
	}

	return response.Success(c, "Ok", book)
}

var _ interfaces.IBookController = (*bookController)(nil)

func NewBookController(
	creator interfaces.ICreateBook,
	updater interfaces.IUpdateBook,
	deleter interfaces.IDeleteBook,
	getter interfaces.IGetBookById,
	allGetter interfaces.IGetBookAll,
	byAuthorGetter interfaces.IGetBookByAuthor,
) *bookController {
	return &bookController{
		creator:        creator,
		updater:        updater,
		deleter:        deleter,
		getter:         getter,
		allGetter:      allGetter,
		byAuthorGetter: byAuthorGetter,
	}
}
