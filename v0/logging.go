package wdk

import (
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func init() {
	logger = log.New()
	logger.SetLevel(log.FatalLevel)
}

func Logger() *log.Logger {
	return logger
}

func SetLogger(log *log.Logger) {
	logger = log
}

