package servers

import (
	"context"
	"log/slog"

	"github.com/aakash19here/mini-dog/proto/generated/minidogpb"
)

func (s *CollectorService) SendLog(ctx context.Context, req *minidogpb.LogEntryRequest) (*minidogpb.LogEntryResponse, error) {
	// save the request inputs somewhere
	slog.Info("received event", "event", req)

	return &minidogpb.LogEntryResponse{
		Ok: true,
	}, nil
}
