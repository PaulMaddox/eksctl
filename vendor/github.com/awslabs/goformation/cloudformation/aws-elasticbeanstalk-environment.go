package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSElasticBeanstalkEnvironment AWS CloudFormation Resource (AWS::ElasticBeanstalk::Environment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html
type AWSElasticBeanstalkEnvironment struct {

	// ApplicationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-applicationname
	ApplicationName *StringIntrinsic `json:"ApplicationName,omitempty"`

	// CNAMEPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-cnameprefix
	CNAMEPrefix *StringIntrinsic `json:"CNAMEPrefix,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-description
	Description *StringIntrinsic `json:"Description,omitempty"`

	// EnvironmentName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-name
	EnvironmentName *StringIntrinsic `json:"EnvironmentName,omitempty"`

	// OptionSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-optionsettings
	OptionSettings []AWSElasticBeanstalkEnvironment_OptionSetting `json:"OptionSettings,omitempty"`

	// PlatformArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn
	PlatformArn *StringIntrinsic `json:"PlatformArn,omitempty"`

	// SolutionStackName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-solutionstackname
	SolutionStackName *StringIntrinsic `json:"SolutionStackName,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-elasticbeanstalk-environment-tags
	Tags []Tag `json:"Tags,omitempty"`

	// TemplateName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-templatename
	TemplateName *StringIntrinsic `json:"TemplateName,omitempty"`

	// Tier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-tier
	Tier *AWSElasticBeanstalkEnvironment_Tier `json:"Tier,omitempty"`

	// VersionLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-versionlabel
	VersionLabel *StringIntrinsic `json:"VersionLabel,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkEnvironment) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::Environment"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSElasticBeanstalkEnvironment) MarshalJSON() ([]byte, error) {
	type Properties AWSElasticBeanstalkEnvironment
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSElasticBeanstalkEnvironment) UnmarshalJSON(b []byte) error {
	type Properties AWSElasticBeanstalkEnvironment
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSElasticBeanstalkEnvironment(*res.Properties)
	}

	return nil
}

// GetAllAWSElasticBeanstalkEnvironmentResources retrieves all AWSElasticBeanstalkEnvironment items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticBeanstalkEnvironmentResources() map[string]AWSElasticBeanstalkEnvironment {
	results := map[string]AWSElasticBeanstalkEnvironment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSElasticBeanstalkEnvironment:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ElasticBeanstalk::Environment" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSElasticBeanstalkEnvironment
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSElasticBeanstalkEnvironmentWithName retrieves all AWSElasticBeanstalkEnvironment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticBeanstalkEnvironmentWithName(name string) (AWSElasticBeanstalkEnvironment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSElasticBeanstalkEnvironment:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ElasticBeanstalk::Environment" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSElasticBeanstalkEnvironment
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSElasticBeanstalkEnvironment{}, errors.New("resource not found")
}
