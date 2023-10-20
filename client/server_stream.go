package main

import (
  "context"
  "io"
  "log"
  "fmt"
  "github.com/fatih/color"

  pb "github.com/yehdar/watchdogs/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
  // log.Printf("Server streaming RPC started")
  color.Cyan("Server streaming RPC started")
  stream, err := client.SayHelloServerStreaming(context.Background(), names)
  if err != nil {
    log.Fatalf("could not send names: %v", err)
  }

  for {
    message, err := stream.Recv()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatalf("error while streaming %v", err)
    }
    log.Println(message)
  }
  log.Printf("Streaming finished")
  fmt.Printf("\n")
}
