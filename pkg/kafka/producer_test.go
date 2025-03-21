package kafka

import (
	"encoding/json"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/goshell/internal/config"
	"github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/pkg/logging"
)

func TestKafkaConnection(t *testing.T) {
	logger := logging.Logger()
	config.Read("../../configs/goshell", logger)
	kafkaConfig := make(map[string]kafka.ConfigValue)
	kafkaConfig["bootstrap.servers"] = viper.GetString("kafka.bootstrapServers")
	kafkaConfig["acks"] = "all"

	kafkaProducer := NewProducer(kafkaConfig, logging.Logger())

	defer kafkaProducer.Close()

	if kafkaProducer == nil {
		t.Error("Failed to start kafka producer")
	}

	resultsTopic := "trooper-cep-results"

	// Produce some records in transaction
	for i := 0; i < 1000; i++ {
		output := goshell.Output{
			Agentid:  "345345-36456-65645-24324",
			Hostname: "SONU07",
			Scriptid: "3131-42423-545-6542",
			Score:    "3",
			Output:   "Test output ....",
		}
		record, er := json.Marshal(output)

		if er != nil {
			t.Error("Failed to encdode data")
		}

		kafkaProducer.Create(resultsTopic, record)

	}

}
