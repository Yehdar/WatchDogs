package main

import (
  "io"
  "log"
  "github.com/fatih/color"

  pb "github.com/yehdar/watchdogs/proto"
)

func (s * helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
  color.Cyan("\nClient streaming RPC started")
  var messages []string
  for {
    req, err := stream.Recv()
    if err == io.EOF {
      return stream.SendAndClose(&pb.MessagesList{Messages: messages})
    }
    if err != nil {
      return err
    }
    log.Printf("Got request with name: %v", req.Name)
    messages = append(messages, "Hello", req.Name)
  }
}
