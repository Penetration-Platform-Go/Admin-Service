package model

import (
	"context"
	"io"
	"time"

	pb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
)

// QueryViews handle
func QueryViews() ([]View, error) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.QueryViews(ctx, &pb.Condition{Value: []*pb.Value{}})
	if err != nil {
		return nil, err
	}
	var result []View
	for {
		feature, err := stream.Recv()
		if err == io.EOF || feature == nil {
			break
		}
		result = append(result, View{
			Date:   feature.Key,
			Number: feature.Value,
		})
	}
	return result, err
}
