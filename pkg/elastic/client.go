package elastic

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-logr/zerologr"
)

/*
* Returns a new elasticsearch client
* @Param logger
* @Return pointer to client instance
 */
func NewClient(logger zerologr.Logger) *elasticsearch.Client {
	//Setup elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://es01:9200",
		},
	}
	esClient, err := elasticsearch.NewClient(cfg)

	if err != nil {
		logger.Error(err, "Failed to start es client")
	}

	return esClient
}
