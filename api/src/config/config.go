package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ConnectionUrl = ""
	Port          = ""
)

func Load() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("erro ao carregar as envs ", err)
	}

	ConnectionUrl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	Port = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
}
