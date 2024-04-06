package engine

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alwindoss/manna"
	"github.com/alwindoss/manna/internal/handler"
	"github.com/alwindoss/manna/internal/storage"
	"github.com/alwindoss/manna/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/template/html/v2"
)

type Config struct {
	Port int
}

func Run(cfg *Config) error {
	// viewEngine := html.New("./views", ".html")
	viewEngine := html.NewFileSystem(http.FS(manna.ViewFS), ".html")

	app := fiber.New(fiber.Config{
		Views:       viewEngine,
		ViewsLayout: "views/layouts/default.layout",
	})
	// CSRF Config
	// app.Use(csrf.New())
	app.Use(csrf.New(csrf.Config{
		// KeyLookup:      "header:X-Csrf-Token",
		KeyLookup:      "form:_csrf",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		ContextKey:     "token",
		Next: func(c *fiber.Ctx) bool {
			return false
		},
		// Extractor: func(c *fiber.Ctx) (string, error) {
		// 	return "", nil
		// },
	}))

	sessionStore := session.New()
	repo := storage.NewInMemRepository()
	userSvc := user.NewService(repo)

	handlerCfg := handler.Config{
		UserSvc:      userSvc,
		SessionStore: sessionStore,
	}
	if "" == "" {
	}
	pageHdlrs := handler.NewPageHandler(&handlerCfg)
	apiHdlrs := handler.NewAPIHandler(&handlerCfg)

	app.Get("/", pageHdlrs.ShowHomePage)
	app.Get("/about", pageHdlrs.ShowAboutPage)
	app.Get("/contact", pageHdlrs.ShowContactPage)
	app.Get("/login", pageHdlrs.ShowLoginPage)

	app.Post("/login", apiHdlrs.Login)
	app.Get("/logout", apiHdlrs.Logout)

	addr := fmt.Sprintf(":%d", cfg.Port)
	return app.Listen(addr)
}
