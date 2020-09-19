package query

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic"
)

//Model - Document to be stored in index
type Model struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsEvent  int    `json:"isEvent"`
	Category string `json:"category,omitempty"`
	Owner    string `json:"owner,omitempty"`
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
