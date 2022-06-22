// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the code signing configuration for the specified function.
func (c *Client) GetFunctionCodeSigningConfig(ctx context.Context, params *GetFunctionCodeSigningConfigInput, optFns ...func(*Options)) (*GetFunctionCodeSigningConfigOutput, error) {
	if params == nil {
		params = &GetFunctionCodeSigningConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFunctionCodeSigningConfig", params, optFns, c.addOperationGetFunctionCodeSigningConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFunctionCodeSigningConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFunctionCodeSigningConfigInput struct {

	// The name of the Lambda function. Name formats
	//
	// * Function name - MyFunction.
	//
	// *
	// Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	// *
	// Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	noSmithyDocumentSerde
}

type GetFunctionCodeSigningConfigOutput struct {

	// The The Amazon Resource Name (ARN) of the code signing configuration.
	//
	// This member is required.
	CodeSigningConfigArn *string

	// The name of the Lambda function. Name formats
	//
	// * Function name - MyFunction.
	//
	// *
	// Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	// *
	// Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFunctionCodeSigningConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFunctionCodeSigningConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFunctionCodeSigningConfig{}, middleware.After)
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
	if err = addOpGetFunctionCodeSigningConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFunctionCodeSigningConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFunctionCodeSigningConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetFunctionCodeSigningConfig",
	}
}
