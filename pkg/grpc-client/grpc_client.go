package grpc_client

import (
	"fmt"
	"time"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	interceptor "github.com/0x5w4/kredit-plus/pkg/logger-interceptor"
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	backoffLinear  = 100 * time.Millisecond
	backoffRetries = 3
)

func NewGrpcClient(port string, li interceptor.LoggerInterceptor, logger *logger.AppLogger, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(
		opts,
		grpc.WithUnaryInterceptor(li.ClientLoggerInterceptor()),
		grpc.WithUnaryInterceptor(grpcRetry.UnaryClientInterceptor([]grpcRetry.CallOption{
			grpcRetry.WithBackoff(grpcRetry.BackoffLinear(backoffLinear)),
			grpcRetry.WithCodes(codes.NotFound, codes.Aborted),
			grpcRetry.WithMax(backoffRetries),
		}...)),
	)

	conn, err := grpc.NewClient(fmt.Sprintf(":%v", port), opts...)
	if err != nil {
		return nil, errors.Wrap(err, "grpc.NewClient")
	}

	return conn, nil
}
