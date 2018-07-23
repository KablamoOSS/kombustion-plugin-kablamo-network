package main

import (
	"github.com/KablamoOSS/kombustion-plugin-network/parsers/vpc"
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

// Parsers functions
var Parsers = map[string]func(
	name string,
	data string,
) []byte{
	"VPC": api.RegisterParser(vpc.ParseNetworkVPC),
}

// Register plugin
func Register() []byte {
	return api.RegisterPlugin(types.Config{
		Name:    name,
		Version: version,
		Prefix:  "Kablamo::Network",
		Help: types.Help{
			Description: "A VPC & Network Plugin",
			Types: []types.TypeMapping{
				{
					Name:        "Kablamo::Network::Network",
					Description: "Creates a VPC and various components.",
					Config:      vpc.NetworkVPCConfig{},
				},
			},
		},
	})
}

func main() {}
