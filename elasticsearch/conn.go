package elasticsearch

import (
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

//Model - Document to be stored in index
type Model struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Category    string `json:"category,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Description string `json:"description,omitempty"`
}

// GetESClient - gives the connection to elastic search server
func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://elastic:Wethefolks@123@lb-uzja7dl7ropiy.centralindia.cloudapp.azure.com:9200"), elastic.SetHealthcheck(false), elastic.SetSniff(false))

	fmt.Println("ES initialized")

	return client, err
}
