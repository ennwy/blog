package server

import (
	"context"

	"github.com/ennwy/blog/internal/app"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Ctx    context.Context
	App    *fiber.App
	Queue  app.Queue
	Addres string
}

var _ app.Server = (*Server)(nil)

func New() *Server {
	server := &Server{
		App: fiber.New(),
	}

	server.App.Post("/posts", server.Posts)

	return server
}

func (s *Server) Start() error {
	return s.App.Listen(s.Addres)
}

func (s *Server) Stop() error {
	return s.App.Server().Shutdown()
}
