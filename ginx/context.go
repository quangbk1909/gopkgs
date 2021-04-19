package ginx

import (
	"context"
	"gitlab.id.vin/platform/gopkgs/meta"

	"github.com/gin-gonic/gin"
)

func RequestContextFromGinCtx(ginCtx *gin.Context) context.Context {
	ctx := ginCtx.Request.Context()
	ctx = ContextWithRequestMeta(ctx, ginCtx)
	return ctx
}

func ContextWithRequestMeta(ctx context.Context, ginCtx *gin.Context) context.Context {
	reqMeta := getRequestMeta(ginCtx)
	return meta.ContextWithRequestMeta(ctx, &reqMeta)
}

func getRequestMeta(ginCtx *gin.Context) meta.RequestMeta {
	return meta.RequestMeta{
		URI:             ginCtx.Request.RequestURI,
		Method:          ginCtx.Request.Method,
		RequestID:       ginCtx.Request.Header.Get(headerRequestID),
		DeviceID:        ginCtx.Request.Header.Get(headerDeviceID),
		RFDeviceID:      ginCtx.Request.Header.Get(headerRFDeviceID),
		DeviceSessionID: ginCtx.Request.Header.Get(headerDeviceSessionID),
		DeviceRequestID: ginCtx.Request.Header.Get(headerDeviceRequestID),
	}
}
