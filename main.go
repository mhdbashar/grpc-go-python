package main

import (
    "context"
    "fmt"
    "log"

    "github.com/mhdbashar/grpc-go-python/greet"
    "google.golang.org/grpc"

)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := greet.NewGreeterClient(conn)

    name := "Mhdbashar"
    response, err := client.SayHello(context.Background(), &greet.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("Error calling SayHello: %v", err)
    }

    fmt.Printf("Response: %s\n", response.Message)
}
