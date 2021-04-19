package trcontext

import "context"

type keyUserMeta struct{}
type keyRequestMeta struct{}
type keyAppMeta struct{}
type keyKafkaMeta struct{}

func WithUserMeta(ctx context.Context, userMeta interface{}) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, keyUserMeta{}, userMeta)
}

func UserMetaFromContext(ctx context.Context) interface{} {
	if ctx == nil {
		return nil
	}
	return ctx.Value(keyUserMeta{})
}

func WithRequestMeta(ctx context.Context, requestMeta interface{}) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, keyRequestMeta{}, requestMeta)
}

func RequestMetaFromContext(ctx context.Context) interface{} {
	if ctx == nil {
		return nil
	}
	return ctx.Value(keyRequestMeta{})
}

func WithAppMeta(ctx context.Context, appMeta interface{}) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, keyAppMeta{}, appMeta)
}

func AppMetaFromContext(ctx context.Context) interface{} {
	if ctx == nil {
		return nil
	}
	return ctx.Value(keyAppMeta{})
}

func WithKafkaMeta(ctx context.Context, kafkaMeta interface{}) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, keyKafkaMeta{}, kafkaMeta)
}

func KafkaMetaFromContext(ctx context.Context) interface{} {
	if ctx == nil {
		return nil
	}
	return ctx.Value(keyKafkaMeta{})
}
