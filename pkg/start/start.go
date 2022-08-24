package start

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/uwindsorcss/site/pkg/config"
)

func Run() {
	// init logger with dev and debug
	atom := zap.NewAtomicLevel()
	atom.SetLevel(zap.DebugLevel)

	k, err := config.Load("config.hcl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "start.Run: %s", err)
		return
	}

	c, err := config.UnmarshalConfig(k)
	if err != nil {
		fmt.Fprintf(os.Stderr, "start.Run: %s", err)
		return
	}

	var zLog *zap.Logger
	var log *zap.SugaredLogger

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
	defer zLog.Sync()

	log = zLog.Sugar()

	// turn debug off
	if !c.Debug {
		atom.SetLevel(zap.InfoLevel)
	}

	log.Debugf("%#v", c)
}
