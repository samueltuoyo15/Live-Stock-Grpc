package main

import (
	"context"
	"log/slog"
	"os"
	"time"
	"github.com/samueltuoyo15/Live-Stock-Grpc/client/clients"
	"github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
	"github.com/samueltuoyo15/Live-Stock-Grpc/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main(){
  utils.InitLogger(true)
  
  connection, err := grpc.Dial("http://localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials))
  if err != nil {
    slog.Error("Could not connect to Grpc Server", err)
    os.Exit(1)
  }
  defer connection.Close()
  
  client := stockpb.NewStockServiceClient(connection)
  stockClients := clients.NewStockClient(client)
  
  slog.Info("Starting Grpc Client Demo")
  
  // unary rpc
  stockClient.GetStock("AAPL")
  
  // server streaming
  stockClient.WatchStock("GOOGL")
  
  // client streaming 
  stockClient.UploadStockHistory([] *stockpb.StockRequest{
    { Symbol: "MSFT", Price: 280.0 },
    { Symbol: "MSFT", Price: 291.5 },
    { Symbol: "MSFT", Price: 299.5 },
  })
  
  // bi-directional straeaming 
  stockClient.ChatStock([] *stockpb.StockRequest{
    { Symbol: "AMZN" },
    { Symbol: "TSLA" },
    { Symbol: "NVDA" },
  })
}