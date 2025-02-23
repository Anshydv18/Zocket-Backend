package env

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func LoadEnv() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("⚠️ Warning: .env file not found")
		} else {
			log.Println("✅ .env file loaded successfully.")
		}
	})
}

func Get(key string) string {
	LoadEnv()
	return os.Getenv(key)
}
