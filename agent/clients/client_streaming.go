package clients

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"time"

	"github.com/aakash19here/mini-dog/proto/generated/minidogpb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Client_Streaming(client minidogpb.LogCollectorClient) {
	ctx := context.Background()
	messages := []string{
		"User logged in",
		"Payment initiated",
		"Video Uploaded",
	}

	//initialize stream connection
	stream, err := client.SubmitLogs(ctx)

	if err != nil {
		slog.Error("Could not initialize stream", "error", err)
	}

	for range 4 {
		logRequest := &minidogpb.LogEntryRequest{
			Id:      uuid.NewString(),
			AgentId: fmt.Sprint("Minidog-Frontend"),
			Level:   minidogpb.LogLevel_INFO,
			// 0 <= n < 3
			Msg:       fmt.Sprintf("%s", messages[rand.IntN(3)]),
			Timestamp: timestamppb.Now(),
		}

		if err := stream.Send(logRequest); err != nil {
			slog.Error("Failed to send a metric", "error", err)
			break
		}

		fmt.Println("Request Sent Successfully")
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	fmt.Println("\n📊 Summary:")
	fmt.Printf("Total Events: %d\n", res.TotalRecordsProcessed)
	fmt.Printf("Is successful: %v\n", res.IsSuccessful)

}
