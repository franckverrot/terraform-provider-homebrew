package homebrew

import (
	"log"
)

// Config stores Homebrew's API configuration
type Config struct {
	Login              string
	Host               string
	PubKey             string
	HomebrewBinaryPath string
}

// Client returns a new Client for accessing Homebrew.
func (c *Config) Client() (*PackageEndpoint, error) {
	client := &Client{}
	client.Init(c)
	log.Printf("[INFO] Homebrew Client configured.")

	return client.PackageEndpoint, nil
}
