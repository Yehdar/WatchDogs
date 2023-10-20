package main

import (
  "context"
  "github.com/fatih/color"

  pb "github.com/yehdar/watchdogs/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error){
  color.Cyan("Unary RPC started")
  return &pb.HelloResponse{
    Message: "Hello my friend!",
  }, nil
}
