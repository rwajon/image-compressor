package controllers

import (
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rwajon/image-compressor/helpers"
)

type FileController struct {
	Dir     string
	BaseDir string
}

func (f FileController) ListFiles(c *fiber.Ctx) error {
	dir := f.BaseDir + "/" + f.Dir

	helpers.CreateDir(dir)

	files, err := ioutil.ReadDir(dir)
	file_names := make([]string, 0)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		file_names = append(file_names, f.Name())
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res := fiber.Map{
		"status":  200,
		"message": "Files directory: " + f.Dir,
		"files":   file_names,
	}
	return c.JSON(res)
}
