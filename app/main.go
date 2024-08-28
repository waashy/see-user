package main

import (
	log "log/slog"

	config "github.com/waashy/see-user/app/model/config"
	"github.com/waashy/utils/config/parser"
	database "github.com/waashy/utils/database"
	dbConfig "github.com/waashy/utils/database/config"
	wasshylogger "github.com/waashy/utils/logger"
)

const CONFIG_FILE_PATH = "../config/runConfig.json"

var (
	logger    *log.Logger
	appConfig config.AppConfig
)

func init() {
	log.Info("See-User service initiating")

	err := parser.ConfigParser(CONFIG_FILE_PATH, &appConfig)
	if err != nil {
		log.Error("Failed to parse the configuration", "Err", err)
		return
	}
	log.Info("config parser loaded")

	logger, err = wasshylogger.NewLogger(appConfig.LogLevel)
	if err != nil {
		log.Error("Failed to parse the configuration", "Err", err)
		return
	}
	log.Info("logger lodded")
}

func main() {

	logger.Info("See-User service Starting")

	logger.Info("Establishing database connection")
	db, err := database.NewDatabase(dbConfig.DBConfig{})
	if err != nil {
		return
	}
	logger.Info("Database connection established")

}
