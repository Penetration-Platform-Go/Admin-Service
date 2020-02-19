package model

import (
	"context"
	"io"
	"time"

	pb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
)

// QueryAllProjects Query Project By username
func QueryAllProjects() ([]Project, error) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.QueryProjects(ctx, &pb.Condition{Value: []*pb.Value{}})
	if err != nil {
		return nil, err
	}
	var results []Project
	for {
		feature, err := stream.Recv()
		if err == io.EOF || feature == nil {
			break
		}
		results = append(results, Project{
			ID:    feature.Id,
			User:  feature.User,
			IP:    feature.Ip,
			Title: feature.Title,
			Score: feature.Score,
			Map:   feature.Map,
		})

	}
	return results, nil
}

// DeleteProjectByID delete project
func DeleteProjectByID(id string) (bool, string) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, _ := client.DeleteProject(ctx, &pb.Condition{Value: []*pb.Value{
		{Key: "_id", Value: id},
	},
	})
	return result.IsVaild, result.Value
}

// DeleteProjectByUsername delete project
func DeleteProjectByUsername(username string) (bool, string) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, _ := client.DeleteProject(ctx, &pb.Condition{Value: []*pb.Value{
		{Key: "user", Value: username},
	},
	})
	return result.IsVaild, result.Value
}

// EvaluateProject set score for project
func EvaluateProject(id string, score int32) (bool, string) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, _ := client.UpdateProject(ctx, &pb.UpdateMessage{
		Condition: &pb.Condition{
			Value: []*pb.Value{
				{Key: "_id", Value: id},
			},
		},
		Key: []string{
			"score",
		},
		Value: &pb.ProjectInformation{
			Score: score,
		},
	})
	return result.IsVaild, result.Value
}

// QueryProjectByTitle Query Project By Title
func QueryProjectByTitle(title string) ([]Project, error) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.QueryProjects(ctx, &pb.Condition{Value: []*pb.Value{
		{Key: "title", Value: title},
	}})
	if err != nil {
		return nil, err
	}
	var results []Project
	for {
		feature, err := stream.Recv()
		if err == io.EOF || feature == nil {
			break
		}
		results = append(results, Project{
			ID:    feature.Id,
			User:  feature.User,
			IP:    feature.Ip,
			Title: feature.Title,
			Score: feature.Score,
			Map:   feature.Map,
		})

	}
	return results, nil
}

// QueryProjectsByUsername Query Project By username
func QueryProjectsByUsername(username string) ([]Project, error) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.QueryProjects(ctx, &pb.Condition{Value: []*pb.Value{
		{Key: "user", Value: username},
	}})
	if err != nil {
		return nil, err
	}
	var results []Project
	for {
		feature, err := stream.Recv()
		if err == io.EOF || feature == nil {
			break
		}
		results = append(results, Project{
			ID:    feature.Id,
			User:  feature.User,
			IP:    feature.Ip,
			Title: feature.Title,
			Score: feature.Score,
			Map:   feature.Map,
		})

	}
	return results, nil
}
