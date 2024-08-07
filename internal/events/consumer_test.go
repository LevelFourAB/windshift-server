package events_test

import (
	"context"
	"time"

	"github.com/levelfourab/windshift-server/internal/events"

	"github.com/cockroachdb/errors"
	"github.com/nats-io/nats.go/jetstream"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap/zaptest"
)

var _ = Describe("Consumers", func() {
	var manager *events.Manager
	var js jetstream.JetStream

	BeforeEach(func() {
		var err error
		natsConn := GetNATS()
		js, err = jetstream.New(natsConn)
		Expect(err).ToNot(HaveOccurred())

		manager, err = events.NewManager(
			zaptest.NewLogger(GinkgoT()),
			otel.Tracer("tests"),
			js,
		)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("Configuration issues", func() {
		It("consumer with no stream fails", func(ctx context.Context) {
			_, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Subjects: []string{"test"},
			})
			Expect(err).To(HaveOccurred())
		})

		It("consumer with invalid stream name fails", func(ctx context.Context) {
			_, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream:   "invalid name",
				Subjects: []string{"test"},
			})
			Expect(err).To(HaveOccurred())
		})

		It("consumer with empty name fails", func(ctx context.Context) {
			_, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream:   "test",
				Subjects: []string{"test"},
			})
			Expect(err).To(HaveOccurred())

			_, err = manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream:   "test",
				Name:     "",
				Subjects: []string{"test"},
			})
			Expect(err).To(HaveOccurred())
		})

		It("consumer with invalid name fails", func(ctx context.Context) {
			_, err := manager.EnsureStream(ctx, &events.StreamConfig{
				Name: "test",
				Subjects: []string{
					"test",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream: "test",
				Name:   "invalid name",
				Subjects: []string{
					"test",
				},
			})
			Expect(err).To(HaveOccurred())
		})

		It("consumer with zero subjects fails", func(ctx context.Context) {
			_, err := manager.EnsureStream(ctx, &events.StreamConfig{
				Name: "test",
				Subjects: []string{
					"test",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream: "test",
			})
			Expect(err).To(HaveOccurred())
		})

		It("consumer with invalid subject fails", func(ctx context.Context) {
			_, err := manager.EnsureStream(ctx, &events.StreamConfig{
				Name: "test",
				Subjects: []string{
					"test",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream: "test",
				Subjects: []string{
					"invalid subject",
				},
			})
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Ephemeral", func() {
		It("can create a subscription", func(ctx context.Context) {
			_, err := manager.EnsureStream(ctx, &events.StreamConfig{
				Name: "test",
				Subjects: []string{
					"test",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = js.Consumer(ctx, "test", "test")
			Expect(err).To(HaveOccurred())
			Expect(errors.Is(err, jetstream.ErrConsumerNotFound)).To(BeTrue())

			s, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream: "test",
				Subjects: []string{
					"test",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = js.Consumer(ctx, "test", s.ID)
			Expect(err).ToNot(HaveOccurred())
		})

		It("can create consumer with multiple subjects", func(ctx context.Context) {
			_, err := manager.EnsureStream(ctx, &events.StreamConfig{
				Name: "test",
				Subjects: []string{
					"test",
					"test.>",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = js.Consumer(ctx, "test", "test")
			Expect(err).To(HaveOccurred())
			Expect(errors.Is(err, jetstream.ErrConsumerNotFound)).To(BeTrue())

			s, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream: "test",
				Subjects: []string{
					"test",
					"test.2",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			c, err := js.Consumer(ctx, "test", s.ID)
			Expect(err).ToNot(HaveOccurred())
			Expect(c.CachedInfo().Config.FilterSubject).To(BeEmpty())
			Expect(c.CachedInfo().Config.FilterSubjects).To(ConsistOf("test", "test.2"))
		})
	})

	Describe("Durable", func() {
		It("can create a subscription", func(ctx context.Context) {
			_, err := manager.EnsureStream(ctx, &events.StreamConfig{
				Name: "test",
				Subjects: []string{
					"test",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = js.Consumer(ctx, "test", "test")
			Expect(err).To(HaveOccurred())
			Expect(errors.Is(err, jetstream.ErrConsumerNotFound)).To(BeTrue())

			s, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream: "test",
				Name:   "test",
				Subjects: []string{
					"test",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = js.Consumer(ctx, "test", s.ID)
			Expect(err).ToNot(HaveOccurred())
		})

		It("can update subject of subscription", func(ctx context.Context) {
			_, err := manager.EnsureStream(ctx, &events.StreamConfig{
				Name: "test",
				Subjects: []string{
					"test",
					"test.>",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			s, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream: "test",
				Name:   "test",
				Subjects: []string{
					"test",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = js.Consumer(ctx, "test", s.ID)
			Expect(err).ToNot(HaveOccurred())

			s, err = manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream: "test",
				Name:   "test",
				Subjects: []string{
					"test.2",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			c, err := js.Consumer(ctx, "test", s.ID)
			Expect(err).ToNot(HaveOccurred())
			Expect(c.CachedInfo().Config.FilterSubject).To(Equal("test.2"))
		})

		It("can create consumer with multiple subjects", func(ctx context.Context) {
			_, err := manager.EnsureStream(ctx, &events.StreamConfig{
				Name: "test",
				Subjects: []string{
					"test",
					"test.>",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = js.Consumer(ctx, "test", "test")
			Expect(err).To(HaveOccurred())
			Expect(errors.Is(err, jetstream.ErrConsumerNotFound)).To(BeTrue())

			s, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream: "test",
				Name:   "test",
				Subjects: []string{
					"test",
					"test.2",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			c, err := js.Consumer(ctx, "test", s.ID)
			Expect(err).ToNot(HaveOccurred())
			Expect(c.CachedInfo().Config.FilterSubject).To(BeEmpty())
			Expect(c.CachedInfo().Config.FilterSubjects).To(ConsistOf("test", "test.2"))
		})

		It("can update from one subject to multiple", func(ctx context.Context) {
			_, err := manager.EnsureStream(ctx, &events.StreamConfig{
				Name: "test",
				Subjects: []string{
					"test",
					"test.>",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			s, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream: "test",
				Name:   "test",
				Subjects: []string{
					"test",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = js.Consumer(ctx, "test", s.ID)
			Expect(err).ToNot(HaveOccurred())

			s, err = manager.EnsureConsumer(ctx, &events.ConsumerConfig{
				Stream: "test",
				Name:   "test",
				Subjects: []string{
					"test.2",
					"test.3",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			c, err := js.Consumer(ctx, "test", s.ID)
			Expect(err).ToNot(HaveOccurred())
			Expect(c.CachedInfo().Config.FilterSubject).To(BeEmpty())
			Expect(c.CachedInfo().Config.FilterSubjects).To(ConsistOf("test.2", "test.3"))
		})

		Describe("From", func() {
			It("defaults to delivering new messages", func(ctx context.Context) {
				_, err := manager.EnsureStream(ctx, &events.StreamConfig{
					Name: "test",
					Subjects: []string{
						"test",
					},
				})
				Expect(err).ToNot(HaveOccurred())

				s, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
					Stream: "test",
					Name:   "test",
					Subjects: []string{
						"test",
					},
				})
				Expect(err).ToNot(HaveOccurred())

				c, err := js.Consumer(ctx, "test", s.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(c.CachedInfo().Config.DeliverPolicy).To(Equal(jetstream.DeliverNewPolicy))
			})

			It("can set to specific start time", func(ctx context.Context) {
				_, err := manager.EnsureStream(ctx, &events.StreamConfig{
					Name: "test",
					Subjects: []string{
						"test",
					},
				})
				Expect(err).ToNot(HaveOccurred())

				t := time.Now()
				s, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
					Stream: "test",
					Name:   "test",
					Subjects: []string{
						"test",
					},
					From: &events.StreamPointer{
						Time: t,
					},
				})
				Expect(err).ToNot(HaveOccurred())

				c, err := js.Consumer(ctx, "test", s.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(c.CachedInfo().Config.DeliverPolicy).To(Equal(jetstream.DeliverByStartTimePolicy))
				Expect(*c.CachedInfo().Config.OptStartTime).To(BeTemporally("~", t, time.Millisecond))
			})

			It("can set to specific start ID", func(ctx context.Context) {
				_, err := manager.EnsureStream(ctx, &events.StreamConfig{
					Name: "test",
					Subjects: []string{
						"test",
					},
				})
				Expect(err).ToNot(HaveOccurred())

				s, err := manager.EnsureConsumer(ctx, &events.ConsumerConfig{
					Stream: "test",
					Name:   "test",
					Subjects: []string{
						"test",
					},
					From: &events.StreamPointer{
						ID: 1,
					},
				})
				Expect(err).ToNot(HaveOccurred())

				c, err := js.Consumer(ctx, "test", s.ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(c.CachedInfo().Config.DeliverPolicy).To(Equal(jetstream.DeliverByStartSequencePolicy))
				Expect(c.CachedInfo().Config.OptStartSeq).To(Equal(uint64(1)))
			})
		})
	})
})
