package servers

import (
  "fmt"
  "os"
  "net/http"
  "encoding/json"
  "io"
  "log/slog"
  "github.com/joho/godotenv"
  )
 
type AlphaVantageResponse struct {
  GlobalQuote struct{
    Symbol string `json:"01. symbol"`
    Price string `json:"05. price"`
  } `json:"Global Quote"`
}

var (
  apiKey string
  baseUrl = "https://alphavantage.co/query"
  )
  

func init() {
    if err := godotenv.Load(); err != nil { 
      slog.Error("Error loading .env file", "error", err)
    }
    apiKey = os.Getenv("ALPHA_VANTAGE_API_KEY") 
}

func GetRealTimePrice(symbol string) (float64, error){
 if apiKey == "" {
		return 0, fmt.Errorf("Alpha Vantage API key not configured")
	}
	
  url := fmt.Sprintf("%s?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", baseUrl, symbol, apiKey)
  
  resp, err := http.Get(url)
  if err != nil {
    return 0, fmt.Errorf("Api Request Failed: %v", err)
  }
  defer resp.Body.Close()
  
  body, _ := io.ReadAll(resp.Body)
  var result AlphaVantageResponse
  
  if err := json.Unmarshal(body, &result); err != nil {
    return 0, fmt.Errorf("Failed to pass response", "%v", err)
  }
  
  if result.GlobalQuote.Price == "" {
    fmt.Errorf("No price data for: %v", symbol)
  }
  // to convert the price string to float
  var price float64 
  fmt.Sscanf(result.GlobalQuote.Price, "%f", &price)
  return price, nil
}