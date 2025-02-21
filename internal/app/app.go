package app

import (
	"fmt"
	"github.com/osamikoyo/IM-catalog/internal/config"
	"github.com/osamikoyo/IM-catalog/internal/data"
	"github.com/osamikoyo/IM-catalog/internal/server"
	"github.com/osamikoyo/IM-catalog/pkg/loger"
	"github.com/osamikoyo/IM-catalog/pkg/proto/pb"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	loger loger.Logger
	cfg *config.Config
	server *server.Server
	gRPC *grpc.Server
}

func Init() (*App, error) {
	cfg, err := config.Load("config.toml")
	if err != nil {
		return nil, err
	}

	storage, err := data.New(cfg)
	if err != nil {
		return nil, err
	}

	serve := &server.Server{
		Storage: storage,
	}

	return &App{
		loger: loger.New(),
		gRPC: grpc.NewServer(),
		server: serve,
		cfg: cfg,
	}, nil
}

func (a *App) Run() error {
	pb.RegisterCatalogServiceServer(a.gRPC, a.server)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port))
	if err != nil {
		return err
	}

	a.loger.Info().Msg("Catalog service starting...")
	a.loger.Info().Str("addr", lis.Addr().String()).Msg("service started! :3")

	return a.gRPC.Serve(lis)
}