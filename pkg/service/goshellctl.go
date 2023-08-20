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
	"github.com/go-logr/zerologr"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	respository "github.com/v1gn35h7/goshell/internal/repository"
	"github.com/v1gn35h7/goshell/pkg/goshell"
)

type Service struct {
	logger       zerologr.Logger
	srvcInstance *Service
}

func NewService(logr zerologr.Logger) *Service {
	return &Service{
		logger: logr,
	}
}

func (srvc *Service) StartService() (bool, error) {
	srvc.logger.Info("Starting GoshellCtl service...")

	kconfig := getClientConfig()

	rst := viper.GetString("kafka.consumers[0].topic")
	fmt.Println(rst)

	resultsTopic := "trooper-cep-results"

	// kbrokers := viper.GetString("kafka.bootstrapServers")
	// kproducers := strings.Split(kbrokers, ",")

	// cpus := runtime.GOMAXPROCS(runtime.NumCPU())

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
			// consumerName := "consumer-" + strconv.Itoa(index)
			// kafka.StartKafkaConsumer(ctx, consumerName, srvc.logger, kconfig, kproducers, "trooper-scripts-results")
			consumer, err := kafka.NewConsumer(&kconfig)

			if err != nil {
				srvc.logger.Error(err, "Error starting the consumer", "consumerId:", i)
			}

			err = consumer.SubscribeTopics([]string{resultsTopic}, nil)

			if err != nil {
				srvc.logger.Error(err, "Error subscribing to topic")
				os.Exit(1)
			}

			// Process messages
			run := true
			for run {
				select {
				case sig := <-sigchan:
					srvc.logger.Info("Msg", "Caught signal terminating", "SIgnal", sig)
					run = false
				default:
					record, err := consumer.ReadMessage(10000 * time.Millisecond)
					if err != nil {
						// Errors are informational and automatically handled by the consumer
						srvc.logger.Error(err, "Failed to read message from kafka")
						continue
					}
					output := goshell.Output{}
					err = json.Unmarshal(record.Value, &output)
					if err != nil {
						srvc.logger.Error(err, "Failed to unmarshal kafka paylod")
					}
					fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
						*record.TopicPartition.Topic, string(record.Key), output)

					if output.Output != "" {
						output.Id = uuid.NewString()
						respository.ResultsRepository(srvc.logger).AddResults(output)
					}

				}
			}
		}(i, ctc)
	}

	wg.Wait()

	return true, nil
}

func getClientConfig() kafka.ConfigMap {
	conf := make(kafka.ConfigMap)

	conf["bootstrap.servers"] = viper.GetString("kafka.bootstrapServers")
	conf["auto.offset.reset"] = "earliest"
	conf["group.id"] = "results-consumer-grp"

	return conf
}
