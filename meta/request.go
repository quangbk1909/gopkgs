package meta

import (
	"context"

	"gitlab.id.vin/platform/gopkgs/internal/trcontext"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	keyRequestMeta     = "request_meta"
	keyURI             = "uri"
	keyMethod          = "method"
	keyRequestID       = "request_id"
	keyDeviceID        = "device_id"
	keyRFDeviceID      = "rf_device_id"
	keyDeviceSessionID = "device_session_id"
	keyDeviceRequestID = "device_request_id"
)

type RequestMeta struct {
	URI             string
	Method          string
	RequestID       string
	DeviceID        string
	RFDeviceID      string
	DeviceSessionID string
	DeviceRequestID string
}

func (m *RequestMeta) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	if m == nil {
		return nil
	}
	encoder.AddString(keyURI, m.URI)
	encoder.AddString(keyMethod, m.Method)
	encoder.AddString(keyRequestID, m.RequestID)
	encoder.AddString(keyDeviceID, m.DeviceID)
	encoder.AddString(keyRFDeviceID, m.RFDeviceID)
	encoder.AddString(keyDeviceSessionID, m.DeviceSessionID)
	encoder.AddString(keyDeviceRequestID, m.DeviceRequestID)
	return nil
}

func ContextWithRequestMeta(ctx context.Context, r *RequestMeta) context.Context {
	return trcontext.WithRequestMeta(ctx, r)
}

func RequestMetaFromContext(ctx context.Context) *RequestMeta {
	if ctx == nil {
		return nil
	}
	r, _ := trcontext.RequestMetaFromContext(ctx).(*RequestMeta)
	return r
}

func ExtractRequestMetaZapFields(ctx context.Context) []zap.Field {
	if ctx == nil {
		return nil
	}
	requestMeta := RequestMetaFromContext(ctx)
	if requestMeta == nil {
		return nil
	}
	return []zap.Field{zap.Object(keyRequestMeta, requestMeta)}
}
