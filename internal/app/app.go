package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arsnazarenko/basketball-service/config"
	"github.com/arsnazarenko/basketball-service/pkg/httpserver"
	"github.com/arsnazarenko/basketball-service/pkg/logger"
	"github.com/arsnazarenko/basketball-service/pkg/postgres"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Postgresql
	_, err := postgres.New(cfg.Postgres.PostgresURL)
	if err != nil {
		panic(err)
	}
	l.Info("app.Run: PostgreSQL is started and connected by URL: %s", cfg.Postgres.PostgresURL)

	handler := gin.New()
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Content-Type", "Access-Control-Allow-Credentials", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	l.Info("server is start")

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))

		err = httpServer.Shutdown()
		if err != nil {
			l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
		}
	}

}
