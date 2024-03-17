package api

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/jamesread/SpaghettiCannon/internal/config"

	"net"

	pb "github.com/jamesread/SpaghettiCannon/gen/grpc"
	"google.golang.org/grpc"
)

type SpaghettiCannonApiService struct {
	cfg *config.Config
}

func (api *SpaghettiCannonApiService) GetReadyz(ctx context.Context, req *pb.GetReadyzRequest) (*pb.GetReadyzResponse, error) {
	ret := &pb.GetReadyzResponse{
		IsReady: true,
	}

	return ret, nil
}

func Start(cfg *config.Config) {
	api := &SpaghettiCannonApiService{}

	lis, err := net.Listen("tcp", cfg.ListenAddressGrpc)

	if err != nil {
		log.Fatalf("Failed to listen - %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterSpaghettiCannonApiServiceServer(grpcServer, api)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Could not start gRPC Server - %v", err)
	}
}
