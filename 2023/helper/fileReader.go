package helper

import (
	"log"
	"os"
)

func ReadFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
