package logger_interceptor

import (
	"context"
	"time"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type LoggerInterceptor interface {
	ServerLoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error)
	ClientLoggerInterceptor() func(ctx context.Context, method string, req interface{}, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error
}

type loggerInterceptor struct {
	logger logger.AppLogger
}

func NewLoggerInterceptor(logger logger.AppLogger) *loggerInterceptor {
	return &loggerInterceptor{logger: logger}
}

func (li *loggerInterceptor) ServerLoggerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)
	res, err := handler(ctx, req)

	gmf := &logger.GrpcFields{
		Method:   info.FullMethod,
		Request:  req,
		Response: res,
		Duration: time.Since(start),
		Metadata: md,
	}
	if err != nil {
		gmf.Error = err
		li.logger.GrpcErrorLogger(gmf)
		return res, err
	}

	li.logger.GrpcInfoLogger(gmf)
	return res, err
}

func (li *loggerInterceptor) ClientLoggerInterceptor() func(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return func(ctx context.Context, method string, req interface{}, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, res, cc, opts...)
		md, _ := metadata.FromIncomingContext(ctx)

		gmf := &logger.GrpcFields{
			Method:   method,
			Request:  req,
			Response: res,
			Duration: time.Since(start),
			Metadata: md,
		}
		if err != nil {
			gmf.Error = err
			li.logger.GrpcErrorLogger(gmf)
			return err
		}

		li.logger.GrpcErrorLogger(gmf)
		return err
	}
}
