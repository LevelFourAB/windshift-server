package v1alpha1_test

import (
	"context"

	eventsv1alpha1 "github.com/levelfourab/windshift-server/internal/proto/windshift/events/v1alpha1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Consumers", func() {
	var service eventsv1alpha1.EventsServiceClient

	BeforeEach(func(ctx context.Context) {
		service, _ = GetClient()

		_, err := service.EnsureStream(ctx, &eventsv1alpha1.EnsureStreamRequest{
			Name: "test",
			Source: &eventsv1alpha1.EnsureStreamRequest_Subjects_{
				Subjects: &eventsv1alpha1.EnsureStreamRequest_Subjects{
					Subjects: []string{
						"test",
						"events.>",
					},
				},
			},
		})
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("Ephemeral", func() {
		It("can create a consumer", func(ctx context.Context) {
			_, err := service.EnsureConsumer(ctx, &eventsv1alpha1.EnsureConsumerRequest{
				Stream: "test",
				Subjects: []string{
					"test",
				},
			})
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("Durable", func() {
		It("can update subject of consumer", func(ctx context.Context) {
			subID := "test-sub"
			_, err := service.EnsureConsumer(ctx, &eventsv1alpha1.EnsureConsumerRequest{
				Stream: "test",
				Name:   &subID,
				Subjects: []string{
					"test",
				},
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = service.EnsureConsumer(ctx, &eventsv1alpha1.EnsureConsumerRequest{
				Stream: "test",
				Name:   &subID,
				Subjects: []string{
					"events.test",
				},
			})
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
