package cvar

const (
	TypeString = "string"
	TypeInt    = "int"
	TypeBool   = "bool"

	KeySiteTitle         = "site_title"
	KeyShowFooter        = "show_footer"
	KeyShowNewVersions   = "show_new_versions"
	KeyShowVersionNumber = "show_version_number"

	CategorySite = "Site"
	CategoryUI   = "UI"
)

type Def struct {
	Key         string
	MainType    string
	Title       string
	Description string
	Category    string
	ValueString string
	Ordinal     int
	ValueInt    int
}

type Seed struct {
	SiteTitle         string
	ShowFooter        bool
	ShowNewVersions   bool
	ShowVersionNumber bool
}

func boolInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

func Defaults(seed Seed) []Def {
	return []Def{
		{
			Key: KeySiteTitle, MainType: TypeString, ValueString: seed.SiteTitle,
			Title: "Site title", Description: "Shown in the header and browser tab.",
			Category: CategorySite, Ordinal: 10,
		},
		{
			Key: KeyShowFooter, MainType: TypeBool, ValueInt: boolInt(seed.ShowFooter),
			Title: "Show footer", Description: "Show the page footer under the main content.",
			Category: CategoryUI, Ordinal: 20,
		},
		{
			Key: KeyShowNewVersions, MainType: TypeBool, ValueInt: boolInt(seed.ShowNewVersions),
			Title: "Show new versions", Description: "Hint in the footer when a newer version is available.",
			Category: CategoryUI, Ordinal: 30,
		},
		{
			Key: KeyShowVersionNumber, MainType: TypeBool, ValueInt: boolInt(seed.ShowVersionNumber),
			Title: "Show version number", Description: "Show the installed version next to the product name in the footer.",
			Category: CategoryUI, Ordinal: 40,
		},
	}
}
