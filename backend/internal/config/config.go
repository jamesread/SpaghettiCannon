package config

type Config struct {
	ListenAddressSingleHTTPFrontend string
	ListenAddressRest               string
	ListenAddressWebUI              string
	ShowFooter                      bool
	ShowNewVersions                 bool
	ExternalRestAddress             string
	WebUIDir                        string
}

func DefaultConfig() *Config {
	return &Config{
		WebUIDir:                        "./webui/",
		ExternalRestAddress:             ".",
		ListenAddressSingleHTTPFrontend: "0.0.0.0:4337",
		ListenAddressRest:               "localhost:4339",
		ListenAddressWebUI:              "localhost:4340",
	}
}
