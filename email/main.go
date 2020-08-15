package main

import (
	. "email/internal/logger"
)

func init() {
	SetupLogger()
}

func main() {
	Logger.Info("Starting application...")

	application, err := SetupApplication()
	if err != nil {
		Logger.Fatal("Setup Application Error", err)
		panic("Error to start application")
	}

	application.CloudEventHandler.StartReceiver()

	Logger.Infof("Terminating application...")
}
