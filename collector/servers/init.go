package servers

import "github.com/aakash19here/mini-dog/proto/generated/minidogpb"

type CollectorService struct {
	minidogpb.UnimplementedLogCollectorServer
}

func NewCollectorService() *CollectorService {
	return &CollectorService{}
}
