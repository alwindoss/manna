package handler

import "github.com/gofiber/fiber/v2"

func NewPageHandler(svc interface{}) PageHandler {
	return &pageHandler{}
}

type PageHandler interface {
	ShowHomePage(c *fiber.Ctx) error
	ShowAboutPage(c *fiber.Ctx) error
	ShowContactPage(c *fiber.Ctx) error
}

type pageHandler struct {
}

// ShowAboutPage implements PageHandler
func (*pageHandler) ShowAboutPage(c *fiber.Ctx) error {
	return c.Render("about.page", fiber.Map{
		"Title": "About Page",
	})
}

// ShowContactPage implements PageHandler
func (*pageHandler) ShowContactPage(c *fiber.Ctx) error {
	return c.Render("contact.page", fiber.Map{
		"Title": "Contact Page",
	})
}

// ShowHomePage implements PageHandler
func (*pageHandler) ShowHomePage(c *fiber.Ctx) error {
	return c.Render("home.page", fiber.Map{
		"Title": "Home Page",
	})
}
