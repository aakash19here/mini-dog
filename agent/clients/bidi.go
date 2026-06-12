package clients

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand/v2"
	"time"

	"github.com/aakash19here/mini-dog/proto/generated/minidogpb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func BidiStreaming(client minidogpb.LogCollectorClient) {
	stream, err := client.StreamLogs(context.Background())

	if err != nil {
		slog.Error("could not initiate stream", "error", err)
		return
	}

	ctx := stream.Context()
	done := make(chan struct{})
	messages := []string{
		"User logged in",
		"Payment initiated",
		"Video Uploaded",
	}

	//sending
	go func() {
		for range 3 {
			logRequest := &minidogpb.LogEntryRequest{
				Id:      uuid.NewString(),
				AgentId: fmt.Sprint("Minidog-Frontend"),
				Level:   minidogpb.LogLevel_INFO,
				// 0 <= n < 3
				Msg:       fmt.Sprintf("%s", messages[rand.IntN(3)]),
				Timestamp: timestamppb.Now(),
			}

			if err := stream.Send(logRequest); err != nil {
				slog.Error("can not send", "error", err)
			}

			slog.Info("sent")
			time.Sleep(3 * time.Second)
		}

		close(done)

		if err := stream.CloseSend(); err != nil {
			log.Println("is it here ????", err)
		}
	}()

	//receiving
	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				close(done)
				return
			}

			if err != nil {
				slog.Error("can not receive", "error", err)
			}

			fmt.Println("Received from server", resp.Ok)
		}
	}()

	// third goroutine closes done channel
	// if context is done

	go func() {
		// get some clarity one this
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
	}()

	<-done
}
