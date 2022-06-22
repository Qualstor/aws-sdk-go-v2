// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the code signing configuration. Changes to the code signing configuration
// take effect the next time a user tries to deploy a code package to the function.
func (c *Client) UpdateCodeSigningConfig(ctx context.Context, params *UpdateCodeSigningConfigInput, optFns ...func(*Options)) (*UpdateCodeSigningConfigOutput, error) {
	if params == nil {
		params = &UpdateCodeSigningConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCodeSigningConfig", params, optFns, c.addOperationUpdateCodeSigningConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCodeSigningConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCodeSigningConfigInput struct {

	// The The Amazon Resource Name (ARN) of the code signing configuration.
	//
	// This member is required.
	CodeSigningConfigArn *string

	// Signing profiles for this code signing configuration.
	AllowedPublishers *types.AllowedPublishers

	// The code signing policy.
	CodeSigningPolicies *types.CodeSigningPolicies

	// Descriptive name for this code signing configuration.
	Description *string

	noSmithyDocumentSerde
}

type UpdateCodeSigningConfigOutput struct {

	// The code signing configuration
	//
	// This member is required.
	CodeSigningConfig *types.CodeSigningConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCodeSigningConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCodeSigningConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCodeSigningConfig{}, middleware.After)
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
	if err = addOpUpdateCodeSigningConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCodeSigningConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCodeSigningConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "UpdateCodeSigningConfig",
	}
}
