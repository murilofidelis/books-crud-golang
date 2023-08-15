package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	JdbcUrl   = ""
	Port      = 0
	SecretKey []byte
)

func LoadConfig() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	JdbcUrl = fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable&search_path=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SCHEMA"),
	)
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
