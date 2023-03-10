package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NewAPIHandler(c *Config) APIHandler {
	return &apiHandler{
		cfg: c,
	}
}

type APIHandler interface {
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	Signup(c *fiber.Ctx) error
}

type apiHandler struct {
	// svc user.Service
	cfg *Config
}

// Logout implements APIHandler
func (*apiHandler) Logout(c *fiber.Ctx) error {
	return c.Render("views/home.page", fiber.Map{
		"Title":         "Home",
		"Authenticated": false,
		"csrfToken":     c.Locals("token"),
	})
}

// Login implements APIHandler
func (h *apiHandler) Login(c *fiber.Ctx) error {

	sess, err := h.cfg.SessionStore.Get(c)
	if err != nil {
		return err
	}
	sess.Set("user_name", "Alwin Doss")
	fmt.Println("User Name:", sess.Get("user_name"))
	fmt.Println("Session ID:", sess.Get("session_id"))
	fmt.Println("Keys:", sess.Keys())
	// err = sess.Save()
	// if err != nil {
	// 	return err
	// }

	return c.Render("views/admin/home.page", fiber.Map{
		"Title":         "Admin | Home",
		"Authenticated": true,
		"csrfToken":     c.Locals("token"),
	})
}

// Signup implements APIHandler
func (*apiHandler) Signup(c *fiber.Ctx) error {
	panic("unimplemented")
}
