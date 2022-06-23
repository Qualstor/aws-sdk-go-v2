// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates one or more targets with an event window. Only one type of target
// (instance IDs, Dedicated Host IDs, or tags) can be specified with an event
// window. For more information, see Define event windows for scheduled events
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html) in the
// Amazon EC2 User Guide.
func (c *Client) AssociateInstanceEventWindow(ctx context.Context, params *AssociateInstanceEventWindowInput, optFns ...func(*Options)) (*AssociateInstanceEventWindowOutput, error) {
	if params == nil {
		params = &AssociateInstanceEventWindowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateInstanceEventWindow", params, optFns, c.addOperationAssociateInstanceEventWindowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateInstanceEventWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateInstanceEventWindowInput struct {

	// One or more targets associated with the specified event window.
	//
	// This member is required.
	AssociationTarget *types.InstanceEventWindowAssociationRequest

	// The ID of the event window.
	//
	// This member is required.
	InstanceEventWindowId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type AssociateInstanceEventWindowOutput struct {

	// Information about the event window.
	InstanceEventWindow *types.InstanceEventWindow

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateInstanceEventWindowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssociateInstanceEventWindow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateInstanceEventWindow{}, middleware.After)
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
	if err = addOpAssociateInstanceEventWindowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateInstanceEventWindow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateInstanceEventWindow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssociateInstanceEventWindow",
	}
}
