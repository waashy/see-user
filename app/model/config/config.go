package model

import (
	server "github.com/waashy/see-user/api/server"
	dbConfig "github.com/waashy/utils/database/config"
)

type AppConfig struct {
	Server   *server.ServerConfig `json:"server_config"`
	LogLevel string               `json:"log_level"` // DEBUG | ERROR | WARN | INFO
	DBConfig *dbConfig.DBConfig   `json:"db_config"`
}
