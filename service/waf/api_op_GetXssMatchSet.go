// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/waf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Returns the XssMatchSet that is specified by XssMatchSetId.
func (c *Client) GetXssMatchSet(ctx context.Context, params *GetXssMatchSetInput, optFns ...func(*Options)) (*GetXssMatchSetOutput, error) {
	if params == nil {
		params = &GetXssMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetXssMatchSet", params, optFns, c.addOperationGetXssMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetXssMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get an XssMatchSet.
type GetXssMatchSetInput struct {

	// The XssMatchSetId of the XssMatchSet that you want to get. XssMatchSetId is
	// returned by CreateXssMatchSet and by ListXssMatchSets.
	//
	// This member is required.
	XssMatchSetId *string

	noSmithyDocumentSerde
}

// The response to a GetXssMatchSet request.
type GetXssMatchSetOutput struct {

	// Information about the XssMatchSet that you specified in the GetXssMatchSet
	// request. For more information, see the following topics:
	//
	// * XssMatchSet:
	// Contains Name, XssMatchSetId, and an array of XssMatchTuple objects
	//
	// *
	// XssMatchTuple: Each XssMatchTuple object contains FieldToMatch and
	// TextTransformation
	//
	// * FieldToMatch: Contains Data and Type
	XssMatchSet *types.XssMatchSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetXssMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetXssMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetXssMatchSet{}, middleware.After)
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
	if err = addOpGetXssMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetXssMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetXssMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "GetXssMatchSet",
	}
}
