package util

import (
	"fmt"
	"log"
	"os"

	"chit-chat/internal/di/config"
)

const logFile = "chitchat.log"

func Initialize() (*config.Configuration, *log.Logger, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("load config: %w", err)
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("open file %s: %w", logFile, err)
	}

	logger := log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

	return cfg, logger, nil
}
