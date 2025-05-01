package main

import (
 "log"
 "net"
 "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
 "github.com/samueltuoyo15/Live-Stock-Grpc/server/servers"
 "google.golang.org/grpc"
)

func main(){
  listen, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }
  
  // to create a grpc server
  grpcServer := grpc.NewServer()
  stockpb.RegisterStockServiceServer(grpcServer, &servers.StockServiceServer{})
  log.Println("Server is listening on port :8080")
  if err := grpcServer.Serve(listen); err != nil {
    log.Fatalf("Failed to serve: %v", err)
  }
}

