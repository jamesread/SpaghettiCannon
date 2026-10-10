package cvar

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultsHaveMetadata(t *testing.T) {
	defs := Defaults(Seed{
		SiteTitle:         "SpaghettiCannon",
		ShowFooter:        true,
		ShowNewVersions:   true,
		ShowVersionNumber: true,
	})
	require.NotEmpty(t, defs)

	seen := map[string]bool{}
	for _, def := range defs {
		assertDefMetadata(t, def, seen)
	}
	require.True(t, seen[KeySiteTitle], "missing site_title")
}

func assertDefMetadata(t *testing.T, def Def, seen map[string]bool) {
	t.Helper()
	require.NotEmpty(t, def.Key)
	require.False(t, seen[def.Key], "duplicate key %s", def.Key)
	seen[def.Key] = true
	require.NotEmpty(t, def.Title, "%s: title", def.Key)
	require.NotEmpty(t, def.Description, "%s: description", def.Key)
	require.NotEmpty(t, def.Category, "%s: category", def.Key)
	require.NotEmpty(t, def.MainType, "%s: main type", def.Key)
	require.Positive(t, def.Ordinal, "%s: ordinal", def.Key)
}
