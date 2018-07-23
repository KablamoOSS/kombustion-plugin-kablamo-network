package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// InstanceGroupConfigConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html
type InstanceGroupConfigConfiguration struct {
	Classification          interface{} `yaml:"Classification,omitempty"`
	ConfigurationProperties interface{} `yaml:"ConfigurationProperties,omitempty"`
	Configurations          interface{} `yaml:"Configurations,omitempty"`
}

// InstanceGroupConfigConfiguration validation
func (resource InstanceGroupConfigConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
