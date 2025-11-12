package authlib

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if isPublicMethod(info.FullMethod) {
		return handler(ctx, req)
	}
	token, err := extractTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	userID, err := GetUserIDFromToken(token)
	if err != nil {
		return nil, err
	}
	newCtx := context.WithValue(ctx, userIDKey{}, userID)
	return handler(newCtx, req)
}

func isPublicMethod(fullMethod string) bool {
	publicMethods := map[string]bool{
		"auth.AuthService/Register":    true,
		"auth.AuthService/Login":       true,
		"/grpc.health.v1.Health/Check": true,
		"/grpc.health.v1.Health/Watch": true,
	}
	return publicMethods[fullMethod]
}

func extractTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata not found")
	}
	authHeaders := md.Get("authorization")
	if len(authHeaders) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization header not found")
	}
	token := strings.TrimPrefix(authHeaders[0], "Bearer ")
	token = strings.TrimSpace(token)
	if token == "" {
		return "", status.Errorf(codes.Unauthenticated, "authorization header not found")
	}
	return token, nil
}

type userIDKey struct{}
