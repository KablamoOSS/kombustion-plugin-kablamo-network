package common

import (
	"fmt"
	"strings"
)

//Subnet object
type Subnet struct {
	CIDR       string `yaml:"CIDR"`
	AZ         string `yaml:"AZ"`
	NetACL     string `yaml:"NetACL"`
	RouteTable string `yaml:"RouteTable"`
}

//NATGateway object
type NATGateway struct {
	Subnet     string `yaml:"Subnet"`
	Routetable string `yaml:"Routetable"`
}

//Routetable an array of Route(s)
type Routetable struct {
	routes []Route `yaml:"routes"`
}

//Route Object
type Route struct {
	RouteName string `yaml:"RouteName"`
	RouteCIDR string `yaml:"RouteCIDR"`
	RouteGW   string `yaml:"RouteGW"`
}

type Tag struct {
	Key   string `yaml:"Key"`
	Value string `yaml:"Value"`
}

//NetworkVPCConfig Main Object and construct
type NetworkVPCConfig struct {
	Properties struct {
		CIDR *string `yaml:"CIDR"`
		DHCP struct {
			Name           string `yaml:"Name"`
			DNSServers     string `yaml:"DNSServers"`
			NTPServers     string `yaml:"NTPServers,omitempty"`
			NTBType        string `yaml:"NTBType,omitempty"`
			Domainname     string `yaml:"Domainname,omitempty"`
			Netbiosservers string `yaml:"Netbiosservers,omitempty"`
		} `yaml:"DHCP"`
		Details struct {
			VPCName string `yaml:"VPCName"`
			VPCDesc string `yaml:"VPCDesc"`
			Region  string `yaml:"Region"`
		} `yaml:"Details"`
		Subnets     map[string]Subnet      `yaml:"Subnets,omitempty"`
		NatGateways map[string]NATGateway  `yaml:"NATGateways,omitempty"`
		RouteTables map[string][]Route     `yaml:"RouteTables,omitempty"`
		NetworkACLs map[string]interface{} `yaml:"NetworkACLs,omitempty"`
		Tags        interface{}            `yaml:"Tags"`
	} `yaml:"Properties"`
}

func SplitStrArray(asset string) []string {
	if len(asset) > 0 {
		strArray := strings.Split(asset, ",")
		return strArray
	} else {
		return nil
	}
}

func GenMap(asset map[string]string) []map[string]string {
	arrayMap := make([]map[string]string, 0)
	arrayMap = append(arrayMap, asset)
	return arrayMap
}

func GenTags(Tags map[string]string) []Tag {
	arrayTags := make([]Tag, 0)
	for k, v := range Tags {
		arrayTags = append(arrayTags, Tag{Key: k, Value: v})
	}
	return arrayTags
}

// Validate - input Config validation
func (this NetworkVPCConfig) Validate() (errors []error) {
	if this.Properties.CIDR == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'CIDR'"))
	}
	return
}
