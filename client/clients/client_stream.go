package clients

import (
	"context"
	"log/slog"
	"time"
	"github.com/samueltuoyo15/Live-Stock-Grpc/proto/generated/stockpb"
)

func (sc *StockClient) UploadStockHistory(stocks []*stockpb.StockRequest) {
	slog.Info("Starting client streaming RPC")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	stream, err := sc.client.UploadStockHistory(ctx)
	if err != nil {
		slog.Error("error calling upload stock history", "error", err)
		return 
	}
	
	for _, stock := range stocks {
		slog.Info("sending stock symbol", "symbol", stock.GetSymbol())
		if err := stream.Send(stock); err != nil {
			slog.Error("error sending stock", "error", err)
			return 
		}
		time.Sleep(500 * time.Millisecond)
	}
	
	res, err := stream.CloseAndRecv()
	if err != nil {
		slog.Error("error receiving response", "error", err)
		return 
	}
	
	slog.Info("upload summary",
		"message", res.GetMessage(),
		"count", res.GetCount())
}