package service

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/go-logr/zerologr"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/goshell/pkg/kclient"
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

	eventsTopic := viper.GetString("kafka.consumers[0].topic")
	fmt.Println(eventsTopic)

	kbrokers := viper.GetString("kafka.bootstrapServers")
	kproducers := strings.Split(kbrokers, ",")

	cpus := runtime.GOMAXPROCS(runtime.NumCPU())

	wg := &sync.WaitGroup{}

	contxt := context.Background()
	ctc, cancel := context.WithCancel(contxt)

	// Clean up all goroutines
	defer cancel()

	for i := 0; i < cpus; i++ {
		wg.Add(1)
		go func(index int, ctx context.Context) {
			defer wg.Done()
			consumerName := "consumer-" + strconv.Itoa(index)
			kclient.StartKafkaConsumer(ctx, consumerName, srvc.logger, kconfig, kproducers, "trooper-events")
		}(i, ctc)
	}

	wg.Wait()

	return true, nil
}

func getClientConfig() *sarama.Config {
	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version, _ = sarama.ParseKafkaVersion("0.11")

	assignor := "range"
	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategySticky}
	case "roundrobin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}

	if false {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	return config
}
