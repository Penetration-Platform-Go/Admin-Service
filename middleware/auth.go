package middleware

import (
	"context"
	"time"

	pb "github.com/Penetration-Platform-Go/gRPC-Files/Auth-Service"
	"github.com/gin-gonic/gin"
)

// Auth Middleware
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		client := pb.NewAuthClient(AuthGrpcClient)
		r, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, err := client.GetUsernameByToken(r, &pb.Token{
			JWT: ctx.Request.Header.Get("Authentication"),
		})
		if err != nil || result.Username != "admin" {
			ctx.Status(403)
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}
