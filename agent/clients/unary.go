package clients

import (
	"context"
	"log/slog"
	"time"

	"github.com/aakash19here/mini-dog/proto/generated/minidogpb"
	"github.com/google/uuid"
)

func Unary(client minidogpb.LogCollectorClient) {
	log := &minidogpb.LogEntryRequest{
		Id:      uuid.New().String(),
		AgentId: "agent_01",
		Level:   minidogpb.LogLevel_DEBUG,
		Msg:     "Just checking in",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err := client.SendLog(ctx, log)

	if err != nil {
		slog.Error("could not send logs to the server", "error", err)
	}

	slog.Info("Log ingested")

}
