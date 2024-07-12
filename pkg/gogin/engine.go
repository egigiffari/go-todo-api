package gogin

import (
	"fmt"

	"github.com/egigiffari/go-todo-api.git/constants"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Listen()
}

type server struct {
	config *Config
	engine *gin.Engine
}

func New(config *Config) Server {
	if config.Env == constants.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery(), gin.CustomRecovery(Recovery))

	if config.Debug {
		engine.Use(gin.Logger())
	}

	return &server{
		config: config,
		engine: engine,
	}
}

func (s *server) Listen() {
	s.engine.Run(fmt.Sprintf(":%s", s.config.Port))
}
