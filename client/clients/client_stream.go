package clients

import (
	"context"
	"log/slog"
	"time"
	"github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
)

func (sc *StockClient) UploadStockHistory(stocks []*stockpb.StockRequest) {
  slog.Info("Starting Client Streaming Rpc")
  
  ctx, cancel := context.WithTimeout(context.Background, 10*time.Second)
  defer cancel()
  
  stream, err := sc.client.UploadStockHistory(ctx)
  if err != nil {
    slog.Error("Error calling upload stock history", err)
    return 
  }
  
  for _, stock := range stocks {
    slog.Info("Sending Stock", "symbol", stock.GetSymbol())
    if err := stream.Send(stock); err != nil {
    slog.Error("Error sending stock", err)
    return 
    }
    time.Sleep(500, time.Millisecond)
  }
  
  res, err := stream.CloseAndRecv()
  if err != nil {
    slog.Error("Error Receiving Response", err)
    return 
  }
  
  slog.Info("Upload Summary",
  "message", res.GetMessage(),
  "count", res.GetCount())
}