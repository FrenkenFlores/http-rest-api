package apiserver

import (
	"log"

	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
}

func New(new_config *Config) *APIserver {
	return &APIserver{
		logger: logrus.New(),
		config: new_config,
	}
}

func (s *APIserver) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		log.Fatal(err)
	}
	s.logger.Info("Starting API server")
	return nil
}

func (s *APIserver) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	s.logger.SetLevel(level)
	return nil
}
