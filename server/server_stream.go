package main

import (
  "log"
  "time"
  "github.com/fatih/color"
  
  pb "github.com/yehdar/watchdogs/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
  color.Cyan("\nServer streaming RPC started")
  log.Printf("got request with names : %v", req.Names)

  for _, name := range req.Names{
    res := &pb.HelloResponse {
      Message: "Hello " + name,
    }
    if err := stream.Send(res); err != nil {
      return err
    }
    time.Sleep(2 * time.Second)
  }
  return nil
}
