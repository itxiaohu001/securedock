package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Host   string
	Port   int
}

func NewServer(host string, port int) *Server {
	return &Server{
		Engine: gin.Default(),
		Host:   host,
		Port:   port,
	}
}

func (s *Server) SetupRoutes() {
	// API版本v1
	v1 := s.Engine.Group("/api/v1")

	// 镜像扫描相关路由
	scan := v1.Group("/scan")
	{
		// TODO: 实现具体的handler
		scan.POST("/image", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "image scan endpoint"})
		})
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	return s.Engine.Run(addr)
}
