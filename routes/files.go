package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rwajon/image-compressor/controllers"
)

func filesRoutes(router fiber.Router) {
	f := controllers.FileController{
		BaseDir: "./uploaded_files",
		Dir:     time.Now().Format("01-02-2006"), // MM-DD-YYYY
	}
	router.Get("/", f.ListFiles)
	router.Post("/", f.UploadFiles)
}
