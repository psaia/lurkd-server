package github

import (
	"errors"
	"fmt"
	"log"

	"github.com/psaia/lurkd-server/providers"
	"github.com/psaia/lurkd-server/providers/utils"
)

// Github provider
type Github struct {
	ModuleName string
}

type shape []struct {
	Name string
}

// GetVersion of the module.
func (p *Github) GetVersion() (string, error) {
	data := make(shape, 0)
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases", p.ModuleName)
	err := utils.GetJSON(url, &data)

	if err != nil {
		log.Printf("failed grabbing the version for module (%s). error: %v", p.ModuleName, err)
		return "", fmt.Errorf("could not get version for module %s", p.ModuleName)
	}

	return utils.FormatVersion(data[0].Name), nil
}

// Initialize to set the current version.
func (p *Github) Initialize() (string, error) {
	if p.ModuleName == "" {
		return "", errors.New("no module name provided")
	}

	v, err := p.GetVersion()

	if err != nil {
		return "", err
	}

	// Save the version here.

	return v, nil
}

// Check the current release.
func (p *Github) Check(current providers.ReleaseLabel) (providers.VersionResult, error) {
	r := providers.VersionResult{}
	return r, nil
}
