// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update template sync configuration parameters, except for the templateName and
// templateType.
func (c *Client) UpdateTemplateSyncConfig(ctx context.Context, params *UpdateTemplateSyncConfigInput, optFns ...func(*Options)) (*UpdateTemplateSyncConfigOutput, error) {
	if params == nil {
		params = &UpdateTemplateSyncConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTemplateSyncConfig", params, optFns, c.addOperationUpdateTemplateSyncConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTemplateSyncConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTemplateSyncConfigInput struct {

	// The repository branch.
	//
	// This member is required.
	Branch *string

	// The name of the repository (for example, myrepos/myrepo).
	//
	// This member is required.
	RepositoryName *string

	// The repository provider.
	//
	// This member is required.
	RepositoryProvider types.RepositoryProvider

	// The synced template name.
	//
	// This member is required.
	TemplateName *string

	// The synced template type.
	//
	// This member is required.
	TemplateType types.TemplateType

	// A subdirectory path to your template bundle version. When included, limits the
	// template bundle search to this repository directory.
	Subdirectory *string

	noSmithyDocumentSerde
}

type UpdateTemplateSyncConfigOutput struct {

	// The template sync configuration detail data that's returned by Proton.
	TemplateSyncConfig *types.TemplateSyncConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTemplateSyncConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateTemplateSyncConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateTemplateSyncConfig{}, middleware.After)
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
	if err = addOpUpdateTemplateSyncConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTemplateSyncConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTemplateSyncConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "UpdateTemplateSyncConfig",
	}
}
