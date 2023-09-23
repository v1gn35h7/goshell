package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-logr/logr"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/goshell/pkg/elastic"
	"github.com/v1gn35h7/goshell/pkg/goshell"
)

func IndexerService(logger logr.Logger) {
	// Start kafka consumers
	conf := make(kafka.ConfigMap)
	conf["bootstrap.servers"] = viper.GetString("kafka.bootstrapServers")
	conf["auto.offset.reset"] = "earliest"
	conf["group.id"] = "cep-consumer-grp"

	wg := &sync.WaitGroup{}

	contxt := context.Background()
	ctc, cancel := context.WithCancel(contxt)

	// Clean up all goroutines
	defer cancel()

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(index int, ctx context.Context) {
			defer wg.Done()
			consumer, err := kafka.NewConsumer(&conf)

			if err != nil {
				logger.Error(err, "Error starting the consumer", "consumerId:", index*4654)
			}

			err = consumer.SubscribeTopics([]string{"trooper-cep-results"}, nil)

			if err != nil {
				logger.Error(err, "Error subscribing to topic")
				os.Exit(1)
			}

			partitions := make([]kafka.TopicPartition, 0)
			t := "trooper-cep-results"
			m := ""
			partitions = append(partitions, kafka.TopicPartition{&t, 0, 0, &m, nil, new(int32)})
			partitions = append(partitions, kafka.TopicPartition{&t, 1, 0, &m, nil, new(int32)})

			offset, err := consumer.Committed(partitions, 100)

			if err != nil {
				logger.Error(err, "Failed to read topic offset")
			} else {
				logger.Info("Topic offset", "offset", offset)
				consumer.Seek(kafka.TopicPartition{&t, 0, 0, &m, nil, new(int32)}, 100)
				consumer.Seek(kafka.TopicPartition{&t, 1, 0, &m, nil, new(int32)}, 100)
			}

			// consume messages
			run := true
			for run {
				select {
				case sig := <-sigchan:
					logger.Info("Msg", "Caught signal terminating", "SIgnal", sig)
					run = false
				default:
					record, err := consumer.ReadMessage(10000 * time.Millisecond)
					if err != nil {
						// Errors are informational and automatically handled by the consumer
						logger.Error(err, "Failed to read message from kafka")
						continue
					}
					output := goshell.Output{}
					err = json.Unmarshal(record.Value, &output)
					if err != nil {
						logger.Error(err, "Failed to unmarshal kafka paylod")
					}
					fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
						*record.TopicPartition.Topic, string(record.Key), output)

					if output.Output != "" {
						elastic.IndexDocument(logger, elastic.NewClient(logger), record.Value)
					}
				}
			}
		}(i, ctc)
	}

	wg.Wait()
}
