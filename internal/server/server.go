package server

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/cronui/cronui/ent"
	"github.com/cronui/cronui/internal/config"
)

func New(conf *config.Server, db *ent.Client) *Server {
	app := fiber.New(fiber.Config{
		IdleTimeout:           time.Second * 5,
		DisableStartupMessage: true,
	})

	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(logger.New())
	app.Use(cors.New())

	server := &Server{
		app:  app,
		conf: conf,
		db:   db,
	}

	app.Get("/install", server.handleInstallGet)
	app.Post("/install", server.handleInstallPost)
	app.Post("/auth/token", server.handleAuthToken)

	return server
}

type Server struct {
	app  *fiber.App
	conf *config.Server
	db   *ent.Client
}

func (s *Server) Listen() error {
	log.Printf("CronUI listening on %s\n", s.conf.GetListenAddr())
	go s.app.Listen(s.conf.GetListenAddr())
	return nil
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
