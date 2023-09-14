package constructors

import (
	"context"
	"log"

	"github.com/go-bricks/bricks/v2/logger/naive"

	"github.com/go-bricks/bricks/v2/interfaces/cfg"
	confkeys "github.com/go-bricks/bricks/v2/interfaces/cfg/keys"
	logInt "github.com/go-bricks/bricks/v2/interfaces/log"

	"go.uber.org/fx"
)

// FxGroupLoggerContextExtractors defines group name
const FxGroupLoggerContextExtractors = "loggerContextExtractors"
const (
	application = "app"
	hostname    = "host"
	gitCommit   = "git"
)
const compensateDefaultLogger = 1

type loggerDeps struct {
	fx.In

	Config            cfg.Config
	LoggerBuilder     logInt.Builder            `optional:"true"`
	ContextExtractors []logInt.ContextExtractor `group:"loggerContextExtractors"`
}

// DefaultLogger is a constructor that will create a logger with some default values on top of provided ones
func DefaultLogger(deps loggerDeps) logInt.Logger {
	var logLevel = logInt.InfoLevel
	if levelValue := deps.Config.Get(confkeys.LogLevel); levelValue.IsSet() {
		logLevel = logInt.ParseLevel(levelValue.String())
	}

	builder := deps.getLogBuilder().SetLevel(logLevel).IncrementSkipFrames(compensateDefaultLogger)
	return logger.CreateBricksLogger(builder, append(deps.ContextExtractors, deps.selfStaticFieldsContextExtractor)...)
}

func (d loggerDeps) selfStaticFieldsContextExtractor(_ context.Context) map[string]interface{} {
	output := make(map[string]interface{})
	info := bricks.GetBuildInformation()
	appName := d.Config.Get(confkeys.ApplicationName).String()
	if len(appName) > 0 && d.Config.Get(confkeys.LogIncludeName).Bool() {
		output[application] = appName
	}
	if len(info.Hostname) > 0 && d.Config.Get(confkeys.LogIncludeHost).Bool() {
		output[hostname] = info.Hostname
	}
	if len(info.GitCommit) > 0 && d.Config.Get(confkeys.LogIncludeGitSHA).Bool() {
		output[gitCommit] = info.GitCommit
	}
	return output
}

func (d loggerDeps) getLogBuilder() logInt.Builder {
	if d.LoggerBuilder != nil {
		return d.LoggerBuilder
	}
	log.Printf("[Bricks] WARNING \tNo logger builder provided, using default logger. Some features will not be supported")
	return naive.Builder()
}
