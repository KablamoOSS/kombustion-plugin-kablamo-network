package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeliveryStreamProcessor Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processor.html
type DeliveryStreamProcessor struct {
	Type       interface{} `yaml:"Type"`
	Parameters interface{} `yaml:"Parameters"`
}

// DeliveryStreamProcessor validation
func (resource DeliveryStreamProcessor) Validate() []error {
	errs := []error{}

	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.Parameters == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Parameters'"))
	}
	return errs
}