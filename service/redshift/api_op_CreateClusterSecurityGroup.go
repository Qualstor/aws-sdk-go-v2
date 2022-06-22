// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon Redshift security group. You use security groups to control
// access to non-VPC clusters. For information about managing security groups, go
// to Amazon Redshift Cluster Security Groups
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) CreateClusterSecurityGroup(ctx context.Context, params *CreateClusterSecurityGroupInput, optFns ...func(*Options)) (*CreateClusterSecurityGroupOutput, error) {
	if params == nil {
		params = &CreateClusterSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateClusterSecurityGroup", params, optFns, c.addOperationCreateClusterSecurityGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateClusterSecurityGroupInput struct {

	// The name for the security group. Amazon Redshift stores the value as a lowercase
	// string. Constraints:
	//
	// * Must contain no more than 255 alphanumeric characters or
	// hyphens.
	//
	// * Must not be "Default".
	//
	// * Must be unique for all security groups
	// that are created by your Amazon Web Services account.
	//
	// Example:
	// examplesecuritygroup
	//
	// This member is required.
	ClusterSecurityGroupName *string

	// A description for the security group.
	//
	// This member is required.
	Description *string

	// A list of tag instances.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateClusterSecurityGroupOutput struct {

	// Describes a security group.
	ClusterSecurityGroup *types.ClusterSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateClusterSecurityGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateClusterSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateClusterSecurityGroup{}, middleware.After)
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
	if err = addOpCreateClusterSecurityGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateClusterSecurityGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateClusterSecurityGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateClusterSecurityGroup",
	}
}
