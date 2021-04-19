package zapx

import (
	"context"
	"errors"
	"os"

	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
)

var (
	zapStackTraceLogger *zap.Logger
)

func init() {
	zapCfg := zap.NewProductionConfig()
	zapCfg.DisableStacktrace = true
	if !isEnableSampling() {
		zapCfg.Sampling = nil
	}
	l, err := zapCfg.Build(
		zap.WrapCore((&apmzap.Core{}).WrapCore),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(l)
	zapStackTraceLogger = zap.L().WithOptions(zap.AddStacktrace(zap.DebugLevel))
}

func ReplaceNop() {
	zap.ReplaceGlobals(zap.NewNop())
	zapStackTraceLogger = zap.NewNop()
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = buildFieldsWithContext(ctx, fields)
	logger(ctx, nil).Info(msg, fields...)
}

func Error(ctx context.Context, msg string, err error, fields ...zap.Field) {
	fields = buildFieldsWithContext(ctx, fields)
	fields = buildFieldsWithError(fields, err)
	logger(ctx, err).Error(msg, fields...)
}

func Fatal(ctx context.Context, msg string, err error, fields ...zap.Field) {
	fields = buildFieldsWithContext(ctx, fields)
	fields = buildFieldsWithError(fields, err)
	logger(ctx, err).Fatal(msg, fields...)
}

func Panic(ctx context.Context, msg string, err error, fields ...zap.Field) {
	fields = buildFieldsWithContext(ctx, fields)
	fields = buildFieldsWithError(fields, err)
	logger(ctx, err).Panic(msg, fields...)
}

func logger(ctx context.Context, err error) *zap.Logger {
	traceFields := apmzap.TraceContext(ctx)
	if err == nil {
		return zap.L().With(traceFields...)
	}
	var d debugger
	if !errors.As(err, &d) {
		return zapStackTraceLogger.With(traceFields...)
	}
	return zap.L().With(traceFields...)
}

func buildFieldsWithContext(ctx context.Context, fields []zap.Field) []zap.Field {
	if ctx == nil {
		return fields
	}
	for _, e := range fieldExtractors {
		fields = append(e.ExtractZapFields(ctx), fields...)
	}
	return fields
}

func buildFieldsWithError(fields []zap.Field, err error) []zap.Field {
	if err == nil {
		return fields
	}
	var d debugger
	if errors.As(err, &d) {
		fields = append(fields, zap.String("root_caller", d.Caller()), zap.String("root_stack", d.StackTrace()))
	}
	fields = append(fields, zap.Error(err))
	return fields
}

func isEnableSampling() bool {
	return os.Getenv("ZAP_X_ENABLE_SAMPLING") == "true"
}
