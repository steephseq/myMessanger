package config

import "api-gateway/internal/models"

func Load() *models.Config {
	return &models.Config{
		Address: ":8080",
		Services: []models.Service{
			{
				Name:  "auth-service",
				Port:  "50051",
				Paths: []string{"/login", "/register", "/"},
			},
			{
				Name: "user_service",
				Port: "50052",
				Paths: []string{"/createUser","/deleteUser","",
			/*
				{
					Name:  "profile-service",
					Port:  "50052",
					Paths: []string{"/profile", "/setProfileData", "/myProfileHP"},
				},
				{
					Name: "chats-service",
					Port: "50053",
					Paths: []string{
						"/chats",
						"/messages",
						"/createChat",
						"/create121Chat",
						"/chatExists",
						"/editMessage",
						"/addUsers",
						"/removeUserFromChat",
						"/newAdmin",
						"/howCanIDoMessage",
						"/deleteMessage",
						"/createEmptyMessage",
						"/howCanIDoUser",
						"/howCanIDoGroup",
					},
				},
				{
					Name: "ws-service",
					Port: "50054",
					Paths: []string{
						"/ws/onlineStatus",
						"/ws/profile",
						"/ws",
					},
				},
				{
					Name: "user-service",
				},
			*/
		},
	}
}
