package main

import (
	log "log/slog"

	"github.com/waashy/utils/config/parser"
	database "github.com/waashy/utils/database"
	dbConfig "github.com/waashy/utils/database/config"
	wasshylogger "github.com/waashy/utils/logger"

	config "github.com/waashy/see-user/app/model/config"
	userdao "github.com/waashy/see-user/database/dao/user"
	userservice "github.com/waashy/see-user/service/user"
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
	/////////////////////////////// DATABASE INITIATION ///////////////////////////////
	logger.Info("Establishing database connection")
	db, err := database.NewDatabase(dbConfig.DBConfig{})
	if err != nil {
		logger.Error("initiating database failed", "ERR", err)
		return
	}
	logger.Info("Database connection established")

	///////////////////////////////////// USER ////////////////////////////////////////
	/////////////////////////////// USER DAO INITIATION ///////////////////////////////
	logger.Info("initiating user dao")
	userDao, err := userdao.NewUserDao(db, logger)
	if err != nil {
		logger.Error("initiating user dao failed", "ERR", err)
		return
	}
	logger.Info("initiated user dao")

	/////////////////////////////// USER SERVICE INITIATION ///////////////////////////////
	logger.Info("initiating user service")
	userService, err := userservice.NewUserService(userDao, logger)
	if err != nil {
		logger.Error("initiating user service failed", "ERR", err)
		return
	}
	logger.Info("initiated user service")

	/////////////////////////////// START PROCESSES ///////////////////////////////
	if err := userService.Start(); err != nil {
		logger.Info("failed to start user service processes")
		return
	}

}
