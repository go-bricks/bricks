package providers

import (
	"github.com/go-bricks/bricks/constructors"
	"github.com/go-bricks/bricks/middleware/interceptors/server"
	"github.com/go-bricks/bricks/providers/groups"
	"go.uber.org/fx"
)

// TODO Revisit Monitoring

// MonitorFxOption adds default metric client to the graph
func MonitorFxOption() fx.Option {
	return fx.Provide(constructors.DefaultMonitor)
}

// MonitorGRPCInterceptorFxOption adds Unary Server Interceptor that will notify metric provider of every call
func MonitorGRPCInterceptorFxOption() fx.Option {
	return fx.Provide(
		fx.Annotated{
			Group:  groups.UnaryServerInterceptors,
			Target: server.MonitorGRPCInterceptor,
		})
}
