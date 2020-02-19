package model

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "github.com/Penetration-Platform-Go/gRPC-Files/Mysql-Service"
)

// QueryAllUsers query from mysql
func QueryAllUsers() ([]User, error) {
	client := pb.NewMysqlClient(MysqlGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.QueryUsers(ctx, &pb.Condition{Value: []*pb.Value{}})
	if err != nil {
		return nil, err
	}
	var results []User
	for {
		feature, err := stream.Recv()
		if err == io.EOF || feature == nil {
			break
		}
		results = append(results, User{
			Username: feature.Username,
			Password: feature.Password,
			Email:    feature.Email,
			Nickname: feature.Nickname,
			Photo:    feature.Photo,
		})

	}
	return results, nil
}

// QueryUserByUsername query from mysql
func QueryUserByUsername(username string) ([]User, error) {
	client := pb.NewMysqlClient(MysqlGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.QueryUsers(ctx, &pb.Condition{Value: []*pb.Value{
		{Key: "username", Value: username},
	}})
	if err != nil {
		return nil, err
	}
	var results []User
	for {
		feature, err := stream.Recv()
		if err == io.EOF || feature == nil {
			break
		}
		results = append(results, User{
			Username: feature.Username,
			Password: feature.Password,
			Email:    feature.Email,
			Nickname: feature.Nickname,
			Photo:    feature.Photo,
		})

	}
	return results, nil
}

// DeleteUser delete user from mysql
func DeleteUser(username string) bool {
	DeleteProjectByUsername(username)
	client := pb.NewMysqlClient(MysqlGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := client.DeleteUser(ctx, &pb.Condition{Value: []*pb.Value{
		{Key: "username", Value: username},
	}})
	fmt.Println(result.Value)
	return result.IsVaild
}
