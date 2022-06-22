// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your Device Defender detect custom metrics. Requires permission to access
// the ListCustomMetrics
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListCustomMetrics(ctx context.Context, params *ListCustomMetricsInput, optFns ...func(*Options)) (*ListCustomMetricsOutput, error) {
	if params == nil {
		params = &ListCustomMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomMetrics", params, optFns, c.addOperationListCustomMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomMetricsInput struct {

	// The maximum number of results to return at one time. The default is 25.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCustomMetricsOutput struct {

	// The name of the custom metric.
	MetricNames []string

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCustomMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCustomMetrics{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomMetrics(options.Region), middleware.Before); err != nil {
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

// ListCustomMetricsAPIClient is a client that implements the ListCustomMetrics
// operation.
type ListCustomMetricsAPIClient interface {
	ListCustomMetrics(context.Context, *ListCustomMetricsInput, ...func(*Options)) (*ListCustomMetricsOutput, error)
}

var _ ListCustomMetricsAPIClient = (*Client)(nil)

// ListCustomMetricsPaginatorOptions is the paginator options for ListCustomMetrics
type ListCustomMetricsPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 25.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomMetricsPaginator is a paginator for ListCustomMetrics
type ListCustomMetricsPaginator struct {
	options   ListCustomMetricsPaginatorOptions
	client    ListCustomMetricsAPIClient
	params    *ListCustomMetricsInput
	nextToken *string
	firstPage bool
}

// NewListCustomMetricsPaginator returns a new ListCustomMetricsPaginator
func NewListCustomMetricsPaginator(client ListCustomMetricsAPIClient, params *ListCustomMetricsInput, optFns ...func(*ListCustomMetricsPaginatorOptions)) *ListCustomMetricsPaginator {
	if params == nil {
		params = &ListCustomMetricsInput{}
	}

	options := ListCustomMetricsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomMetricsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomMetricsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomMetrics page.
func (p *ListCustomMetricsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomMetricsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListCustomMetrics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListCustomMetrics",
	}
}
