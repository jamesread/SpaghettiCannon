package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/jamesread/SpaghettiCannon/internal/config"

	"connectrpc.com/connect"

	pb "github.com/jamesread/SpaghettiCannon/gen/proto"
)

type SpaghettiCannonApiService struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *SpaghettiCannonApiService {
	return &SpaghettiCannonApiService{
		cfg: cfg,
	}
}

func (api *SpaghettiCannonApiService) GetReadyz(ctx context.Context, req *connect.Request[pb.GetReadyzRequest]) (*connect.Response[pb.GetReadyzResponse], error) {
	res := connect.NewResponse(&pb.GetReadyzResponse{
		IsReady: true,
	})

	log.Info("GetReadyz called")

	return res, nil
}

/*
func Start(cfg *config.Config) {
	api := &SpaghettiCannonApiService{}

	lis, err := net.Listen("tcp", cfg.ListenAddressGrpc)

	if err != nil {
		log.Fatalf("Failed to listen - %v", err)
	}


	if err != nil {
		log.Fatalf("Could not start gRPC Server - %v", err)
	}
}
*/
