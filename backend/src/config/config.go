package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DatabaseConnectionString = ""
	Port                     = 0
	DateLayout               = "2006-01-02"
	AbsolutePath             = ""
	UAParserRegexesPath      = ""
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, b, _, _ := runtime.Caller(0)
	AbsolutePath = filepath.Join(filepath.Dir(b), "../..")

	UAParserRegexesPath = fmt.Sprintf("%s/regexes.yaml", AbsolutePath)

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 8000
	}

	DatabaseConnectionString = fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	os.Setenv("TZ", "UTC")
}
