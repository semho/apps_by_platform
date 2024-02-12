package main

import (
	"platform/config"
	"platform/logging"
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
	var cfg config.Configuration
	var err error
	cfg, err = config.Load("config.json")
	if err != nil {
		panic(err)
	}
	var logger = logging.NewDefaultLogger(cfg)
	writeMessage(logger, cfg)
}
