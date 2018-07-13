package outputs

import (
	"github.com/KablamoOSS/kombustion-plugin-network/common"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ParseNetworkVPC -
func ParseNetworkVPC(name string, data string) (cf types.TemplateObject, errs []error) {
	var config common.NetworkVPCConfig

	err := yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		errs = append(errs, err)
		return
	}

	// create a group of objects (each to be validated)
	cf = make(types.TemplateObject)

	return
}
