package api

import (
	"context"
	"database/sql"
	"fmt"

	"connectrpc.com/connect"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	pb "github.com/jamesread/SpaghettiCannon/gen/SpaghettiCannon/clientapi/v1"
	"github.com/jamesread/SpaghettiCannon/internal/buildinfo"
	"github.com/jamesread/SpaghettiCannon/internal/config"
	"github.com/jamesread/SpaghettiCannon/internal/store"
)

type SpaghettiCannonApiService struct {
	cfg   *config.Config
	store store.Store
}

func NewServer(cfg *config.Config) *SpaghettiCannonApiService {
	api := &SpaghettiCannonApiService{cfg: cfg}

	db, err := openDB(cfg)
	if err != nil {
		log.WithError(err).Fatal("database connection failed")
	}

	st := store.NewMySQL(db)
	api.store = st

	ctx := context.Background()
	if err := assertMigration(ctx, st, config.RequiredMigration); err != nil {
		log.Fatal(err)
	}
	logDatabaseMigration(ctx, st)
	if err := ensureDefaultCvars(ctx, st, cfg); err != nil {
		log.WithError(err).Fatal("cvars")
	}

	return api
}

func openDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User, cfg.Database.Pass, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	log.WithFields(log.Fields{
		"name": cfg.Database.Name,
		"host": cfg.Database.Host,
	}).Infof("Connected to database")

	return db, nil
}

func assertMigration(ctx context.Context, st store.Store, required string) error {
	ok, err := st.HasMigration(ctx, required)
	if err != nil {
		return fmt.Errorf("connected to the database, but the migrations table could not be queried. run sql-migrate up: %w", err)
	}
	if ok {
		return nil
	}
	latest, _ := st.LatestMigration(ctx)
	display := latest
	if display == "" {
		display = "null"
	}
	return fmt.Errorf("requires database version %s but the database is at version %s; run database migrations", required, display)
}

func logDatabaseMigration(ctx context.Context, st store.Store) {
	latest, err := st.LatestMigration(ctx)
	if err != nil {
		log.Warnf("Could not read current database migration: %v", err)
		return
	}
	if latest == "" {
		latest = "none"
	}
	log.Infof("Database migration: %s", latest)
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

func (api *SpaghettiCannonApiService) Init(ctx context.Context, req *connect.Request[pb.InitRequest]) (*connect.Response[pb.InitResponse], error) {
	showVersion := api.showVersionNumber(ctx)
	currentVersion := ""
	availableVersion := ""
	showNew := false

	if showVersion {
		currentVersion = buildinfo.Version
		showNew = api.showNewVersions(ctx)
		if showNew {
			availableVersion = "none"
		}
	}

	res := connect.NewResponse(&pb.InitResponse{
		ShowFooter:        api.showFooter(ctx),
		ShowNewVersions:   showVersion && showNew,
		AvailableVersion:  availableVersion,
		CurrentVersion:    currentVersion,
		PageTitle:         api.siteTitle(ctx),
		ShowVersionNumber: showVersion,
	})

	log.Info("Init called")

	return res, nil
}
