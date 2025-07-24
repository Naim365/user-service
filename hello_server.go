// package main

// import (
// 	"context"

// 	pb "user-service/proto"
// )

// type Server struct {
// 	pb.UnimplementedHelloServiceServer
// }

// func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
// 	return &pb.HelloResponse{Message: "Hello from user-service to " + req.From}, nil
// }

// func startHelloServer() {
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterHelloServiceServer(s, &server{})
// 	log.Println("user-service gRPC server running on :50051")
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }