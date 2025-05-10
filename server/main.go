package main

import (
 "net"
 "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
 "github.com/samueltuoyo15/Live-Stock-Grpc/server/servers"
 "github.com/samueltuoyo15/Live-Stock-Grpc/utils"
 "google.golang.org/grpc"
)

func main(){
  logger := utils.InitLogger(true)
  listen, err := net.Listen("tcp", ":8080")
  if err != nil {
    logger.Error("Failed to listen: %v", "error", err)
  }
  
  grpcServer := grpc.NewServer()
  stockpb.RegisterStockServiceServer(grpcServer, servers.NewStockServer(logger))
  logger.Info("Grpc Server is listening on port :8080")
  if err := grpcServer.Serve(listen); err != nil {
    logger.Error("Failed to serve: %v", "error", err)
  }
}

