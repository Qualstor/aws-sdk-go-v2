// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified Call Analytics category with new rules. Note that the
// UpdateCallAnalyticsCategory operation overwrites all existing rules contained in
// the specified category. You cannot append additional rules onto an existing
// category. To create a new category, see .
func (c *Client) UpdateCallAnalyticsCategory(ctx context.Context, params *UpdateCallAnalyticsCategoryInput, optFns ...func(*Options)) (*UpdateCallAnalyticsCategoryOutput, error) {
	if params == nil {
		params = &UpdateCallAnalyticsCategoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCallAnalyticsCategory", params, optFns, c.addOperationUpdateCallAnalyticsCategoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCallAnalyticsCategoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCallAnalyticsCategoryInput struct {

	// The name of the Call Analytics category you want to update. Category names are
	// case sensitive.
	//
	// This member is required.
	CategoryName *string

	// The rules used for the updated Call Analytics category. The rules you provide in
	// this field replace the ones that are currently being used in the specified
	// category.
	//
	// This member is required.
	Rules []types.Rule

	noSmithyDocumentSerde
}

type UpdateCallAnalyticsCategoryOutput struct {

	// Provides you with the properties of the Call Analytics category you specified in
	// your UpdateCallAnalyticsCategory request.
	CategoryProperties *types.CategoryProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCallAnalyticsCategoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCallAnalyticsCategory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCallAnalyticsCategory{}, middleware.After)
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
	if err = addOpUpdateCallAnalyticsCategoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCallAnalyticsCategory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCallAnalyticsCategory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "UpdateCallAnalyticsCategory",
	}
}
