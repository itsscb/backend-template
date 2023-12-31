package gapi

import (
	"fmt"
	"log/slog"
	"os"

	db "github.com/itsscb/backend-template/bff/db/sqlc"
	"github.com/itsscb/backend-template/bff/pb"
	"github.com/itsscb/backend-template/bff/token"
	"github.com/itsscb/backend-template/bff/util"
)

// Server serves gRPC requests for backend service
type Server struct {
	pb.UnimplementedBackendServer
	store      db.Store
	config     util.Config
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenPrivateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	logLevel := slog.LevelError
	var logSource bool
	if config.Environment == "development" {
		logLevel = slog.LevelDebug
		logSource = true
	}

	opts := slog.HandlerOptions{
		Level:     logLevel,
		AddSource: logSource,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &opts))

	if config.LogOutput == "json" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &opts))
	}

	slog.SetDefault(logger)

	return server, nil
}
