package middleware

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

// AuthGrpcClient for connection auth grpc service
var AuthGrpcClient *grpc.ClientConn

func init() {
	// get user service address
	AUTHADDRESS := "localhost:8081"
	conn, err := grpc.Dial(AUTHADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	AuthGrpcClient = conn
}
