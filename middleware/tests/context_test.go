package tests

import (
	"context"

	"github.com/go-bricks/bricks/v2/interfaces/cfg"
	confkeys "github.com/go-bricks/bricks/v2/interfaces/cfg/keys"
	mock_cfg "github.com/go-bricks/bricks/v2/interfaces/cfg/mock"
	contextMiddleware "github.com/go-bricks/bricks/v2/middleware/context"
	"go.uber.org/fx"
	"google.golang.org/grpc/metadata"
)

func (s *middlewareSuite) TestLoggerGRPCIncomingContextExtractor() {
	md := metadata.MD{
		"one-of-a-kind": []string{"1"},
		"another-kind":  []string{"not extracted"},
	}
	ctxWithValues := metadata.NewIncomingContext(context.Background(), md)
	extracted := s.logExtractor(ctxWithValues)
	s.Contains(extracted, "one-of-a-kind")
	s.NotContains(extracted, "another-kind")
}

func (s *middlewareSuite) testLoggerGRPCIncomingContextExtractorBeforeTest() fx.Option {
	s.cfgMock.EXPECT().Get(confkeys.LoggerIncomingGRPCMetadataHeadersExtractor).DoAndReturn(func(key string) cfg.Value {
		value := mock_cfg.NewMockValue(s.ctrl)
		value.EXPECT().IsSet().Return(true)
		value.EXPECT().StringSlice().Return([]string{
			"one", "two",
		})
		return value
	})
	return fx.Options(
		fx.Provide(contextMiddleware.LoggerGRPCIncomingContextExtractor),
		fx.Populate(&s.logExtractor),
	)
}
