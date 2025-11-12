package main

import "context"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcAuthServerEndpoint := "auth-service:50051"
	grpcUserServerEndpoint := "user-service:50052"
}
