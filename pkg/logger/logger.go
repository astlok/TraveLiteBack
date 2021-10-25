package logger

import (
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
)

func InitLogger(stream string, level string) error {
	if stream == "stdout" {
		log.SetOutput(os.Stdout)
	} else {
		logFileStream, err := os.Open(stream)
		if err != nil {
			return err
		}
		log.SetOutput(logFileStream)
	}

	logrusLevel, err := log.ParseLevel(level)
	if err != nil {
		return err
	}
	log.SetLevel(logrusLevel)

	return nil
}

func LoggingRequest(reqID uint64, url *url.URL, method string) {
	log.WithFields(log.Fields{
		"request_id": reqID,
		"url":        url,
		"method":     method,
	}).Info()
}

func LoggingResponse(reqID uint64, code int) {
	log.WithFields(log.Fields{
		"request_id": reqID,
		"reply_code": code,
	}).Info()
}

func LoggingError(reqID uint64, err error) {
	log.WithFields(log.Fields{
		"request_id": reqID,
		"error":      err,
	}).Error()
}