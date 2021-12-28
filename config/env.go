package config

import (
  "log"
	"github.com/joho/godotenv"
)

func Getenvfile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("file .env tidak ada")
	}
}
