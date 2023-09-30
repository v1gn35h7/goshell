package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-logr/zerologr"
)

type producer struct {
	Instance *kafka.Producer
	logger   zerologr.Logger
}

func NewProducer(conf kafka.ConfigMap, logger zerologr.Logger) *producer {
	fmt.Println("kafka client ...")
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		logger.Error(err, "Failed to creater prdoucer instance")
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s\n",
						*ev.TopicPartition.Topic, string(ev.Key))
				}
			}
		}
	}()

	return &producer{
		Instance: p,
		logger:   logger,
	}
}

func (p *producer) Create(topic string, record []byte) {
	err := p.Instance.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            record,
		Value:          record,
	}, nil)

	if err != nil {
		p.logger.Error(err, "Error writing to kafka")
	}

}

func (p *producer) Close() {
	p.Instance.Flush(1 * 1000)
	p.Instance.Close()
}
