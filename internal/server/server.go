package server

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"

	"github.com/cronui/cronui/internal/config"
)

func New(conf *config.Server, db *gorm.DB) *Server {
	app := fiber.New(fiber.Config{
		IdleTimeout:           time.Second * 5,
		DisableStartupMessage: true,
	})

	app.Use(compress.New())
	app.Use(logger.New())
	app.Use(cors.New())

	server := &Server{
		app:    app,
		listen: conf.GetListenAddr(),
		db:     db,
	}

	app.Get("/install", server.handleInstallGet)
	app.Post("/install", server.handleInstallPost)

	return server
}

type Server struct {
	app    *fiber.App
	listen string
	db     *gorm.DB
}

func (s *Server) Listen() error {
	log.Printf("GoMVN self-hosted repository listening on %s\n", s.listen)
	go s.app.Listen(s.listen)
	return nil
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
