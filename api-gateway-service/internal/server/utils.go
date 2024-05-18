package server

import (
	"context"

	grpcClient "github.com/0x5w4/kredit-plus/pkg/grpc-client"
	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	readerService "github.com/0x5w4/kredit-plus/reader-service/proto/reader"

	v1 "github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/delivery/http/v1"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/service"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/middlewares"
	loggerInterceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
)

func (s *Server) setupLogger() error {
	var err error
	s.logger, err = loggerClient.NewAppLogger(s.cfg.Logger)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setupLoggerInterceptor() {
	s.loggerInterceptor = loggerInterceptor.NewLoggerInterceptor(s.logger)
}

func (s *Server) setupMiddleware() error {
	var err error
	s.mw = middlewares.NewMiddlewareManager(s.logger, s.cfg)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setupGrpcClient() error {
	var err error
	s.grpcClientConn, err = grpcClient.NewGrpcClient(s.cfg.GrpcClient.Port, s.loggerInterceptor, s.logger)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setupReaderServiceClient() error {
	var err error
	s.readerClient = readerService.NewReaderServiceClient(s.grpcClientConn)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setupService(kafkaProducer *kafkaClient.Producer) {
	s.service = service.NewKreditService(s.logger, s.cfg, kafkaProducer, s.readerClient)
}

func (s *Server) setupHttpHandler(cancel context.CancelFunc) error {
	kreditHandlers := v1.NewKreditHandler(s.echo.Group(s.cfg.Http.KreditPath), s.logger, s.mw, s.cfg, s.service, s.v)
	kreditHandlers.MapRoutes()

	go func() {
		if err := s.runHttpServer(); err != nil {
			s.logger.SLogger.Errorf(" s.runHttpServer: %v", err)
			cancel()
		}
	}()
	s.logger.SLogger.Infof("API Gateway is listening on PORT: %s", s.cfg.Http.Port)

	return nil
}
