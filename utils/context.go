package utils

import (
	"context"
	"encoding/json"

	"go.opentelemetry.io/otel/trace"
)

func GetStrFromContext(ctx context.Context, key string) (string, bool) {
	value := ctx.Value(key)
	if value == nil {
		return "", false
	}
	strValue, ok := value.(string)
	return strValue, ok
}

func GetInt64FromContext(ctx context.Context, key string) (int64, bool) {
	value := ctx.Value(key)
	if value == nil {
		return 0, false
	}
	i64Value, err := value.(json.Number).Int64()
	return i64Value, err == nil
}

func NewContextCopySpanAndTrace(ctx context.Context) context.Context {
	newCtx := context.Background()
	span := trace.SpanFromContext(ctx)
	return trace.ContextWithSpan(newCtx, span)
}
