package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	fp, err := filepath.Abs(".env")
	if strings.Contains(fp, "tests\\") {
		fp, err = filepath.Abs("../../.env")
	}

	if err != nil {
		log.Print(err)
	}
	err = godotenv.Load(fp)
	if err != nil {
		log.Print(err)
		log.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
