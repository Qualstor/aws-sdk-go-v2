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

// Returns hook-related information for the change set and a list of changes that
// CloudFormation makes when you run the change set.
func (c *Client) DescribeChangeSetHooks(ctx context.Context, params *DescribeChangeSetHooksInput, optFns ...func(*Options)) (*DescribeChangeSetHooksOutput, error) {
	if params == nil {
		params = &DescribeChangeSetHooksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeChangeSetHooks", params, optFns, c.addOperationDescribeChangeSetHooksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeChangeSetHooksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeChangeSetHooksInput struct {

	// The name or Amazon Resource Name (ARN) of the change set that you want to
	// describe.
	//
	// This member is required.
	ChangeSetName *string

	// If specified, lists only the hooks related to the specified LogicalResourceId.
	LogicalResourceId *string

	// A string, provided by the DescribeChangeSetHooks response output, that
	// identifies the next page of information that you want to retrieve.
	NextToken *string

	// If you specified the name of a change set, specify the stack name or stack ID
	// (ARN) of the change set you want to describe.
	StackName *string

	noSmithyDocumentSerde
}

type DescribeChangeSetHooksOutput struct {

	// The change set identifier (stack ID).
	ChangeSetId *string

	// The change set name.
	ChangeSetName *string

	// List of hook objects.
	Hooks []types.ChangeSetHook

	// Pagination token, null or empty if no more results.
	NextToken *string

	// The stack identifier (stack ID).
	StackId *string

	// The stack name.
	StackName *string

	// Provides the status of the change set hook.
	Status types.ChangeSetHooksStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeChangeSetHooksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeChangeSetHooks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeChangeSetHooks{}, middleware.After)
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
	if err = addOpDescribeChangeSetHooksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChangeSetHooks(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeChangeSetHooks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeChangeSetHooks",
	}
}
