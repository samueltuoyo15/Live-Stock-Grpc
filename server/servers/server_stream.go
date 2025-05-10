package servers

import(
  "time"
  "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
 )

func(s *StockServer) WatchStock(req *stockpb.StockRequest, stream stockpb.StockService_WatchStockServer) error {
  s.Logger.Info("Streaming Stock updates for: %s", req.GetSymbol())
  
  for i := 0; i < 10; i++ {
    price, err := GetRealTimePrice(req.GetSymbol()) 
     if err != nil {
      s.Logger.Error("Failed to get real-time price", 
        "symbol", req.GetSymbol(), 
        "error", err)
        return err 
        }
    resp := &stockpb.StockResponse {
      Symbol: req.GetSymbol(),
      Price: price,
      Timestamp: time.Now().Format(time.RFC3339),
    }
    if err := stream.Send(resp); err != nil {
      return err
    }
    time.Sleep(1 * time.Second)
  }
  return nil
}
