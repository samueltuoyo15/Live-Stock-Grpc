package clients

import (
	"context"
	"io"
	"log/slog"
	"time"
	"github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
)

func(sc *StockClient) WatchStock(symbol string){
  slog.Info("Starting Server Streaming Rpc")
  
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  
  stream, err := sc.clients.WatchStock(ctx, &stockpb.StockRequest{
    Symbol: symbol,
  })
  if err != nil {
    slog.Error("Error calling watch stock", err)
    return 
  }
  
  for {
    res, err := stream.Recv()
    if err == io.EOF {
      slog.Info("Stream Ended")
      break
    }
    if err != nip {
      slog.Error("Error receiving stream")
      break
    }
      slog.Info("Stock Response",
       "symbol": res.GetSymbol(),
       "price": res.GetPrice(),
       "timestamp": res.GetTimestamp())
  }
}