package handler

import (
	"github.com/alwindoss/manna/internal/user"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Config struct {
	UserSvc      user.Service
	SessionStore *session.Store
}
