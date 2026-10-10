package api

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	pb "github.com/jamesread/SpaghettiCannon/gen/SpaghettiCannon/clientapi/v1"
	"github.com/jamesread/SpaghettiCannon/internal/config"
	"github.com/jamesread/SpaghettiCannon/internal/cvar"
	"github.com/jamesread/SpaghettiCannon/internal/store"
)

func ensureDefaultCvars(ctx context.Context, st store.Store, cfg *config.Config) error {
	seed := cvar.Seed{
		SiteTitle:         cfg.SiteTitle,
		ShowFooter:        cfg.ShowFooter,
		ShowNewVersions:   cfg.ShowNewVersions,
		ShowVersionNumber: cfg.ShowVersionNumber,
	}
	for _, def := range cvar.Defaults(seed) {
		if err := st.InsertCvarIfMissing(ctx, store.CvarRow{
			Key: def.Key, MainType: def.MainType,
			Title: def.Title, Description: def.Description,
			Category: def.Category, Ordinal: def.Ordinal,
			ValueInt: def.ValueInt, ValueString: def.ValueString,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (api *SpaghettiCannonApiService) siteTitle(ctx context.Context) string {
	row, err := api.store.FindCvar(ctx, cvar.KeySiteTitle)
	if err != nil || row == nil || row.ValueString == "" {
		return api.cfg.SiteTitle
	}
	return row.ValueString
}

func (api *SpaghettiCannonApiService) boolCvar(ctx context.Context, key string, fallback bool) bool {
	row, err := api.store.FindCvar(ctx, key)
	if err != nil || row == nil {
		return fallback
	}
	return row.ValueInt != 0
}

func (api *SpaghettiCannonApiService) showFooter(ctx context.Context) bool {
	return api.boolCvar(ctx, cvar.KeyShowFooter, api.cfg.ShowFooter)
}

func (api *SpaghettiCannonApiService) showNewVersions(ctx context.Context) bool {
	return api.boolCvar(ctx, cvar.KeyShowNewVersions, api.cfg.ShowNewVersions)
}

func (api *SpaghettiCannonApiService) showVersionNumber(ctx context.Context) bool {
	return api.boolCvar(ctx, cvar.KeyShowVersionNumber, api.cfg.ShowVersionNumber)
}

func toProtoCvar(row *store.CvarRow) *pb.Cvar {
	return &pb.Cvar{
		Key: row.Key, MainType: row.MainType,
		ValueInt: int32(row.ValueInt), ValueString: row.ValueString,
		Title: row.Title, Description: row.Description,
		Category: row.Category, Ordinal: int32(row.Ordinal),
	}
}

func (api *SpaghettiCannonApiService) ListCvars(ctx context.Context, _ *connect.Request[pb.ListCvarsRequest]) (*connect.Response[pb.ListCvarsResponse], error) {
	rows, err := api.store.ListCvars(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	out := make([]*pb.Cvar, 0, len(rows))
	for i := range rows {
		out = append(out, toProtoCvar(&rows[i]))
	}
	return connect.NewResponse(&pb.ListCvarsResponse{Cvars: out}), nil
}

func validateCvarUpdate(row *store.CvarRow, valueInt int32, valueString string) (int, string, error) {
	switch row.MainType {
	case cvar.TypeString:
		return validateStringCvar(valueString)
	case cvar.TypeInt:
		return int(valueInt), "", nil
	case cvar.TypeBool:
		return validateBoolCvar(valueInt), "", nil
	default:
		return 0, "", fmt.Errorf("unsupported cvar type")
	}
}

func validateStringCvar(valueString string) (int, string, error) {
	if valueString == "" {
		return 0, "", fmt.Errorf("value required")
	}
	if len(valueString) > 255 {
		return 0, "", fmt.Errorf("value too long")
	}
	return 0, valueString, nil
}

func validateBoolCvar(valueInt int32) int {
	if valueInt != 0 {
		return 1
	}
	return 0
}

func (api *SpaghettiCannonApiService) requireCvar(ctx context.Context, key string) (*store.CvarRow, error) {
	row, err := api.store.FindCvar(ctx, key)
	if err != nil || row == nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("cvar not found"))
	}
	return row, nil
}

func (api *SpaghettiCannonApiService) UpdateCvar(ctx context.Context, req *connect.Request[pb.UpdateCvarRequest]) (*connect.Response[pb.Cvar], error) {
	row, findErr := api.requireCvar(ctx, req.Msg.Key)
	if findErr != nil {
		return nil, findErr
	}
	valueInt, valueString, valErr := validateCvarUpdate(row, req.Msg.ValueInt, req.Msg.ValueString)
	if valErr != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, valErr)
	}
	if updErr := api.store.UpdateCvar(ctx, row.Key, valueInt, valueString); updErr != nil {
		return nil, connect.NewError(connect.CodeInternal, updErr)
	}
	updated, _ := api.store.FindCvar(ctx, row.Key)
	return connect.NewResponse(toProtoCvar(updated)), nil
}
