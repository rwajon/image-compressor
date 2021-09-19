package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/rwajon/image-compressor/routes"
)

func index(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.Render("index", fiber.Map{
		"title": "Image compressor REST API 👋",
	})
}

func main() {
	engine := html.NewFileSystem(http.Dir("./views"), ".html")

	engine.Delims("{{", "}}")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(logger.New())

	api_v1 := app.Group("/api").Group("/v1")

	app.Get("/", index)
	app.Get("/monitor", monitor.New())

	routes.Routes(api_v1)

	log.Fatal(app.Listen(":3000"))
}
