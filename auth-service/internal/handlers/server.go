package handler

import (
	"auth-service/internal/models"
	"auth-service/internal/services"
	"context"
	"log"
	"pkg/proto/authpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
	authService *services.AuthService
}

func NewAuthServer(authService *services.AuthService) *AuthServer {
	return &AuthServer{
		authService: authService,
	}
}

func (s *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	email := req.GetEmail()
	password := req.GetPassword()

	if email == "" || password == "" {
		return nil, status.Error(codes.InvalidArgument, "empty email or password")
	}

	userLogin := models.UserLogin{
		Email:    email,
		Password: password,
	}

	token, userID, err := s.authService.LoginUser(userLogin)
	if err != nil {
		log.Printf("Login: failed to login,error:%v", err)

		switch err.Error() {
		case "empty email or password":
			return nil, status.Error(codes.InvalidArgument, "email and password are required")
		case "invalid password or email":
			return nil, status.Error(codes.Unauthenticated, "invalid credentials")
		case "failed to get UserID":
			return nil, status.Error(codes.NotFound, "user not found")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &authpb.LoginResponse{
		Token:   token,
		UserId:  userID,
		Success: true,
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	err := s.authService.RegisterUser(models.UserRegister{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Username: req.GetUsername(),
	})
	if err != nil {
		log.Printf("Register: failed to register,error:%v", err)
		switch err.Error() {
		case "empty email or password":
			return nil, status.Error(codes.InvalidArgument, "email and password are required")
		case "user already exists":
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}
	return &authpb.RegisterResponse{
		Success: true,
	}, nil
}
