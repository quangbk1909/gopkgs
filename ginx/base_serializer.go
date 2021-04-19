package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseSerializer struct {
}

func (s *BaseSerializer) Health(ctx *gin.Context) {
	BuildSuccessResponse(ctx, http.StatusOK, nil)
}
