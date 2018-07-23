package vpc

import "fmt"

// Validate - input Config validation
func (config NetworkVPCConfig) Validate() (errors []error) {
	// Enforce required props
	if config.Properties.CIDR == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'CIDR'"))
	}
	if config.Properties.DHCP.Name == "" {
		errors = append(errors, fmt.Errorf("Missing required field 'DHCP.Name'"))
	}
	return
}
