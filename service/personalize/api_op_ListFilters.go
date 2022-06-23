// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all filters that belong to a given dataset group.
func (c *Client) ListFilters(ctx context.Context, params *ListFiltersInput, optFns ...func(*Options)) (*ListFiltersOutput, error) {
	if params == nil {
		params = &ListFiltersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFilters", params, optFns, c.addOperationListFiltersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFiltersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFiltersInput struct {

	// The ARN of the dataset group that contains the filters.
	DatasetGroupArn *string

	// The maximum number of filters to return.
	MaxResults *int32

	// A token returned from the previous call to ListFilters for getting the next set
	// of filters (if they exist).
	NextToken *string

	noSmithyDocumentSerde
}

type ListFiltersOutput struct {

	// A list of returned filters.
	Filters []types.FilterSummary

	// A token for getting the next set of filters (if they exist).
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFiltersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFilters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFilters{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFilters(options.Region), middleware.Before); err != nil {
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

// ListFiltersAPIClient is a client that implements the ListFilters operation.
type ListFiltersAPIClient interface {
	ListFilters(context.Context, *ListFiltersInput, ...func(*Options)) (*ListFiltersOutput, error)
}

var _ ListFiltersAPIClient = (*Client)(nil)

// ListFiltersPaginatorOptions is the paginator options for ListFilters
type ListFiltersPaginatorOptions struct {
	// The maximum number of filters to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFiltersPaginator is a paginator for ListFilters
type ListFiltersPaginator struct {
	options   ListFiltersPaginatorOptions
	client    ListFiltersAPIClient
	params    *ListFiltersInput
	nextToken *string
	firstPage bool
}

// NewListFiltersPaginator returns a new ListFiltersPaginator
func NewListFiltersPaginator(client ListFiltersAPIClient, params *ListFiltersInput, optFns ...func(*ListFiltersPaginatorOptions)) *ListFiltersPaginator {
	if params == nil {
		params = &ListFiltersInput{}
	}

	options := ListFiltersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFiltersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFiltersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFilters page.
func (p *ListFiltersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFiltersOutput, error) {
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

	result, err := p.client.ListFilters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFilters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListFilters",
	}
}
