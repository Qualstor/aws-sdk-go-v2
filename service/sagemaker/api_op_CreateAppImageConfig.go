// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a configuration for running a SageMaker image as a KernelGateway app.
// The configuration specifies the Amazon Elastic File System (EFS) storage volume
// on the image, and a list of the kernels in the image.
func (c *Client) CreateAppImageConfig(ctx context.Context, params *CreateAppImageConfigInput, optFns ...func(*Options)) (*CreateAppImageConfigOutput, error) {
	if params == nil {
		params = &CreateAppImageConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppImageConfig", params, optFns, c.addOperationCreateAppImageConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppImageConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppImageConfigInput struct {

	// The name of the AppImageConfig. Must be unique to your account.
	//
	// This member is required.
	AppImageConfigName *string

	// The KernelGatewayImageConfig.
	KernelGatewayImageConfig *types.KernelGatewayImageConfig

	// A list of tags to apply to the AppImageConfig.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAppImageConfigOutput struct {

	// The Amazon Resource Name (ARN) of the AppImageConfig.
	AppImageConfigArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppImageConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAppImageConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAppImageConfig{}, middleware.After)
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
	if err = addOpCreateAppImageConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppImageConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAppImageConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateAppImageConfig",
	}
}
