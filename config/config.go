package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort      string
	SecretKey     string
	RedisAddr     string
	Smtp          Smtp
	Authorization Authorization
	Postgres      PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Smtp struct {
	Sender   string
	Password string
}

type Authorization struct {
	HeaderKey  string
	PayloadKey string
}

func New(path string) Config {
	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		HttpPort:  conf.GetString("HTTP_PORT"),
		SecretKey: conf.GetString("SECRET_KEY"),
		RedisAddr: conf.GetString("REDIS_ADDR"),
		Postgres: PostgresConfig{
			Host:     conf.GetString("POSTGRES_HOST"),
			Port:     conf.GetString("POSTGRES_PORT"),
			User:     conf.GetString("POSTGRES_USER"),
			Password: conf.GetString("POSTGRES_PASSWORD"),
			Database: conf.GetString("POSTGRES_DATABASE"),
		},
		Smtp: Smtp{
			Sender:   conf.GetString("SMTP_SENDER"),
			Password: conf.GetString("SMTP_PASSWORD"),
		},
		Authorization: Authorization{
			HeaderKey:  conf.GetString("AUTH_HEADER_KEY"),
			PayloadKey: conf.GetString("AUTH_PAYLOAD_KEY"),
		},
	}

	return cfg
}
