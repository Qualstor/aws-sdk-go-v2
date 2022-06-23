// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the content and settings of a message template for messages sent
// through the in-app channel.
func (c *Client) GetInAppTemplate(ctx context.Context, params *GetInAppTemplateInput, optFns ...func(*Options)) (*GetInAppTemplateOutput, error) {
	if params == nil {
		params = &GetInAppTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInAppTemplate", params, optFns, c.addOperationGetInAppTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInAppTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInAppTemplateInput struct {

	// The name of the message template. A template name must start with an
	// alphanumeric character and can contain a maximum of 128 characters. The
	// characters can be alphanumeric characters, underscores (_), or hyphens (-).
	// Template names are case sensitive.
	//
	// This member is required.
	TemplateName *string

	// The unique identifier for the version of the message template to update,
	// retrieve information about, or delete. To retrieve identifiers and other
	// information for all the versions of a template, use the Template Versions
	// resource. If specified, this value must match the identifier for an existing
	// template version. If specified for an update operation, this value must match
	// the identifier for the latest existing version of the template. This restriction
	// helps ensure that race conditions don't occur. If you don't specify a value for
	// this parameter, Amazon Pinpoint does the following:
	//
	// * For a get operation,
	// retrieves information about the active version of the template.
	//
	// * For an update
	// operation, saves the updates to (overwrites) the latest existing version of the
	// template, if the create-new-version parameter isn't used or is set to false.
	//
	// *
	// For a delete operation, deletes the template, including all versions of the
	// template.
	Version *string

	noSmithyDocumentSerde
}

type GetInAppTemplateOutput struct {

	// In-App Template Response.
	//
	// This member is required.
	InAppTemplateResponse *types.InAppTemplateResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInAppTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInAppTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInAppTemplate{}, middleware.After)
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
	if err = addOpGetInAppTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInAppTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetInAppTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetInAppTemplate",
	}
}
