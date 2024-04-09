package wdk

import (
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger = log.New()

func Logger() *log.Logger {
	return logger
}

func SetLogger(log *log.Logger) {
	logger = log
}
