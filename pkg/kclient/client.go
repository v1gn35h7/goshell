package kclient

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/go-logr/zerologr"
	respository "github.com/v1gn35h7/goshell/internal/repository"
	"github.com/v1gn35h7/goshell/pkg/goshell"
)

func StartKafkaConsumer(ctx context.Context, consumerName string, logger zerologr.Logger, config *sarama.Config, brokers []string, topic string) {
	logger.Info(fmt.Sprintf("Starting kafka consumer %s", consumerName))

	consumer, err := sarama.NewConsumer(brokers, config)

	if err != nil {
		logger.Error(err, "Kafka consumer initialization failed")
	}

	pconsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		logger.Error(err, "Kafka partition consumer initialization failed")
	}

	for {

		select {
		case msg := <-pconsumer.Messages():
			logger.Info("Message received... ", "Key:", msg.Key, "Value:", string(msg.Value), "Offset:", msg.Offset)
			event := goshell.Events{
				Id:   string(msg.Key),
				Name: string(msg.Value),
			}
			respository.EventRepository(logger).AddEvents(event)
		case <-ctx.Done():
			return
		default:
			continue
		}

	}

}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready  chan bool
	logger zerologr.Logger
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			msg := fmt.Sprintf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			consumer.logger.Info(msg)
			session.MarkMessage(message, "")

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}
