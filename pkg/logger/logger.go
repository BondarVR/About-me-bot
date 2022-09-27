package logger

import (
	"github.com/gemnasium/logrus-graylog-hook"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

type LogrusLogger struct {
	logrus *logrus.Logger
	entry  *logrus.Entry
}

type Config struct {
	LogLevel    string
	LogServer   string
	ServiceName string
}

func New(cfg Config) (*LogrusLogger, error) {
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse log level")
	}

	logger := &LogrusLogger{
		logrus: logrus.New(),
	}

	logger.logrus.SetLevel(level)

	customFormatter := &logrus.JSONFormatter{
		TimestampFormat: time.Layout,
	}

	logger.logrus.SetFormatter(customFormatter)

	if cfg.LogServer != "" {
		logger.logrus.AddHook(
			graylog.NewGraylogHook(cfg.LogServer, map[string]interface{}{}),
		)
	}

	logger.entry = logger.logrus.WithFields(logrus.Fields{
		"service_name": cfg.ServiceName,
	})

	return logger, nil
}
