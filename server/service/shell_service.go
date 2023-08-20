package service

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	respository "github.com/v1gn35h7/goshell/internal/repository"
	"github.com/v1gn35h7/goshell/pkg/goshell"
	intKafka "github.com/v1gn35h7/goshell/pkg/kafka"
	"github.com/v1gn35h7/goshell/pkg/logging"
)

type shellService interface {
	ExecuteCmd(cmd string) (string, error)
	ConnectToRemoteHost(hostId string) (bool, error)
	GetScripts(asset goshell.Asset) ([]*goshell.Script, error)
	SaveScripts(scriptPayload goshell.Script) (bool, error)
	//EndpointHeartBeat(hostId string) ([]execu)
	SendFragment(payload goshell.Fragment) (int32, error)
	SearchResults(query string) ([]*goshell.Output, error)
}

func (srvc service) ExecuteCmd(cmd string) (string, error) {
	return "ellow!!", nil
}

func (srvc service) ConnectToRemoteHost(hostId string) (bool, error) {
	return true, nil
}

func (srvc service) GetScripts(asset goshell.Asset) ([]*goshell.Script, error) {
	respository.AssetsRepository(logging.Logger()).UpdateAsset(asset)
	return respository.ScriptsRepository(logging.Logger()).GetScripts(asset.Agentid)
}

func (srvc service) SaveScripts(scriptPayload goshell.Script) (bool, error) {
	scriptPayload.Id = uuid.NewString()
	return respository.ScriptsRepository(logging.Logger()).AddScripts(scriptPayload)
}

func (srvc service) SendFragment(fragment goshell.Fragment) (int32, error) {
	kafkaConfig := make(map[string]kafka.ConfigValue)
	kafkaConfig["bootstrap.servers"] = viper.GetString("kafka.bootstrapServers")
	kafkaConfig["acks"] = "all"

	kafkaProducer := intKafka.NewProducer(kafkaConfig, logging.Logger())

	resultsTopic := viper.GetString("kafka.producers.results.topic")

	// Produce some records in transaction
	for _, output := range fragment.Outputs {
		record, er := json.Marshal(output)

		if er != nil {
			srvc.logger.Log("Error", "Filed while serializing", "er", er)
		}
		kafkaProducer.Create(resultsTopic, record)
	}

	// Close the producer
	kafkaProducer.Close()

	return int32(1), nil
}

func (srvc service) SearchResults(query string) ([]*goshell.Output, error) {
	return respository.ResultsRepository(logging.Logger()).SearchResults(query)
}
