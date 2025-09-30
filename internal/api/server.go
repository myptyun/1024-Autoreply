package api

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Server encapsulates the HTTP server and dependencies.
// Additional fields (config, services) will be added as features grow.
type Server struct {
	app    *fiber.App
	logger *zap.Logger
	addr   string
}

func NewServer() *Server {
	logger, _ := zap.NewProduction()
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	addr := ":8080"
	if v := os.Getenv("APP_ADDR"); v != "" {
		addr = v
	}

	srv := &Server{
		app:    app,
		logger: logger,
		addr:   addr,
	}

	// routes
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "time": time.Now().UTC()})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Modbus RTU Server is running. See /healthz\n")
	})

	return srv
}

func (s *Server) Start() error {
	s.logger.Info("http server starting", zap.String("addr", s.addr))
	return s.app.Listen(s.addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := s.app.Shutdown(); err != nil {
		return fmt.Errorf("fiber shutdown: %w", err)
	}
	_ = s.logger.Sync()
	return nil
}