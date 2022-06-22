// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// For the time range passed in, returns the number of open reactive insight that
// were created, the number of open proactive insights that were created, and the
// Mean Time to Recover (MTTR) for all closed reactive insights.
func (c *Client) DescribeAccountOverview(ctx context.Context, params *DescribeAccountOverviewInput, optFns ...func(*Options)) (*DescribeAccountOverviewOutput, error) {
	if params == nil {
		params = &DescribeAccountOverviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountOverview", params, optFns, c.addOperationDescribeAccountOverviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountOverviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountOverviewInput struct {

	// The start of the time range passed in. The start time granularity is at the day
	// level. The floor of the start time is used. Returned information occurred after
	// this day.
	//
	// This member is required.
	FromTime *time.Time

	// The end of the time range passed in. The start time granularity is at the day
	// level. The floor of the start time is used. Returned information occurred before
	// this day. If this is not specified, then the current day is used.
	ToTime *time.Time

	noSmithyDocumentSerde
}

type DescribeAccountOverviewOutput struct {

	// The Mean Time to Recover (MTTR) for all closed insights that were created during
	// the time range passed in.
	//
	// This member is required.
	MeanTimeToRecoverInMilliseconds *int64

	// An integer that specifies the number of open proactive insights in your Amazon
	// Web Services account that were created during the time range passed in.
	//
	// This member is required.
	ProactiveInsights int32

	// An integer that specifies the number of open reactive insights in your Amazon
	// Web Services account that were created during the time range passed in.
	//
	// This member is required.
	ReactiveInsights int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccountOverviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccountOverview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccountOverview{}, middleware.After)
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
	if err = addOpDescribeAccountOverviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountOverview(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccountOverview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "DescribeAccountOverview",
	}
}
