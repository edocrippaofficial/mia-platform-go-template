package main

import (
	"echotonic/config"
	"echotonic/controllers"
	"echotonic/router"
	"echotonic/services"

	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	entrypoint(make(chan os.Signal, 1))
	os.Exit(0)
}

func entrypoint(sigtermCh chan os.Signal) {
	envs := config.MustGetEnvs()
	log := config.MustGetLogger(envs.LogLevel)

	srv, err := startHttpServer(envs, log)
	if err != nil {
		log.WithError(err).Fatal("Unable to start HTTP server")
	}

	// Block and wait for the SIGTERM signal
	waitForShutdown(sigtermCh)
	handleGracefulShutdown(srv, log, envs)
}

func startHttpServer(envs config.Envs, log *logrus.Logger) (*http.Server, error) {
	router := router.NewRouter(log)

	svcs := services.NewServices()
	for _, ctr := range controllers.NewControllers(svcs) {
		ctr.RegisterRoutes(router)
	}

	srv := &http.Server{
		Handler: router.Handler,
		Addr:    fmt.Sprintf("0.0.0.0:%d", envs.HttpPort),
	}

	go func(srv *http.Server, log *logrus.Logger, envs config.Envs) {
		log.WithFields(map[string]any{"port": envs.HttpPort, "pid": os.Getpid()}).Info("Starting server")

		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}(srv, log, envs)

	return srv, nil
}

func waitForShutdown(sigtermCh chan os.Signal) {
	signal.Notify(sigtermCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigtermCh
}

func handleGracefulShutdown(srv *http.Server, log *logrus.Logger, envs config.Envs) {
	log.Info("Received SIGTERM signal")
	time.Sleep(time.Duration(envs.DelayShutdownSeconds) * time.Second)
	log.Info("Gracefully shutting down...")
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err.Error())
	}
}
