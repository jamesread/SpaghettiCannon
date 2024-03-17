package httpservers

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"

	gw "github.com/jamesread/SpaghettiCannon/gen/grpc"

	config "github.com/jamesread/SpaghettiCannon/internal/config"
)

var (
	cfg *config.Config
)

func SetGlobalRestConfig(config *config.Config) {
	cfg = config
}

func startRestAPIServer(globalConfig *config.Config) error {
	cfg = globalConfig

	log.WithFields(log.Fields{
		"address": cfg.ListenAddressGrpc,
	}).Info("Starting REST API")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// The JSONPb.EmitDefaults is necssary, so "empty" fields are returned in JSON.
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   false, // eg: canExec for js instead of can_exec from protobuf
					EmitUnpopulated: true,
				},
			},
		}),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := gw.RegisterSpaghettiCannonApiServiceHandlerFromEndpoint(ctx, mux, cfg.ListenAddressGrpc, opts)

	if err != nil {
		log.Errorf("Could not register REST API Handler %v", err)

		return err
	}

	return http.ListenAndServe(cfg.ListenAddressRest, mux)
}
