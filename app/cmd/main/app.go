package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"syscall"
	"template/app/internal/config"
	template2 "template/app/internal/template"
	"template/app/internal/template/db"
	"template/app/pkg/client/postgresql"
	"template/app/pkg/logging"
	"template/app/pkg/shutdown"
	"time"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")
	logger.Println("config initializing")
	cfg := config.GetConfig()

	logger.Println("router initializing")
	router := mux.NewRouter()

	postgresqlClient, err := postgresql.NewClient(context.Background(), cfg.Storage)
	if err != nil {
		logger.Println("failed to connect to postgresql")
		logger.Error(err)
	}

	userStorage := db.NewStorage(postgresqlClient, logger)

	userService, err := template2.NewService(userStorage, logger)

	if err != nil {
		logger.Fatal(err)
	}

	templateHandler := template2.Handler{
		Logger:      logger,
		UserService: userService,
		Validator:   validator.New(),
	}

	templateHandler.Register(router)

	logger.Println("Start microservice")

	start(router, logger, cfg)
}

func start(router http.Handler, logger logging.Logger, cfg *config.Config) {
	var server *http.Server

	server = &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Listen.Port),
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown.Graceful(
		[]os.Signal{
			syscall.SIGABRT,
			syscall.SIGQUIT,
			syscall.SIGHUP,
			syscall.SIGTERM,
			os.Interrupt,
		},
		server,
	)

	logger.Println("application initialized and started")

	if err := server.ListenAndServe(); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err)
		}
	}

}
