package vpc

import (
	"fmt"
	"strings"

	"github.com/KablamoOSS/kombustion-plugin-network/common"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	cfResources "github.com/KablamoOSS/kombustion/pkg/parsers/resources"

	kombustionTypes "github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

//ParseNetworkVPC parser builder.
func ParseNetworkVPC(name string,
	data string,
) (
	conditions kombustionTypes.TemplateObject,
	metadata kombustionTypes.TemplateObject,
	mappings kombustionTypes.TemplateObject,
	outputs kombustionTypes.TemplateObject,
	parameters kombustionTypes.TemplateObject,
	resources kombustionTypes.TemplateObject,
	transform kombustionTypes.TemplateObject,
	errors []error,
) {
	// Parse the config data
	var config NetworkVPCConfig

	err := yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		errors = append(errors, err)
		return
	}

	// validate the config
	validateErrs := config.Validate()

	if validateErrs != nil {
		errors = append(errors, validateErrs...)
		return
	}

	// create a group of objects (each to be validated)
	resources = make(kombustionTypes.TemplateObject)

	resources[fmt.Sprintf("%s%s", name, config.Properties.Details.VPCName)] = cfResources.NewEC2VPC(
		cfResources.EC2VPCProperties{
			CidrBlock:          config.Properties.CIDR,
			EnableDnsHostnames: true,
			EnableDnsSupport:   true,
			InstanceTenancy:    "default",
			Tags:               common.GenTags(map[string]string{"Name": config.Properties.Details.VPCName}),
		},
	)

	resources[fmt.Sprintf("%s%s", name, config.Properties.DHCP.Name)] = cfResources.NewEC2DHCPOptions(
		cfResources.EC2DHCPOptionsProperties{
			DomainName:         config.Properties.DHCP.Name,
			NetbiosNodeType:    config.Properties.DHCP.NTBType,
			DomainNameServers:  common.SplitStrArray(config.Properties.DHCP.DNSServers),
			NetbiosNameServers: common.SplitStrArray(config.Properties.DHCP.Netbiosservers),
			NtpServers:         common.SplitStrArray(config.Properties.DHCP.NTPServers),
			Tags:               common.GenTags(map[string]string{"Name": config.Properties.DHCP.Name}),
		},
	)

	resources[fmt.Sprintf("%s%s%s", name, config.Properties.DHCP.Name, "VPCDHCPOptionsAssociation")] =
		cfResources.NewEC2VPCDHCPOptionsAssociation(
			cfResources.EC2VPCDHCPOptionsAssociationProperties{
				DhcpOptionsId: map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, config.Properties.DHCP.Name)},
				VpcId:         map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, config.Properties.Details.VPCName)},
			},
		)

	resources[fmt.Sprintf("%s%s", name, "InternetGateway")] = cfResources.NewEC2InternetGateway(
		cfResources.EC2InternetGatewayProperties{
			Tags: common.GenTags(map[string]string{"Name": "IGW"}),
		},
	)

	resources[fmt.Sprintf("%s%s", name, "InternetGatewayVPCGatewayAttachment")] =
		cfResources.NewEC2VPCGatewayAttachment(
			cfResources.EC2VPCGatewayAttachmentProperties{
				InternetGatewayId: map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, "InternetGateway")},
				VpcId:             map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, config.Properties.Details.VPCName)},
			},
		)

	resources[fmt.Sprintf("%s%s", name, "VPNGatewayVPCGatewayAttachment")] =
		cfResources.NewEC2VPCGatewayAttachment(
			cfResources.EC2VPCGatewayAttachmentProperties{
				VpnGatewayId: map[string]interface{}{"Ref": "VGW"},
				VpcId:        map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, config.Properties.Details.VPCName)},
			},
		)

	for routetable, settings := range config.Properties.RouteTables {
		resources[fmt.Sprintf("%s%s", name, routetable)] = cfResources.NewEC2RouteTable(
			cfResources.EC2RouteTableProperties{
				VpcId: map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, config.Properties.Details.VPCName)},
				Tags:  common.GenTags(map[string]string{"Name": routetable}),
			},
		)

		for _, routeinfo := range settings {
			resources[fmt.Sprintf("%s%s", name, routeinfo.RouteName)] = cfResources.NewEC2Route(
				cfResources.EC2RouteProperties{
					DestinationCidrBlock: routeinfo.RouteCIDR,
					GatewayId:            map[string]string{"Ref": fmt.Sprintf("%s%s", name, routeinfo.RouteGW)},
					RouteTableId:         map[string]string{"Ref": fmt.Sprintf("%s%s", name, routetable)},
				},
			)
		}

		resources[fmt.Sprintf("%s%s%s", name, routetable, "RoutePropagation")] =
			cfResources.NewEC2VPNGatewayRoutePropagation(
				cfResources.EC2VPNGatewayRoutePropagationProperties{
					RouteTableIds: common.GenMap(map[string]string{"Ref": fmt.Sprintf("%s%s", name, routetable)}),
					VpnGatewayId:  map[string]string{"Ref": "VGW"},
				},
				fmt.Sprintf("%s%s", name, "VPNGatewayVPCGatewayAttachment"),
			)
	}

	for networkacl, settings := range config.Properties.NetworkACLs {
		resources[fmt.Sprintf("%s%s", name, networkacl)] = cfResources.NewEC2NetworkAcl(
			cfResources.EC2NetworkAclProperties{
				VpcId: map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, config.Properties.Details.VPCName)},
				Tags:  common.GenTags(map[string]string{"Name": networkacl}),
			},
		)

		for aclentry, acl := range settings.(map[interface{}]interface{}) {
			ports := properties.NetworkAclEntryPortRange{
				From: strings.Split(acl.(string), ",")[5],
				To:   strings.Split(acl.(string), ",")[6],
			}
			resources[fmt.Sprintf("%s%s", name, aclentry.(string))] = cfResources.NewEC2NetworkAclEntry(
				cfResources.EC2NetworkAclEntryProperties{
					NetworkAclId: map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, networkacl)},
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
		resources[fmt.Sprintf("%s%s", name, subnet)] = cfResources.NewEC2Subnet(
			cfResources.EC2SubnetProperties{
				VpcId:     map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, config.Properties.Details.VPCName)},
				CidrBlock: settings.CIDR,
				AvailabilityZone: map[interface{}]interface{}{"Fn::Select": []interface{}{
					settings.AZ, map[string]string{"Fn::GetAZs": ""}},
				},
				Tags: common.GenTags(map[string]string{"Name": subnet}),
			},
		)

		resources[fmt.Sprintf("%s%s%s", name, subnet, "SubnetNetworkAclAssociation")] =
			cfResources.NewEC2SubnetNetworkAclAssociation(
				cfResources.EC2SubnetNetworkAclAssociationProperties{
					NetworkAclId: map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, settings.NetACL)},
					SubnetId:     map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, subnet)},
				},
			)

		resources[fmt.Sprintf("%s%s%s", name, subnet, "SubnetRouteTableAssociation")] =
			cfResources.NewEC2SubnetRouteTableAssociation(
				cfResources.EC2SubnetRouteTableAssociationProperties{
					RouteTableId: map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, settings.RouteTable)},
					SubnetId:     map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, subnet)},
				},
			)
	}

	for natgw, settings := range config.Properties.NatGateways {
		resources[fmt.Sprintf("%s%s%s", name, "EIP", natgw)] = cfResources.NewEC2EIP(
			cfResources.EC2EIPProperties{
				Domain: "vpc",
			},
		)

		resources[fmt.Sprintf("%s%s", name, natgw)] = cfResources.NewEC2NatGateway(
			cfResources.EC2NatGatewayProperties{
				AllocationId: map[string]interface{}{"Fn::GetAtt": []string{fmt.Sprintf("%s%s", name, "EIP"+natgw), "AllocationId"}},
				SubnetId:     map[string]interface{}{"Ref": fmt.Sprintf("%s%s", name, settings.Subnet)},
				Tags:         common.GenTags(map[string]string{"Name": natgw}),
			},
		)

		resources[fmt.Sprintf("%s%s%s", name, natgw, "Route")] = cfResources.NewEC2Route(
			cfResources.EC2RouteProperties{
				DestinationCidrBlock: "0.0.0.0/0",
				RouteTableId:         map[string]string{"Ref": fmt.Sprintf("%s%s", name, settings.Routetable)},
				NatGatewayId:         map[string]string{"Ref": fmt.Sprintf("%s%s", name, natgw)},
			},
		)
	}

	return
}
