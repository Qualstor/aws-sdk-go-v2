// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specifies the configuration data for a registered CloudFormation extension, in
// the given account and region. To view the current configuration data for an
// extension, refer to the ConfigurationSchema element of DescribeType. For more
// information, see Configuring extensions at the account level
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-register.html#registry-set-configuration)
// in the CloudFormation User Guide. It's strongly recommended that you use dynamic
// references to restrict sensitive configuration definitions, such as third-party
// credentials. For more details on dynamic references, see Using dynamic
// references to specify template values (https://docs.aws.amazon.com/) in the
// CloudFormation User Guide.
func (c *Client) SetTypeConfiguration(ctx context.Context, params *SetTypeConfigurationInput, optFns ...func(*Options)) (*SetTypeConfigurationOutput, error) {
	if params == nil {
		params = &SetTypeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetTypeConfiguration", params, optFns, c.addOperationSetTypeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetTypeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetTypeConfigurationInput struct {

	// The configuration data for the extension, in this account and region. The
	// configuration data must be formatted as JSON, and validate against the schema
	// returned in the ConfigurationSchema response element of API_DescribeType. For
	// more information, see Defining account-level configuration data for an extension
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-model.html#resource-type-howto-configuration)
	// in the CloudFormation CLI User Guide.
	//
	// This member is required.
	Configuration *string

	// An alias by which to refer to this extension configuration data. Conditional:
	// Specifying a configuration alias is required when setting a configuration for a
	// resource type extension.
	ConfigurationAlias *string

	// The type of extension. Conditional: You must specify ConfigurationArn, or Type
	// and TypeName.
	Type types.ThirdPartyType

	// The Amazon Resource Name (ARN) for the extension, in this account and region.
	// For public extensions, this will be the ARN assigned when you activate the type
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ActivateType.html)
	// in this account and region. For private extensions, this will be the ARN
	// assigned when you register the type
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterType.html)
	// in this account and region. Do not include the extension versions suffix at the
	// end of the ARN. You can set the configuration for an extension, but not for a
	// specific extension version.
	TypeArn *string

	// The name of the extension. Conditional: You must specify ConfigurationArn, or
	// Type and TypeName.
	TypeName *string

	noSmithyDocumentSerde
}

type SetTypeConfigurationOutput struct {

	// The Amazon Resource Name (ARN) for the configuration data, in this account and
	// region. Conditional: You must specify ConfigurationArn, or Type and TypeName.
	ConfigurationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetTypeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetTypeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetTypeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSetTypeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetTypeConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSetTypeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "SetTypeConfiguration",
	}
}
