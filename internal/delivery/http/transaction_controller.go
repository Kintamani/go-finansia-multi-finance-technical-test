package http

import (
	"go-finansia-multi-finance-technical-test/internal/delivery/http/middleware"
	"go-finansia-multi-finance-technical-test/internal/model"
	"go-finansia-multi-finance-technical-test/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TransactionController struct {
	UseCase *usecase.TransactionUseCase
	Log     *logrus.Logger
}

func NewTransactionController(useCase *usecase.TransactionUseCase, log *logrus.Logger) *TransactionController {
	return &TransactionController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *TransactionController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateTransactionRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserId = auth.ID
	request.ContactId = ctx.Params("contactId")

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create transaction")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.TransactionResponse]{Data: response})
}

func (c *TransactionController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")

	request := &model.ListTransactionRequest{
		UserId:    auth.ID,
		ContactId: contactId,
	}

	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list transactions")
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.TransactionResponse]{Data: responses})
}

func (c *TransactionController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")
	transactionId := ctx.Params("transactionId")

	request := &model.GetTransactionRequest{
		UserId:    auth.ID,
		ContactId: contactId,
		ID:        transactionId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get transaction")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.TransactionResponse]{Data: response})
}

func (c *TransactionController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateTransactionRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserId = auth.ID
	request.ContactId = ctx.Params("contactId")
	request.ID = ctx.Params("transactionId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update transaction")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.TransactionResponse]{Data: response})
}

func (c *TransactionController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")
	transactionId := ctx.Params("transactionId")

	request := &model.DeleteTransactionRequest{
		UserId:    auth.ID,
		ContactId: contactId,
		ID:        transactionId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete transaction")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
