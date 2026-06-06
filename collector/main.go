package main

import (
	"log/slog"
	"net"

	"github.com/aakash19here/mini-dog/collector/servers"
	"github.com/aakash19here/mini-dog/proto/generated/minidogpb"
	"github.com/aakash19here/mini-dog/utils"
	"google.golang.org/grpc"
)

func main() {
	utils.InitLogger(true)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		slog.Error("tcp connection could not be established", "error", err)
		return
	}

	srv := grpc.NewServer()

	minidogpb.RegisterLogCollectorServer(srv, servers.NewCollectorService())

	if err := srv.Serve(lis); err != nil {
		slog.Error("failed to serve", "error", err)
		return
	}
}
