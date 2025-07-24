package main

import (
	"context"
	"log"
	"net"
	"time"
	pb "user-service/proto"

	"google.golang.org/grpc"

	"encoding/json"
	"net/http"
	"os"
)

// func callMainHandler(w http.ResponseWriter, r *http.Request) {
// 	mainServiceURL := os.Getenv("MAIN_SERVICE_URL")
// 	if mainServiceURL == "" {
// 		http.Error(w, "MAIN_SERVICE_URL not set", http.StatusInternalServerError)
// 		return
// 	}

// 	resp, err := http.Get(fmt.Sprintf("%s/api/ping", mainServiceURL))
// 	if err != nil {
// 		http.Error(w, "Failed to call main-service: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		http.Error(w, "Failed to read main-service response: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	var mainServiceResp map[string]interface{}
// 	if err := json.Unmarshal(body, &mainServiceResp); err != nil {
// 		http.Error(w, "Failed to parse main-service response: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	response := map[string]interface{}{
// 		"message": "Successfully called main-service",
// 		"data":    mainServiceResp,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

func callMainServiceHello(w http.ResponseWriter, r *http.Request) {
	mainServiceHost := os.Getenv("MAIN_SERVICE_HOST")
	if mainServiceHost == "" {
		mainServiceHost = "localhost:50052"
	}
	conn, err := grpc.Dial(mainServiceHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{From: "user-service"})
	if err != nil {
		log.Printf("could not greet: %+v", err)
		return
	}
	log.Printf("Greeting from main-service: %s", resp.Message)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello from user-service to " + req.From}, nil
}


func main() {
	// database, err := db.Connect()
	// if err != nil {
	// 	log.Fatal("DB connection failed:", err)
	// }

	
	// repo := repository.NewUserRepository(database)
	// svc := service.NewUserService(repo)
	// handler := handler.NewUserHandler(svc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &Server{})

	// Start gRPC server in a goroutine
	go func() {
		log.Println("gRPC server running on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// Set up HTTP handler
	http.HandleFunc("/api/hello", callMainServiceHello)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// Start HTTP server
	log.Printf("user-service running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}

