package meta

import (
	"context"

	"gitlab.id.vin/platform/gopkgs/internal/trcontext"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	keyKafkaMeta = "kafka_meta"
	keyTopic     = "topic"
	keyPartition = "partition"
	keyOffset    = "offset"
	keyEvent     = "event"
)

type KafkaMeta struct {
	Topic     string
	Partition int32
	Offset    int64
	Event     string
}

func (m *KafkaMeta) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	if m == nil {
		return nil
	}
	encoder.AddString(keyTopic, m.Topic)
	encoder.AddInt32(keyPartition, m.Partition)
	encoder.AddInt64(keyOffset, m.Offset)
	if len(m.Event) > 0 {
		encoder.AddString(keyEvent, m.Event)
	}
	return nil
}

func ContextWithKafkaMeta(ctx context.Context, r *KafkaMeta) context.Context {
	return trcontext.WithKafkaMeta(ctx, r)
}

func KafkaMetaFromContext(ctx context.Context) *KafkaMeta {
	if ctx == nil {
		return nil
	}
	r, _ := trcontext.KafkaMetaFromContext(ctx).(*KafkaMeta)
	return r
}

func ExtractKafkaMetaZapFields(ctx context.Context) []zap.Field {
	if ctx == nil {
		return nil
	}
	kafkaMeta := KafkaMetaFromContext(ctx)
	if kafkaMeta == nil {
		return nil
	}
	return []zap.Field{zap.Object(keyKafkaMeta, kafkaMeta)}
}
