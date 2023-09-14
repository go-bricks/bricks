package logger

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	logInt "github.com/go-bricks/bricks/interfaces/log"
	mock_log "github.com/go-bricks/bricks/interfaces/log/mock"
	"github.com/go-bricks/bricks/logger/naive"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

const (
	skipFramesExpect = 1
)

type wrapperSuite struct {
	suite.Suite
}

func TestWrapper(t *testing.T) {
	suite.Run(t, new(wrapperSuite))
}

func (s *wrapperSuite) TestLogLevels() {
	var output bytes.Buffer
	builder := naive.Builder().SetWriter(&output).SetLevel(logInt.TraceLevel)
	logger := CreateBricksLogger(builder)
	logger.Trace(nil, "trace line")
	s.Contains(output.String(), "trace line")
	logger.Debug(nil, "debug line")
	s.Contains(output.String(), "debug line")
	logger.Info(nil, "info line")
	s.Contains(output.String(), "info line")
	logger.Warn(nil, "warn line")
	s.Contains(output.String(), "warn line")
	logger.Error(nil, "error line")
	s.Contains(output.String(), "error line")
	logger.Custom(nil, logInt.ErrorLevel, 0, "custom line")
	s.Contains(output.String(), "custom line")
}

func (s *wrapperSuite) TestSetLevel() {
	var output bytes.Buffer
	builder := naive.Builder().SetWriter(&output).SetLevel(logInt.InfoLevel)
	logger := CreateBricksLogger(builder)
	logger.Debug(nil, "no debug line")
	s.Empty(output.String(), "debug printed")
	logger.Info(nil, "info printed")
	s.Contains(output.String(), "info printed")
}

func (s *wrapperSuite) TestIncludeCallerNoAdditionalSkipFrames() {
	var output bytes.Buffer
	builder := naive.Builder().SetWriter(&output).IncludeCaller().SetLevel(logInt.TraceLevel)
	logger := CreateBricksLogger(builder)
	logger.Trace(nil, "trace line")
	s.Contains(output.String(), "logger/wrapper_test.go")
	output.Reset()
	logger.WithField("nothing", "printed").Trace(nil, "trace line")
	s.Contains(output.String(), "logger/wrapper_test.go")
}

func (s *wrapperSuite) TestIncludeCallerAddSkipFrames() {
	var output bytes.Buffer
	builder := naive.Builder().SetWriter(&output).IncludeCaller().IncrementSkipFrames(1).SetLevel(logInt.InfoLevel)
	logger := CreateBricksLogger(builder)
	logger.Info(nil, "info line")
	s.Contains(output.String(), "reflect/value.go") // to much
}

func (s *wrapperSuite) TestIncludeCallerWithFields() {
	var output bytes.Buffer
	builder := naive.Builder().SetWriter(&output).IncludeCaller().SetLevel(logInt.InfoLevel)
	logger := CreateBricksLogger(builder)
	logger.WithField("field", "value").Info(nil, "info line")
	s.Contains(output.String(), "logger/wrapper_test.go")
	output.Reset()
	logger.WithField("field", "value").Custom(nil, logInt.InfoLevel, 0, "info line")
	s.Contains(output.String(), "logger/wrapper_test.go")
}

func (s *wrapperSuite) TestFields() {
	controller := gomock.NewController(s.T())
	mockLogger := mock_log.NewMockLogger(controller)
	mockBuilder := mock_log.NewMockBuilder(controller)
	mockBuilder.EXPECT().Build().Return(mockLogger).After(
		mockBuilder.EXPECT().IncrementSkipFrames(skipFramesExpect).Return(mockBuilder),
	)
	logger := CreateBricksLogger(mockBuilder)
	mockLogger.EXPECT().Custom(gomock.Not((context.Context)(nil)), logInt.InfoLevel, 0, "info line").After(
		mockLogger.EXPECT().WithField("field", "value").Return(mockLogger),
	)
	logger.WithField("field", "value").Info(nil, "info line")
	controller.Finish()
}

func (s *wrapperSuite) TestFieldsNotOverlapping() {
	controller := gomock.NewController(s.T())
	mockLogger := mock_log.NewMockLogger(controller)
	mockBuilder := mock_log.NewMockBuilder(controller)
	mockBuilder.EXPECT().Build().Return(mockLogger).After(
		mockBuilder.EXPECT().IncrementSkipFrames(skipFramesExpect).Return(mockBuilder),
	)
	logger := CreateBricksLogger(mockBuilder)
	firstFieldCall := mockLogger.EXPECT().Custom(gomock.Not((context.Context)(nil)), logInt.InfoLevel, 0, "info line").After(
		mockLogger.EXPECT().WithField("first", "value").Return(mockLogger),
	)
	mockLogger.EXPECT().Custom(gomock.Not((context.Context)(nil)), logInt.ErrorLevel, 0, "error line").After(
		mockLogger.EXPECT().WithField("second", "value").Return(mockLogger).After(
			firstFieldCall,
		),
	)

	logger.WithField("first", "value").Info(nil, "info line")
	logger.WithField("second", "value").Error(nil, "error line")
	controller.Finish()
}

func (s *wrapperSuite) TestError() {
	notRelevantError := fmt.Errorf("missing error")
	err := fmt.Errorf("this is an error")
	controller := gomock.NewController(s.T())
	mockLogger := mock_log.NewMockLogger(controller)
	mockBuilder := mock_log.NewMockBuilder(controller)
	mockBuilder.EXPECT().Build().Return(mockLogger).After(
		mockBuilder.EXPECT().IncrementSkipFrames(skipFramesExpect).Return(mockBuilder),
	)
	logger := CreateBricksLogger(mockBuilder)
	mockLogger.EXPECT().Custom(gomock.Not((context.Context)(nil)), logInt.InfoLevel, 0, "info line").After(
		mockLogger.EXPECT().WithError(err).Return(mockLogger),
	)
	logger.WithError(notRelevantError).WithError(err).Info(nil, "info line")
	controller.Finish()
}

func (s *wrapperSuite) TestErrorsNotOverlapping() {
	err1 := fmt.Errorf("this is one error")
	err2 := fmt.Errorf("here is another")
	controller := gomock.NewController(s.T())
	mockLogger := mock_log.NewMockLogger(controller)
	mockBuilder := mock_log.NewMockBuilder(controller)
	mockBuilder.EXPECT().Build().Return(mockLogger).After(
		mockBuilder.EXPECT().IncrementSkipFrames(skipFramesExpect).Return(mockBuilder),
	)
	logger := CreateBricksLogger(mockBuilder)
	firstCall := mockLogger.EXPECT().Custom(gomock.Not((context.Context)(nil)), logInt.InfoLevel, 0, "info line").After(
		mockLogger.EXPECT().WithError(err1).Return(mockLogger),
	)
	mockLogger.EXPECT().Custom(gomock.Not((context.Context)(nil)), logInt.WarnLevel, 0, "warn line").After(
		mockLogger.EXPECT().WithError(err2).Return(mockLogger).After(
			firstCall,
		),
	)
	logger.WithError(err1).Info(nil, "info line")
	logger.WithError(err2).Warn(nil, "warn line")
	controller.Finish()
}

func (s *wrapperSuite) TestConfiguration() {
	builder := naive.Builder().SetLevel(logInt.InfoLevel)
	logger := CreateBricksLogger(builder)
	s.Equal(logInt.InfoLevel, logger.Configuration().Level())
}

func (s *wrapperSuite) TestContextExtractors() {
	contextExtractor := func(ctx context.Context) map[string]interface{} {
		s.Require().NotNil(ctx)
		return map[string]interface{}{
			"one": "two",
		}
	}
	controller := gomock.NewController(s.T())
	mockLogger := mock_log.NewMockLogger(controller)
	mockBuilder := mock_log.NewMockBuilder(controller)
	mockBuilder.EXPECT().Build().Return(mockLogger).After(
		mockBuilder.EXPECT().IncrementSkipFrames(skipFramesExpect).Return(mockBuilder),
	)
	mockLogger.EXPECT().Custom(gomock.Not((context.Context)(nil)), logInt.InfoLevel, 0, "info line").After(
		mockLogger.EXPECT().WithField("one", "three").Return(mockLogger).After(
			mockLogger.EXPECT().WithField("one", "two").Return(mockLogger),
		),
	).After(
		mockLogger.EXPECT().WithError(gomock.Not((error)(nil))).Return(mockLogger),
	)

	logger := CreateBricksLogger(mockBuilder, contextExtractor)
	logger.WithError(fmt.Errorf("an error")).WithField("one", "three").Info(nil, "info line")
	controller.Finish()
}

func (s *wrapperSuite) TestPanicExtractor() {
	contextExtractor := func(ctx context.Context) map[string]interface{} {
		panic("oops")
	}
	controller := gomock.NewController(s.T())
	mockLogger := mock_log.NewMockLogger(controller)
	mockBuilder := mock_log.NewMockBuilder(controller)
	mockBuilder.EXPECT().Build().Return(mockLogger).After(
		mockBuilder.EXPECT().IncrementSkipFrames(skipFramesExpect).Return(mockBuilder),
	)
	mockLogger.EXPECT().Custom(gomock.Not((context.Context)(nil)), logInt.InfoLevel, 1, "info line").After(
		mockLogger.EXPECT().Error(gomock.Not((context.Context)(nil)), "one of the context extractors panicked").After(
			mockLogger.EXPECT().WithField("__panic__", "oops").Return(mockLogger),
		),
	)

	logger := CreateBricksLogger(mockBuilder, contextExtractor)
	logger.Info(nil, "info line")
	controller.Finish()
}
