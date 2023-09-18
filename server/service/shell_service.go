package service

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

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
	data := make([]*goshell.Output, 0)
	client := elastic.NewElasticClient(logging.Logger())
	// Build the request body.
	var buf bytes.Buffer
	esquery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"output": query,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(esquery); err != nil {
		log.Printf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("result*"),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)

	if err != nil {
		log.Printf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		output := &goshell.Output{
			Id:       hit.(map[string]interface{})["_id"].(string),
			Agentid:  "--",
			Hostname: "--",
			Scriptid: "--",
			Output:   hit.(map[string]interface{})["_source"].(map[string]interface{})["output"].(string),
			Score:    hit.(map[string]interface{})["_source"].(map[string]interface{})["score"].(string),
		}
		data = append(data, output)
	}

	return data, nil
}
