package main

import (
	"platform/config"
	"platform/logging"
	"platform/services"
)

func writeMessage(logger logging.Logger, cfg config.Configuration) {
	section, ok := cfg.GetSection("main")
	if !ok {
		logger.Panic("Config section not found")
	}
	message, ok := section.GetString("message")
	if !ok {
		logger.Panic("Connot find configuration setting")
	}
	logger.Info(message)
}

func main() {
	services.RegisterDefaultServices()

	services.Call(writeMessage)
}
