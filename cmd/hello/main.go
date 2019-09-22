package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/alexpil/hello/internal/diagnostics"
	"github.com/alexpil/hello/internal/handlers"
)

func main() {
	log := logrus.New()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set")
	}

	diagPort := os.Getenv("DIAG_PORT")
	if port == "" {
		log.Fatal("Diagnostics port is not set")
	}

	log.Info("Application is starting...")
	log.Infof("Version: %s, hash: %s, build time: %s", diagnostics.Version, diagnostics.Hash, diagnostics.BuildTime)

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", handlers.Hello(log)).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Infof("The server is running on port %s", port)

	// Diagnostics server
	diagServer := diagnostics.NewServer(log, diagPort)

	go func() {
		err := diagServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Infof("Diagnostics server is running on port %s", diagPort)

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			log.Info("Got SIGINT...")
		case syscall.SIGTERM:
			log.Info("Got SIGTERM...")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Error(err)
	}
	log.Info("Server shutted down properly")
}
