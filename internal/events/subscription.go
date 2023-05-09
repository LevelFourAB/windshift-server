package events

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type SubscriptionConfig struct {
	Name     string
	Stream   string
	Subjects []string

	Timeout time.Duration

	MaxDeliveryAttempts uint

	DeliverFromID    uint64
	DeliverFromTime  time.Time
	DeliverFromFirst bool
}

type Subscription struct {
	ID string
}

func EnsureSubscription(_ context.Context, js nats.JetStreamContext, config *SubscriptionConfig) (*Subscription, error) {
	if len(config.Subjects) != 1 {
		return nil, errors.New("only one subject can be specified")
	}

	var name string
	var err error
	if config.Name == "" {
		// If the name is not specified, we create an ephemeral consumer
		name, err = declareEphemeralConsumer(js, config)
		if err != nil {
			return nil, err
		}
	} else {
		name, err = declareDurableConsumer(js, config)
		if err != nil {
			return nil, err
		}
	}

	return &Subscription{
		ID: name,
	}, nil
}

func declareEphemeralConsumer(js nats.JetStreamContext, config *SubscriptionConfig) (string, error) {
	consumerConfig := &nats.ConsumerConfig{
		Name:              uuid.New().String(),
		InactiveThreshold: 1 * time.Hour,
	}

	setConsumerSettings(consumerConfig, config, false)

	_, err := js.AddConsumer(config.Stream, consumerConfig)
	if err != nil {
		return "", errors.Wrap(err, "could not create consumer")
	}
	return consumerConfig.Name, nil
}

func declareDurableConsumer(js nats.JetStreamContext, config *SubscriptionConfig) (string, error) {
	ci, err := js.ConsumerInfo(config.Stream, config.Name)
	if err != nil {
		if errors.Is(err, nats.ErrConsumerNotFound) {
			// Consumer does not exist, create it
			consumerConfig := &nats.ConsumerConfig{
				Durable:           config.Name,
				InactiveThreshold: 30 * 24 * time.Hour,
			}

			setConsumerSettings(consumerConfig, config, false)

			_, err = js.AddConsumer(config.Stream, consumerConfig)
			if err != nil {
				return "", errors.Wrap(err, "could not create consumer")
			}
			return config.Name, nil
		}

		return "", errors.Wrap(err, "could not get consumer info")
	}

	// For updates certain fields can not be set, so we only set what we can
	consumerConfig := ci.Config
	setConsumerSettings(&consumerConfig, config, true)

	// Perform the update
	_, err = js.UpdateConsumer(config.Stream, &consumerConfig)
	if err != nil {
		return "", errors.Wrap(err, "could not update consumer")
	}
	return config.Name, nil
}

// setConsumerSettings sets the shared settings for both ephemeral and durable
// consumers.
func setConsumerSettings(c *nats.ConsumerConfig, qc *SubscriptionConfig, update bool) {
	c.AckPolicy = nats.AckExplicitPolicy
	// TODO: With NATS 2.10 multiple subjects can be specified
	c.FilterSubject = qc.Subjects[0]

	// If a timeout is specified set it or use the default
	if qc.Timeout > 0 {
		c.AckWait = qc.Timeout
	} else {
		c.AckWait = 30 * time.Second
	}

	// If the max delivery attempts is specified set it
	if qc.MaxDeliveryAttempts > 0 {
		c.MaxDeliver = int(qc.MaxDeliveryAttempts)
	}

	if !update {
		// When creating a consumer we can specify where to start from
		if !qc.DeliverFromTime.IsZero() {
			c.DeliverPolicy = nats.DeliverByStartTimePolicy
			c.OptStartTime = &qc.DeliverFromTime
		} else if qc.DeliverFromID > 0 {
			c.DeliverPolicy = nats.DeliverByStartSequencePolicy
			c.OptStartSeq = qc.DeliverFromID
		} else if qc.DeliverFromFirst {
			c.DeliverPolicy = nats.DeliverAllPolicy
		} else {
			c.DeliverPolicy = nats.DeliverNewPolicy
		}
	}
}
