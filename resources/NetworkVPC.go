package resources

import (
	"strings"

	"github.com/KablamoOSS/kombustion-plugin-kablamo-network/common"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/pkg/parsers/resources"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
	yaml "gopkg.in/yaml.v2"
)

//ParseNetworkVPC parser builder.
func ParseNetworkVPC(ctx map[string]interface{}, name string, data string) (cf kombustionTypes.TemplateObject) {
	// Parse the config data
	var config common.NetworkVPCConfig
	if err := yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	// validate the config
	config.Validate()

	// create a group of objects (each to be validated)
	cf = make(kombustionTypes.TemplateObject)

	cf[config.Properties.Details.VPCName] = resources.NewEC2VPC(
		resources.EC2VPCProperties{
			CidrBlock:          config.Properties.CIDR,
			EnableDnsHostnames: true,
			EnableDnsSupport:   true,
			InstanceTenancy:    "default",
			Tags:               common.GenTags(map[string]string{"Name": config.Properties.Details.VPCName}),
		},
	)

	cf[config.Properties.DHCP.Name] = resources.NewEC2DHCPOptions(
		resources.EC2DHCPOptionsProperties{
			DomainName:         config.Properties.DHCP.Name,
			NetbiosNodeType:    config.Properties.DHCP.NTBType,
			DomainNameServers:  common.SplitStrArray(config.Properties.DHCP.DNSServers),
			NetbiosNameServers: common.SplitStrArray(config.Properties.DHCP.Netbiosservers),
			NtpServers:         common.SplitStrArray(config.Properties.DHCP.NTPServers),
			Tags:               common.GenTags(map[string]string{"Name": config.Properties.DHCP.Name}),
		},
	)

	cf[config.Properties.DHCP.Name+"VPCDHCPOptionsAssociation"] = resources.NewEC2VPCDHCPOptionsAssociation(
		resources.EC2VPCDHCPOptionsAssociationProperties{
			DhcpOptionsId: map[string]interface{}{"Ref": config.Properties.DHCP.Name},
			VpcId:         map[string]interface{}{"Ref": config.Properties.Details.VPCName},
		},
	)

	cf["InternetGateway"] = resources.NewEC2InternetGateway(
		resources.EC2InternetGatewayProperties{
			Tags: common.GenTags(map[string]string{"Name": "IGW"}),
		},
	)

	cf["InternetGatewayVPCGatewayAttachment"] = resources.NewEC2VPCGatewayAttachment(
		resources.EC2VPCGatewayAttachmentProperties{
			InternetGatewayId: map[string]interface{}{"Ref": "InternetGateway"},
			VpcId:             map[string]interface{}{"Ref": config.Properties.Details.VPCName},
		},
	)

	cf["VPNGatewayVPCGatewayAttachment"] = resources.NewEC2VPCGatewayAttachment(
		resources.EC2VPCGatewayAttachmentProperties{
			VpnGatewayId: map[string]interface{}{"Ref": "VGW"},
			VpcId:        map[string]interface{}{"Ref": config.Properties.Details.VPCName},
		},
	)

	for routetable, settings := range config.Properties.RouteTables {
		cf[routetable] = resources.NewEC2RouteTable(
			resources.EC2RouteTableProperties{
				VpcId: map[string]interface{}{"Ref": config.Properties.Details.VPCName},
				Tags:  common.GenTags(map[string]string{"Name": routetable}),
			},
		)

		for _, routeinfo := range settings {
			cf[routeinfo.RouteName] = resources.NewEC2Route(
				resources.EC2RouteProperties{
					DestinationCidrBlock: routeinfo.RouteCIDR,
					GatewayId:            map[string]string{"Ref": routeinfo.RouteGW},
					RouteTableId:         map[string]string{"Ref": routetable},
				},
			)
		}

		cf[routetable+"RoutePropagation"] = resources.NewEC2VPNGatewayRoutePropagation(
			resources.EC2VPNGatewayRoutePropagationProperties{
				RouteTableIds: common.GenMap(map[string]string{"Ref": routetable}),
				VpnGatewayId:  map[string]string{"Ref": "VGW"},
			},
			"VPNGatewayVPCGatewayAttachment",
		)
	}

	for networkacl, settings := range config.Properties.NetworkACLs {
		cf[networkacl] = resources.NewEC2NetworkAcl(
			resources.EC2NetworkAclProperties{
				VpcId: map[string]interface{}{"Ref": config.Properties.Details.VPCName},
				Tags:  common.GenTags(map[string]string{"Name": networkacl}),
			},
		)

		for aclentry, acl := range settings.(map[interface{}]interface{}) {
			ports := properties.NetworkAclEntryPortRange{
				From: strings.Split(acl.(string), ",")[5],
				To:   strings.Split(acl.(string), ",")[6],
			}
			cf[aclentry.(string)] = resources.NewEC2NetworkAclEntry(
				resources.EC2NetworkAclEntryProperties{
					NetworkAclId: map[string]interface{}{"Ref": networkacl},
					RuleNumber:   strings.Split(acl.(string), ",")[0],
					Protocol:     strings.Split(acl.(string), ",")[1],
					RuleAction:   strings.Split(acl.(string), ",")[2],
					Egress:       strings.Split(acl.(string), ",")[3],
					CidrBlock:    strings.Split(acl.(string), ",")[4],
					PortRange:    &ports,
				},
			)
		}
	}

	for subnet, settings := range config.Properties.Subnets {
		cf[subnet] = resources.NewEC2Subnet(
			resources.EC2SubnetProperties{
				VpcId:            map[string]interface{}{"Ref": config.Properties.Details.VPCName},
				CidrBlock:        settings.CIDR,
				AvailabilityZone: map[interface{}]interface{}{"Fn::Select": []interface{}{settings.AZ, map[string]string{"Fn::GetAZs": ""}}},
				Tags:             common.GenTags(map[string]string{"Name": subnet}),
			},
		)

		cf[subnet+"SubnetNetworkAclAssociation"] = resources.NewEC2SubnetNetworkAclAssociation(
			resources.EC2SubnetNetworkAclAssociationProperties{
				NetworkAclId: map[string]interface{}{"Ref": settings.NetACL},
				SubnetId:     map[string]interface{}{"Ref": subnet},
			},
		)

		cf[subnet+"SubnetRouteTableAssociation"] = resources.NewEC2SubnetRouteTableAssociation(
			resources.EC2SubnetRouteTableAssociationProperties{
				RouteTableId: map[string]interface{}{"Ref": settings.RouteTable},
				SubnetId:     map[string]interface{}{"Ref": subnet},
			},
		)
	}

	for natgw, settings := range config.Properties.NatGateways {
		cf["EIP"+natgw] = resources.NewEC2EIP(
			resources.EC2EIPProperties{
				Domain: "vpc",
			},
		)

		cf[natgw] = resources.NewEC2NatGateway(
			resources.EC2NatGatewayProperties{
				AllocationId: map[string]interface{}{"Fn::GetAtt": []string{"EIP" + natgw, "AllocationId"}},
				SubnetId:     map[string]interface{}{"Ref": settings.Subnet},
				Tags:         common.GenTags(map[string]string{"Name": natgw}),
			},
		)

		cf[natgw+"Route"] = resources.NewEC2Route(
			resources.EC2RouteProperties{
				DestinationCidrBlock: "0.0.0.0/0",
				RouteTableId:         map[string]string{"Ref": settings.Routetable},
				NatGatewayId:         map[string]string{"Ref": natgw},
			},
		)
	}

	/* 	for k, resource := range cf {
	   		if errs := resource.Validate(); len(errs) > 0 {
	   			for _, err = range errs {
	   				log.Println("WARNING: KablamoNetworkConfig - ", err)
	   			}
	   			return
	   		}
	   	}
		   return */
	return
}
