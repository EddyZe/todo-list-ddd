package main

import (
	"todo/internal/adaptar/http/handlers/gin/auth/sessionauth"
	userHandler "todo/internal/adaptar/http/handlers/gin/user"
	"todo/internal/application/authmanager"
	userUseCase "todo/internal/application/user"
	"todo/internal/infrastructure/config"
	server "todo/internal/infrastructure/server/gin"
	userRepository "todo/internal/infrastructure/storage/inmemory/user"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	userRepo := userRepository.NewInMemory()

	userService := userUseCase.NewUseCase(userRepo)
	userH := userHandler.New(userService)

	authManager := authmanager.New(userRepo)
	authHandler := sessionauth.New(authManager, userService)

	s := server.New(&cfg.Server, authHandler, userH)

	if err := s.Run(); err != nil {
		panic(err)
	}
}
