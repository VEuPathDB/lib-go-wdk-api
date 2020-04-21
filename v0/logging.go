package wdk

import (
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

func init() {
	tmp := log.New()
	tmp.SetLevel(log.FatalLevel)
	logger = tmp.WithField("sourceLib", "wdk")
}

func Logger() *log.Logger {
	return logger.Logger
}

func SetLogger(log *log.Logger) {
	logger = log.WithField("sourceLib", "wdk")
}

