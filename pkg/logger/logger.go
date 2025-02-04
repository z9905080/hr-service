package logger

import (
	"context"
	"github.com/z9905080/hr_service/environment"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"
	"os"
)

type logger struct {
	zapSugar *zap.SugaredLogger
	zap      *zap.Logger
}

func (l *logger) GetSugarLogger(ctx context.Context) *zap.SugaredLogger {
	if ctx != nil {
		span := trace.SpanFromContext(ctx)
		if span != nil {
			return l.zapSugar.With(
				zap.String("trace_id", span.SpanContext().TraceID().String()),
				zap.String("span_id", span.SpanContext().SpanID().String()),
			)
		}
	}
	return l.zapSugar
}
func (l *logger) InfoF(ctx context.Context, format string, args ...interface{}) {
	l.GetSugarLogger(ctx).Infof(format, args...)
}

func (l *logger) ErrorF(ctx context.Context, format string, args ...interface{}) {
	l.GetSugarLogger(ctx).Errorf(format, args...)
}

func (l *logger) WarnF(ctx context.Context, format string, args ...interface{}) {
	l.GetSugarLogger(ctx).Warnf(format, args...)
}

func (l *logger) DebugF(ctx context.Context, format string, args ...interface{}) {
	l.GetSugarLogger(ctx).Debugf(format)
}

func (l *logger) Info(ctx context.Context, args ...interface{}) {
	l.GetSugarLogger(ctx).Info(args...)
}

func (l *logger) Error(ctx context.Context, args ...interface{}) {
	l.GetSugarLogger(ctx).Error(args...)
}

func (l *logger) Warn(ctx context.Context, args ...interface{}) {
	l.GetSugarLogger(ctx).Warn(args...)
}

func (l *logger) Debug(ctx context.Context, args ...interface{}) {
	l.GetSugarLogger(ctx).Debug(args...)
}

func NewLogger(config environment.Config) InfLogger {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		NameKey:      "logger",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
	})

	level := config.LogLevel
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(
			encoder,
			zapcore.Lock(os.Stdout),
			zapLevel,
		),
	)

	s := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &logger{
		zap:      s,
		zapSugar: s.Sugar(),
	}
}
