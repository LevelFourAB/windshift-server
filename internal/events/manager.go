package events

import (
	"github.com/cockroachdb/errors"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Manager struct {
	logger        *zap.Logger
	tracer        trace.Tracer
	w3cPropagator propagation.TextMapPropagator

	jetStream nats.JetStreamContext
}

func NewManager(
	logger *zap.Logger,
	tracer trace.Tracer,
	conn *nats.Conn,
) (*Manager, error) {
	js, err := conn.JetStream()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create JetStream channel")
	}

	m := &Manager{
		logger:        logger,
		tracer:        tracer,
		w3cPropagator: propagation.TraceContext{},
		jetStream:     js,
	}

	return m, nil
}
