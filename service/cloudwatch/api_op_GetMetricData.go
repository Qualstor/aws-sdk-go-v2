// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// You can use the GetMetricData API to retrieve CloudWatch metric values. The
// operation can also include a CloudWatch Metrics Insights query, and one or more
// metric math functions. A GetMetricData operation that does not include a query
// can retrieve as many as 500 different metrics in a single request, with a total
// of as many as 100,800 data points. You can also optionally perform metric math
// expressions on the values of the returned statistics, to create new time series
// that represent new insights into your data. For example, using Lambda metrics,
// you could divide the Errors metric by the Invocations metric to get an error
// rate time series. For more information about metric math expressions, see Metric
// Math Syntax and Functions
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax)
// in the Amazon CloudWatch User Guide. If you include a Metrics Insights query,
// each GetMetricData operation can include only one query. But the same
// GetMetricData operation can also retrieve other metrics. Metrics Insights
// queries can query only the most recent three hours of metric data. For more
// information about Metrics Insights, see Query your metrics with CloudWatch
// Metrics Insights
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/query_with_cloudwatch-metrics-insights.html).
// Calls to the GetMetricData API have a different pricing structure than calls to
// GetMetricStatistics. For more information about pricing, see Amazon CloudWatch
// Pricing (https://aws.amazon.com/cloudwatch/pricing/). Amazon CloudWatch retains
// metric data as follows:
//
// * Data points with a period of less than 60 seconds are
// available for 3 hours. These data points are high-resolution metrics and are
// available only for custom metrics that have been defined with a
// StorageResolution of 1.
//
// * Data points with a period of 60 seconds (1-minute)
// are available for 15 days.
//
// * Data points with a period of 300 seconds
// (5-minute) are available for 63 days.
//
// * Data points with a period of 3600
// seconds (1 hour) are available for 455 days (15 months).
//
// Data points that are
// initially published with a shorter period are aggregated together for long-term
// storage. For example, if you collect data using a period of 1 minute, the data
// remains available for 15 days with 1-minute resolution. After 15 days, this data
// is still available, but is aggregated and retrievable only with a resolution of
// 5 minutes. After 63 days, the data is further aggregated and is available with a
// resolution of 1 hour. If you omit Unit in your request, all data that was
// collected with any unit is returned, along with the corresponding units that
// were specified when the data was reported to CloudWatch. If you specify a unit,
// the operation returns only data that was collected with that unit specified. If
// you specify a unit that does not match the data collected, the results of the
// operation are null. CloudWatch does not perform unit conversions. Using Metrics
// Insights queries with metric math You can't mix a Metric Insights query and
// metric math syntax in the same expression, but you can reference results from a
// Metrics Insights query within other Metric math expressions. A Metrics Insights
// query without a GROUP BY clause returns a single time-series (TS), and can be
// used as input for a metric math expression that expects a single time series. A
// Metrics Insights query with a GROUP BY clause returns an array of time-series
// (TS[]), and can be used as input for a metric math expression that expects an
// array of time series.
func (c *Client) GetMetricData(ctx context.Context, params *GetMetricDataInput, optFns ...func(*Options)) (*GetMetricDataOutput, error) {
	if params == nil {
		params = &GetMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMetricData", params, optFns, c.addOperationGetMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMetricDataInput struct {

	// The time stamp indicating the latest data to be returned. The value specified is
	// exclusive; results include data points up to the specified time stamp. For
	// better performance, specify StartTime and EndTime values that align with the
	// value of the metric's Period and sync up with the beginning and end of an hour.
	// For example, if the Period of a metric is 5 minutes, specifying 12:05 or 12:30
	// as EndTime can get a faster response from CloudWatch than setting 12:07 or 12:29
	// as the EndTime.
	//
	// This member is required.
	EndTime *time.Time

	// The metric queries to be returned. A single GetMetricData call can include as
	// many as 500 MetricDataQuery structures. Each of these structures can specify
	// either a metric to retrieve, a Metrics Insights query, or a math expression to
	// perform on retrieved data.
	//
	// This member is required.
	MetricDataQueries []types.MetricDataQuery

	// The time stamp indicating the earliest data to be returned. The value specified
	// is inclusive; results include data points with the specified time stamp.
	// CloudWatch rounds the specified time stamp as follows:
	//
	// * Start time less than
	// 15 days ago - Round down to the nearest whole minute. For example, 12:32:34 is
	// rounded down to 12:32:00.
	//
	// * Start time between 15 and 63 days ago - Round down
	// to the nearest 5-minute clock interval. For example, 12:32:34 is rounded down to
	// 12:30:00.
	//
	// * Start time greater than 63 days ago - Round down to the nearest
	// 1-hour clock interval. For example, 12:32:34 is rounded down to 12:00:00.
	//
	// If
	// you set Period to 5, 10, or 30, the start time of your request is rounded down
	// to the nearest time that corresponds to even 5-, 10-, or 30-second divisions of
	// a minute. For example, if you make a query at (HH:mm:ss) 01:05:23 for the
	// previous 10-second period, the start time of your request is rounded down and
	// you receive data from 01:05:10 to 01:05:20. If you make a query at 15:07:17 for
	// the previous 5 minutes of data, using a period of 5 seconds, you receive data
	// timestamped between 15:02:15 and 15:07:15. For better performance, specify
	// StartTime and EndTime values that align with the value of the metric's Period
	// and sync up with the beginning and end of an hour. For example, if the Period of
	// a metric is 5 minutes, specifying 12:05 or 12:30 as StartTime can get a faster
	// response from CloudWatch than setting 12:07 or 12:29 as the StartTime.
	//
	// This member is required.
	StartTime *time.Time

	// This structure includes the Timezone parameter, which you can use to specify
	// your time zone so that the labels of returned data display the correct time for
	// your time zone.
	LabelOptions *types.LabelOptions

	// The maximum number of data points the request should return before paginating.
	// If you omit this, the default of 100,800 is used.
	MaxDatapoints *int32

	// Include this value, if it was returned by the previous GetMetricData operation,
	// to get the next set of data points.
	NextToken *string

	// The order in which data points should be returned. TimestampDescending returns
	// the newest data first and paginates when the MaxDatapoints limit is reached.
	// TimestampAscending returns the oldest data first and paginates when the
	// MaxDatapoints limit is reached.
	ScanBy types.ScanBy

	noSmithyDocumentSerde
}

type GetMetricDataOutput struct {

	// Contains a message about this GetMetricData operation, if the operation results
	// in such a message. An example of a message that might be returned is Maximum
	// number of allowed metrics exceeded. If there is a message, as much of the
	// operation as possible is still executed. A message appears here only if it is
	// related to the global GetMetricData operation. Any message about a specific
	// metric returned by the operation appears in the MetricDataResult object returned
	// for that metric.
	Messages []types.MessageData

	// The metrics that are returned, including the metric name, namespace, and
	// dimensions.
	MetricDataResults []types.MetricDataResult

	// A token that marks the next batch of returned results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetMetricData{}, middleware.After)
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
	if err = addOpGetMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMetricData(options.Region), middleware.Before); err != nil {
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

// GetMetricDataAPIClient is a client that implements the GetMetricData operation.
type GetMetricDataAPIClient interface {
	GetMetricData(context.Context, *GetMetricDataInput, ...func(*Options)) (*GetMetricDataOutput, error)
}

var _ GetMetricDataAPIClient = (*Client)(nil)

// GetMetricDataPaginatorOptions is the paginator options for GetMetricData
type GetMetricDataPaginatorOptions struct {
	// The maximum number of data points the request should return before paginating.
	// If you omit this, the default of 100,800 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetMetricDataPaginator is a paginator for GetMetricData
type GetMetricDataPaginator struct {
	options   GetMetricDataPaginatorOptions
	client    GetMetricDataAPIClient
	params    *GetMetricDataInput
	nextToken *string
	firstPage bool
}

// NewGetMetricDataPaginator returns a new GetMetricDataPaginator
func NewGetMetricDataPaginator(client GetMetricDataAPIClient, params *GetMetricDataInput, optFns ...func(*GetMetricDataPaginatorOptions)) *GetMetricDataPaginator {
	if params == nil {
		params = &GetMetricDataInput{}
	}

	options := GetMetricDataPaginatorOptions{}
	if params.MaxDatapoints != nil {
		options.Limit = *params.MaxDatapoints
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetMetricDataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetMetricDataPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetMetricData page.
func (p *GetMetricDataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetMetricDataOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxDatapoints = limit

	result, err := p.client.GetMetricData(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "GetMetricData",
	}
}
