package goindexer

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/goshell/internal/config"
	"github.com/v1gn35h7/goshell/pkg/goshell"
	fkafka "github.com/v1gn35h7/goshell/pkg/kafka"
	"github.com/v1gn35h7/goshell/pkg/logging"
)

func NewSeedCommand() *cobra.Command {
	var startCmd = &cobra.Command{
		Use:   "seed",
		Short: "goshellctl seed",
		Long:  "Seeds Goshell cli process",
		Run: func(cmd *cobra.Command, args []string) {
			// Set-up logger
			logger := logging.Logger()
			logger.Info("Logger initated...")

			// Seed records
			seedRecords(logger)

		},
	}
	return startCmd
}

func seedRecords(logger logr.Logger) {
	config.Read(configPath, logger)
	kafkaConfig := make(map[string]kafka.ConfigValue)
	kafkaConfig["bootstrap.servers"] = viper.GetString("kafka.bootstrapServers")
	kafkaConfig["acks"] = "all"

	kafkaProducer := fkafka.NewProducer(kafkaConfig, logging.Logger())

	defer kafkaProducer.Close()

	if kafkaProducer == nil {
		logger.Info("Failed to start kafka producer")
	}

	resultsTopic := "trooper-cep-results"

	// Produce some records in transaction
	for {
		output := goshell.Output{
			Agentid:  uuid.NewString(),
			Hostname: "SONU07",
			Scriptid: "3131-42423-545-6542",
			Score:    "3",
			Output:   "Test output ....",
		}
		record, er := json.Marshal(output)

		if er != nil {
			logger.Info("Failed to encdode data")
		}

		kafkaProducer.Create(resultsTopic, record)

	}

}
