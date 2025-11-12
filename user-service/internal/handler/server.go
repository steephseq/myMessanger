package handler

import (
	"context"
	"pkg/proto/userpb"
	"time"
	"user-service/internal/models"
	"user-service/internal/services"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
	userService *services.UserService
}

func NewUserServer(userService *services.UserService) *UserServer {
	return &UserServer{
		userService: userService,
	}
}

func (s *UserServer) CreateUserHandler(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	lastSeen := time.Now()
	email := req.GetEmail()
	password := req.GetPassword()
	name := req.GetName()
	username := req.GetUsername()
	avatarURL := req.GetAvatarUrl()
	bio := req.GetBio()

	user := models.User{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  password,
		Name:      name,
		Username:  username,
		LastSeen:  lastSeen,
		AvatarURL: avatarURL,
		Bio:       bio,
	}

	if err := s.userService.CreateUser(user); err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}
	return &userpb.CreateUserResponse{
		Success: true,
		UserId:  user.ID,
	}, nil
}
