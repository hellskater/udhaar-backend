package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/hellskater/udhaar-backend/internal/repository"
	"github.com/hellskater/udhaar-backend/internal/repository/gorm"
	"github.com/hellskater/udhaar-backend/internal/service"
	"github.com/hellskater/udhaar-backend/pkg/utils/gormzap"
	"github.com/labstack/echo"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func serveCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "serve",
		Short: "Start Udhaar API server",
		Run: func(_ *cobra.Command, _ []string) {
			// Logger
			logger := getLogger()
			defer logger.Sync()

			// Database
			logger.Info("connecting database...")
			engine, err := config.getDatabase()
			if err != nil {
				logger.Fatal("failed to connect database", zap.Error(err))
			}

			engine.Logger = gormzap.New(logger.Named("gorm"))
			db, err := engine.DB()
			if err != nil {
				logger.Fatal("failed to get *sql.DB", zap.Error(err))
			}
			defer db.Close()
			logger.Info("database connection was established")

			// Repository
			logger.Info("setting up repository...")
			repo, _, err := gorm.NewGormRepository(engine, logger, true)
			if err != nil {
				logger.Fatal("failed to initialize repository", zap.Error(err))
			}
			logger.Info("repository was set up")

			// Server creation
			server, err := newServer(engine, repo, logger, &config)
			if err != nil {
				logger.Fatal("failed to create server", zap.Error(err))
			}

			go func() {
				if err := server.Start(fmt.Sprintf(":%d", config.Port)); err != nil {
					logger.Info("shutting down the server")
				}
			}()

			logger.Info("udhaar started")
			waitSIGINT()
			logger.Info("udhaar shutting down...")

			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				logger.Warn("abnormal shutdown", zap.Error(err))
			}
			logger.Info("udhaar shutdown")
		},
	}
	return &cmd
}

type Server struct {
	L      *zap.Logger
	SS     *service.Services
	Router *echo.Echo
	Repo   repository.Repository
}

func (s *Server) Start(address string) error {
	return s.Router.Start(address)
}

func (s *Server) Shutdown(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := s.Router.Shutdown(ctx)
		s.L.Info("Router shutdown")
		return err
	})
	return eg.Wait()
}
