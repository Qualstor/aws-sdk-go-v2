// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the resources associated to a custom line item.
func (c *Client) ListResourcesAssociatedToCustomLineItem(ctx context.Context, params *ListResourcesAssociatedToCustomLineItemInput, optFns ...func(*Options)) (*ListResourcesAssociatedToCustomLineItemOutput, error) {
	if params == nil {
		params = &ListResourcesAssociatedToCustomLineItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourcesAssociatedToCustomLineItem", params, optFns, c.addOperationListResourcesAssociatedToCustomLineItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourcesAssociatedToCustomLineItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourcesAssociatedToCustomLineItemInput struct {

	// The ARN of the custom line item for which the resource associations will be
	// listed.
	//
	// This member is required.
	Arn *string

	// The billing period for which the resource associations will be listed.
	BillingPeriod *string

	// (Optional) A ListResourcesAssociatedToCustomLineItemFilter that can specify the
	// types of resources that should be retrieved.
	Filters *types.ListResourcesAssociatedToCustomLineItemFilter

	// (Optional) The maximum number of resource associations to be retrieved.
	MaxResults *int32

	// (Optional) The pagination token returned by a previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourcesAssociatedToCustomLineItemOutput struct {

	// The custom line item ARN for which the resource associations are listed.
	Arn *string

	// A list of ListResourcesAssociatedToCustomLineItemResponseElement for each
	// resource association retrieved.
	AssociatedResources []types.ListResourcesAssociatedToCustomLineItemResponseElement

	// The pagination token to be used in subsequent requests to retrieve additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourcesAssociatedToCustomLineItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListResourcesAssociatedToCustomLineItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListResourcesAssociatedToCustomLineItem{}, middleware.After)
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
	if err = addOpListResourcesAssociatedToCustomLineItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourcesAssociatedToCustomLineItem(options.Region), middleware.Before); err != nil {
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

// ListResourcesAssociatedToCustomLineItemAPIClient is a client that implements the
// ListResourcesAssociatedToCustomLineItem operation.
type ListResourcesAssociatedToCustomLineItemAPIClient interface {
	ListResourcesAssociatedToCustomLineItem(context.Context, *ListResourcesAssociatedToCustomLineItemInput, ...func(*Options)) (*ListResourcesAssociatedToCustomLineItemOutput, error)
}

var _ ListResourcesAssociatedToCustomLineItemAPIClient = (*Client)(nil)

// ListResourcesAssociatedToCustomLineItemPaginatorOptions is the paginator options
// for ListResourcesAssociatedToCustomLineItem
type ListResourcesAssociatedToCustomLineItemPaginatorOptions struct {
	// (Optional) The maximum number of resource associations to be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourcesAssociatedToCustomLineItemPaginator is a paginator for
// ListResourcesAssociatedToCustomLineItem
type ListResourcesAssociatedToCustomLineItemPaginator struct {
	options   ListResourcesAssociatedToCustomLineItemPaginatorOptions
	client    ListResourcesAssociatedToCustomLineItemAPIClient
	params    *ListResourcesAssociatedToCustomLineItemInput
	nextToken *string
	firstPage bool
}

// NewListResourcesAssociatedToCustomLineItemPaginator returns a new
// ListResourcesAssociatedToCustomLineItemPaginator
func NewListResourcesAssociatedToCustomLineItemPaginator(client ListResourcesAssociatedToCustomLineItemAPIClient, params *ListResourcesAssociatedToCustomLineItemInput, optFns ...func(*ListResourcesAssociatedToCustomLineItemPaginatorOptions)) *ListResourcesAssociatedToCustomLineItemPaginator {
	if params == nil {
		params = &ListResourcesAssociatedToCustomLineItemInput{}
	}

	options := ListResourcesAssociatedToCustomLineItemPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourcesAssociatedToCustomLineItemPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourcesAssociatedToCustomLineItemPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourcesAssociatedToCustomLineItem page.
func (p *ListResourcesAssociatedToCustomLineItemPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourcesAssociatedToCustomLineItemOutput, error) {
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

	result, err := p.client.ListResourcesAssociatedToCustomLineItem(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResourcesAssociatedToCustomLineItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "ListResourcesAssociatedToCustomLineItem",
	}
}
