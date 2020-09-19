package elastic

import (
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

// GetESClient - gives the connection to elastic search server
func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))

	fmt.Println("ES initialized")

	return client, err
}
