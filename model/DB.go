package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	Host     string `yaml:"db.host"`
	Port     string `yaml:"db.port"`
	Username string `yaml:"db.username"`
	Password string `yaml:"db.password"`
	DBName   string `yaml:"db.dbname"`
	SSLMode  string `yaml:"db.sslmode"`
}

func NewDataBase(cfg DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
