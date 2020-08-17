package main

import (
	"github.com/joho/godotenv"
	"net/http"
	"os"
	. "purchase/internal/logger"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error to start application: cannot load environment variables: " + err.Error())
	}

	SetupLogger()
}

func main() {
	Logger.Info("Starting application...")

	application, err := SetupApplication()
	if err != nil {
		Logger.Fatal("Setup Application Error", err)
		panic("Error to start application")
	}
	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		port = "8080"
	}

	handler := application.SystemRoutes.SetupHandler()
	server := &http.Server{
		Handler:      handler,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	Logger.Infof("Application listening on port %s", port)
	Logger.Fatal(server.ListenAndServe())
}
