// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels an update on the specified stack. If the call completes successfully,
// the stack rolls back the update and reverts to the previous stack configuration.
// You can cancel only stacks that are in the UPDATE_IN_PROGRESS state.
func (c *Client) CancelUpdateStack(ctx context.Context, params *CancelUpdateStackInput, optFns ...func(*Options)) (*CancelUpdateStackOutput, error) {
	if params == nil {
		params = &CancelUpdateStackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelUpdateStack", params, optFns, c.addOperationCancelUpdateStackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelUpdateStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CancelUpdateStack action.
type CancelUpdateStackInput struct {

	// The name or the unique stack ID that's associated with the stack.
	//
	// This member is required.
	StackName *string

	// A unique identifier for this CancelUpdateStack request. Specify this token if
	// you plan to retry requests so that CloudFormation knows that you're not
	// attempting to cancel an update on a stack with the same name. You might retry
	// CancelUpdateStack requests to ensure that CloudFormation successfully received
	// them.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type CancelUpdateStackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelUpdateStackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCancelUpdateStack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCancelUpdateStack{}, middleware.After)
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
	if err = addOpCancelUpdateStackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelUpdateStack(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelUpdateStack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "CancelUpdateStack",
	}
}
