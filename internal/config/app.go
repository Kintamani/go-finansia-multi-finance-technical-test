package config

import (
	"go-finansia-multi-finance-technical-test/internal/delivery/http"
	"go-finansia-multi-finance-technical-test/internal/delivery/http/middleware"
	"go-finansia-multi-finance-technical-test/internal/delivery/http/route"
	"go-finansia-multi-finance-technical-test/internal/repository"
	"go-finansia-multi-finance-technical-test/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	contactRepository := repository.NewContactRepository(config.Log)
	addressRepository := repository.NewAddressRepository(config.Log)
	contactLimitRepository := repository.NewContactLimitRepository(config.Log)
	transactionRepository := repository.NewTransactionRepository(config.Log)

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	contactUseCase := usecase.NewContactUseCase(config.DB, config.Log, config.Validate, contactRepository)
	addressUseCase := usecase.NewAddressUseCase(config.DB, config.Log, config.Validate, contactRepository, addressRepository)
	contactLimitUseCase := usecase.NewContactLimitUseCase(config.DB, config.Log, config.Validate, contactRepository, contactLimitRepository, transactionRepository)
	transactionUseCase := usecase.NewTransactionUseCase(config.DB, config.Log, config.Validate, contactRepository, contactLimitRepository, transactionRepository)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	contactController := http.NewContactController(contactUseCase, config.Log)
	addressController := http.NewAddressController(addressUseCase, config.Log)
	contactLimitController := http.NewContactLimitController(contactLimitUseCase, config.Log)
	transactionController := http.NewTransactionController(transactionUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:                    config.App,
		UserController:         userController,
		ContactController:      contactController,
		AddressController:      addressController,
		ContactLimitController: contactLimitController,
		TransactionController:  transactionController,
		AuthMiddleware:         authMiddleware,
	}
	routeConfig.Setup()
}
