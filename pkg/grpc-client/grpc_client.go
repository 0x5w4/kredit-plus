package grpc_client

import (
	"context"
	"fmt"
	"time"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	interceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	backoffLinear  = 100 * time.Millisecond
	backoffRetries = 3
)

type GrpcClient struct {
	Connection *grpc.ClientConn
	Config     Config
	Logger     logger.AppLogger
}

type Config struct {
	Network string
	Port    string
	Tls     bool
}

func NewGrpcClient(ctx context.Context, cfg Config, li interceptor.LoggerInterceptor, logger logger.AppLogger, opts ...grpc.DialOption) (*GrpcClient, error) {
	if cfg.Tls {
		certFile := "ssl/certificates/ca.crt" // => file path location your certFile
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			logger.SLogger.Fatalf("credentials.NewClientTLSFromFile: %v", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	opts = append(
		opts,
		grpc.WithUnaryInterceptor(li.ClientLoggerInterceptor()),
		grpc.WithUnaryInterceptor(grpcRetry.UnaryClientInterceptor([]grpcRetry.CallOption{
			grpcRetry.WithBackoff(grpcRetry.BackoffLinear(backoffLinear)),
			grpcRetry.WithCodes(codes.NotFound, codes.Aborted),
			grpcRetry.WithMax(backoffRetries),
		}...)),
	)

	conn, err := grpc.DialContext(ctx, fmt.Sprintf(":%v", cfg.Port), opts...)
	if err != nil {
		return nil, errors.Wrap(err, "grpc.DialContext")
	}
	return &GrpcClient{
		Connection: conn,
		Config:     cfg,
		Logger:     logger,
	}, nil
}

func (c *GrpcClient) Close(ctx context.Context) {
	if err := c.Connection.Close(); err != nil {
		c.Logger.SLogger.Fatalf("grpcClient.Connection.Close: %v", err)
	}

	go func() {
		<-ctx.Done()
	}()
}
