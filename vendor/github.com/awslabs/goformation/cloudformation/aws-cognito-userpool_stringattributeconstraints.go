package cloudformation

// AWSCognitoUserPool_StringAttributeConstraints AWS CloudFormation Resource (AWS::Cognito::UserPool.StringAttributeConstraints)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html
type AWSCognitoUserPool_StringAttributeConstraints struct {

	// MaxLength AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html#cfn-cognito-userpool-stringattributeconstraints-maxlength
	MaxLength *StringIntrinsic `json:"MaxLength,omitempty"`

	// MinLength AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html#cfn-cognito-userpool-stringattributeconstraints-minlength
	MinLength *StringIntrinsic `json:"MinLength,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_StringAttributeConstraints) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.StringAttributeConstraints"
}
