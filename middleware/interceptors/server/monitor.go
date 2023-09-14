package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-bricks/bricks/v2/interfaces/log"
	"github.com/go-bricks/bricks/v2/interfaces/monitor"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const (
	gRPCCodeTagName = "code"
	grpcNamePrefix  = "grpc_"
)

type gRPCMetricInterceptorsDeps struct {
	fx.In

	Logger  log.Logger
	Metrics monitor.Metrics `optional:"true"`
}

// MonitorGRPCInterceptor sends gRPC method invocation metrics to the configured Metrics server (Prometheus, Datadog)
func MonitorGRPCInterceptor(deps gRPCMetricInterceptorsDeps) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		if deps.Metrics != nil {
			_, methodName := utils.SplitMethodAndPackage(info.FullMethod)
			// fetch one from registry or create new
			timer := deps.Metrics.WithTags(monitor.Tags{
				gRPCCodeTagName: gRPCCodeTagValue(err),
			}).Timer(grpcNamePrefix+methodName, fmt.Sprintf("time api calls for %s", info.FullMethod))

			timer.Record(time.Since(start))
		}
		return
	}
}

func gRPCCodeTagValue(err error) string {
	s, _ := status.FromError(err)
	return strconv.Itoa(int(s.Code()))
}
