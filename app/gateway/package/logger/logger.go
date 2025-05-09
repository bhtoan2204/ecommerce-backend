package logger

import (
	"context"
	"gateway/package/settings"
	"os"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey string

const loggerKey = contextKey("logger")

type LoggerZap struct {
	*otelzap.Logger
}

func (l *LoggerZap) WithFields(fields ...zapcore.Field) *LoggerZap {
	clone := l.Logger.Clone(otelzap.WithExtraFields(fields...))
	return &LoggerZap{Logger: clone}
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encodeConfig.TimeKey = "timestamp"
	return zapcore.NewJSONEncoder(encodeConfig)
}

func NewLogger(config settings.LogConfig) *LoggerZap {
	logLevel := config.LogLevel
	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "dpanic":
		level = zapcore.DPanicLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.FilePath,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&hook),
		),
		level,
	)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	otelLogger := otelzap.New(zapLogger)
	return &LoggerZap{Logger: otelLogger}
}

func WithLogger(ctx context.Context, logger *LoggerZap) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func DefaultLogger() *LoggerZap {
	defaultConfig := settings.LogConfig{
		LogLevel:   "info",
		FilePath:   "app.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return NewLogger(defaultConfig)
}

func FromContext(ctx context.Context) *LoggerZap {
	if logger, ok := ctx.Value(loggerKey).(*LoggerZap); ok {
		return logger
	}
	return DefaultLogger()
}
