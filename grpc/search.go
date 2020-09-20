package grpc

import (
	context "context"
	"search_searvice/elasticsearch"
	"search_searvice/elasticsearch/query"
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
			IsEvent:  int32(element.IsEvent),
		}

		items = append(items, &item)

	}

	return &SearchResponse{Items: items}, nil
}
