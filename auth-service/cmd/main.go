package main

import (
	"auth-service/config"
	"auth-service/internal/clients"
	handler "auth-service/internal/handlers"
	"auth-service/internal/services"
	"auth-service/internal/token"
	"auth-service/storage"
	"fmt"
	"net"
	"pkg/proto/authpb"

	"log"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Auth-service: failed to load env variables")
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Auth-service: failed to load config")
	}
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)
	repo, err := storage.NewPostgresRepo(connString)
	if err != nil {
		log.Fatalf("Auth-service: failed to create db,error:%v", err)
	}

	jwtManager, err := token.NewJWTManager()
	if err != nil {
		log.Fatalf("Auth-service: failed to create jwt manager,error:%v", err)
	}

	grpcServer := grpc.NewServer()

	clientsRegistry, err := clients.NewRegistry(cfg)
	if err != nil {
		log.Fatalf("Auth-service: failed to init registry,error:%v", err)
	}
	defer func() {
		if err := clientsRegistry.Close(); err != nil {
			log.Printf("Auth-service: failed to close registry,error:%v", err)
		}
	}()

	authService := services.NewAuthService(repo, jwtManager, clientsRegistry.User)
	if authService == nil {
		log.Fatalf("Auth-service: failed to create auth service,error:%v", err)
	}

	authServer := handler.NewAuthServer(authService)
	authpb.RegisterAuthServiceServer(grpcServer, authServer)

	lis, err := net.Listen("tcp", cfg.Server.Port)
	if err != nil {
		log.Fatalf("Auth-service: failed to listen,error:%v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Auth-service: failed to serve,error:%v", err)
	}
}
