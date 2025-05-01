package servers

import(
  "time"
  "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
 )

func(s *StockServer) WatchStock(req *stockpb.StockRequest, stream stockpb.StockService_WatchStockServer) error {
  s.logger.Infof("Streaming Stock updates for: %s", req.Symbol)
  
  for i := 0; i < 10; i++ {
    price := 100.0 + float64(i)
    
    resp := &stockpb.StockResponse {
      Symbol: req.Symbol,
      Price: price,
      Timestamp: time.Now().Format(time.RFC3339),
    }
    if err := stream.Send(resp); err != nil {
      return error 
    }
    time.Sleep(1 * time.Second)
  }
  return nil
}
