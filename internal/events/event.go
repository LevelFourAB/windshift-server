package events

import (
	"context"
	"time"

	"github.com/levelfourab/windshift-server/internal/events/flowcontrol"

	"github.com/cockroachdb/errors"
	"github.com/nats-io/nats.go/jetstream"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"
)

// Headers contains information about an event.
type Headers struct {
	// PublishedAt is the time the event was published by the producer.
	PublishedAt time.Time
	// IdempotencyKey is the idempotency key of the event. An idempotency key
	// is used to ensure that an event is only published once. May be empty.
	IdempotencyKey *string
	// TraceParent is the trace parent of the event. May be empty.
	TraceParent *string
	// TraceState is the trace state of the event. May be empty.
	TraceState *string
}

// Event represents a single event consumed from a stream. It is received via
// NATS and should be processed within a certain deadline, using Ack() or
// Reject(shouldRetry) to acknowledge the event. If the deadline is exceeded,
// the event will be redelivered. To extend the deadline, use Ping().
type Event struct {
	span      trace.Span
	logger    *zap.Logger
	msg       jetstream.Msg
	onProcess func(flowcontrol.ProcessType)

	// Context is the context of this event. It will be valid until the event
	// expires, is acknowledged or rejected.
	Context context.Context

	// Subject is the subject the event was published to.
	Subject string

	// ConsumerSeq is the sequence number of the event.
	ConsumerSeq uint64

	// StreamSeq is the sequence number of the event in the event stream. Can
	// be used for resuming from a certain point in time. For example with an
	// ephemeral consumer, the consumer can store the last seen StreamSeq and
	// resume from there on the next run.
	StreamSeq uint64

	// DeliveryAttempt is the number of times the event has been delivered to
	// a consumer. The first delivery is 1.
	DeliveryAttempt uint64

	// Headers contains the headers of the event.
	Headers *Headers

	// Data is the protobuf message published by the producer.
	Data *anypb.Any
}

func newEvent(
	ctx context.Context,
	span trace.Span,
	logger *zap.Logger,
	msg jetstream.Msg,
	md *jetstream.MsgMetadata,
	onProcess func(flowcontrol.ProcessType),
) (*Event, error) {
	headers := &Headers{
		PublishedAt: md.Timestamp,
	}

	natsHeaders := msg.Headers()

	// Get the published header
	publishTimeHeader := natsHeaders.Get("WS-Published-Time")
	if publishTimeHeader != "" {
		publishedTime, err := time.Parse(time.RFC3339Nano, publishTimeHeader)
		if err != nil {
			return nil, errors.Wrap(err, "could not parse header")
		}

		headers.PublishedAt = publishedTime
	}

	// Get the idempotency key header
	idempotencyKeyHeader := natsHeaders.Get("Nats-Msg-Id")
	if idempotencyKeyHeader != "" {
		headers.IdempotencyKey = &idempotencyKeyHeader
	}

	// Get the trace parent header
	traceParentHeader := natsHeaders.Get("WS-Trace-Parent")
	if traceParentHeader != "" {
		headers.TraceParent = &traceParentHeader
	}

	// Get the trace state header
	traceStateHeader := natsHeaders.Get("WS-Trace-State")
	if traceStateHeader != "" {
		headers.TraceState = &traceStateHeader
	}

	data := &anypb.Any{
		TypeUrl: "type.googleapis.com/" + natsHeaders.Get("WS-Data-Type"),
		Value:   msg.Data(),
	}

	return &Event{
		span:            span,
		logger:          logger,
		msg:             msg,
		onProcess:       onProcess,
		Context:         ctx,
		Subject:         msg.Subject(),
		ConsumerSeq:     md.Sequence.Stream,
		StreamSeq:       md.Sequence.Consumer,
		DeliveryAttempt: md.NumDelivered,
		Headers:         headers,
		Data:            data,
	}, nil
}

// DiscardData discards the data of the event. This should be called if the
// event data is not needed anymore. Acknowledging or rejecting the event will
// continue working after this.
func (e *Event) DiscardData() {
	e.Data = nil
}

// Ping extends the deadline of the event. This should be called periodically
// to prevent the event from being redelivered.
func (e *Event) Ping() error {
	e.logger.Debug("Pinging event", zap.Uint64("streamSeq", e.StreamSeq))
	err := e.msg.InProgress()
	if err != nil {
		e.span.RecordError(err)
		return errors.Wrap(err, "could not ping message")
	}
	e.span.AddEvent("pinged")
	e.onProcess(flowcontrol.ProcessTypePing)
	return nil
}

// Ack acknowledges the event. The event will be removed from the consumer.
func (e *Event) Ack() error {
	defer e.span.End()

	e.logger.Debug("Acknowledging event", zap.Uint64("streamSeq", e.StreamSeq))
	err := e.msg.Ack()
	if err != nil {
		e.span.RecordError(err)
		return errors.Wrap(err, "could not acknowledge message")
	}

	e.span.SetStatus(codes.Ok, "")
	e.onProcess(flowcontrol.ProcessTypeAck)
	return nil
}

// Reject rejects the event.
func (e *Event) Reject() error {
	defer e.span.End()

	// The event should be redelivered if possible
	e.logger.Debug("Rejecting event", zap.Uint64("streamSeq", e.StreamSeq))
	err := e.msg.Nak()
	if err != nil {
		e.span.RecordError(err)
		return errors.Wrap(err, "could not reject message")
	}

	e.span.SetStatus(codes.Error, "event rejected")
	e.onProcess(flowcontrol.ProcessTypeReject)
	return nil
}

// RejectWithDelay rejects the event with a delay. The event will be redelivered
// after the delay.
func (e *Event) RejectWithDelay(delay time.Duration) error {
	defer e.span.End()

	// The event should be redelivered if possible
	e.logger.Debug("Rejecting event with delay", zap.Uint64("streamSeq", e.StreamSeq), zap.Duration("delay", delay))
	err := e.msg.NakWithDelay(delay)
	if err != nil {
		e.span.RecordError(err)
		return errors.Wrap(err, "could not reject message")
	}

	e.span.SetStatus(codes.Error, "event rejected")
	e.onProcess(flowcontrol.ProcessTypeReject)
	return nil
}

// RejectPermanently permanently rejects the event. The event will be removed
// and no redelivery will be attempted.
func (e *Event) RejectPermanently() error {
	defer e.span.End()

	// This is a permanent rejection, terminate the event
	e.logger.Debug("Permanently rejecting event", zap.Uint64("streamSeq", e.StreamSeq))
	err := e.msg.Term()
	if err != nil {
		e.span.RecordError(err)
		return errors.Wrap(err, "could not permanently reject message")
	}

	e.span.SetStatus(codes.Error, "event permanently rejected")
	e.onProcess(flowcontrol.ProcessTypePermanentReject)
	return nil
}
