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

// Updates a Lambda function's code. If code signing is enabled for the function,
// the code package must be signed by a trusted publisher. For more information,
// see Configuring code signing
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-trustedcode.html).
// If the function's package type is Image, you must specify the code package in
// ImageUri as the URI of a container image
// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) in the Amazon
// ECR registry. If the function's package type is Zip, you must specify the
// deployment package as a .zip file archive
// (https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip).
// Enter the Amazon S3 bucket and key of the code .zip file location. You can also
// provide the function code inline using the ZipFile field. The code in the
// deployment package must be compatible with the target instruction set
// architecture of the function (x86-64 or arm64). The function's code is locked
// when you publish a version. You can't modify the code of a published version,
// only the unpublished version. For a function defined as a container image,
// Lambda resolves the image tag to an image digest. In Amazon ECR, if you update
// the image tag to a new image, Lambda does not automatically update the function.
func (c *Client) UpdateFunctionCode(ctx context.Context, params *UpdateFunctionCodeInput, optFns ...func(*Options)) (*UpdateFunctionCodeOutput, error) {
	if params == nil {
		params = &UpdateFunctionCodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFunctionCode", params, optFns, c.addOperationUpdateFunctionCodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFunctionCodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFunctionCodeInput struct {

	// The name of the Lambda function. Name formats
	//
	// * Function name - my-function.
	//
	// *
	// Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// *
	// Partial ARN - 123456789012:function:my-function.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// The instruction set architecture that the function supports. Enter a string
	// array with one of the valid values (arm64 or x86_64). The default value is
	// x86_64.
	Architectures []types.Architecture

	// Set to true to validate the request parameters and access permissions without
	// modifying the function code.
	DryRun bool

	// URI of a container image in the Amazon ECR registry. Do not use for a function
	// defined with a .zip file archive.
	ImageUri *string

	// Set to true to publish a new version of the function after updating the code.
	// This has the same effect as calling PublishVersion separately.
	Publish bool

	// Only update the function if the revision ID matches the ID that's specified. Use
	// this option to avoid modifying a function that has changed since you last read
	// it.
	RevisionId *string

	// An Amazon S3 bucket in the same Amazon Web Services Region as your function. The
	// bucket can be in a different Amazon Web Services account. Use only with a
	// function defined with a .zip file archive deployment package.
	S3Bucket *string

	// The Amazon S3 key of the deployment package. Use only with a function defined
	// with a .zip file archive deployment package.
	S3Key *string

	// For versioned objects, the version of the deployment package object to use.
	S3ObjectVersion *string

	// The base64-encoded contents of the deployment package. Amazon Web Services SDK
	// and Amazon Web Services CLI clients handle the encoding for you. Use only with a
	// function defined with a .zip file archive deployment package.
	ZipFile []byte

	noSmithyDocumentSerde
}

// Details about a function's configuration.
type UpdateFunctionCodeOutput struct {

	// The instruction set architecture that the function supports. Architecture is a
	// string array with one of the valid values. The default architecture value is
	// x86_64.
	Architectures []types.Architecture

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string

	// The size of the function's deployment package, in bytes.
	CodeSize int64

	// The function's dead letter queue.
	DeadLetterConfig *types.DeadLetterConfig

	// The function's description.
	Description *string

	// The function's environment variables
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html).
	Environment *types.EnvironmentResponse

	// The size of the function’s /tmp directory in MB. The default value is 512, but
	// can be any whole number between 512 and 10240 MB.
	EphemeralStorage *types.EphemeralStorage

	// Connection settings for an Amazon EFS file system
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html).
	FileSystemConfigs []types.FileSystemConfig

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string

	// The name of the function.
	FunctionName *string

	// The function that Lambda calls to begin executing your function.
	Handler *string

	// The function's image configuration values.
	ImageConfigResponse *types.ImageConfigResponse

	// The KMS key that's used to encrypt the function's environment variables. This
	// key is only returned if you've configured a customer managed key.
	KMSKeyArn *string

	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string

	// The status of the last update that was performed on the function. This is first
	// set to Successful after function creation completes.
	LastUpdateStatus types.LastUpdateStatus

	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string

	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode types.LastUpdateStatusReasonCode

	// The function's  layers
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	Layers []types.Layer

	// For Lambda@Edge functions, the ARN of the main function.
	MasterArn *string

	// The amount of memory available to the function at runtime.
	MemorySize *int32

	// The type of deployment package. Set to Image for container image and set Zip for
	// .zip file archive.
	PackageType types.PackageType

	// The latest updated revision of the function or alias.
	RevisionId *string

	// The function's execution role.
	Role *string

	// The runtime environment for the Lambda function.
	Runtime types.Runtime

	// The ARN of the signing job.
	SigningJobArn *string

	// The ARN of the signing profile version.
	SigningProfileVersionArn *string

	// The current state of the function. When the state is Inactive, you can
	// reactivate the function by invoking it.
	State types.State

	// The reason for the function's current state.
	StateReason *string

	// The reason code for the function's current state. When the code is Creating, you
	// can't invoke or modify the function.
	StateReasonCode types.StateReasonCode

	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int32

	// The function's X-Ray tracing configuration.
	TracingConfig *types.TracingConfigResponse

	// The version of the Lambda function.
	Version *string

	// The function's networking configuration.
	VpcConfig *types.VpcConfigResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFunctionCodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFunctionCode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFunctionCode{}, middleware.After)
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
	if err = addOpUpdateFunctionCodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFunctionCode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFunctionCode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "UpdateFunctionCode",
	}
}
