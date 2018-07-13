package main

import (
	"github.com/KablamoOSS/kombustion-plugin-network/common"
	"github.com/KablamoOSS/kombustion-plugin-network/resources"
	"github.com/KablamoOSS/kombustion/pkg/plugins/api"
	"github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
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
		name = "kombustion-plugin-network"
	}
}

// Resources for this plugin
var Resources = map[string]func(
	name string,
	data string,
) []byte{
	"Network::VPC": api.RegisterResource(resources.ParseNetworkVPC),
}

// Outputs for this plugin
var Outputs = map[string]func(
	name string,
	data string,
) []byte{}

// Mappings for this plugin
var Mappings = map[string]func(
	name string,
	data string,
) []byte{}

// Help -
var Help = types.Help{
	Description: "Helper function for Kablamo VPC",
	TypeMappings: []types.TypeMapping{
		{
			Name:        "Network::VPC",
			Description: "Creates a complete VPC network with subnets, route tables, routes & NACL's",
			Config:      common.NetworkVPCConfig{},
		},
	},
}

// Register plugin
func Register() []byte {
	return api.RegisterPlugin(types.Config{
		Name:    name,
		Version: version,
		Prefix:  "Kablamo",
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
