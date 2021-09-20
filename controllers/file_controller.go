package controllers

import (
	"io"
	"io/ioutil"
	"log"
	"strings"

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
		"status": 200,
		"files":  file_names,
	}
	return c.JSON(res)
}

func (f FileController) UploadFiles(c *fiber.Ctx) error {
	dir := f.BaseDir + "/" + f.Dir
	helpers.CreateDir(dir)

	fileheader, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}

	file, err := fileheader.Open()

	if err != nil {
		panic(err)
	}

	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	filename, err := helpers.CompressImage(buffer, 40, dir)
	if err != nil {
		panic(err)
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	res := fiber.Map{
		"status":  201,
		"message": "File uploaded",
		"url":     c.BaseURL() + "/" + strings.Replace(f.Dir, ".", "", -1) + "/" + filename,
	}
	return c.JSON(res)
}
