package config

type Config struct {
	ListenAddressSingleHTTPFrontend string
	ListenAddressGrpc string
	ListenAddressRest string
	ListenAddressWebUI string
	ShowFooter bool
	ShowNewVersions bool
	ExternalRestAddress string
	WebUIDir string
}

func DefaultConfig() *Config {
	return &Config {
		WebUIDir: "./webui/",
		ExternalRestAddress: ".",
		ListenAddressSingleHTTPFrontend: "0.0.0.0:4337",
		ListenAddressGrpc: "0.0.0.0:4338",
		ListenAddressRest: "0.0.0.0:4339",
		ListenAddressWebUI: "0.0.0.0:4340",
	}
}
