package query

import (
	"context"
	"encoding/json"
	"fmt"
	"search_searvice/elasticsearch"

	elastic "github.com/olivere/elastic/v7"
)

//InsertData - function to insert data into elastic database
func InsertData(ctx context.Context, esclient *elastic.Client, name, id, category, owner string) error {
	var isEvent int
	if category != "" {
		isEvent = 1
	}

	newEntry := elasticsearch.Model{
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
		return err
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")

	return nil
}
