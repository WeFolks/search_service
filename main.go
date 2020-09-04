package main

import (
	"context"
	"encoding/json"
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
	ctx := context.Background()
	esclient, err := GetESClient()

	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail")
	}

	// insertData(ctx, esclient, "sarthak", "112", "", "")
	// insertData(ctx, esclient, "Anshal", "362", "", "")
	// insertData(ctx, esclient, "shaskdh", "789", "", "")
	// insertData(ctx, esclient, "party", "112", "party", "sarthak")
	// insertData(ctx, esclient, "swim", "112", "party", "shaskdh")
	// insertData(ctx, esclient, "fcf", "112", "football", "Anshal")
	// insertData(ctx, esclient, "vish", "112", "", "")
	// insertData(ctx, esclient, "poker213", "112", "poker", "vish")
	// insertData(ctx, esclient, "mudit", "112", "", "")
	// insertData(ctx, esclient, "meeting", "112", "gmeet", "mudit")
	insertData(ctx, esclient, "sart choudhary", "112", "", "")
}

func insertData(ctx context.Context, esclient *elastic.Client, name, id, category, owner string) {
	var isEvent int
	if category != "" {
		isEvent = 1
	}

	newEntry := Model{
		ID:       id,
		Name:     name,
		Category: category,
		Owner:    owner,
		IsEvent:  isEvent,
	}

	dataJSON, err := json.Marshal(newEntry)
	js := string(dataJSON)
	_, err = esclient.Index().Index("search_data").BodyJson(js).Do(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
}

func queryData(ctx context.Context, esclient *elastic.Client, name, category string) {
	if category != "" {
		queryEvent(ctx, esclient, name, category)
	} else {
		queryUser(ctx, esclient, name)
	}
}

func queryEvent(ctx context.Context, esclient *elastic.Client, name, category string) {
	multiQuery := elastic.NewMultiMatchQuery(name, "name").Type("phrase_prefix")

	matchQuery := elastic.NewMatchQuery("category", category)
	query := elastic.NewBoolQuery().Must(multiQuery, matchQuery)

	searchResult, err := esclient.Search().Index("search_data").Query(query).Do(ctx)

	if err != nil {
		panic(err)
	}

	for _, hit := range searchResult.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}

func queryUser(ctx context.Context, esclient *elastic.Client, name string) {
	multiQuery := elastic.NewMultiMatchQuery(name, "name", "owner").Type("phrase_prefix")

	searchResult, err := esclient.Search().Index("search_data").Query(multiQuery).Do(ctx)

	if err != nil {
		panic(err)
	}

	for _, hit := range searchResult.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}
