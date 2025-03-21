package service

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	respository "github.com/v1gn35h7/goshell/internal/repository"
	"github.com/v1gn35h7/goshell/pkg/elastic"
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

func (s service) ExecuteCmd(cmd string) (string, error) {
	return "ellow!!", nil
}

func (s service) ConnectToRemoteHost(hostId string) (bool, error) {
	return true, nil
}

func (s service) GetScripts(asset goshell.Asset) ([]*goshell.Script, error) {
	respository.AssetsRepository(logging.Logger()).Update(asset)
	return respository.ScriptsRepository(logging.Logger()).List(asset.Agentid)
}

func (s service) SaveScripts(scriptPayload goshell.Script) (bool, error) {
	scriptPayload.Id = uuid.NewString()
	return respository.ScriptsRepository(logging.Logger()).Save(scriptPayload)
}

func (s service) SendFragment(fragment goshell.Fragment) (int32, error) {
	kafkaConfig := make(map[string]kafka.ConfigValue)
	kafkaConfig["bootstrap.servers"] = viper.GetString("kafka.bootstrapServers")
	kafkaConfig["acks"] = "all"

	kafkaProducer := intKafka.NewProducer(kafkaConfig, logging.Logger())

	resultsTopic := viper.GetString("kafka.producers.results.topic")

	// Produce some records in transaction
	for _, output := range fragment.Outputs {
		record, er := json.Marshal(output)

		if er != nil {
			s.logger.Log("Error", "Filed while serializing", "er", er)
		}
		kafkaProducer.Create(resultsTopic, record)
	}

	// Close the producer
	kafkaProducer.Close()

	return int32(1), nil
}

func (s service) SearchResults(query string) ([]*goshell.Output, error) {
	data := make([]*goshell.Output, 0)
	client := elastic.NewClient(logging.Logger())

	hits, err := elastic.Search(query, client, logging.Logger())

	if err != nil {
		return data, err
	}

	for _, hit := range hits {
		output := &goshell.Output{
			Id:       hit.(map[string]interface{})["_id"].(string),
			Agentid:  hit.(map[string]interface{})["_source"].(map[string]interface{})["agentid"].(string),
			Hostname: hit.(map[string]interface{})["_source"].(map[string]interface{})["hostname"].(string),
			Scriptid: hit.(map[string]interface{})["_source"].(map[string]interface{})["scriptid"].(string),
			Output:   hit.(map[string]interface{})["_source"].(map[string]interface{})["output"].(string),
			Score:    hit.(map[string]interface{})["_source"].(map[string]interface{})["score"].(string),
		}
		data = append(data, output)
	}

	return data, nil
}
