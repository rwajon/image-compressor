package config

import (
	"os"
)

type Config struct {
	Port    string
	BaseDir string
}

func Get() Config {
	port := os.Getenv("PORT")
	base_dir := os.Getenv("BASE_DIR")

	if port == "" {
		port = "3000"
	}
	if base_dir == "" {
		base_dir = "./image_compressor_uploaded_files"
	}

	return Config{
		Port:    port,
		BaseDir: base_dir,
	}
}
