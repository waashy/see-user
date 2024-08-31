package server

import (
	"context"
	"encoding/json"
	log "log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/waashy/see-user/api/handler"
	waashyLogger "github.com/waashy/utils/logger"
)

type ServerHandlerMap struct {
	APIPath string
	Handler handler.APIHandler
}

type Server struct {
	Config   *ServerConfig
	FiberApp *fiber.App
	Handlers []*ServerHandlerMap
	logger   *log.Logger
}

func AddServerHandler(apipath string, handler handler.APIHandler) *ServerHandlerMap {
	return &ServerHandlerMap{APIPath: apipath, Handler: handler}
}

func (s *Server) Setup() *Server {
	if s.FiberApp == nil {
		log.Error("Server setup failed!")
	}
	//middlewares.AddMiddlewares(s.FiberApp)

	for _, eachhandlermap := range s.Handlers {
		apigroup := s.FiberApp.Group(eachhandlermap.APIPath)
		eachhandlermap.Handler.RegisterRoutes(apigroup)
	}
	return s
}

func NewServer(Config *ServerConfig, FiberApp *fiber.App, Handlers []*ServerHandlerMap, logger *log.Logger) *Server {
	serv := &Server{
		Config,
		FiberApp,
		Handlers,
		logger,
	}
	serv.Setup()
	return serv
}

func (s *Server) StartServer() <-chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.FiberApp.Listen(":" + strconv.Itoa(s.Config.Port)); err != nil {
			log.Error("failed to start server", "ERR", err)
		}
	}()
	return quit
}

func NewFiberApplication(appLogger *log.Logger) *fiber.App {
	config := fiber.Config{
		ReadTimeout: 2 * time.Second,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}

	app := fiber.New(config)

	app.Use(requestid.New(),
		cors.New(),
		waashyLogger.RequestLogger(appLogger),
	)

	return app
}

func (s *Server) ShutdownGracefully() {
	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	shutdownChan := make(chan error, 1)
	go func() { shutdownChan <- s.FiberApp.Shutdown() }()

	select {
	case <-timeout.Done():
		log.Error("Server Shutdown Timed out before shutdown.")
	case err := <-shutdownChan:
		if err != nil {
			log.Error("Error while shutting down server", "ERR", err)
		} else {
			log.Info("Server Shutdown Successful")
		}
	}
}
