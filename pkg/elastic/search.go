package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-logr/logr"
)

func Search(query string, esClient *elasticsearch.Client, logger logr.Logger) ([]interface{}, error) {
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
		logger.Error(err, "Error encoding query")
		return nil, err
	}

	// Perform the search request.
	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex("result*"),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)

	if err != nil {
		logger.Error(err, "Error getting response")
		return nil, err

	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			logger.Error(err, "Error parsing the response body")
			return nil, err
		} else {
			// Print the response status and error information.
			logger.Info("Got response from es",
				"status", res.Status(),
				"query", query,
				"type", e["error"].(map[string]interface{})["type"],
				"reason", e["error"].(map[string]interface{})["reason"])
			return nil, fmt.Errorf("%s", e["error"].(map[string]interface{})["reason"])
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logger.Error(err, "Error parsing the response body")
		return nil, err
	}

	// Print the response status, number of results, and request duration.
	// For debugging only
	// log.Printf(
	// 	"[%s] %d hits; took: %dms",
	// 	res.Status(),
	// 	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	// 	int(r["took"].(float64)),
	// )

	return r["hits"].(map[string]interface{})["hits"].([]interface{}), nil
}
