package model

import (
	dbConfig "github.com/waashy/utils/database/config"
)

type ServerConfig struct {
	Port int `json:"port"`
}

type AppConfig struct {
	Server   ServerConfig      `json:"server_config"`
	LogLevel string            `json:"log_level"` // DEBUG | ERROR | WARN | INFO
	DBConfig dbConfig.DBConfig `json:"db_config"`
}
