// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakera2iruntime

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/sagemakera2iruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a human loop, provided that at least one activation condition is met.
func (c *Client) StartHumanLoop(ctx context.Context, params *StartHumanLoopInput, optFns ...func(*Options)) (*StartHumanLoopOutput, error) {
	if params == nil {
		params = &StartHumanLoopInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartHumanLoop", params, optFns, c.addOperationStartHumanLoopMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartHumanLoopOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartHumanLoopInput struct {

	// The Amazon Resource Name (ARN) of the flow definition associated with this human
	// loop.
	//
	// This member is required.
	FlowDefinitionArn *string

	// An object that contains information about the human loop.
	//
	// This member is required.
	HumanLoopInput *types.HumanLoopInput

	// The name of the human loop.
	//
	// This member is required.
	HumanLoopName *string

	// Attributes of the specified data. Use DataAttributes to specify if your data is
	// free of personally identifiable information and/or free of adult content.
	DataAttributes *types.HumanLoopDataAttributes

	noSmithyDocumentSerde
}

type StartHumanLoopOutput struct {

	// The Amazon Resource Name (ARN) of the human loop.
	HumanLoopArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartHumanLoopMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartHumanLoop{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartHumanLoop{}, middleware.After)
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
	if err = addOpStartHumanLoopValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartHumanLoop(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartHumanLoop(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "StartHumanLoop",
	}
}
