package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DeploymentGroupRevisionLocation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html
type DeploymentGroupRevisionLocation struct {
	RevisionType   interface{}                    `yaml:"RevisionType,omitempty"`
	S3Location     *DeploymentGroupS3Location     `yaml:"S3Location,omitempty"`
	GitHubLocation *DeploymentGroupGitHubLocation `yaml:"GitHubLocation,omitempty"`
}

// DeploymentGroupRevisionLocation validation
func (resource DeploymentGroupRevisionLocation) Validate() []error {
	errors := []error{}

	return errors
}
