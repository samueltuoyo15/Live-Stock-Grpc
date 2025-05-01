package servers

import "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"

type StockServiceServer struct {
  logger Logger
  stockpb.UnimplementedStockServiceServer
}

