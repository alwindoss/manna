package handler

import (
	"github.com/alwindoss/manna/internal/user"
	"github.com/gofiber/fiber/v2"
)

func NewAPIHandler(s user.Service) APIHandler {
	return &apiHandler{
		svc: s,
	}
}

type APIHandler interface {
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	Signup(c *fiber.Ctx) error
}

type apiHandler struct {
	svc user.Service
}

// Logout implements APIHandler
func (*apiHandler) Logout(c *fiber.Ctx) error {
	return c.Render("views/home.page", fiber.Map{
		"Title":         "Home",
		"Authenticated": false,
	})
}

// Login implements APIHandler
func (*apiHandler) Login(c *fiber.Ctx) error {
	return c.Render("views/admin/home.page", fiber.Map{
		"Title":         "Admin | Home",
		"Authenticated": true,
	})
}

// Signup implements APIHandler
func (*apiHandler) Signup(c *fiber.Ctx) error {
	panic("unimplemented")
}
