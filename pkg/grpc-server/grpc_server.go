package grpc_server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	interceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

type GrpcServer struct {
	Server   *grpc.Server
	Listener net.Listener
	Config   Config
	Logger   *logger.AppLogger
}

type Config struct {
	Port string
}

const (
	maxConnectionIdle = 300
	gRPCTimeout       = 15
	maxConnectionAge  = 300
	gRPCTime          = 600
	maxReqSize        = 10e6
)

func NewGrpcServer(cfg Config, li interceptor.LoggerInterceptor, logger *logger.AppLogger, opts ...grpc.ServerOption) (*GrpcServer, error) {
	opts = append(
		opts,
		grpc.MaxRecvMsgSize(maxReqSize),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Second,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Second,
			Time:              gRPCTime * time.Second,
		}),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcTags.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(),
			li.ServerLoggerInterceptor,
		)),
	)

	server := grpc.NewServer(opts...)

	return &GrpcServer{
		Server: server,
		Config: cfg,
		Logger: logger,
	}, nil
}

func (s *GrpcServer) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", s.Config.Port))
	if err != nil {
		return errors.Wrap(err, "net.Listen")
	}
	s.Listener = listener

	go func() {
		if err := s.Server.Serve(s.Listener); err != nil {
			s.Logger.SLogger.Fatalf("grpcServer.Server.Serve: %v", err)
		}
	}()

	return nil
}

func (s *GrpcServer) Stop(ctx context.Context) {
	if err := s.Listener.Close(); err != nil {
		s.Logger.SLogger.Fatalf("grpcServer.Listener.Close: %v", err)
	}

	go func() {
		defer s.Server.GracefulStop()
		<-ctx.Done()
	}()
}
