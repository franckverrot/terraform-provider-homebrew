package homebrew

// Client is holding endpoints
type Client struct {
	PackageEndpoint *PackageEndpoint
}

// Init is initializing a client
func (client *Client) Init(config *Config) {
	client.PackageEndpoint = &PackageEndpoint{
		Commander: &SSHExecutor{
			Config: config,
		},
	}
}
