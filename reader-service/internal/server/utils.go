package server

import (
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
)

func (s *Server) setupLogger() error {
	var err error

	s.appLogger, err = loggerClient.NewAppLogger(s.cfg.Logger)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setupLoggerInterceptor() {
	s.loggerInterceptor = loggerInterceptor.NewLoggerInterceptor(s.appLogger)
}
