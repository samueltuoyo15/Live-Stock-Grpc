package clients

import (
	"context"
	"io"
	"log/slog"
	"sync"
	"time"
	"github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
)

func (sc *StockClient) ChatStock(stocks []*stockpb.StockRequest) {
  slog.Info("Starting Bidirectional Streaming Rpc")
  
  ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
  defer cancel()
  
  stream, err := sc.client.ChatStock(ctx)
  if err != nil {
    slog.Error("Error calling Chat Stock", err)
    return 
  }
  
  var wg sync.WaitGroup()
  wg.Add(2)
  
  // go routines to send response 
  go func(){
    defer wg.Done()
    for _, stock := range stocks {
     slog.Info("Sending Stock symbol", "symbol", stock.GetSymbol())
     if err := stream.Send(stock); err != nil {
     slog.Error("Error sending stock", err)
     return 
    }
    time.Sleep(1, time.Second)
   }
  stream.CloseSend()
  }()
  
  // go routines to revceive response 
  go func(){
    defer wg.Done()
    for {
      res, err := stream.Recv()
      if err == io.EOF{
        slog.Info("Stream Ended")
        return 
      }
      if err != nil{
        slog.Error("Error receiving response", err)
        return 
      }
    	slog.Info("Received Stock Update", 
				"symbol", res.GetSymbol(),
				"price", res.GetPrice(),
				"timestamp", res.GetTimestamp())
	  	}
    }
  }()
  wg.Wait()
  log.Info("Chat Stock Completed")
}

