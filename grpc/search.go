package grpc

import (
	context "context"

	"github.com/WeFolks/search_service/elasticsearch"
	"github.com/WeFolks/search_service/elasticsearch/query"
)

//Server - an interface
type Server struct {
}

//GetItems - function to query data
func (s *Server) GetItems(ctx context.Context, message *SearchRequest) (*SearchResponse, error) {
	client, err := elasticsearch.GetESClient()

	if err != nil {
		return nil, err
	}

	result, err := query.Data(ctx, client, message.Name, message.Category)

	if err != nil {
		return nil, err
	}

	items := []*Item{}
	for _, element := range result {
		item := Item{
			Id:       element.ID,
			Name:     element.Name,
			Category: element.Category,
			Owner:    element.Owner,
			Type:     int32(element.Type),
		}

		items = append(items, &item)

	}

	return &SearchResponse{Items: items}, nil
}

//AddItem - rpc function to insert data in elastic search
func (s *Server) AddItem(ctx context.Context, message *Item) (*Response, error) {
	client, err := elasticsearch.GetESClient()

	if err != nil {
		response := Response{
			Error: 1,
		}

		return &response, err
	}

	err = query.InsertData(ctx, client, message.Name, message.Id, message.Category, message.Owner, int(message.Type))

	if err != nil {
		response := Response{
			Error: 1,
		}

		return &response, err
	}

	response := Response{
		Error: 0,
	}

	return &response, nil
}
