package engine

import (
	"fmt"

	"github.com/alwindoss/manna/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type Config struct {
	Port int
}

func Run(cfg *Config) error {
	viewEngine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:       viewEngine,
		ViewsLayout: "layouts/default.layout",
	})
	hdlrs := handler.NewPageHandler(nil)

	app.Get("/", hdlrs.ShowHomePage)
	app.Get("/about", hdlrs.ShowAboutPage)
	app.Get("/contact", hdlrs.ShowContactPage)
	addr := fmt.Sprintf(":%d", cfg.Port)
	return app.Listen(addr)
}