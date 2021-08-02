package server

import (
	"fmt"
	"log"

	mooc "github.com/alfonsovgs/go-hexagonal-architecture/internal"
	"github.com/alfonsovgs/go-hexagonal-architecture/internal/platform/server/handler/courses"
	"github.com/alfonsovgs/go-hexagonal-architecture/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	courseRepository mooc.CourseRepository
}

func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("#{host}:#{port}"),
	}

	srv.registerRoutes()

	return srv
}

func (s *Server) Run() error {
	log.Println("Server running", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
}
