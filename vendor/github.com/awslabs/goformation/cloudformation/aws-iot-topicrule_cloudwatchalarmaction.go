package cloudformation

// AWSIoTTopicRule_CloudwatchAlarmAction AWS CloudFormation Resource (AWS::IoT::TopicRule.CloudwatchAlarmAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchalarmaction.html
type AWSIoTTopicRule_CloudwatchAlarmAction struct {

	// AlarmName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchalarmaction.html#cfn-iot-topicrule-cloudwatchalarmaction-alarmname
	AlarmName *StringIntrinsic `json:"AlarmName,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchalarmaction.html#cfn-iot-topicrule-cloudwatchalarmaction-rolearn
	RoleArn *StringIntrinsic `json:"RoleArn,omitempty"`

	// StateReason AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchalarmaction.html#cfn-iot-topicrule-cloudwatchalarmaction-statereason
	StateReason *StringIntrinsic `json:"StateReason,omitempty"`

	// StateValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-cloudwatchalarmaction.html#cfn-iot-topicrule-cloudwatchalarmaction-statevalue
	StateValue *StringIntrinsic `json:"StateValue,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_CloudwatchAlarmAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.CloudwatchAlarmAction"
}
