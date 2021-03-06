package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSAutoScalingLaunchConfiguration AWS CloudFormation Resource (AWS::AutoScaling::LaunchConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html
type AWSAutoScalingLaunchConfiguration struct {

	// AssociatePublicIpAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cf-as-launchconfig-associatepubip
	AssociatePublicIpAddress bool `json:"AssociatePublicIpAddress,omitempty"`

	// BlockDeviceMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-blockdevicemappings
	BlockDeviceMappings []AWSAutoScalingLaunchConfiguration_BlockDeviceMapping `json:"BlockDeviceMappings,omitempty"`

	// ClassicLinkVPCId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-classiclinkvpcid
	ClassicLinkVPCId *StringIntrinsic `json:"ClassicLinkVPCId,omitempty"`

	// ClassicLinkVPCSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-classiclinkvpcsecuritygroups
	ClassicLinkVPCSecurityGroups []*StringIntrinsic `json:"ClassicLinkVPCSecurityGroups,omitempty"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-ebsoptimized
	EbsOptimized bool `json:"EbsOptimized,omitempty"`

	// IamInstanceProfile AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-iaminstanceprofile
	IamInstanceProfile *StringIntrinsic `json:"IamInstanceProfile,omitempty"`

	// ImageId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-imageid
	ImageId *StringIntrinsic `json:"ImageId,omitempty"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instanceid
	InstanceId *StringIntrinsic `json:"InstanceId,omitempty"`

	// InstanceMonitoring AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instancemonitoring
	InstanceMonitoring bool `json:"InstanceMonitoring,omitempty"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instancetype
	InstanceType *StringIntrinsic `json:"InstanceType,omitempty"`

	// KernelId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-kernelid
	KernelId *StringIntrinsic `json:"KernelId,omitempty"`

	// KeyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-keyname
	KeyName *StringIntrinsic `json:"KeyName,omitempty"`

	// LaunchConfigurationName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-autoscaling-launchconfig-launchconfigurationname
	LaunchConfigurationName *StringIntrinsic `json:"LaunchConfigurationName,omitempty"`

	// PlacementTenancy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-placementtenancy
	PlacementTenancy *StringIntrinsic `json:"PlacementTenancy,omitempty"`

	// RamDiskId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-ramdiskid
	RamDiskId *StringIntrinsic `json:"RamDiskId,omitempty"`

	// SecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-securitygroups
	SecurityGroups []*StringIntrinsic `json:"SecurityGroups,omitempty"`

	// SpotPrice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-spotprice
	SpotPrice *StringIntrinsic `json:"SpotPrice,omitempty"`

	// UserData AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-userdata
	UserData *StringIntrinsic `json:"UserData,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingLaunchConfiguration) AWSCloudFormationType() string {
	return "AWS::AutoScaling::LaunchConfiguration"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSAutoScalingLaunchConfiguration) MarshalJSON() ([]byte, error) {
	type Properties AWSAutoScalingLaunchConfiguration
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
func (r *AWSAutoScalingLaunchConfiguration) UnmarshalJSON(b []byte) error {
	type Properties AWSAutoScalingLaunchConfiguration
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
		*r = AWSAutoScalingLaunchConfiguration(*res.Properties)
	}

	return nil
}

// GetAllAWSAutoScalingLaunchConfigurationResources retrieves all AWSAutoScalingLaunchConfiguration items from an AWS CloudFormation template
func (t *Template) GetAllAWSAutoScalingLaunchConfigurationResources() map[string]AWSAutoScalingLaunchConfiguration {
	results := map[string]AWSAutoScalingLaunchConfiguration{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSAutoScalingLaunchConfiguration:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AutoScaling::LaunchConfiguration" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAutoScalingLaunchConfiguration
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

// GetAWSAutoScalingLaunchConfigurationWithName retrieves all AWSAutoScalingLaunchConfiguration items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAutoScalingLaunchConfigurationWithName(name string) (AWSAutoScalingLaunchConfiguration, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSAutoScalingLaunchConfiguration:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AutoScaling::LaunchConfiguration" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAutoScalingLaunchConfiguration
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSAutoScalingLaunchConfiguration{}, errors.New("resource not found")
}
