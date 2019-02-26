package homebrew

import (
	"fmt"
	"strings"
)

// PackageEndpoint is what instruments the SSH commands
type PackageEndpoint struct {
	Commander *SSHExecutor
}

// Package data structure
// If the endpoint's API arguments start diverging from this base data structure,
// which should really only hold a name, and eventually a version, let's create
// a PackageParams data structure so we don't couple the API and the base structure.
type Package struct {
	Name string
}

// Get retrieves a package
func (endpoint *PackageEndpoint) Get(name string) (*Package, error) {
	result, err := endpoint.Commander.Command("info", name)

	if err != nil {
		return nil, fmt.Errorf("Command info error: %+v", err)
	}

	if strings.Contains(result.Output, "Not installed") {
		return nil, fmt.Errorf("Package %s is not installed", name)
	}

	return &Package{Name: name}, nil
}

// Install installs a package
func (endpoint *PackageEndpoint) Install(params *Package) (*Package, error) {
	_, err := endpoint.Commander.Command("install", params.Name)

	if err != nil {
		return nil, fmt.Errorf("Command install error: %+v", err)
	}

	return &Package{Name: params.Name}, nil
}

// Uninstall uninstalls a package
func (endpoint *PackageEndpoint) Uninstall(params *Package) (*Package, error) {
	_, err := endpoint.Commander.Command("uninstall", params.Name)

	if err != nil {
		return nil, fmt.Errorf("Command uninstall error: %+v", err)
	}

	return &Package{Name: params.Name}, nil
}
