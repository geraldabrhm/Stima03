package config

import (
	"os"
	"path/filepath"
	"github.com/joho/godotenv" // ? Kenapa gabisa diimport
	"github.com/apex/log" // ? Kenapa gabisa diimport 
)

const (
	POSTGRES_USER        = "POSTGRES_USER"
	POSTGRES_PASSWORD    = "POSTGRES_PASSWORD"
	POSTGRES_DB          = "POSTGRES_DB"
	CLIENT_URL           = "CLIENT_URL"
	SERVER_PORT          = "SERVER_PORT"
	// JWT_KEY              = "JWT_KEY"
	RUN_MIGRATION        = "RUN_MIGRATION"
	POSTGRES_SERVER_HOST = "POSTGRES_SERVER_HOST"
	ENVIRONMENT           = "ENV"
)

type ConfigType map[string]string

var Config = ConfigType{
	POSTGRES_USER:        "",
	POSTGRES_PASSWORD:    "",
	POSTGRES_DB:          "",
	CLIENT_URL:           "",
	SERVER_PORT:          "",
	// JWT_KEY:              "",
	RUN_MIGRATION:        "",
	POSTGRES_SERVER_HOST: "localhost",
}

func initConfig() {
	_, exist := os.LookupEnv(ENVIRONMENT)

	var path string
	
	if exist {
		path, _ = filepath.Abs("../.env")
	}

	if err := godotenv.Load(path); err != nil {
		log.WithField("Problem: ", err.Error()).Fatal("Env file not found")
	}

	req := map[string]bool {
		POSTGRES_USER:     true,
		POSTGRES_PASSWORD: true,
		POSTGRES_DB:       true,
		CLIENT_URL:        true,
		SERVER_PORT:       true,
		RUN_MIGRATION:     true,
	}

	for key := range Config {
		envVal, exists := os.LookupEnv(key)
		if !exists {
			if req[key] {
				log.Fatal(key + " is not found in Env")
			}
			continue
		}
		if _, ok := Config[key]; ok {
			Config[key] = envVal
		}
	}

	log.Info("Configuration settle")
}