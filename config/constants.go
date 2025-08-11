package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
  URL string

  POSTGRES_DB string
  POSTGRES_USER string
  POSTGRES_PASSWORD string
  POSTGRES_HOST string
  POSTGRES_PORT string
)

func LoadConfig(path ...string) {
  envPath := ".env"
  if len(path) > 0 && path[0] != "" {
    envPath = path[0]
  }
  
  err := godotenv.Load(envPath)
  if err != nil {
    fmt.Println("Erro ao carregar o arquivo .env")
    os.Exit(1)
  }

  URL = os.Getenv("URL")

  POSTGRES_DB = os.Getenv("POSTGRES_DB")
  POSTGRES_USER = os.Getenv("POSTGRES_USER")
  POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
  POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
  POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
}
