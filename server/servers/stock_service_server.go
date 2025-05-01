package servers

import "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"

type StockServiceServer struct {
  stockpb.UnimplementedStockServiceServer
}

