package main

import (
	"log/slog"
	"os"

	"github.com/aakash19here/mini-dog/agent/clients"
	"github.com/aakash19here/mini-dog/proto/generated/minidogpb"
	"github.com/aakash19here/mini-dog/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	utils.InitLogger(true)

	if len(os.Args) < 2 {
		slog.Error("Use case: go run agent/main.go <unary | server | client | bi>")
		return
	}

	// create a connection with grpc server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("could not connect with the grpc server", "error", err)
		return
	}

	defer conn.Close()

	// create a new client
	client := minidogpb.NewLogCollectorClient(conn)

	switch os.Args[1] {
	case "unary":
		clients.Unary(client)
	case "client":
		clients.Client_Streaming(client)
	case "bidi":
		clients.BidiStreaming(client)
	}
}
