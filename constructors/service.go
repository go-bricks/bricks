package constructors

import (
	"context"
	"fmt"
	"github.com/go-bricks/bricks/http/server/health"
	"github.com/go-bricks/bricks/interfaces/http/server"
	"github.com/go-bricks/bricks/interfaces/log"
	"github.com/go-bricks/bricks/utils"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type webServiceDependencies struct {
	fx.In

	LifeCycle         fx.Lifecycle
	Logger            log.Logger
	WebServiceBuilder server.GRPCWebServiceBuilder
}

// Service should be invoked by FX, it will build the entire dependencies graph and add lifecycle hooks
func Service(deps webServiceDependencies) (server.WebService, error) {
	webService, err := deps.WebServiceBuilder.Build()
	if err != nil {
		return nil, err
	}
	deps.LifeCycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go webService.Run(ctx) // this should exit only when service was shutdown
			return deps.pingService(ctx, webService)
		},
		OnStop: func(ctx context.Context) error {
			return webService.Stop(ctx)
		},
	})
	return webService, nil
}

func (deps webServiceDependencies) pingService(ctx context.Context, service server.WebService) (err error) {
	err = fmt.Errorf("failed to check internal service health")
	ports := service.Ports()
	tlsCfg := service.TlsConfig()
	if grpcAddress := deps.getGRPCAddress(ports); len(grpcAddress) > 0 {
		var conn *grpc.ClientConn
		transportOption := grpc.WithInsecure()
		if tlsCfg != nil && tlsCfg.EnableTLS {
			deps.Logger.Debug(ctx, "grpc TLS enabled")
			tlsCredentials, err := utils.LoadTLSCredentials(tlsCfg.CaCertFile, tlsCfg.ClientKeyFile, tlsCfg.ClientCertFile)
			if err != nil {
				deps.Logger.Error(ctx, fmt.Sprintf("error to configure tls connection(check certificate files), err: %s", err.Error()))
				return err
			} else {
				transportOption = grpc.WithTransportCredentials(tlsCredentials)
			}
		} else {
			deps.Logger.Debug(ctx, "grpc TLS disabled")
		}
		if conn, err = grpc.DialContext(ctx, grpcAddress, transportOption); err == nil {
			defer conn.Close()
			healthClient := health.NewHealthClient(conn)
			_, err = healthClient.Check(ctx, &health.HealthCheckRequest{})
		}
	}
	if err == nil {
		for _, info := range ports {
			deps.Logger.Debug(ctx, "Service is accepting %s calls on %s", info.Type, info.Address)
		}
		deps.Logger.Debug(ctx, "Service is up")
	}
	return
}
func (deps webServiceDependencies) getGRPCAddress(ports []server.ListenInfo) string {
	for _, info := range ports {
		if info.Type == server.GRPCServer {
			return info.Address
		}
	}
	return ""
}
