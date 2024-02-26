package main

import (
	"api/app"
	"api/logger"
)

func main() {

	logger.Info("Starting the application....")

	app.Start()

}
