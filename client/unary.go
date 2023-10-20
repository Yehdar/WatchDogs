package main

import (
  "context"
  "log"
  "time"
  "fmt"
  "github.com/fatih/color"

  pb "github.com/yehdar/watchdogs/proto"
)

func callSayHello(client pb.GreetServiceClient) {
  // log.Printf("Unary RPC started")
  color.Cyan("Unary RPC started:")
  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  res, err := client.SayHello(ctx, &pb.NoParam{})
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }
  log.Printf("%s", res.Message)
  fmt.Printf("\n")
}
