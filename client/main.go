package main

import (
  "log"

  pb "github.com/yehdar/watchdogs/proto"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
)

const (
  port = ":8080"
)

func main() {
  conn, err := grpc.Dial("localhost" + port, grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()

  client := pb.NewGreetServiceClient(conn)

  names := &pb.NamesList {
    Names: []string{"Bruce Wayne", "Clark Kent", "Dick Grayson"},
  }

  //  unary rpc
  //  callSayHello(client)
    
  // server streaming rpc
  // callSayHelloServerStream(client, names)
  
  // client streaming rpc
  // callSayHelloClientStream(client, names)

  // bidirectional streaming rpc
  callSayHelloBidirectionalStream(client, names)
}
