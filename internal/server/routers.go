package server

import (
	"net/http"

	"github.com/ennwy/blog/internal/blog"
	"github.com/gofiber/fiber/v2"

	"github.com/goccy/go-json"
)

func (s *Server) Posts(ctx *fiber.Ctx) error {
	post := blog.Post{
		Author: ctx.Params("author"),
		Title:  ctx.Params("title"),
		Body:   ctx.Params("author"),
	}

	bytesPost, err := json.Marshal(post)
	if err != nil {
		return err
	}

	if err := s.Queue.Push(s.Ctx, "queue:new-post", bytesPost); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity)
		return err
	}

	return nil
}
