package helpers

import (
	"log"
	"os"
)

func CreateDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}
