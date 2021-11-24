package config

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
	PgTimeout  				int
}

func Get() *Config {
/*	pgTimeout, err := strconv.Atoi(os.Getenv("PG_TIMEOUT"))
	if err != nil {
		log.Fatal("fail to load config", err)
	}*/

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	config := Config {
			DbUser:     os.Getenv("DATABASE_USER"),
			DbPassword: os.Getenv("DATABASE_PASSWORD"),
			DbName:     os.Getenv("DATABASE_NAME"),
			DriverName: os.Getenv("DRIVER_NAME"),
			DbHost:     os.Getenv("DATABASE_HOST"),
			DbPort: 	os.Getenv("DATABASE_PORT"),
	}

	configBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Configuration:", string(configBytes))

	return &config
}
