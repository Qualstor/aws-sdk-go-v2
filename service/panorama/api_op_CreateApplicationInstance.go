// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an application instance and deploys it to a device.
func (c *Client) CreateApplicationInstance(ctx context.Context, params *CreateApplicationInstanceInput, optFns ...func(*Options)) (*CreateApplicationInstanceOutput, error) {
	if params == nil {
		params = &CreateApplicationInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApplicationInstance", params, optFns, c.addOperationCreateApplicationInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApplicationInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApplicationInstanceInput struct {

	// A device's ID.
	//
	// This member is required.
	DefaultRuntimeContextDevice *string

	// The application's manifest document.
	//
	// This member is required.
	ManifestPayload types.ManifestPayload

	// The ID of an application instance to replace with the new instance.
	ApplicationInstanceIdToReplace *string

	// A description for the application instance.
	Description *string

	// Setting overrides for the application manifest.
	ManifestOverridesPayload types.ManifestOverridesPayload

	// A name for the application instance.
	Name *string

	// The ARN of a runtime role for the application instance.
	RuntimeRoleArn *string

	// Tags for the application instance.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateApplicationInstanceOutput struct {

	// The application instance's ID.
	//
	// This member is required.
	ApplicationInstanceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateApplicationInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApplicationInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApplicationInstance{}, middleware.After)
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
	if err = addOpCreateApplicationInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApplicationInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApplicationInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "panorama",
		OperationName: "CreateApplicationInstance",
	}
}
