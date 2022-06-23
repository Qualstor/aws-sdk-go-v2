// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all findings-generating solutions (products) that you are subscribed to
// receive findings from in Security Hub.
func (c *Client) ListEnabledProductsForImport(ctx context.Context, params *ListEnabledProductsForImportInput, optFns ...func(*Options)) (*ListEnabledProductsForImportOutput, error) {
	if params == nil {
		params = &ListEnabledProductsForImportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnabledProductsForImport", params, optFns, c.addOperationListEnabledProductsForImportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnabledProductsForImportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnabledProductsForImportInput struct {

	// The maximum number of items to return in the response.
	MaxResults int32

	// The token that is required for pagination. On your first call to the
	// ListEnabledProductsForImport operation, set the value of this parameter to NULL.
	// For subsequent calls to the operation, to continue listing data, set the value
	// of this parameter to the value returned from the previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEnabledProductsForImportOutput struct {

	// The pagination token to use to request the next page of results.
	NextToken *string

	// The list of ARNs for the resources that represent your subscriptions to
	// products.
	ProductSubscriptions []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnabledProductsForImportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEnabledProductsForImport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEnabledProductsForImport{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnabledProductsForImport(options.Region), middleware.Before); err != nil {
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

// ListEnabledProductsForImportAPIClient is a client that implements the
// ListEnabledProductsForImport operation.
type ListEnabledProductsForImportAPIClient interface {
	ListEnabledProductsForImport(context.Context, *ListEnabledProductsForImportInput, ...func(*Options)) (*ListEnabledProductsForImportOutput, error)
}

var _ ListEnabledProductsForImportAPIClient = (*Client)(nil)

// ListEnabledProductsForImportPaginatorOptions is the paginator options for
// ListEnabledProductsForImport
type ListEnabledProductsForImportPaginatorOptions struct {
	// The maximum number of items to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnabledProductsForImportPaginator is a paginator for
// ListEnabledProductsForImport
type ListEnabledProductsForImportPaginator struct {
	options   ListEnabledProductsForImportPaginatorOptions
	client    ListEnabledProductsForImportAPIClient
	params    *ListEnabledProductsForImportInput
	nextToken *string
	firstPage bool
}

// NewListEnabledProductsForImportPaginator returns a new
// ListEnabledProductsForImportPaginator
func NewListEnabledProductsForImportPaginator(client ListEnabledProductsForImportAPIClient, params *ListEnabledProductsForImportInput, optFns ...func(*ListEnabledProductsForImportPaginatorOptions)) *ListEnabledProductsForImportPaginator {
	if params == nil {
		params = &ListEnabledProductsForImportInput{}
	}

	options := ListEnabledProductsForImportPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnabledProductsForImportPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnabledProductsForImportPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnabledProductsForImport page.
func (p *ListEnabledProductsForImportPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnabledProductsForImportOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListEnabledProductsForImport(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEnabledProductsForImport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "ListEnabledProductsForImport",
	}
}
