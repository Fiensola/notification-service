package grpc

import (
	"context"
	"log"
	"net"

	pb "github.com/fiensola/notification-service/gen/go/notification"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *Server) SendNotification(
	ctx context.Context,
	req *pb.SendRequest,
) (*pb.SendResponse, error) {
	log.Printf("Received: %v", req)
	return &pb.SendResponse{
		Success:   true,
		MessageId: "msg-123",
	}, nil
}

func StartGrpcServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(grpcServer, &Server{})

	log.Printf("gRPC server starting on %s", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
