// +build plugin

package outputs

import (
	"github.com/KablamoOSS/kombustion-plugin-kablamo-network/common"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
	yaml "gopkg.in/yaml.v2"
)

func ParseNetworkVPC(ctx map[string]interface{}, name string, data string) (cf kombustionTypes.TemplateObject, err error) {
	var config common.NetworkVPCConfig
	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	// create a group of objects (each to be validated)
	cf = make(kombustionTypes.TemplateObject)

	return
}
