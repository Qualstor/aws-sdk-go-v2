// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an workgroup in Amazon Redshift Serverless.
func (c *Client) CreateWorkgroup(ctx context.Context, params *CreateWorkgroupInput, optFns ...func(*Options)) (*CreateWorkgroupOutput, error) {
	if params == nil {
		params = &CreateWorkgroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkgroup", params, optFns, c.addOperationCreateWorkgroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkgroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkgroupInput struct {

	// The name of the namespace to associate with the workgroup.
	//
	// This member is required.
	NamespaceName *string

	// The name of the created workgroup.
	//
	// This member is required.
	WorkgroupName *string

	// The base data warehouse capacity of the workgroup in Redshift Processing Units
	// (RPUs).
	BaseCapacity *int32

	// An array of parameters to set for more control over a serverless database. The
	// options are datestyle, enable_user_activity_logging, query_group, search_path,
	// and max_query_execution_time.
	ConfigParameters []types.ConfigParameter

	// The value that specifies whether to turn on enhanced virtual private cloud (VPC)
	// routing, which forces Amazon Redshift Serverless to route traffic through your
	// VPC instead of over the internet.
	EnhancedVpcRouting *bool

	// A value that specifies whether the workgroup can be accessed from a public
	// network.
	PubliclyAccessible *bool

	// An array of security group IDs to associate with the workgroup.
	SecurityGroupIds []string

	// An array of VPC subnet IDs to associate with the workgroup.
	SubnetIds []string

	// A array of tag instances.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateWorkgroupOutput struct {

	// The created workgroup object.
	Workgroup *types.Workgroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkgroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWorkgroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWorkgroup{}, middleware.After)
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
	if err = addOpCreateWorkgroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkgroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkgroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-serverless",
		OperationName: "CreateWorkgroup",
	}
}
