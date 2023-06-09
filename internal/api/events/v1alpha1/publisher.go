package v1alpha1

import (
	"context"
	"time"

	"windshift/service/internal/events"
	eventsv1alpha1 "windshift/service/internal/proto/windshift/events/v1alpha1"
)

func (e *EventsServiceServer) PublishEvent(ctx context.Context, req *eventsv1alpha1.PublishEventRequest) (*eventsv1alpha1.PublishEventResponse, error) {
	now := time.Now()
	config := &events.PublishConfig{
		Subject:       req.Subject,
		Data:          req.Data,
		PublishedTime: &now,
	}

	if req.Timestamp != nil {
		publishedAt := req.Timestamp.AsTime()
		config.PublishedTime = &publishedAt
	}

	if req.IdempotencyKey != nil {
		config.IdempotencyKey = *req.IdempotencyKey
	}

	if req.ExpectedLastId != nil {
		config.ExpectedSubjectSeq = req.ExpectedLastId
	}

	ack, err := e.events.Publish(ctx, config)
	if err != nil {
		return nil, err
	}
	return &eventsv1alpha1.PublishEventResponse{
		Id: ack.ID,
	}, nil
}
