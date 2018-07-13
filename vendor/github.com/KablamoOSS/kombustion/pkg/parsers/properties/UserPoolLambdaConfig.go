package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// UserPoolLambdaConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-lambdaconfig.html
type UserPoolLambdaConfig struct {
	CreateAuthChallenge         interface{} `yaml:"CreateAuthChallenge,omitempty"`
	CustomMessage               interface{} `yaml:"CustomMessage,omitempty"`
	DefineAuthChallenge         interface{} `yaml:"DefineAuthChallenge,omitempty"`
	PostAuthentication          interface{} `yaml:"PostAuthentication,omitempty"`
	PostConfirmation            interface{} `yaml:"PostConfirmation,omitempty"`
	PreAuthentication           interface{} `yaml:"PreAuthentication,omitempty"`
	PreSignUp                   interface{} `yaml:"PreSignUp,omitempty"`
	VerifyAuthChallengeResponse interface{} `yaml:"VerifyAuthChallengeResponse,omitempty"`
}

// UserPoolLambdaConfig validation
func (resource UserPoolLambdaConfig) Validate() []error {
	errs := []error{}

	return errs
}