package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TopicRuleDynamoDBv2Action Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbv2action.html
type TopicRuleDynamoDBv2Action struct {
	RoleArn interface{}            `yaml:"RoleArn,omitempty"`
	PutItem *TopicRulePutItemInput `yaml:"PutItem,omitempty"`
}

// TopicRuleDynamoDBv2Action validation
func (resource TopicRuleDynamoDBv2Action) Validate() []error {
	errors := []error{}

	return errors
}
