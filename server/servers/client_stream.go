package servers


import (
  "io"
  "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
  )
  
func(s *StockServer) UploadStockHistory(stream stockpb.StockService_UploadStockHistoryServer) error {
  count := 0
  for {
    req, err := stream.Recv()
    if err == io.EOF{
      s.Logger.Info("Received %d stock history entries", "count", count)
      return stream.SendAndClose(&stockpb.StockSummary{
        Message: "Upload Complete",
        Count: int32(count),
      })
    }
    if err != nil {
      return err
    }
    s.Logger.Info("Received stock", "symbol", req.GetSymbol())
    count++
  }
}

