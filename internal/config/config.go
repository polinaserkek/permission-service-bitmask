package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type Config struct {
	Port int
	Host string
	/*uncomment later when there ll be db*/
	// Db           string
	// Username     string
	Storage_type string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки .env")
	}
}
