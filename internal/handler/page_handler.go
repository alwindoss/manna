package handler

import (
	"github.com/alwindoss/manna/internal/user"
	"github.com/gofiber/fiber/v2"
)

func NewPageHandler(s user.Service) PageHandler {
	return &pageHandler{
		svc: s,
	}
}

type PageHandler interface {
	ShowHomePage(c *fiber.Ctx) error
	ShowAboutPage(c *fiber.Ctx) error
	ShowContactPage(c *fiber.Ctx) error
	ShowLoginPage(c *fiber.Ctx) error
}

type pageHandler struct {
	svc user.Service
}

// ShowLoginPage implements PageHandler
func (h *pageHandler) ShowLoginPage(c *fiber.Ctx) error {
	return c.Render("views/login.page", fiber.Map{
		"Title":         "Login",
		"Authenticated": false,
	})
}

// ShowAboutPage implements PageHandler
func (h *pageHandler) ShowAboutPage(c *fiber.Ctx) error {
	return c.Render("views/about.page", fiber.Map{
		"Title":         "About Page",
		"Authenticated": false,
	})
}

// ShowContactPage implements PageHandler
func (h *pageHandler) ShowContactPage(c *fiber.Ctx) error {
	return c.Render("views/contact.page", fiber.Map{
		"Title":         "Contact Page",
		"Authenticated": false,
	})
}

// ShowHomePage implements PageHandler
func (h *pageHandler) ShowHomePage(c *fiber.Ctx) error {
	return c.Render("views/home.page", fiber.Map{
		"Title":         "Home Page",
		"Authenticated": false,
	})
}
