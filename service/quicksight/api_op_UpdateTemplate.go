// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a template from an existing Amazon QuickSight analysis or another
// template.
func (c *Client) UpdateTemplate(ctx context.Context, params *UpdateTemplateInput, optFns ...func(*Options)) (*UpdateTemplateOutput, error) {
	if params == nil {
		params = &UpdateTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTemplate", params, optFns, c.addOperationUpdateTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTemplateInput struct {

	// The ID of the Amazon Web Services account that contains the template that you're
	// updating.
	//
	// This member is required.
	AwsAccountId *string

	// The entity that you are using as a source when you update the template. In
	// SourceEntity, you specify the type of object you're using as source:
	// SourceTemplate for a template or SourceAnalysis for an analysis. Both of these
	// require an Amazon Resource Name (ARN). For SourceTemplate, specify the ARN of
	// the source template. For SourceAnalysis, specify the ARN of the source analysis.
	// The SourceTemplate ARN can contain any Amazon Web Services account and any
	// Amazon QuickSight-supported Amazon Web Services Region;. Use the
	// DataSetReferences entity within SourceTemplate or SourceAnalysis to list the
	// replacement datasets for the placeholders listed in the original. The schema in
	// each dataset must match its placeholder.
	//
	// This member is required.
	SourceEntity *types.TemplateSourceEntity

	// The ID for the template.
	//
	// This member is required.
	TemplateId *string

	// The name for the template.
	Name *string

	// A description of the current template version that is being updated. Every time
	// you call UpdateTemplate, you create a new version of the template. Each version
	// of the template maintains a description of the version in the VersionDescription
	// field.
	VersionDescription *string

	noSmithyDocumentSerde
}

type UpdateTemplateOutput struct {

	// The Amazon Resource Name (ARN) for the template.
	Arn *string

	// The creation status of the template.
	CreationStatus types.ResourceStatus

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The ID for the template.
	TemplateId *string

	// The ARN for the template, including the version information of the first
	// version.
	VersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTemplate{}, middleware.After)
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
	if err = addOpUpdateTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateTemplate",
	}
}
