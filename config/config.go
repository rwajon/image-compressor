package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	BaseDir string
}

func Get() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	baseDir := os.Getenv("BASE_DIR")

	if port == "" {
		port = "3000"
	}
	if baseDir == "" {
		baseDir = "./image_compressor_uploaded_files"
	}

	return Config{
		Port:    port,
		BaseDir: baseDir,
	}
}
