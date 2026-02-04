package usecase

import (
	"context"
	"errors"
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

type ContactLimitUseCase struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	Validate               *validator.Validate
	ContactRepository      *repository.ContactRepository
	ContactLimitRepository *repository.ContactLimitRepository
	TransactionRepository  *repository.TransactionRepository
}

func NewContactLimitUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	contactRepository *repository.ContactRepository, contactLimitRepository *repository.ContactLimitRepository,
	transactionRepository *repository.TransactionRepository,
) *ContactLimitUseCase {
	return &ContactLimitUseCase{
		DB:                     db,
		Log:                    logger,
		Validate:               validate,
		ContactRepository:      contactRepository,
		ContactLimitRepository: contactLimitRepository,
		TransactionRepository:  transactionRepository,
	}
}

func (c *ContactLimitUseCase) Create(ctx context.Context, request *model.CreateContactLimitRequest) (*model.ContactLimitResponse, error) {
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

	existing := new(entity.ContactLimit)
	if err := c.ContactLimitRepository.FindByContactIdAndTenor(tx, existing, contact.ID, request.Tenor); err == nil {
		return nil, fiber.ErrConflict
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.Log.WithError(err).Error("failed to check existing limit")
		return nil, fiber.ErrInternalServerError
	}

	usedAmount, err := c.TransactionRepository.SumOtrByContactAndTenor(tx, contact.ID, request.Tenor)
	if err != nil {
		c.Log.WithError(err).Error("failed to calculate used limit")
		return nil, fiber.ErrInternalServerError
	}
	if request.LimitAmount < usedAmount {
		return nil, fiber.NewError(fiber.StatusBadRequest, "limit amount is lower than current usage")
	}

	limit := &entity.ContactLimit{
		ID:          uuid.NewString(),
		ContactId:   contact.ID,
		Tenor:       request.Tenor,
		LimitAmount: request.LimitAmount,
	}

	if err := c.ContactLimitRepository.Create(tx, limit); err != nil {
		c.Log.WithError(err).Error("failed to create contact limit")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ContactLimitToResponse(limit, usedAmount), nil
}

func (c *ContactLimitUseCase) Update(ctx context.Context, request *model.UpdateContactLimitRequest) (*model.ContactLimitResponse, error) {
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
	if err := c.ContactLimitRepository.FindByIdAndContactId(tx, limit, request.ID, contact.ID); err != nil {
		c.Log.WithError(err).Error("failed to find contact limit")
		return nil, fiber.ErrNotFound
	}

	if request.Tenor != limit.Tenor {
		existing := new(entity.ContactLimit)
		if err := c.ContactLimitRepository.FindByContactIdAndTenor(tx, existing, contact.ID, request.Tenor); err == nil {
			return nil, fiber.ErrConflict
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.Log.WithError(err).Error("failed to check existing limit")
			return nil, fiber.ErrInternalServerError
		}
	}

	usedAmount, err := c.TransactionRepository.SumOtrByContactAndTenor(tx, contact.ID, request.Tenor)
	if err != nil {
		c.Log.WithError(err).Error("failed to calculate used limit")
		return nil, fiber.ErrInternalServerError
	}
	if request.LimitAmount < usedAmount {
		return nil, fiber.NewError(fiber.StatusBadRequest, "limit amount is lower than current usage")
	}

	limit.Tenor = request.Tenor
	limit.LimitAmount = request.LimitAmount

	if err := c.ContactLimitRepository.Update(tx, limit); err != nil {
		c.Log.WithError(err).Error("failed to update contact limit")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ContactLimitToResponse(limit, usedAmount), nil
}

func (c *ContactLimitUseCase) Get(ctx context.Context, request *model.GetContactLimitRequest) (*model.ContactLimitResponse, error) {
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
	if err := c.ContactLimitRepository.FindByIdAndContactId(tx, limit, request.ID, contact.ID); err != nil {
		c.Log.WithError(err).Error("failed to find contact limit")
		return nil, fiber.ErrNotFound
	}

	usedAmount, err := c.TransactionRepository.SumOtrByContactAndTenor(tx, contact.ID, limit.Tenor)
	if err != nil {
		c.Log.WithError(err).Error("failed to calculate used limit")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ContactLimitToResponse(limit, usedAmount), nil
}

func (c *ContactLimitUseCase) Delete(ctx context.Context, request *model.DeleteContactLimitRequest) error {
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

	limit := new(entity.ContactLimit)
	if err := c.ContactLimitRepository.FindByIdAndContactId(tx, limit, request.ID, contact.ID); err != nil {
		c.Log.WithError(err).Error("failed to find contact limit")
		return fiber.ErrNotFound
	}

	if err := c.ContactLimitRepository.Delete(tx, limit); err != nil {
		c.Log.WithError(err).Error("failed to delete contact limit")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *ContactLimitUseCase) List(ctx context.Context, request *model.ListContactLimitRequest) ([]model.ContactLimitResponse, error) {
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

	limits, err := c.ContactLimitRepository.FindAllByContactId(tx, contact.ID)
	if err != nil {
		c.Log.WithError(err).Error("failed to list contact limits")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.ContactLimitResponse, len(limits))
	for i, limit := range limits {
		usedAmount, err := c.TransactionRepository.SumOtrByContactAndTenor(tx, contact.ID, limit.Tenor)
		if err != nil {
			c.Log.WithError(err).Error("failed to calculate used limit")
			return nil, fiber.ErrInternalServerError
		}
		responses[i] = *converter.ContactLimitToResponse(&limit, usedAmount)
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return responses, nil
}
