package servers

import (
	"io"
	"log/slog"

	"github.com/aakash19here/mini-dog/proto/generated/minidogpb"
)

func (s *CollectorService) StreamLogs(stream minidogpb.LogCollector_StreamLogsServer) error {
	ctx := stream.Context()

	for {
		select {
		//get clarity on this
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		_, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				slog.Info("exit")
				return nil
			}
			slog.Error("something went wrong receiving data", "error", err)
			continue
		}

		resp := &minidogpb.LogEntryResponse{Ok: true}

		if err := stream.Send(resp); err != nil {
			slog.Error("send error", "error", err)
		}
	}
}
