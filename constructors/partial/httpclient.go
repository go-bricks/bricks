package partial

import (
	"github.com/go-bricks/bricks/v2/http/client"
	clientInt "github.com/go-bricks/bricks/v2/interfaces/http/client"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

const (
	// FxGroupRESTClientInterceptors defines group name
	FxGroupRESTClientInterceptors = "restClientInterceptors"
	// FxGroupGRPCUnaryClientInterceptors defines group name
	FxGroupGRPCUnaryClientInterceptors = "grpcUnaryClientInterceptors"
)

type httpClientBuilderDeps struct {
	fx.In

	Interceptors []clientInt.HTTPClientInterceptor `group:"restClientInterceptors"`
}

// HTTPClientBuilder creates an injectable http.Client builder that can be predefined with Interceptors
//
// This function returns a closure that will always create a new builder. That way every usage can add different
// interceptors without influencing others
func HTTPClientBuilder(deps httpClientBuilderDeps) clientInt.NewHTTPClientBuilder {
	return func() clientInt.HTTPClientBuilder {
		return client.HTTPClientBuilder().AddInterceptors(deps.Interceptors...)
	}
}

// GRPC
type grpcClientConnectionBuilderDeps struct {
	fx.In

	Interceptors []grpc.UnaryClientInterceptor `group:"grpcUnaryClientInterceptors"`
}

// GRPCClientConnectionBuilder creates an injectable grpc.ClientConn that can be predefined with Interceptors
// or/and additional options later
func GRPCClientConnectionBuilder(deps grpcClientConnectionBuilderDeps) clientInt.GRPCClientConnectionBuilder {
	interceptors := grpc.WithChainUnaryInterceptor(deps.Interceptors...)
	return client.GRPCClientConnBuilder().AddOptions(interceptors)
}
