package main

import (
	"todo/internal/adaptar/http/handlers/gin/auth/sessionauth"
	taskHandler "todo/internal/adaptar/http/handlers/gin/task"
	userHandler "todo/internal/adaptar/http/handlers/gin/user"
	"todo/internal/application/authmanager"
	taskUseCase "todo/internal/application/task"
	userUseCase "todo/internal/application/user"
	"todo/internal/infrastructure/config"
	server "todo/internal/infrastructure/server/gin"
	taskRepository "todo/internal/infrastructure/storage/inmemory/task"
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

	taskRepo := taskRepository.NewInMemory()
	taskService := taskUseCase.NewUseCase(taskRepo, userService)
	th := taskHandler.NewHandler(taskService)

	s := server.New(&cfg.Server, authHandler, userH, th)

	if err := s.Run(); err != nil {
		panic(err)
	}
}
