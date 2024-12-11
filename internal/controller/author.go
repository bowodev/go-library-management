package controller

import (
	"strconv"

	"github.com/bowodev/go-library-management/internal/controller/request"
	"github.com/bowodev/go-library-management/internal/controller/response"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/bowodev/go-library-management/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type authorController struct {
	creator   interfaces.ICreateAuthor
	updater   interfaces.IUpdateAuthor
	deleter   interfaces.IDeleteAuthor
	getter    interfaces.IGetAuthorById
	allGetter interfaces.IGetAuthorAll
}

// Create implements interfaces.IAuthorController.
func (a *authorController) Create(c *fiber.Ctx) error {
	var request request.CreateAuthor

	if err := c.BodyParser(&request); err != nil {
		return response.BadRequest[any](c, err.Error(), nil)
	}

	authorDto, err := request.ToDTO()
	if err != nil {
		return response.BadRequest[any](c, err.Error(), nil)
	}

	author, err := a.creator.Do(c.Context(), authorDto)
	if err != nil {
		if utils.IsUnprocessableEntityErrors(err) {
			return response.UnprocessableEntity[any](c, err.Error(), nil)
		}
		return response.InternalServerError[any](c, err.Error(), nil)
	}

	return response.Success(c, "Ok", author)
}

// Delete implements interfaces.IAuthorController.
func (a *authorController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	if id < 1 {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	if err := a.deleter.Do(c.Context(), id); err != nil {
		return response.InternalServerError[any](c, err.Error(), nil)
	}

	return response.Success[any](c, "Ok", nil)
}

// GetAll implements interfaces.IAuthorController.
func (a *authorController) GetAll(c *fiber.Ctx) error {
	authors, err := a.allGetter.Do(c.Context())
	if err != nil {
		return response.InternalServerError[any](c, err.Error(), nil)
	}

	return response.Success(c, "Ok", authors)
}

// GetById implements interfaces.IAuthorController.
func (a *authorController) GetById(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	if id < 1 {
		return response.BadRequest[any](c, "invalid id", nil)
	}

	author, err := a.getter.Do(c.Context(), id)
	if err != nil {
		return response.InternalServerError[any](c, err.Error(), nil)
	}

	return response.Success(c, "Ok", author)
}

// Update implements interfaces.IAuthorController.
func (a *authorController) Update(c *fiber.Ctx) error {
	ctx := c.Context()
	req := request.UpdateAuthor{}

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

	authorDto, err := req.ToDTO()
	if err != nil {
		return response.BadRequest[any](c, err.Error(), nil)
	}

	author, err := a.updater.Do(ctx, authorDto)
	if err != nil {
		if utils.IsUnprocessableEntityErrors(err) {
			return response.UnprocessableEntity[any](c, err.Error(), nil)
		}
		return response.InternalServerError[any](c, err.Error(), nil)
	}

	return response.Success(c, "Ok", author)
}

var _ interfaces.IAuthorController = (*authorController)(nil)

func NewAuthorController(
	creator interfaces.ICreateAuthor,
	updater interfaces.IUpdateAuthor,
	deleter interfaces.IDeleteAuthor,
	getter interfaces.IGetAuthorById,
	allGetter interfaces.IGetAuthorAll,
) *authorController {
	return &authorController{
		creator:   creator,
		updater:   updater,
		deleter:   deleter,
		getter:    getter,
		allGetter: allGetter,
	}
}
