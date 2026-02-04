package route

import (
	"go-finansia-multi-finance-technical-test/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                    *fiber.App
	UserController         *http.UserController
	ContactController      *http.ContactController
	AddressController      *http.AddressController
	ContactLimitController *http.ContactLimitController
	TransactionController  *http.TransactionController
	AuthMiddleware         fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	guest := c.App.Group("/api")
	guest.Post("/users", c.UserController.Register)
	guest.Post("/users/_login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	auth := c.App.Group("/api", c.AuthMiddleware)
	auth.Delete("/users", c.UserController.Logout)
	auth.Patch("/users/_current", c.UserController.Update)
	auth.Get("/users/_current", c.UserController.Current)

	auth.Get("/contacts", c.ContactController.List)
	auth.Post("/contacts", c.ContactController.Create)
	auth.Put("/contacts/:contactId", c.ContactController.Update)
	auth.Get("/contacts/:contactId", c.ContactController.Get)
	auth.Delete("/contacts/:contactId", c.ContactController.Delete)

	auth.Get("/contacts/:contactId/addresses", c.AddressController.List)
	auth.Post("/contacts/:contactId/addresses", c.AddressController.Create)
	auth.Put("/contacts/:contactId/addresses/:addressId", c.AddressController.Update)
	auth.Get("/contacts/:contactId/addresses/:addressId", c.AddressController.Get)
	auth.Delete("/contacts/:contactId/addresses/:addressId", c.AddressController.Delete)

	auth.Get("/contacts/:contactId/limits", c.ContactLimitController.List)
	auth.Post("/contacts/:contactId/limits", c.ContactLimitController.Create)
	auth.Put("/contacts/:contactId/limits/:limitId", c.ContactLimitController.Update)
	auth.Get("/contacts/:contactId/limits/:limitId", c.ContactLimitController.Get)
	auth.Delete("/contacts/:contactId/limits/:limitId", c.ContactLimitController.Delete)

	auth.Get("/contacts/:contactId/transactions", c.TransactionController.List)
	auth.Post("/contacts/:contactId/transactions", c.TransactionController.Create)
	auth.Put("/contacts/:contactId/transactions/:transactionId", c.TransactionController.Update)
	auth.Get("/contacts/:contactId/transactions/:transactionId", c.TransactionController.Get)
	auth.Delete("/contacts/:contactId/transactions/:transactionId", c.TransactionController.Delete)
}
