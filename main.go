package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/rwajon/image-compressor/config"
	"github.com/rwajon/image-compressor/routes"
)

func isPortOpened(port string) bool {
	conn, err := net.DialTimeout("tcp", ":"+port, 5*time.Second)

	if err != nil {
		return false
	}

	if conn != nil {
		conn.Close()
		return true
	}

	return false
}

func isPortUsed(config *config.Config) {
	for i := isPortOpened(config.Port); i; i = isPortOpened(config.Port) {
		fmt.Print("Error: Port " + config.Port + ": address already in use \nEnter a new port: ")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		text = regexp.MustCompile(`[\r\n]`).ReplaceAllString(text, "")

		if _, err := strconv.Atoi(text); err == nil {
			os.Setenv("PORT", text)
			config.Port = text
		} else {
			os.Exit(1)
		}
	}
}

func getCLIArgs(config *config.Config) {
	for _, v := range os.Args {
		if strings.Contains(v, "--port=") {
			if port, err := strconv.Atoi(strings.Replace(v, "--port=", "", -1)); err == nil {
				os.Setenv("PORT", strconv.Itoa(port))
				config.Port = strconv.Itoa(port)
			}
		}
		if strings.Contains(v, "--base_dir=") {
			if baseDir := strings.Replace(v, "--base_dir=", "", -1); baseDir != "" {
				os.Setenv("BASE_DIR", baseDir)
				config.BaseDir = baseDir
			}
		}
	}
}

func main() {
	appName := "Image compressor REST API"
	config := config.Get()
	engine := html.NewFileSystem(http.Dir("./public"), ".html").Delims("{{", "}}")

	getCLIArgs(&config)
	isPortUsed(&config)

	app := fiber.New(fiber.Config{
		Views:   engine,
		Prefork: true,
		AppName: appName + " v1.0",
	})

	app.Static("/", config.BaseDir)
	app.Static("/assets", "./public/assets")

	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.Render("index", fiber.Map{"title": appName + " v1.0"})
	})
	app.Get("/monitor", monitor.New())

	routes.Routes(app.Group("/api").Group("/v1"))

	log.Fatal(app.Listen(":" + config.Port))
}
