package service

import (
	"github.com/Shopify/sarama"
	"github.com/v1gn35h7/goshell/pkg/kafka"
	"github.com/v1gn35h7/goshell/pkg/logging"
)

type transformationService interface {
	getEventsProto() (bool, error)
	pushEvents() (bool, error)
}

func (srvc service) getEventsProto() (bool, error) {
	// TODO: get proto for parsing
	return true, nil
}

func (srvc service) pushEvents() (bool, error) {
	event := "Test event data"
	client := kafka.NewKafkaClient(logging.Logger())

	producer := client.Producer
	if producer == nil {
		return true, nil
	}
	// Start Kafka transaction
	err := producer.BeginTxn()
	if err != nil {
		srvc.logger.Log(err, "Kafa transtcion failied")
	}

	//eventsTopic := viper.GetString("kafka.producers[0].topic")

	producer.Input() <- &sarama.ProducerMessage{Topic: "trooper-events", Key: sarama.StringEncoder(event), Value: sarama.StringEncoder(event)}

	// commit transaction
	err = producer.CommitTxn()
	if err != nil {
		srvc.logger.Log(err, "Producer: unable to commit tx")
		for {
			if producer.TxnStatus()&sarama.ProducerTxnFlagFatalError != 0 {
				// fatal error. need to recreate producer.
				srvc.logger.Log("Producer: producer is in a fatal state, need to recreate it")
				break
			}
			// If producer is in abortable state, try to abort current transaction.
			if producer.TxnStatus()&sarama.ProducerTxnFlagAbortableError != 0 {
				err = producer.AbortTxn()
				if err != nil {
					// If an error occured just retry it.
					srvc.logger.Log(err, "Producer: unable to abort transaction")
					continue
				}
				break
			}
			// if not you can retry
			err = producer.CommitTxn()
			if err != nil {
				srvc.logger.Log(err, "Producer: unable to commit txn")
				continue
			}
		}
		return true, nil

	}
	return true, nil

}
