package logger

import (
	"context"
	"runtime/debug"
)

// ------------------- 以下为 key value 格式输入 ----------------------------
func Debugw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	traceId := GetTraceId(ctx)
	keysAndValues = append([]interface{}{traceKey, traceId}, keysAndValues...)
	sugarLogger.Debugw(msg, keysAndValues...)
}

func Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	traceId := GetTraceId(ctx)
	keysAndValues = append([]interface{}{traceKey, traceId}, keysAndValues...)
	sugarLogger.Infow(msg, keysAndValues...)

}

func Warnw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	traceId := GetTraceId(ctx)
	keysAndValues = append([]interface{}{traceKey, traceId}, keysAndValues...)
	sugarLogger.Warnw(msg, keysAndValues...)

}

func Errorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	traceId := GetTraceId(ctx)
	keysAndValues = append([]interface{}{traceKey, traceId}, keysAndValues...)
	sugarLogger.Errorw(msg, keysAndValues...)

}

func Fatalw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	traceId := GetTraceId(ctx)
	keysAndValues = append([]interface{}{traceKey, traceId}, keysAndValues...)
	sugarLogger.Fatalw(msg, keysAndValues...)
}

func ErrorwWithTrace(ctx context.Context, msg string, keysAndValues ...interface{}) {
	traceId := GetTraceId(ctx)
	stackInfo := string(debug.Stack())
	keysAndValues = append([]interface{}{traceKey, traceId, "stack", stackInfo}, keysAndValues...)
	sugarLogger.Errorw(msg, keysAndValues...)
}

// ------------------- 以下为 format string 格式输入 ----------------------------
func Debug(ctx context.Context, tag string, args interface{}, v ...interface{}) {
	msg := converArgs(args, v...)
	traceId := GetTraceId(ctx)
	sugarLogger.Debugw(msg, "tag", tag, traceKey, traceId)
}

func Info(ctx context.Context, tag string, args interface{}, v ...interface{}) {
	msg := converArgs(args, v...)
	traceId := GetTraceId(ctx)
	sugarLogger.Infow(msg, "tag", tag, traceKey, traceId)
}

func Warning(ctx context.Context, tag string, args interface{}, v ...interface{}) {
	msg := converArgs(args, v...)
	traceId := GetTraceId(ctx)
	sugarLogger.Warnw(msg, "tag", tag, traceKey, traceId)
}

func Error(ctx context.Context, tag string, args interface{}, v ...interface{}) {
	msg := converArgs(args, v...)
	traceId := GetTraceId(ctx)
	sugarLogger.Errorw(msg, "tag", tag, traceKey, traceId)
}

func Fatal(ctx context.Context, tag string, args interface{}, v ...interface{}) {
	msg := converArgs(args, v...)
	traceId := GetTraceId(ctx)
	sugarLogger.Fatalw(msg, "tag", tag, traceKey, traceId)
}

func ErrorWithTrace(ctx context.Context, tag string, args interface{}, v ...interface{}) {
	msg := converArgs(args, v...)
	traceId := GetTraceId(ctx)
	stackInfo := string(debug.Stack())
	sugarLogger.Errorw(msg, "tag", tag, traceKey, traceId, "stack", stackInfo)
}
