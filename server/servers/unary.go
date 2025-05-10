package servers

import (
  "context"
  "time"
  "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
  )
  
func (s *StockServer) GetStock(ctx context.Context, req *stockpb.StockRequest) (*stockpb.StockResponse, error) {
  s.Logger.Info("Received GetStock request from symbol: %s", req.GetSymbol())
  
  return &stockpb.StockResponse{
    Symbol: req.GetSymbol(),
    Price: 100.0,
    Timestamp: time.Now().Format(time.RFC3339),
  }, nil
}


