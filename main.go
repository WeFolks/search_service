package main

import (
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

//Model - Document to be stored in index
type Model struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsEvent  int    `json:"isEvent"`
	Category string `json:"category,omitempty"`
	Owner    string `json:"owner,omitempty"`
}

//GetESClient - to connect elastic search
func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))

	fmt.Println("ES initialized")

	return client, err
}

func main() {

}
