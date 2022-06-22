// Code generated by smithy-go-codegen DO NOT EDIT.

package rum

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/rum/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the Amazon CloudWatch RUM app monitors in the account.
func (c *Client) ListAppMonitors(ctx context.Context, params *ListAppMonitorsInput, optFns ...func(*Options)) (*ListAppMonitorsOutput, error) {
	if params == nil {
		params = &ListAppMonitorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppMonitors", params, optFns, c.addOperationListAppMonitorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppMonitorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppMonitorsInput struct {

	// The maximum number of results to return in one operation.
	MaxResults *int32

	// Use the token returned by the previous operation to request the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAppMonitorsOutput struct {

	// An array of structures that contain information about the returned app monitors.
	AppMonitorSummaries []types.AppMonitorSummary

	// A token that you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppMonitorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppMonitors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppMonitors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppMonitors(options.Region), middleware.Before); err != nil {
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

// ListAppMonitorsAPIClient is a client that implements the ListAppMonitors
// operation.
type ListAppMonitorsAPIClient interface {
	ListAppMonitors(context.Context, *ListAppMonitorsInput, ...func(*Options)) (*ListAppMonitorsOutput, error)
}

var _ ListAppMonitorsAPIClient = (*Client)(nil)

// ListAppMonitorsPaginatorOptions is the paginator options for ListAppMonitors
type ListAppMonitorsPaginatorOptions struct {
	// The maximum number of results to return in one operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppMonitorsPaginator is a paginator for ListAppMonitors
type ListAppMonitorsPaginator struct {
	options   ListAppMonitorsPaginatorOptions
	client    ListAppMonitorsAPIClient
	params    *ListAppMonitorsInput
	nextToken *string
	firstPage bool
}

// NewListAppMonitorsPaginator returns a new ListAppMonitorsPaginator
func NewListAppMonitorsPaginator(client ListAppMonitorsAPIClient, params *ListAppMonitorsInput, optFns ...func(*ListAppMonitorsPaginatorOptions)) *ListAppMonitorsPaginator {
	if params == nil {
		params = &ListAppMonitorsInput{}
	}

	options := ListAppMonitorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppMonitorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppMonitorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppMonitors page.
func (p *ListAppMonitorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppMonitorsOutput, error) {
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

	result, err := p.client.ListAppMonitors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAppMonitors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rum",
		OperationName: "ListAppMonitors",
	}
}
