package servers

import (
  "context"
  "fmt"
  "time"
  "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
  )
  
func (s *StockServer) GetStock(ctx context.Context, req *stockpb.StockRequest) (*stockpb.StockResponse, error) {
  s.logger.Infof("Received GetStock request from symbol: %s", req.symbol)
  
  return &stockpb.StockResponse{
    Symbol: req.Symbol,
    Price: 100.0,
    Timestamp: time.Now().Format(time.RFC3339),
  }, nil
}


