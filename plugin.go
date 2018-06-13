package main

import (
	"github.com/KablamoOSS/kombustion-plugin-serverless/resources"
	"github.com/KablamoOSS/kombustion/kombustion-plugin-kablamo-network/common"
	"github.com/KablamoOSS/kombustion/plugins/api"
	"github.com/KablamoOSS/kombustion/plugins/api/types"
	"github.com/KablamoOSS/kombustion/types"
)

var (
	version string
	name    string
)

func init() {
	if version == "" {
		version = "BUILT_FROM_SOURCE"
	}
	if name == "" {
		name = "kombustion-plugin-serverless"
	}
}

// Resources for this plugin
var Resources = map[string]func(
	ctx map[string]interface{},
	name string,
	data string,
) []byte{
	"Kablamo::Network::VPC": api.RegisterResource(resources.ParseNetworkVPC),
}

// Outputs for this plugin
var Outputs = map[string]func(
	ctx map[string]interface{},
	name string,
	data string,
) []byte{}

// Mappings for this plugin
var Mappings = map[string]func(
	ctx map[string]interface{},
	name string,
	data string,
) []byte{}

var Help = types.PluginHelp{
	Description: "Helper function for Kablamo VPC",
	TypeMappings: []types.TypeMapping{
		{
			Name:        "Kablamo::Network::VPC",
			Description: "Creates a complete VPC network with subnets, route tables, routes & NACL's",
			Config:      common.NetworkVPCConfig{},
		},
	},
}

// Register plugin
func Register() []byte {
	return api.RegisterPlugin(types.Config{
		Name:               name,
		Version:            version,
		Prefix:             "Kablamo",
		RequiresAWSSession: false,
		Help: types.Help{
			Description: "A VPC & Network Plugin",
			TypeMappings: []types.TypeMapping{
				{
					Name:        "Kablamo::Network::Network",
					Description: "Creates a VPC and various components.",
					Config:      common.NetworkVPCConfig{},
				},
			},
		},
	})
}

func main() {}
