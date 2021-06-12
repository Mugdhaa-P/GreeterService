package main

import (
    "context"
    "net/http"
    "os"
    "fmt"
    "GreeterService/rpc/greeter"
)

func main() {
    client := greeter.NewGreeterProtobufClient("http://localhost:8080", &http.Client{})

	//SET GREETING
	client.SetGreetingForUser(context.Background(), &greeter.Name{Message: "Jane"})

	//GET GREETING
    greeting, err := client.GetGreetingForUser(context.Background(), &greeter.Name{Message: "Jane"})
	
    if err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }
    fmt.Printf(" %+v", greeting)
}