package server

import (
	"context"

	"github.com/0x5w4/kredit-plus/docs"
	echoClient "github.com/0x5w4/kredit-plus/pkg/echo"
	grpcClient "github.com/0x5w4/kredit-plus/pkg/grpc-client"
	kafkaClient "github.com/0x5w4/kredit-plus/pkg/kafka"
	loggerClient "github.com/0x5w4/kredit-plus/pkg/logger"
	readerService "github.com/0x5w4/kredit-plus/reader-service/proto/reader"

	v1 "github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/delivery/http/v1"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/service"
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

func (s *Server) setupGrpcClient() error {
	var err error
	s.grpcClientConn, err = grpcClient.NewGrpcClient(s.cfg.GrpcClient.Port, s.loggerInterceptor, s.logger)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setupReaderServiceClient() error {
	s.readerClient = readerService.NewReaderServiceClient(s.grpcClientConn)
	return nil
}

func (s *Server) setupService(kafkaProducer *kafkaClient.Producer) {
	s.service = service.NewKreditService(s.logger, s.cfg, kafkaProducer, s.readerClient)
}

func (s *Server) setupEcho() {
	s.echo = echoClient.NewEcho()
}

func (s *Server) setupSwagger() {
	docs.SwaggerInfo.BasePath = s.cfg.Http.BasePath
	docs.SwaggerInfo.Host = "localhost" + s.cfg.Http.Port
}

func (s *Server) setupHttpHandler() {
	kreditHandlers := v1.NewKreditHandler(s.echo.Group(s.cfg.Http.KreditPath), s.logger, s.cfg, s.service, s.v)
	kreditHandlers.MapRoutes()
}

func (s *Server) runEcho(cancel context.CancelFunc) {
	if err := s.echo.Start(s.cfg.Http.Port); err != nil {
		s.logger.SLogger.Fatalf("Failed to run echo: %v", err)
		cancel()
	}
}
