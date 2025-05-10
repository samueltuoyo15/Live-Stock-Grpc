package clients

import (
	"github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
)

type StockClient struct {
	client stockpb.StockServiceClient
}

func NewStockClient(client stockpb.StockServiceClient) *StockClient {
	return &StockClient{
		client: client,
	}
}