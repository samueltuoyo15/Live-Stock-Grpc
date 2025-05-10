package servers

import (
  "context"
  "time"
  "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
  )
  
func (s *StockServer) GetStock(ctx context.Context, req *stockpb.StockRequest) (*stockpb.StockResponse, error) {
    s.Logger.Info("Fetching REAL-TIME price", "symbol", req.GetSymbol())
    
    price, err := GetRealTimePrice(req.GetSymbol()) 
    if err != nil {
        s.Logger.Error("Failed to get real-time price", "error", err)
        return nil, err
    }
  return &stockpb.StockResponse{
    Symbol: req.GetSymbol(),
    Price: price,
    Timestamp: time.Now().Format(time.RFC3339),
  }, nil
}


