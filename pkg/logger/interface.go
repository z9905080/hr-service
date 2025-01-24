package logger

import "context"

type InfLogger interface {
	InfoF(ctx context.Context, format string, args ...interface{})
	ErrorF(ctx context.Context, format string, args ...interface{})
	WarnF(ctx context.Context, format string, args ...interface{})
	DebugF(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Debug(ctx context.Context, args ...interface{})
}
