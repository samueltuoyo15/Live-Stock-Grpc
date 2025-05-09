package main

import (
 "log/slog"
 "net"
 "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
 "github.com/samueltuoyo15/Live-Stock-Grpc/server/servers"
 "github.com/samueltuoyo15/Live-Stock-Grpc/utils"
 "google.golang.org/grpc"
)

func main(){
  utils.InitLogger(true)
  listen, err := net.Listen("tcp", ":8080")
  if err != nil {
    slog.Fatalf("Failed to listen: %v", err)
  }
  
  grpcServer := grpc.NewServer()
  stockpb.RegisterStockServiceServer(grpcServer, &servers.StockServiceServer{})
  slog.Println("Server is listening on port :8080")
  if err := grpcServer.Serve(listen); err != nil {
    slog.Fatalf("Failed to serve: %v", err)
  }
}

