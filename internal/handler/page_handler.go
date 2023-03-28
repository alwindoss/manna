package handler

import (
	"github.com/gofiber/fiber/v2"
)

func NewPageHandler(c *Config) PageHandler {
	return &pageHandler{
		cfg: c,
	}
}

type PageHandler interface {
	ShowHomePage(c *fiber.Ctx) error
	ShowAboutPage(c *fiber.Ctx) error
	ShowContactPage(c *fiber.Ctx) error
	ShowLoginPage(c *fiber.Ctx) error
}

type pageHandler struct {
	// svc user.Service
	cfg *Config
}

// ShowLoginPage implements PageHandler
func (h *pageHandler) ShowLoginPage(c *fiber.Ctx) error {
	return c.Render("views/login.page", fiber.Map{
		"Title":         "Login",
		"Authenticated": false,
		"csrfToken":     c.Locals("token"),
	})
}

// ShowAboutPage implements PageHandler
func (h *pageHandler) ShowAboutPage(c *fiber.Ctx) error {
	return c.Render("views/about.page", fiber.Map{
		"Title":         "About Page",
		"Authenticated": false,
		"csrfToken":     c.Locals("token"),
	})
}

// ShowContactPage implements PageHandler
func (h *pageHandler) ShowContactPage(c *fiber.Ctx) error {
	return c.Render("views/contact.page", fiber.Map{
		"Title":         "Contact Page",
		"Authenticated": false,
		"csrfToken":     c.Locals("token"),
	})
}

// ShowHomePage implements PageHandler
func (h *pageHandler) ShowHomePage(c *fiber.Ctx) error {
	return c.Render("views/home.page", fiber.Map{
		"Title":         "Home Page",
		"Authenticated": false,
		"csrfToken":     c.Locals("token"),
	})
}
