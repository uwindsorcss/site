package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/uwindsorcss/site/pkg/config"
)

func ProvideLog(c *config.Config) (*zap.SugaredLogger, error) {
	atom := zap.NewAtomicLevel()
	atom.SetLevel(zap.DebugLevel)

	var zLog *zap.Logger

	// change between prod and dev loggers
	if c.Mode == "prod" {
		zLog = zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.Lock(os.Stdout),
			atom,
		))
	} else {
		zLog = zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.Lock(os.Stdout),
			atom,
		))
	}

	if !c.Debug {
		atom.SetLevel(zap.InfoLevel)
	}

	return zLog.Sugar(), nil
}
