## Kablamo Network Plugin for Kombustion

[![Build Status](https://travis-ci.org/KablamoOSS/kombustion-plugin-network.svg?branch=master)](https://travis-ci.org/KablamoOSS/kombustion-plugin-network)
[![](https://tokei.rs/b1/github/KablamoOSS/kombustion-plugin-network)](https://github.com/KablamoOSS/kombustion-plugin-network)
[![Go Report Card](https://goreportcard.com/badge/github.com/KablamoOSS/kombustion-plugin-network)](https://goreportcard.com/badge/github.com/KablamoOSS/kombustion-plugin-network)
[![Coverage Status](https://coveralls.io/repos/github/KablamoOSS/kombustion-plugin-network/badge.svg?branch=master)](https://coveralls.io/github/KablamoOSS/kombustion-plugin-network?branch=master)

# Install

```sh
$ kombustion add github.com/KablamoOSS/kombustion-plugin-network
```
 
# Examples:
    Private VPC
```sh
# To use this example make sure you have the following plugins
#
#
AWSTemplateFormatVersion: 2010-09-09
Description: Private VPC Template
Parameters:
  VGW: {Description: VPC Gateway, Type: String, Default: vgw-abcdefgh}
Mappings: {}
Resources:

    KABLAMOBUILDVPC:
        Type: Kablamo::Network::VPC
        Properties:
            CIDR: 172.16.0.0/20
            Details: {VPCName: PRIVATEEGRESSVPC, VPCDesc: Private Egress VPC, Region: ap-southeast-2}
            Tags: {Name: PRIVATE-EGRESS-VPC, Template: VPC for private endpoints, egress only}            
            DHCP: {Name: DhcpOptions, DNSServers: 172.16.0.2, NTPServers: 169.254.169.123, NTBType: 2}
            Subnets:
                ReservedMgmt1: {CIDR: 172.16.0.0/26, AZ: 0, NetACL: InternalSubnetAcl, RouteTable: InternalRT1 }
                ReservedMgmt2: {CIDR: 172.16.1.0/26, AZ: 1, NetACL: InternalSubnetAcl, RouteTable: InternalRT2 }
                ReservedMgmt3: {CIDR: 172.16.2.0/26, AZ: 2, NetACL: InternalSubnetAcl, RouteTable: InternalRT3 }
                ReservedNet1: {CIDR: 172.16.0.192/26, AZ: 0, NetACL: RestrictedSubnetAcl, RouteTable: PublicRT }
                ReservedNet2: {CIDR: 172.16.1.192/26, AZ: 1, NetACL: RestrictedSubnetAcl, RouteTable: PublicRT }
                ReservedNet3: {CIDR: 172.16.2.192/26, AZ: 2, NetACL: RestrictedSubnetAcl, RouteTable: PublicRT }
                Internal1: {CIDR: 172.16.3.0/24, AZ: 0, NetACL: InternalSubnetAcl, RouteTable: InternalRT1 }
                Internal2: {CIDR: 172.16.4.0/24, AZ: 1, NetACL: InternalSubnetAcl, RouteTable: InternalRT2 }
                Internal3: {CIDR: 172.16.5.0/24, AZ: 2, NetACL: InternalSubnetAcl, RouteTable: InternalRT3 }
                PerimeterInternal1: {CIDR: 172.16.6.0/24, AZ: 0, NetACL: InternalSubnetAcl, RouteTable: InternalRT1 }
                PerimeterInternal2: {CIDR: 172.16.7.0/24, AZ: 1, NetACL: InternalSubnetAcl, RouteTable: InternalRT2 }
                PerimeterInternal3: {CIDR: 172.16.8.0/24, AZ: 2, NetACL: InternalSubnetAcl, RouteTable: InternalRT3 }
            RouteTables:
                PublicRT:
                  - RouteName: PublicRoute
                    RouteCIDR: 0.0.0.0/0
                    RouteGW: InternetGateway
                InternalRT1:
                InternalRT2:
                InternalRT3:
            NATGateways:
                NATGW1:
                    {Subnet: ReservedNet1, Routetable: InternalRT1}
                NATGW2:
                    {Subnet: ReservedNet2, Routetable: InternalRT2}
                NATGW3:
                    {Subnet: ReservedNet3, Routetable: InternalRT3}
            NetworkACLs:
                RestrictedSubnetAcl: 
                    RestrictedSubnetAclEntryInTCPUnReserved: "90,6,allow,false,0.0.0.0/0,1024,65535"
                    RestrictedSubnetAclEntryInUDPUnReserved: "91,17,allow,false,0.0.0.0/0,1024,65535"
                    RestrictedSubnetAclEntryOutTCPUnReserved: 90,6,allow,true,0.0.0.0/0,1024,65535
                    RestrictedSubnetAclEntryOutUDPUnReserved: 91,17,allow,true,0.0.0.0/0,1024,65535
                    RestrictedSubnetAclEntryOutPuppet: 92,6,allow,true,172.16.0.0/16,8140,8140
                    RestrictedSubnetAclEntryOutHTTP: 101,6,allow,true,0.0.0.0/0,80,80
                    RestrictedSubnetAclEntryOutHTTPS: 102,6,allow,true,0.0.0.0/0,443,443
                    RestrictedSubnetAclEntryOutSSH: 103,6,allow,true,0.0.0.0/0,22,22
                    RestrictedSubnetAclEntryOutAEMIDM: 104,6,allow,true,0.0.0.0/0,1443,1443
                    RestrictedSubnetAclEntryOutAEMGW: 105,6,allow,true,0.0.0.0/0,2443,2443
                    RestrictedSubnetAclEntryInHTTP: 101,6,allow,false,0.0.0.0/0,80,80
                    RestrictedSubnetAclEntryInHTTPS: 102,6,allow,false,0.0.0.0/0,443,443
                    RestrictedSubnetAclEntryIn: 110,-1,allow,false,172.16.0.0/16,1,65535
                    RestrictedSubnetAclEntryOut: 110,-1,allow,true,172.16.0.0/16,1,65535
                    RestrictedSubnetAclEntryNTP: 120,6,allow,true,0.0.0.0/0,123,123
                    RestrictedSubnetAclEntryRoutingIn: 130,-1,allow,false,192.168.0.0/16,1,65535
                    RestrictedSubnetAclEntryRoutingOut: 130,-1,allow,true,192.168.0.0/16,1,65535
                    RestrictedSubnetAclEntryInSquid2: 140,6,allow,false,172.16.0.0/16,3128,3128
                    RestrictedSubnetAclEntryInDNSTCP: 150,6,allow,false,172.16.0.0/16,53,53
                    RestrictedSubnetAclEntryOutDNSTCP: 150,6,allow,true,0.0.0.0/0,53,53
                    RestrictedSubnetAclEntryInDNSUDP: 160,17,allow,false,172.16.0.0/16,53,53
                    RestrictedSubnetAclEntryOutDNSUDP: 160,17,allow,true,0.0.0.0/0,53,53
                    RestrictedSubnetAclEntryInNetBios: 170,6,allow,false,172.16.0.0/16,389,389
                    RestrictedSubnetAclEntryOutNetBios: 170,6,allow,true,172.16.0.0/16,389,389
                    RestrictedSubnetAclEntryInNetBios1: 80,6,allow,false,172.16.0.0/16,137,139
                    RestrictedSubnetAclEntryOutNetBios1: 180,6,allow,true,172.16.0.0/16,137,139
                InternalSubnetAcl:
                    InternalSubnetAclEntryIn: 100,-1,allow,false,172.16.0.0/16,1,65535
                    InternalSubnetAclEntryOut: 100,-1,allow,true,172.16.0.0/16,1,65535
                    InternalSubnetAclEntryInTCPUnreserved: 102,6,allow,false,0.0.0.0/0,1024,65535
                    InternalSubnetAclEntryInUDPUnreserved: 103,17,allow,false,0.0.0.0/0,1024,65535
                    InternalSubnetAclEntryOutHTTP: 102,6,allow,true,0.0.0.0/0,80,80
                    InternalSubnetAclEntryOutHTTPS: 103,6,allow,true,0.0.0.0/0,443,443
                    InternalSubnetAclEntryOutTCPUnreserved: 104,6,allow,true,172.16.0.0/16,1024,65535
                    InternalSubnetAclEntryOutUDPUnreserved: 105,6,allow,true,172.16.0.0/16,1024,65535
                    InternalSubnetAclEntryOutTCPDNS: 110,6,allow,true,0.0.0.0/0,53,53
                    InternalSubnetAclEntryOutUDPDNS: 111,17,allow,true,0.0.0.0/0,53,53
                    InternalSubnetAclEntryOutSSH: 150,6,allow,true,0.0.0.0/0,22,22
                    InternalSubnetAclEntryOutAEMIDM: 160,6,allow,true,0.0.0.0/0,1443,1443
                    InternalSubnetAclEntryOutAEMGW: 161,6,allow,true,0.0.0.0/0,2443,2443
                    InternalSubnetAclEntryInEquinix: 200,-1,allow,false,192.168.0.0/16,1,65535
                    InternalSubnetAclEntryOutEquinix: 200,-1,allow,true,192.168.0.0/16,1,65535
```

Public VPC
```sh
# To use this example make sure you have the following plugins
#
#
AWSTemplateFormatVersion: 2010-09-09
Description: Public VPC Template
Parameters:
  VGW: {Description: VPC Gateway, Type: String, Default: vgw-abcdefgh}
Mappings: {}
Resources:

    KABLAMOBUILDVPC:
        Type: Kablamo::Network::VPC
        Properties:
            CIDR: 172.16.0.0/20
            Details: {VPCName: PUBLICVPC, VPCDesc: Public VPC, Region: ap-southeast-2}
            Tags: {Name: PUBLIC-VPC, Template: VPC for public endpoints}            
            DHCP: {Name: DhcpOptions, DNSServers: 172.16.0.2, NTPServers: 169.254.169.123, NTBType: 2}
            Subnets:
                ReservedMgmt1: {CIDR: 172.16.0.0/26, AZ: 0, NetACL: InternalSubnetAcl, RouteTable: InternalRT1 }
                ReservedMgmt2: {CIDR: 172.16.1.0/26, AZ: 1, NetACL: InternalSubnetAcl, RouteTable: InternalRT2 }
                ReservedMgmt3: {CIDR: 172.16.2.0/26, AZ: 2, NetACL: InternalSubnetAcl, RouteTable: InternalRT3 }
                ReservedNet1: {CIDR: 172.16.0.192/26, AZ: 0, NetACL: RestrictedSubnetAcl, RouteTable: PublicRT }
                ReservedNet2: {CIDR: 172.16.1.192/26, AZ: 1, NetACL: RestrictedSubnetAcl, RouteTable: PublicRT }
                ReservedNet3: {CIDR: 172.16.2.192/26, AZ: 2, NetACL: RestrictedSubnetAcl, RouteTable: PublicRT }
                Internal1: {CIDR: 172.16.3.0/24, AZ: 0, NetACL: InternalSubnetAcl, RouteTable: InternalRT1 }
                Internal2: {CIDR: 172.16.4.0/24, AZ: 1, NetACL: InternalSubnetAcl, RouteTable: InternalRT2 }
                Internal3: {CIDR: 172.16.5.0/24, AZ: 2, NetACL: InternalSubnetAcl, RouteTable: InternalRT3 }
                PerimeterInternal1: {CIDR: 172.16.6.0/24, AZ: 0, NetACL: InternalSubnetAcl, RouteTable: InternalRT1 }
                PerimeterInternal2: {CIDR: 172.16.7.0/24, AZ: 1, NetACL: InternalSubnetAcl, RouteTable: InternalRT2 }
                PerimeterInternal3: {CIDR: 172.16.8.0/24, AZ: 2, NetACL: InternalSubnetAcl, RouteTable: InternalRT3 }
                PerimeterExternal1: {CIDR: 172.16.6.0/24, AZ: 0, NetACL: RestrictedSubnetAcl, RouteTable: PublicRT }
                PerimeterExternal2: {CIDR: 172.16.7.0/24, AZ: 1, NetACL: RestrictedSubnetAcl, RouteTable: PublicRT }
                PerimeterExternal3: {CIDR: 172.16.8.0/24, AZ: 2, NetACL: RestrictedSubnetAcl, RouteTable: PublicRT }
            RouteTables:
                PublicRT:
                  - RouteName: PublicRoute
                    RouteCIDR: 0.0.0.0/0
                    RouteGW: InternetGateway
                InternalRT1:
                InternalRT2:
                InternalRT3:
            NATGateways:
                NATGW1:
                    {Subnet: ReservedNet1, Routetable: InternalRT1}
                NATGW2:
                    {Subnet: ReservedNet2, Routetable: InternalRT2}
                NATGW3:
                    {Subnet: ReservedNet3, Routetable: InternalRT3}
            NetworkACLs:
                RestrictedSubnetAcl: 
                    RestrictedSubnetAclEntryInTCPUnReserved: "90,6,allow,false,0.0.0.0/0,1024,65535"
                    RestrictedSubnetAclEntryInUDPUnReserved: "91,17,allow,false,0.0.0.0/0,1024,65535"
                    RestrictedSubnetAclEntryOutTCPUnReserved: 90,6,allow,true,0.0.0.0/0,1024,65535
                    RestrictedSubnetAclEntryOutUDPUnReserved: 91,17,allow,true,0.0.0.0/0,1024,65535
                    RestrictedSubnetAclEntryOutPuppet: 92,6,allow,true,172.16.0.0/16,8140,8140
                    RestrictedSubnetAclEntryOutHTTP: 101,6,allow,true,0.0.0.0/0,80,80
                    RestrictedSubnetAclEntryOutHTTPS: 102,6,allow,true,0.0.0.0/0,443,443
                    RestrictedSubnetAclEntryOutSSH: 103,6,allow,true,0.0.0.0/0,22,22
                    RestrictedSubnetAclEntryOutAEMIDM: 104,6,allow,true,0.0.0.0/0,1443,1443
                    RestrictedSubnetAclEntryOutAEMGW: 105,6,allow,true,0.0.0.0/0,2443,2443
                    RestrictedSubnetAclEntryInHTTP: 101,6,allow,false,0.0.0.0/0,80,80
                    RestrictedSubnetAclEntryInHTTPS: 102,6,allow,false,0.0.0.0/0,443,443
                    RestrictedSubnetAclEntryIn: 110,-1,allow,false,172.16.0.0/16,1,65535
                    RestrictedSubnetAclEntryOut: 110,-1,allow,true,172.16.0.0/16,1,65535
                    RestrictedSubnetAclEntryNTP: 120,6,allow,true,0.0.0.0/0,123,123
                    RestrictedSubnetAclEntryRoutingIn: 130,-1,allow,false,192.168.0.0/16,1,65535
                    RestrictedSubnetAclEntryRoutingOut: 130,-1,allow,true,192.168.0.0/16,1,65535
                    RestrictedSubnetAclEntryInSquid2: 140,6,allow,false,172.16.0.0/16,3128,3128
                    RestrictedSubnetAclEntryInDNSTCP: 150,6,allow,false,172.16.0.0/16,53,53
                    RestrictedSubnetAclEntryOutDNSTCP: 150,6,allow,true,0.0.0.0/0,53,53
                    RestrictedSubnetAclEntryInDNSUDP: 160,17,allow,false,172.16.0.0/16,53,53
                    RestrictedSubnetAclEntryOutDNSUDP: 160,17,allow,true,0.0.0.0/0,53,53
                    RestrictedSubnetAclEntryInNetBios: 170,6,allow,false,172.16.0.0/16,389,389
                    RestrictedSubnetAclEntryOutNetBios: 170,6,allow,true,172.16.0.0/16,389,389
                    RestrictedSubnetAclEntryInNetBios1: 80,6,allow,false,172.16.0.0/16,137,139
                    RestrictedSubnetAclEntryOutNetBios1: 180,6,allow,true,172.16.0.0/16,137,139
                InternalSubnetAcl:
                    InternalSubnetAclEntryIn: 100,-1,allow,false,172.16.0.0/16,1,65535
                    InternalSubnetAclEntryOut: 100,-1,allow,true,172.16.0.0/16,1,65535
                    InternalSubnetAclEntryInTCPUnreserved: 102,6,allow,false,0.0.0.0/0,1024,65535
                    InternalSubnetAclEntryInUDPUnreserved: 103,17,allow,false,0.0.0.0/0,1024,65535
                    InternalSubnetAclEntryOutHTTP: 102,6,allow,true,0.0.0.0/0,80,80
                    InternalSubnetAclEntryOutHTTPS: 103,6,allow,true,0.0.0.0/0,443,443
                    InternalSubnetAclEntryOutTCPUnreserved: 104,6,allow,true,172.16.0.0/16,1024,65535
                    InternalSubnetAclEntryOutUDPUnreserved: 105,6,allow,true,172.16.0.0/16,1024,65535
                    InternalSubnetAclEntryOutTCPDNS: 110,6,allow,true,0.0.0.0/0,53,53
                    InternalSubnetAclEntryOutUDPDNS: 111,17,allow,true,0.0.0.0/0,53,53
                    InternalSubnetAclEntryOutSSH: 150,6,allow,true,0.0.0.0/0,22,22
                    InternalSubnetAclEntryOutAEMIDM: 160,6,allow,true,0.0.0.0/0,1443,1443
                    InternalSubnetAclEntryOutAEMGW: 161,6,allow,true,0.0.0.0/0,2443,2443
                    InternalSubnetAclEntryInEquinix: 200,-1,allow,false,192.168.0.0/16,1,65535
                    InternalSubnetAclEntryOutEquinix: 200,-1,allow,true,192.168.0.0/16,1,65535
```