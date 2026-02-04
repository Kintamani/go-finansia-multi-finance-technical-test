package http

import (
	"go-finansia-multi-finance-technical-test/internal/delivery/http/middleware"
	"go-finansia-multi-finance-technical-test/internal/model"
	"go-finansia-multi-finance-technical-test/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ContactLimitController struct {
	UseCase *usecase.ContactLimitUseCase
	Log     *logrus.Logger
}

func NewContactLimitController(useCase *usecase.ContactLimitUseCase, log *logrus.Logger) *ContactLimitController {
	return &ContactLimitController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *ContactLimitController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateContactLimitRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserId = auth.ID
	request.ContactId = ctx.Params("contactId")

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create contact limit")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactLimitResponse]{Data: response})
}

func (c *ContactLimitController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")

	request := &model.ListContactLimitRequest{
		UserId:    auth.ID,
		ContactId: contactId,
	}

	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list contact limits")
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.ContactLimitResponse]{Data: responses})
}

func (c *ContactLimitController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")
	limitId := ctx.Params("limitId")

	request := &model.GetContactLimitRequest{
		UserId:    auth.ID,
		ContactId: contactId,
		ID:        limitId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get contact limit")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactLimitResponse]{Data: response})
}

func (c *ContactLimitController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateContactLimitRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserId = auth.ID
	request.ContactId = ctx.Params("contactId")
	request.ID = ctx.Params("limitId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update contact limit")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactLimitResponse]{Data: response})
}

func (c *ContactLimitController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")
	limitId := ctx.Params("limitId")

	request := &model.DeleteContactLimitRequest{
		UserId:    auth.ID,
		ContactId: contactId,
		ID:        limitId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete contact limit")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
