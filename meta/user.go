package meta

import (
	"context"

	"gitlab.id.vin/platform/gopkgs/internal/trcontext"

	"go.uber.org/zap"
)

const (
	keyUserID     = "user_id"
	keyUserIDType = "user_id_type"
)

type UserMeta struct {
	UserID     string
	UserIDType string
}

func ContextWithUserMeta(ctx context.Context, r *UserMeta) context.Context {
	return trcontext.WithUserMeta(ctx, r)
}

func UserMetaFromContext(ctx context.Context) *UserMeta {
	if ctx == nil {
		return nil
	}
	r, _ := trcontext.UserMetaFromContext(ctx).(*UserMeta)
	return r
}

func ExtractUserMetaZapFields(ctx context.Context) []zap.Field {
	userMeta := UserMetaFromContext(ctx)
	if userMeta == nil {
		return nil
	}
	return []zap.Field{
		zap.String(keyUserID, userMeta.UserID),
		zap.String(keyUserIDType, userMeta.UserIDType),
	}
}
