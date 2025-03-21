package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/eapache/go-resiliency/retrier"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
)

func IndexDocument(logger logr.Logger, esClient *elasticsearch.Client, document []byte) {

	r := retrier.New(retrier.ConstantBackoff(3, 100*time.Millisecond), nil)

	err := r.Run(func() error {
		// logger.Info("Retrying index document", "doc", document)
		return saveDocument(logger, esClient, document)
	})

	if err != nil {
		logger.Error(err, "Index retry failed")
	}

}

func saveDocument(logger logr.Logger, esClient *elasticsearch.Client, document []byte) error {

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      "trooper-results-00001",
		DocumentID: uuid.NewString(),
		Body:       bytes.NewReader(document),
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		logger.Error(err, "Error getting response")
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		err := fmt.Errorf("error indexing document")
		logger.Error(err, "Error indexing document", "status", res.Status())
		return err
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}

	return nil
}
