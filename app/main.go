package main

import (
	log "log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/waashy/utils/config/parser"
	database "github.com/waashy/utils/database"
	wasshylogger "github.com/waashy/utils/logger"

	healthHdlr "github.com/waashy/see-user/api/handler/health"
	userHdlr "github.com/waashy/see-user/api/handler/user"
	server "github.com/waashy/see-user/api/server"
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

	initLogger := log.New(log.NewJSONHandler(os.Stdout, &log.HandlerOptions{
		Level: log.LevelInfo,
	}))

	initLogger.Info("See-User service initiating")

	err := parser.ConfigParser(CONFIG_FILE_PATH, &appConfig)
	if err != nil {
		initLogger.Error("failed to parse the configuration", "Err", err)
		return
	}
	initLogger.Info("config parser loaded")

	logger, err = wasshylogger.NewLogger(appConfig.LogLevel)
	if err != nil {
		initLogger.Error("failed to parse the configuration", "Err", err)
		return
	}
	initLogger.Info("logger lodded")
}

func main() {

	logger.Info("See-User service Starting")

	/////////////////////////////// DATABASE INITIATION ///////////////////////////////
	logger.Info("establishing database connection")
	db, err := database.NewDatabase(*appConfig.DBConfig)
	if err != nil {
		logger.Error("initiating database failed", "ERR", err)
		return
	}
	logger.Info("database connection established")

	///////////////////////////////////// SERVICES ////////////////////////////////////////

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

	logger.Info("started user service processes")
	if err := userService.Start(); err != nil {
		logger.Error("failed to start user service processes", "ERR", err)
		return
	}
	logger.Info("started user service processes")

	///////////////////////////////////// HANDLERS ////////////////////////////////////////

	healthHandler := healthHdlr.NewHealthCheckHandler()
	userHandler := userHdlr.NewUserHandler(userService)

	handlers := []*server.ServerHandlerMap{
		server.AddServerHandler("api/v1/health/", healthHandler),
		server.AddServerHandler("api/v1/user/", userHandler),
	}

	///////////////////////////////////// SERVER ////////////////////////////////////////
	app := server.NewFiberApplication(logger)

	server := server.NewServer(appConfig.Server, app, handlers, logger)
	gracefulServerQuit := server.StartServer()

	//////////////////////////////// APPLICATION HOLD ///////////////////////////////////
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	logger.Info("application start up completed", "stop procedure", "press ^C to stop the application")

	select {
	case receivedStopSignal := <-quit:
		logger.Info("recieved stop signal", "stop signal", receivedStopSignal)
	case receivedStopSignal := <-gracefulServerQuit:
		logger.Info("recieved stop signal", "stop signal", receivedStopSignal)
	}

	/////////////////////////////////// STOP SUBPROCESSES /////////////////////////////////////

	server.ShutdownGracefully(logger)

	logger.Info("stopping all sub-processes")
	if err := userService.Stop(); err != nil {
		logger.Error("failed to stop user service processes", "ERR", err)
		return
	}

	logger.Info("all sub-processes succesfully stopped")

	logger.Info("application succesfully stopped")
}
