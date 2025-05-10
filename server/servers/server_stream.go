package servers

import(
  "time"
  "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
 )

func(s *StockServer) WatchStock(req *stockpb.StockRequest, stream stockpb.StockService_WatchStockServer) error {
  s.Logger.Info("Streaming Stock updates for: %s", req.GetSymbol())
  
  for i := 0; i < 10; i++ {
    price := 100.0 + float64(i)
    
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
