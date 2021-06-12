package main

import (
    "net/http"

	"GreeterService/internal/greeterserver"
    pb "GreeterService/rpc/greeter"
)

func main() {
  server := &greeterserver.Server{} // implements Greeter interface
  twirpHandler := pb.NewGreeterServer(server)

  http.ListenAndServe(":8080", twirpHandler)
}