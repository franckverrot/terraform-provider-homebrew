package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-homebrew/homebrew"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: homebrew.Provider})
}
