package utils 

import (
  "os"
  "log/slog"
  )
  
var Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
