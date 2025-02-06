package main

import (
"github.com/Othayz/API/Apis"
"log"
)

func main() {
  server := apis.NewServer()

  server.ConfigRoutes()

  if err := server.Start(); err != nil {
    log.Fatal(err)
  }

}



