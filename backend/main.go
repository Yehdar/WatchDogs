package main

import (
	"context"
	"log"
	"net"

	pb "github.com/yehdar/watchdogs/backend/proto"
	"google.golang.org/grpc"
)

type Server struct {
	comments []string
}

func (s *Server) GetComments(ctx context.Context, req *pb.Empty) (*pb.Comments, error) {
	return &pb.Comments{Comments: s.comments}, nil
}

func (s *Server) AddComment(ctx context.Context, req *pb.CommentRequest) (*pb.Empty, error) {
	s.comments = append(s.comments, req.Text)
	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// pb.RegisterCommentServiceServer(grpcServer, &Server{}) 

	log.Println("gRPC server is running on :50051")
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
