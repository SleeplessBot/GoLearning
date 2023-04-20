package examples

import (
	"context"
	"log"
	"time"

	"golearning/examples/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunGrpcClient() {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewGreetClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHi(ctx, &pb.SayHiRequest{Name: "Jack"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Get reply: %s", r.GetMessage())
}
