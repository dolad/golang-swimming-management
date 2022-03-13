package config

import (
	env "swimming-content-management/utils"
)

type Config struct {
	Environment                string
	Port                       string
	Database                   *Database
	RefreshTokenPrivateKeyPath string
	AccessTokenPrivateKeyPath  string
	AccessTokenPublicKeyPath   string
	RefreshTokenPublicKeyPath  string
	JwtExpiration              int
}

type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

func NewConfig() (*Config, error) {
	env.CheckDotEnv()
	port := env.MustGet("PORT")
	if port == "" {
		port = "5004"
	}

	return &Config{
		Environment: env.MustGet("ENV"),
		Port:        port,
		Database: &Database{
			Host:     env.MustGet("DATABASE_HOST"),
			Port:     env.MustGet("DATABASE_PORT"),
			User:     env.MustGet("DATABASE_USER"),
			DB:       env.MustGet("DATABASE_DB"),
			Password: env.MustGet("DATABASE_PASSWORD"),
		},
		AccessTokenPrivateKeyPath: "./jwtcert/key.pem",
		AccessTokenPublicKeyPath:  "./jwtcert/key.pem.pub",
		JwtExpiration:             60,
	}, nil

}
