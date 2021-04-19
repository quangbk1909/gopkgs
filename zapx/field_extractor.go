package zapx

import (
	"context"

	"go.uber.org/zap"
)

// RegisterFieldExtractors
// RegisterFieldExtractors is not thread-safe
func RegisterFieldExtractors(e ...fieldExtractor) {
	fieldExtractors = append(fieldExtractors, e...)
}

type fieldExtractor interface {
	ExtractZapFields(ctx context.Context) []zap.Field
}

var fieldExtractors []fieldExtractor

type ExtractFunc func(ctx context.Context) []zap.Field

func (f ExtractFunc) ExtractZapFields(ctx context.Context) []zap.Field {
	return f(ctx)
}

type debugger interface {
	StackTrace() string
	Caller() string
}
