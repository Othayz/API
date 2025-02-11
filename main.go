package main

import (
"github.com/Othayz/API/Apis"
"github.com/labstack/gommon/log"
)

func main() {
  server := apis.NewServer()

  server.ConfigRoutes()

  if err := server.Start(); err != nil {
    log.Fatalf("Failed to start server: %s", err.Error())
  }

}



