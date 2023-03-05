package engine

import (
	"fmt"
	"net/http"

	"github.com/alwindoss/manna"
	"github.com/alwindoss/manna/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
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
	pageHdlrs := handler.NewPageHandler(nil)
	apiHdlrs := handler.NewAPIHandler(nil)

	app.Get("/", pageHdlrs.ShowHomePage)
	app.Get("/about", pageHdlrs.ShowAboutPage)
	app.Get("/contact", pageHdlrs.ShowContactPage)
	app.Get("/login", pageHdlrs.ShowLoginPage)

	app.Post("/login", apiHdlrs.Login)
	app.Get("/logout", apiHdlrs.Logout)

	addr := fmt.Sprintf(":%d", cfg.Port)
	return app.Listen(addr)
}
