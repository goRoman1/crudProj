package configs

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DbUser 					string
	DbPassword 				string
	DbName 					string
	DriverName 				string
	DbHost    				string
	DbPort    				string
}



func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	config := Config {
			DbUser:     os.Getenv("POSTGRES_USER"),
			DbPassword: os.Getenv("POSTGRES_PASSWORD"),
			DbName:     os.Getenv("POSTGRES_DB"),
			DriverName: os.Getenv("DRIVER_NAME"),
			DbHost:     os.Getenv("POSTGRES_HOST"),
			DbPort: 	os.Getenv("POSTGRES_PORT"),
	}
	configBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Configuration:", string(configBytes))

	return &config
}
