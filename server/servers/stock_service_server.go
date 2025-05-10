package servers

import (
	"log/slog"
	"github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
)

type StockServer struct {
	Logger *slog.Logger
	stockpb.UnimplementedStockServiceServer
}

func NewStockServer(logger *slog.Logger) *StockServer {
	return &StockServer{
		Logger: logger,
	}
}