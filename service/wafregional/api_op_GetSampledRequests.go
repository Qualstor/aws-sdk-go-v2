// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Gets detailed information about a specified number of requests--a
// sample--that AWS WAF randomly selects from among the first 5,000 requests that
// your AWS resource received during a time range that you choose. You can specify
// a sample size of up to 500 requests, and you can specify any time range in the
// previous three hours. GetSampledRequests returns a time range, which is usually
// the time range that you specified. However, if your resource (such as a
// CloudFront distribution) received 5,000 requests before the specified time range
// elapsed, GetSampledRequests returns an updated time range. This new time range
// indicates the actual period during which AWS WAF selected the requests in the
// sample.
func (c *Client) GetSampledRequests(ctx context.Context, params *GetSampledRequestsInput, optFns ...func(*Options)) (*GetSampledRequestsOutput, error) {
	if params == nil {
		params = &GetSampledRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSampledRequests", params, optFns, c.addOperationGetSampledRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSampledRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSampledRequestsInput struct {

	// The number of requests that you want AWS WAF to return from among the first
	// 5,000 requests that your AWS resource received during the time range. If your
	// resource received fewer requests than the value of MaxItems, GetSampledRequests
	// returns information about all of them.
	//
	// This member is required.
	MaxItems int64

	// RuleId is one of three values:
	//
	// * The RuleId of the Rule or the RuleGroupId of
	// the RuleGroup for which you want GetSampledRequests to return a sample of
	// requests.
	//
	// * Default_Action, which causes GetSampledRequests to return a sample
	// of the requests that didn't match any of the rules in the specified WebACL.
	//
	// This member is required.
	RuleId *string

	// The start date and time and the end date and time of the range for which you
	// want GetSampledRequests to return a sample of requests. You must specify the
	// times in Coordinated Universal Time (UTC) format. UTC format includes the
	// special designator, Z. For example, "2016-09-27T14:50Z". You can specify any
	// time range in the previous three hours.
	//
	// This member is required.
	TimeWindow *types.TimeWindow

	// The WebACLId of the WebACL for which you want GetSampledRequests to return a
	// sample of requests.
	//
	// This member is required.
	WebAclId *string

	noSmithyDocumentSerde
}

type GetSampledRequestsOutput struct {

	// The total number of requests from which GetSampledRequests got a sample of
	// MaxItems requests. If PopulationSize is less than MaxItems, the sample includes
	// every request that your AWS resource received during the specified time range.
	PopulationSize int64

	// A complex type that contains detailed information about each of the requests in
	// the sample.
	SampledRequests []types.SampledHTTPRequest

	// Usually, TimeWindow is the time range that you specified in the
	// GetSampledRequests request. However, if your AWS resource received more than
	// 5,000 requests during the time range that you specified in the request,
	// GetSampledRequests returns the time range for the first 5,000 requests. Times
	// are in Coordinated Universal Time (UTC) format.
	TimeWindow *types.TimeWindow

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSampledRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSampledRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSampledRequests{}, middleware.After)
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
	if err = addOpGetSampledRequestsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSampledRequests(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSampledRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetSampledRequests",
	}
}
