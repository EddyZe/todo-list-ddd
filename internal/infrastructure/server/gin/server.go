package gin

import (
	"fmt"
	"net/http"
	"todo/internal/infrastructure/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(*gin.RouterGroup)
}

type Server struct {
	cfg    *config.ServerConfig
	engine *gin.Engine
}

func New(cfg *config.ServerConfig, handlers ...Handler) *Server {
	store := cookie.NewStore([]byte(cfg.CookieSecret))
	store.Options(sessions.Options{
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	})

	engine := gin.Default()
	engine.Use(sessions.Sessions("todo-list", store))
	engine.Use(gin.Recovery())

	apiV1 := engine.Group("/api/v1")

	for _, handler := range handlers {
		handler.Register(apiV1)
	}

	return &Server{
		engine: engine,
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%v", s.cfg.Port)
	return s.engine.Run(addr)
}
