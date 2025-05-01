package servers

import (
  "io"
  "time"
  "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
  )


func(s *StockServer) ChatStock(stream stockpb.StockService_ChatStockServer) error {
  for {
    req, err := stream.Recv()
    if err == io.EOF {
      s.logger.Info("Chat ended")
      return nil
    }
    if err != nil {
      return err 
    }
    s.loger.Infof("Received chat stock request for: %s", req.Symbol)
    resp := &stockpb.StockResponse{
      Symbol: req.Symbol,
      Price: 200.0,
      Timestamp: time.Now().Format(time.RFC3339),
    }
    if err := stream.Send(resp); err != nil {
      return err 
    }
  }
}

