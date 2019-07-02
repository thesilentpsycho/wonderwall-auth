package wlog

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	if Logger != nil {
		return
	}
	var initLoggerOnce sync.Once
	initLoggerOnce.Do(func() {
		Logger = logrus.New()
	})
}
