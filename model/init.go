package model

import (
	pb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
	"google.golang.org/grpc"
	"log"
	"os"
)

// MongoGrpcClient for connection auth grpc service
var MongoGrpcClient *grpc.ClientConn

// MysqlGrpcClient for connection admin grpc service
var MysqlGrpcClient *grpc.ClientConn

func init() {
	// get user service address
	AUTHADDRESS := "localhost:8083"
	mongoconn, err := grpc.Dial(AUTHADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	MongoGrpcClient = mongoconn

	MYSQLADDRESS := "localhost:8082"
	mysqlconn, err := grpc.Dial(MYSQLADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	MysqlGrpcClient = mysqlconn
}

// Project define
type Project struct {
	ID        string          `json:"id,omitempty"`
	User      string          `json:"user,omitempty"`
	Score     int32           `json:"score,omitempty"`
	Title     string          `json:"title,omitempty"`
	Equipment []*pb.Equipment `json:"equipment,omitempty"`
	Map       *pb.Map         `json:"map,omitempty"`
}

// User define
type User struct {
	Username string `json:"username,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Photo    string `json:"photo,omitempty"`
}

// Info define
type Info struct {
	UserNumber            int     `json:"usernumber,omitempty"`
	AllProjectNumber      int     `json:"allprojectnumber,omitempty"`
	NotRatedProjectNumber int     `json:"notratedprojectnumber,omitempty"`
	AllViews              int32   `json:"allviews,omitempty"`
	ViewsBeforeWeek       []int32 `josn:"viewsbeforeweek,omitempty"`
}

// View define
type View struct {
	Date   string
	Number int32
}
