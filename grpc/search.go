package grpc

import (
	context "context"
	"search_searvice/elastic"
)

//Server - an interface
type Server struct {
}

//GetItems - function to query data
func (s *Server) GetItems(ctx context.Context, message *SearchRequest) (*SearchResponse, error) {
	client, err := elastic.GetESClient()

	if err != nil {
		return nil, err
	}

}
