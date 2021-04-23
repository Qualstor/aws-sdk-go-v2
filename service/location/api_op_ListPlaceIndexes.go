// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Place index resources in your AWS account.
func (c *Client) ListPlaceIndexes(ctx context.Context, params *ListPlaceIndexesInput, optFns ...func(*Options)) (*ListPlaceIndexesOutput, error) {
	if params == nil {
		params = &ListPlaceIndexesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPlaceIndexes", params, optFns, addOperationListPlaceIndexesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPlaceIndexesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPlaceIndexesInput struct {

	// An optional limit for the maximum number of results returned in a single call.
	// Default value: 100
	MaxResults *int32

	// The pagination token specifying which page of results to return in the response.
	// If no token is provided, the default page is the first page. Default value: null
	NextToken *string
}

type ListPlaceIndexesOutput struct {

	// Lists the Place index resources that exist in your AWS account
	//
	// This member is required.
	Entries []types.ListPlaceIndexesResponseEntry

	// A pagination token indicating there are additional pages available. You can use
	// the token in a following request to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPlaceIndexesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPlaceIndexes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPlaceIndexes{}, middleware.After)
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

// ListPlaceIndexesAPIClient is a client that implements the ListPlaceIndexes
// operation.
type ListPlaceIndexesAPIClient interface {
	ListPlaceIndexes(context.Context, *ListPlaceIndexesInput, ...func(*Options)) (*ListPlaceIndexesOutput, error)
}

var _ ListPlaceIndexesAPIClient = (*Client)(nil)

// ListPlaceIndexesPaginatorOptions is the paginator options for ListPlaceIndexes
type ListPlaceIndexesPaginatorOptions struct {
	// An optional limit for the maximum number of results returned in a single call.
	// Default value: 100
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPlaceIndexesPaginator is a paginator for ListPlaceIndexes
type ListPlaceIndexesPaginator struct {
	options   ListPlaceIndexesPaginatorOptions
	client    ListPlaceIndexesAPIClient
	params    *ListPlaceIndexesInput
	nextToken *string
	firstPage bool
}

// NewListPlaceIndexesPaginator returns a new ListPlaceIndexesPaginator
func NewListPlaceIndexesPaginator(client ListPlaceIndexesAPIClient, params *ListPlaceIndexesInput, optFns ...func(*ListPlaceIndexesPaginatorOptions)) *ListPlaceIndexesPaginator {
	if params == nil {
		params = &ListPlaceIndexesInput{}
	}

	options := ListPlaceIndexesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPlaceIndexesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPlaceIndexesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPlaceIndexes page.
func (p *ListPlaceIndexesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPlaceIndexesOutput, error) {
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

	result, err := p.client.ListPlaceIndexes(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}