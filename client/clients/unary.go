package clients

import (
	"context"
	"log/slog"
	"time"
	"github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
)

func(sc *StockClient) GetStock(symbol string){
  slog.Info("Starting Unary rpc")
  
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  
  res, err := sc.client.GetStock(ctx, &stockpb.StockRequest{
    Symbol: symbol,
  })
  if err != nil{
    slog.Error("Error calling GetStock", "error", err)
    return 
  }
  
  slog.Info("Stock Response",
  "symbol", res.GetSymbol(),
  "price", res.GetPrice(),
  "timestamp", res.GetTimestamp())
}
