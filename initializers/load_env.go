package initializers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("❌Error loading the .env file: %v❌", err)
	}

	envPath := filepath.Join(dir, "../.env")

	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}
}
