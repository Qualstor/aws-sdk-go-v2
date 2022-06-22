// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new major or minor version of a service template. A major version of a
// service template is a version that isn't backward compatible. A minor version of
// a service template is a version that's backward compatible within its major
// version.
func (c *Client) CreateServiceTemplateVersion(ctx context.Context, params *CreateServiceTemplateVersionInput, optFns ...func(*Options)) (*CreateServiceTemplateVersionOutput, error) {
	if params == nil {
		params = &CreateServiceTemplateVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServiceTemplateVersion", params, optFns, c.addOperationCreateServiceTemplateVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceTemplateVersionInput struct {

	// An array of environment template objects that are compatible with the new
	// service template version. A service instance based on this service template
	// version can run in environments based on compatible templates.
	//
	// This member is required.
	CompatibleEnvironmentTemplates []types.CompatibleEnvironmentTemplateInput

	// An object that includes the template bundle S3 bucket path and name for the new
	// version of a service template.
	//
	// This member is required.
	Source types.TemplateVersionSourceInput

	// The name of the service template.
	//
	// This member is required.
	TemplateName *string

	// When included, if two identical requests are made with the same client token,
	// Proton returns the service template version that the first request created.
	ClientToken *string

	// A description of the new version of a service template.
	Description *string

	// To create a new minor version of the service template, include a major Version.
	// To create a new major and minor version of the service template, exclude major
	// Version.
	MajorVersion *string

	// An array of supported component sources. Components with supported sources can
	// be attached to service instances based on this service template version. For
	// more information about components, see Proton components
	// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-components.html) in the
	// Proton Administrator Guide.
	SupportedComponentSources []types.ServiceTemplateSupportedComponentSourceType

	// An optional list of metadata items that you can associate with the Proton
	// service template version. A tag is a key-value pair. For more information, see
	// Proton resources and tagging in the Proton Administrator Guide
	// (https://docs.aws.amazon.com/proton/latest/adminguide/resources.html) or Proton
	// User Guide (https://docs.aws.amazon.com/proton/latest/userguide/resources.html).
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateServiceTemplateVersionOutput struct {

	// The service template version summary of detail data that's returned by Proton.
	//
	// This member is required.
	ServiceTemplateVersion *types.ServiceTemplateVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceTemplateVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateServiceTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateServiceTemplateVersion{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateServiceTemplateVersionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateServiceTemplateVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServiceTemplateVersion(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateServiceTemplateVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateServiceTemplateVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateServiceTemplateVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateServiceTemplateVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateServiceTemplateVersionInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateServiceTemplateVersionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateServiceTemplateVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateServiceTemplateVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "CreateServiceTemplateVersion",
	}
}
