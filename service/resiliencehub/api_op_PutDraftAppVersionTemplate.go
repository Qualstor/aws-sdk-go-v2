// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates the app template for a draft version of a Resilience Hub app.
func (c *Client) PutDraftAppVersionTemplate(ctx context.Context, params *PutDraftAppVersionTemplateInput, optFns ...func(*Options)) (*PutDraftAppVersionTemplateOutput, error) {
	if params == nil {
		params = &PutDraftAppVersionTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDraftAppVersionTemplate", params, optFns, c.addOperationPutDraftAppVersionTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDraftAppVersionTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutDraftAppVersionTemplateInput struct {

	// The Amazon Resource Name (ARN) of the application. The format for this ARN is:
	// arn:partition:resiliencehub:region:account:app/app-id. For more information
	// about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	AppArn *string

	// A JSON string that contains the body of the app template.
	//
	// This member is required.
	AppTemplateBody *string

	noSmithyDocumentSerde
}

type PutDraftAppVersionTemplateOutput struct {

	// The Amazon Resource Name (ARN) of the application. The format for this ARN is:
	// arn:partition:resiliencehub:region:account:app/app-id. For more information
	// about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	AppArn *string

	// The version of the application.
	AppVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutDraftAppVersionTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutDraftAppVersionTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutDraftAppVersionTemplate{}, middleware.After)
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
	if err = addOpPutDraftAppVersionTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutDraftAppVersionTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutDraftAppVersionTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "PutDraftAppVersionTemplate",
	}
}
