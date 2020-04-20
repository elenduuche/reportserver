package logging

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Logger implements logging function for the various log level
type Logger interface {
	Info(string, ...zapcore.Field)
	// Infof(string, ...interface{}) string
	// Debug(string)
	// Debugf(string, ...interface{}) string
	// Error(string)
	// Errorf(string, ...interface{}) string
	// Warn(string)
	// Warnf(string, ...interface{}) string
}

type zapLogger struct {
	logger *zap.Logger
}

func (log *zapLogger) Info(msg string, f ...zapcore.Field) {
	if len(f) > 0 {
		log.logger.Info(msg, f...)
		return
	}
	log.logger.Info(msg)
}

//NewLogger Returns a new instance of logger implemented using uber zap library
func NewLogger() Logger {
	// cfg := zap.Config{
	// 	Encoding:         "json",
	// 	Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
	// 	OutputPaths:      []string{"stderr"},
	// 	ErrorOutputPaths: []string{"stderr"},
	// 	EncoderConfig: zapcore.EncoderConfig{
	// 		MessageKey: "message",

	// 		LevelKey:    "level",
	// 		EncodeLevel: zapcore.CapitalLevelEncoder,

	// 		TimeKey:    "time",
	// 		EncodeTime: zapcore.ISO8601TimeEncoder,

	// 		CallerKey:    "caller",
	// 		EncodeCaller: zapcore.ShortCallerEncoder,
	// 	},
	// }
	//logger, err := cfg.Build()
	//logger, err := zap.NewProduction(zap.Development(), zap.AddCallerSkip(2))
	cfg := zap.NewProductionConfig()
	//logPath := os.Getenv("LOGPATH")
	//cfg.OutputPaths = []string{logPath, "stderr"}
	cfg.OutputPaths = []string{"stderr"}
	//cfg.Development = true
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	logger = logger.Named("unaryserver-app")
	zLogger := new(zapLogger)
	zLogger.logger = logger
	return zLogger
}
