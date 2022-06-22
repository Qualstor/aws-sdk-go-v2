// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubstrategy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/migrationhubstrategy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the servers.
func (c *Client) ListServers(ctx context.Context, params *ListServersInput, optFns ...func(*Options)) (*ListServersOutput, error) {
	if params == nil {
		params = &ListServersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServers", params, optFns, c.addOperationListServersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServersInput struct {

	// Specifies the filter value, which is based on the type of server criteria. For
	// example, if serverCriteria is OS_NAME, and the filterValue is equal to
	// WindowsServer, then ListServers returns all of the servers matching the OS name
	// WindowsServer.
	FilterValue *string

	// Specifies the group ID to filter on.
	GroupIdFilter []types.Group

	// The maximum number of items to include in the response. The maximum value is
	// 100.
	MaxResults *int32

	// The token from a previous call that you use to retrieve the next set of results.
	// For example, if a previous call to this action returned 100 items, but you set
	// maxResults to 10. You'll receive a set of 10 results along with a token. You
	// then use the returned token to retrieve the next set of 10.
	NextToken *string

	// Criteria for filtering servers.
	ServerCriteria types.ServerCriteria

	// Specifies whether to sort by ascending (ASC) or descending (DESC) order.
	Sort types.SortOrder

	noSmithyDocumentSerde
}

type ListServersOutput struct {

	// The token you use to retrieve the next set of results, or null if there are no
	// more results.
	NextToken *string

	// The list of servers with detailed information about each server.
	ServerInfos []types.ServerDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListServers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListServers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServers(options.Region), middleware.Before); err != nil {
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

// ListServersAPIClient is a client that implements the ListServers operation.
type ListServersAPIClient interface {
	ListServers(context.Context, *ListServersInput, ...func(*Options)) (*ListServersOutput, error)
}

var _ ListServersAPIClient = (*Client)(nil)

// ListServersPaginatorOptions is the paginator options for ListServers
type ListServersPaginatorOptions struct {
	// The maximum number of items to include in the response. The maximum value is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServersPaginator is a paginator for ListServers
type ListServersPaginator struct {
	options   ListServersPaginatorOptions
	client    ListServersAPIClient
	params    *ListServersInput
	nextToken *string
	firstPage bool
}

// NewListServersPaginator returns a new ListServersPaginator
func NewListServersPaginator(client ListServersAPIClient, params *ListServersInput, optFns ...func(*ListServersPaginatorOptions)) *ListServersPaginator {
	if params == nil {
		params = &ListServersInput{}
	}

	options := ListServersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServers page.
func (p *ListServersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServersOutput, error) {
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

	result, err := p.client.ListServers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-strategy",
		OperationName: "ListServers",
	}
}
