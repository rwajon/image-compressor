package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func index(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.Render("index", fiber.Map{
		"title": "Images compression REST API 👋",
	})
}

func main() {
	engine := html.NewFileSystem(http.Dir("./views"), ".html")

	engine.Delims("{{", "}}")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	api := app.Group("/api", index)
	v1 := api.Group("/v1", index)

	v1.Get("/", index)
	app.Get("/", index)

	log.Fatal(app.Listen(":3000"))
}
