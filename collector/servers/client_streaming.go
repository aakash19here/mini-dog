package servers

import (
	"io"
	"log/slog"

	"github.com/aakash19here/mini-dog/proto/generated/minidogpb"
)

// grpc.ClientStreamingServer[LogEntryRequest, LogSummary]) error
func (s *CollectorService) SubmitLogs(stream minidogpb.LogCollector_SubmitLogsServer) error {
	var count int32 = 0

	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&minidogpb.LogSummary{
					TotalRecordsProcessed: count,
					IsSuccessful:          true,
				})
			}

			return err
		}

		slog.Info("Incoming log...", "req", req)

		count++
	}
}
