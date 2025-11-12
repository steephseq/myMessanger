package main

import (
	"fmt"
	"log"
	"net"
	"pkg/proto/userpb"
	"user-service/config"
	"user-service/internal/handler"
	"user-service/internal/services"
	"user-service/storage"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env variables")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config")
	}

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)
	db, err := storage.NewPostgresRepo(connString)
	if err != nil {
		log.Fatalf("failed to create db,error:%v", err)
	}

	grpcServer := grpc.NewServer()

	userService := services.NewUserService(db)
	userServer := handler.NewUserServer(userService)

	userpb.RegisterUserServiceServer(grpcServer, userServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen,error:%v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve,error:%v", err)
	}

}
