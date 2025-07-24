package main

// func callMainServiceHello(w http.ResponseWriter, r *http.Request) {
// 	mainServiceHost := os.Getenv("MAIN_SERVICE_HOST")
// 	if mainServiceHost == "" {
// 		mainServiceHost = "localhost:50052"
// 	}
// 	conn, err := grpc.Dial(mainServiceHost, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	client := pb.NewHelloServiceClient(conn)

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	resp, err := client.SayHello(ctx, &pb.HelloRequest{From: "user-service"})
// 	if err != nil {
// 		log.Printf("could not greet: %+v", err)
// 		return
// 	}
// 	log.Printf("Greeting from main-service: %s", resp.Message)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(resp)
// }