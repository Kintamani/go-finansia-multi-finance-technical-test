package usecase

import (
	"context"
	"go-finansia-multi-finance-technical-test/internal/entity"
	"go-finansia-multi-finance-technical-test/internal/model"
	"go-finansia-multi-finance-technical-test/internal/model/converter"
	"go-finansia-multi-finance-technical-test/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionUseCase struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	Validate               *validator.Validate
	ContactRepository      *repository.ContactRepository
	ContactLimitRepository *repository.ContactLimitRepository
	TransactionRepository  *repository.TransactionRepository
}

func NewTransactionUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	contactRepository *repository.ContactRepository, contactLimitRepository *repository.ContactLimitRepository,
	transactionRepository *repository.TransactionRepository,
) *TransactionUseCase {
	return &TransactionUseCase{
		DB:                     db,
		Log:                    logger,
		Validate:               validate,
		ContactRepository:      contactRepository,
		ContactLimitRepository: contactLimitRepository,
		TransactionRepository:  transactionRepository,
	}
}

func (c *TransactionUseCase) Create(ctx context.Context, request *model.CreateTransactionRequest) (*model.TransactionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	contact := new(entity.Contact)
	if err := c.ContactRepository.FindByIdAndUserId(tx, contact, request.ContactId, request.UserId); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	limit := new(entity.ContactLimit)
	if err := c.ContactLimitRepository.FindByContactIdAndTenor(tx, limit, contact.ID, request.Tenor); err != nil {
		c.Log.WithError(err).Error("failed to find contact limit")
		return nil, fiber.NewError(fiber.StatusBadRequest, "contact limit not found for tenor")
	}

	usedAmount, err := c.TransactionRepository.SumOtrByContactAndTenor(tx, contact.ID, request.Tenor)
	if err != nil {
		c.Log.WithError(err).Error("failed to calculate used limit")
		return nil, fiber.ErrInternalServerError
	}
	if usedAmount+request.Otr > limit.LimitAmount {
		return nil, fiber.NewError(fiber.StatusBadRequest, "limit exceeded")
	}

	transaction := &entity.Transaction{
		ID:                uuid.NewString(),
		ContactId:         contact.ID,
		ContractNumber:    request.ContractNumber,
		Tenor:             request.Tenor,
		Channel:           request.Channel,
		Otr:               request.Otr,
		AdminFee:          request.AdminFee,
		InstallmentAmount: request.InstallmentAmount,
		InterestAmount:    request.InterestAmount,
		AssetName:         request.AssetName,
	}

	if err := c.TransactionRepository.Create(tx, transaction); err != nil {
		c.Log.WithError(err).Error("failed to create transaction")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.TransactionToResponse(transaction), nil
}

func (c *TransactionUseCase) Update(ctx context.Context, request *model.UpdateTransactionRequest) (*model.TransactionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	contact := new(entity.Contact)
	if err := c.ContactRepository.FindByIdAndUserId(tx, contact, request.ContactId, request.UserId); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	transaction := new(entity.Transaction)
	if err := c.TransactionRepository.FindByIdAndContactId(tx, transaction, request.ID, contact.ID); err != nil {
		c.Log.WithError(err).Error("failed to find transaction")
		return nil, fiber.ErrNotFound
	}

	limit := new(entity.ContactLimit)
	if err := c.ContactLimitRepository.FindByContactIdAndTenor(tx, limit, contact.ID, request.Tenor); err != nil {
		c.Log.WithError(err).Error("failed to find contact limit")
		return nil, fiber.NewError(fiber.StatusBadRequest, "contact limit not found for tenor")
	}

	usedAmount, err := c.TransactionRepository.SumOtrByContactAndTenorExcludeId(tx, contact.ID, request.Tenor, transaction.ID)
	if err != nil {
		c.Log.WithError(err).Error("failed to calculate used limit")
		return nil, fiber.ErrInternalServerError
	}
	if usedAmount+request.Otr > limit.LimitAmount {
		return nil, fiber.NewError(fiber.StatusBadRequest, "limit exceeded")
	}

	transaction.ContractNumber = request.ContractNumber
	transaction.Tenor = request.Tenor
	transaction.Channel = request.Channel
	transaction.Otr = request.Otr
	transaction.AdminFee = request.AdminFee
	transaction.InstallmentAmount = request.InstallmentAmount
	transaction.InterestAmount = request.InterestAmount
	transaction.AssetName = request.AssetName

	if err := c.TransactionRepository.Update(tx, transaction); err != nil {
		c.Log.WithError(err).Error("failed to update transaction")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.TransactionToResponse(transaction), nil
}

func (c *TransactionUseCase) Get(ctx context.Context, request *model.GetTransactionRequest) (*model.TransactionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	contact := new(entity.Contact)
	if err := c.ContactRepository.FindByIdAndUserId(tx, contact, request.ContactId, request.UserId); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	transaction := new(entity.Transaction)
	if err := c.TransactionRepository.FindByIdAndContactId(tx, transaction, request.ID, contact.ID); err != nil {
		c.Log.WithError(err).Error("failed to find transaction")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.TransactionToResponse(transaction), nil
}

func (c *TransactionUseCase) Delete(ctx context.Context, request *model.DeleteTransactionRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return fiber.ErrBadRequest
	}

	contact := new(entity.Contact)
	if err := c.ContactRepository.FindByIdAndUserId(tx, contact, request.ContactId, request.UserId); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return fiber.ErrNotFound
	}

	transaction := new(entity.Transaction)
	if err := c.TransactionRepository.FindByIdAndContactId(tx, transaction, request.ID, contact.ID); err != nil {
		c.Log.WithError(err).Error("failed to find transaction")
		return fiber.ErrNotFound
	}

	if err := c.TransactionRepository.Delete(tx, transaction); err != nil {
		c.Log.WithError(err).Error("failed to delete transaction")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *TransactionUseCase) List(ctx context.Context, request *model.ListTransactionRequest) ([]model.TransactionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	contact := new(entity.Contact)
	if err := c.ContactRepository.FindByIdAndUserId(tx, contact, request.ContactId, request.UserId); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	transactions, err := c.TransactionRepository.FindAllByContactId(tx, contact.ID)
	if err != nil {
		c.Log.WithError(err).Error("failed to list transactions")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.TransactionResponse, len(transactions))
	for i, transaction := range transactions {
		responses[i] = *converter.TransactionToResponse(&transaction)
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return responses, nil
}
