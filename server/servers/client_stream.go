package servers


import (
  "io"
  "github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
  )
  
func(s s.StockServer) UploadStockHistory(stream stockpb.StockService_UploadStockHistoryServer) error {
  count := 0
  for {
    req, err := streamRecv()
    if err == io.EOF{
      s.logger.Infof("Received %d stock history entries", count)
      return StreamAndClose(&stockpb.StockSummary{
        Message: "Upload Complete",
        Count: int32(count),
      })
    }
    if err != nil {
      return err
    }
    s.logger.Infof("Received stock: %s at price %.2f", req.Symbol, req.Price)
    count++
  }
}

