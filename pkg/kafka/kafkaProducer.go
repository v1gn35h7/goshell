package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/go-logr/zerologr"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/pkg/logging"
)

var (
	version string = "2"
)

type KafkaClient struct {
	Config   *sarama.Config
	Producer sarama.AsyncProducer
	logger   zerologr.Logger
}

func NewKafkaClient(logr zerologr.Logger) *KafkaClient {
	sConf := getKafkaClientConfig()

	// Handle panic
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		logr.Info("Kafa connection failed", "Rec", r)
	// 	}
	// }()

	//brokers := []string{"127.0.0.1:29092"}
	kbrokers := viper.GetString("kafka.bootstrapServers")
	kproducers := strings.Split(kbrokers, ",")
	logr.Info("Trying to connect to kafka brokers", "brokers", kproducers)

	producer, err := sarama.NewAsyncProducer(kproducers, sConf)
	if err != nil {
		logr.Error(err, "Failed to initiate kafka producer")
		panic(err)
	}

	return &KafkaClient{
		Config:   sConf,
		Producer: producer,
		logger:   logr,
	}
}

func getKafkaClientConfig() *sarama.Config {
	config := sarama.NewConfig()
	//config.Version = "0.11.0.0"
	config.Producer.Idempotent = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Transaction.Retry.Backoff = 1
	config.Producer.Transaction.ID = "trooper_producer"
	config.Net.MaxOpenRequests = 1
	return config
}

func ProduceRecord(srvc interface{}, fragment goshell.Fragment) {
	kclient := NewKafkaClient(logging.Logger())
	err := kclient.Producer.BeginTxn()

	if err != nil {
		kclient.logger.Error(err, "Error in kafka transaction")
	}

	// Produce some records in transaction
	for _, output := range fragment.Outputs {
		record, er := json.Marshal(output)

		if er != nil {
			kclient.logger.Error(er, "Error ...")
		}
		fmt.Println(record)
		kclient.Producer.Input() <- &sarama.ProducerMessage{Topic: "trooper-scripts-results", Key: nil, Value: sarama.ByteEncoder(record)}
	}

	// commit transaction
	err = kclient.Producer.CommitTxn()
	if err != nil {
		log.Printf("Producer: unable to commit txn %s\n", err)
		for {
			if kclient.Producer.TxnStatus()&sarama.ProducerTxnFlagFatalError != 0 {
				// fatal error. need to recreate producer.
				log.Printf("Producer: producer is in a fatal state, need to recreate it")
				break
			}
			// If producer is in abortable state, try to abort current transaction.
			if kclient.Producer.TxnStatus()&sarama.ProducerTxnFlagAbortableError != 0 {
				err = kclient.Producer.AbortTxn()
				if err != nil {
					// If an error occured just retry it.
					log.Printf("Producer: unable to abort transaction: %+v", err)
					continue
				}
				break
			}
			// if not you can retry
			err = kclient.Producer.CommitTxn()
			if err != nil {
				log.Printf("Producer: unable to commit txn %s\n", err)
				continue
			}
		}
		return
	}

}
