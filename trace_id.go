package logger

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

const traceKey = "trace_id"

// 生成带trace_id 的context
func CtxWithTraceId() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, traceKey, newTraceId())
	return ctx
}

func newTraceId() string {
	uuid := uuid.NewV4()
	// 缩短id长度
	return uuid.String()[24:]
}

// 获取 trace_id
func GetTraceId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	value := ctx.Value(traceKey)
	if value == nil {
		return ""
	}
	traceId := value.(string)
	return traceId
}
