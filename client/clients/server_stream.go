package clients

import (
	"context"
	"io"
	"log/slog"
	"time"
	"fmt"
	"github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
)

func(sc *StockClient) WatchStock(symbol string){
  slog.Info("Starting Server Streaming Rpc")
  
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  
  stream, err := sc.client.WatchStock(ctx, &stockpb.StockRequest{
    Symbol: symbol,
  })
  if err != nil {
    slog.Error("Error calling watch stock", "error", err)
    return 
  }
  
  for {
    res, err := stream.Recv()
    if err == io.EOF {
      slog.Info("Stream Ended")
      break
    }
    if err != nil {
      slog.Error("Error receiving stream", "error", err)
      break
    }
      slog.Info("Stock Response",
       "symbol", res.GetSymbol(),
       "price", fmt.Sprintf("$%.2f", res.GetPrice()),
       "timestamp", res.GetTimestamp())
  }
}