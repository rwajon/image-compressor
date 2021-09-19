package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rwajon/images-compression/controllers"
)

func Files(router fiber.Router) {
	f := controllers.FileController{
		BaseDir: "./uploaded_images",
		Dir:     time.Now().Format("01-02-2006"), // MM-DD-YYYY
	}
	router.Get("/", f.ListFiles)
}
