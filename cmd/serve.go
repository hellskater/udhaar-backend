package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/hellskater/udhaar-backend/internal/repository"
	internalGorm "github.com/hellskater/udhaar-backend/internal/repository/gorm"
	"github.com/hellskater/udhaar-backend/internal/router"
	"github.com/hellskater/udhaar-backend/internal/service"
	"github.com/hellskater/udhaar-backend/internal/service/books"
	"github.com/hellskater/udhaar-backend/pkg/utils/gormzap"
	"github.com/labstack/echo"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func serveCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "serve",
		Short: "Start Udhaar API server",
		Run: func(_ *cobra.Command, _ []string) {
			//----------------- Logger -----------------
			logger := getLogger()
			defer logger.Sync()

			//----------------- Database -----------------
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

			//----------------- Repository -----------------
			logger.Info("setting up repository...")
			repo, _, err := internalGorm.NewGormRepository(engine, logger, true)
			if err != nil {
				logger.Fatal("failed to initialize repository", zap.Error(err))
			}
			logger.Info("repository was set up")

			//----------------- Server Creation -----------------
			server, err := newServer(engine, repo, logger, &config)
			if err != nil {
				logger.Fatal("failed to create server", zap.Error(err))
			}

			//----------------- Start Server -----------------
			go func() {
				if err := server.Start(fmt.Sprintf(":%d", config.Port)); err != nil {
					logger.Info("shutting down the server")
				}
			}()

			logger.Info("udhaar started")

			//----------------- Wait SIGINT -----------------
			waitSIGINT()
			logger.Info("udhaar shutting down...")

			//----------------- Shutdown Server -----------------
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

func newServer(db *gorm.DB, repo repository.Repository, logger *zap.Logger, c *Config) (*Server, error) {
	booksService := books.NewService(repo, logger)
	services := &service.Services{
		Books: booksService,
	}
	routerConfig := provideRouterConfig(c)
	echo := router.Setup(db, repo, services, logger, routerConfig)
	server := &Server{
		L:      logger,
		SS:     services,
		Router: echo,
		Repo:   repo,
	}
	return server, nil
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
