package api

import (
	"context"
	"github.com/jamesread/SpaghettiCannon/internal/config"
	"github.com/jamesread/SpaghettiCannon/internal/buildinfo"
	log "github.com/sirupsen/logrus"

	"fmt"

	"connectrpc.com/connect"

	gormv2logrus "github.com/thomas-tacquet/gormv2-logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	pb "github.com/jamesread/SpaghettiCannon/gen/SpaghettiCannon/clientapi/v1"
)

type SpaghettiCannonApiService struct {
	cfg *config.Config
	dsn string
	db  *gorm.DB
}

func NewServer(cfg *config.Config) *SpaghettiCannonApiService {
	api := &SpaghettiCannonApiService{
		cfg: cfg,
		dsn: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.Database.User, cfg.Database.Pass, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name),
		db:  nil,
	}

	api.connectDB()

	return api
}

func (api *SpaghettiCannonApiService) connectDB() {
	gormLogger := gormv2logrus.NewGormlog(gormv2logrus.WithLogrus(log.New()))
	gormLogger.LogMode(logger.Error)

	gormConfig := &gorm.Config{
		Logger:                 gormLogger,
		CreateBatchSize:        1500,
		SkipDefaultTransaction: true,
	}

	db, err := gorm.Open(mysql.Open(api.dsn), gormConfig)

	if err == nil {
		log.WithFields(log.Fields{
			"name": api.cfg.Database.Name,
			"host": api.cfg.Database.Host,
		}).Infof("Connected to database")
	}

	api.db = db
}

func (api *SpaghettiCannonApiService) GetReadyz(ctx context.Context, req *connect.Request[pb.GetReadyzRequest]) (*connect.Response[pb.GetReadyzResponse], error) {
	res := connect.NewResponse(&pb.GetReadyzResponse{
		IsReady: true,
		Version: buildinfo.Version,
	})

	log.Info("GetReadyz called")

	return res, nil
}

func (api *SpaghettiCannonApiService) AddUpdate(ctx context.Context, req *connect.Request[pb.AddUpdateRequest]) (*connect.Response[pb.AddUpdateResponse], error) {
	res := connect.NewResponse(&pb.AddUpdateResponse{})

	log.WithFields(log.Fields{
		"content": req.Msg.Content,
	}).Info("AddUpdate called")

	return res, nil
}
