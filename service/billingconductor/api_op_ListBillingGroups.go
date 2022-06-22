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

// A paginated call to retrieve a list of billing groups for the given billing
// period. If you don't provide a billing group, the current billing period is
// used.
func (c *Client) ListBillingGroups(ctx context.Context, params *ListBillingGroupsInput, optFns ...func(*Options)) (*ListBillingGroupsOutput, error) {
	if params == nil {
		params = &ListBillingGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBillingGroups", params, optFns, c.addOperationListBillingGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBillingGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBillingGroupsInput struct {

	// The preferred billing period to get billing groups.
	BillingPeriod *string

	// A ListBillingGroupsFilter that specifies the billing group and pricing plan to
	// retrieve billing group information.
	Filters *types.ListBillingGroupsFilter

	// The maximum number of billing groups to retrieve.
	MaxResults *int32

	// The pagination token used on subsequent calls to get billing groups.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBillingGroupsOutput struct {

	// A list of BillingGroupListElement retrieved.
	BillingGroups []types.BillingGroupListElement

	// The pagination token used on subsequent calls to get billing groups.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBillingGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBillingGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBillingGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBillingGroups(options.Region), middleware.Before); err != nil {
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

// ListBillingGroupsAPIClient is a client that implements the ListBillingGroups
// operation.
type ListBillingGroupsAPIClient interface {
	ListBillingGroups(context.Context, *ListBillingGroupsInput, ...func(*Options)) (*ListBillingGroupsOutput, error)
}

var _ ListBillingGroupsAPIClient = (*Client)(nil)

// ListBillingGroupsPaginatorOptions is the paginator options for ListBillingGroups
type ListBillingGroupsPaginatorOptions struct {
	// The maximum number of billing groups to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBillingGroupsPaginator is a paginator for ListBillingGroups
type ListBillingGroupsPaginator struct {
	options   ListBillingGroupsPaginatorOptions
	client    ListBillingGroupsAPIClient
	params    *ListBillingGroupsInput
	nextToken *string
	firstPage bool
}

// NewListBillingGroupsPaginator returns a new ListBillingGroupsPaginator
func NewListBillingGroupsPaginator(client ListBillingGroupsAPIClient, params *ListBillingGroupsInput, optFns ...func(*ListBillingGroupsPaginatorOptions)) *ListBillingGroupsPaginator {
	if params == nil {
		params = &ListBillingGroupsInput{}
	}

	options := ListBillingGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBillingGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBillingGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBillingGroups page.
func (p *ListBillingGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBillingGroupsOutput, error) {
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

	result, err := p.client.ListBillingGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBillingGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "ListBillingGroups",
	}
}
