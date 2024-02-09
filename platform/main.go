package main

import "platform/logging"

func writeMessage(logger logging.Logger) {
	logger.Info("Hello")
}

func main() {
	var logger = logging.NewDefaultLogger(logging.Information)
	writeMessage(logger)
}
