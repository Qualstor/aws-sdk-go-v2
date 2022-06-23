// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/budgets/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the budget names and notifications that are associated with an account.
func (c *Client) DescribeBudgetNotificationsForAccount(ctx context.Context, params *DescribeBudgetNotificationsForAccountInput, optFns ...func(*Options)) (*DescribeBudgetNotificationsForAccountOutput, error) {
	if params == nil {
		params = &DescribeBudgetNotificationsForAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBudgetNotificationsForAccount", params, optFns, c.addOperationDescribeBudgetNotificationsForAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBudgetNotificationsForAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBudgetNotificationsForAccountInput struct {

	// The account ID of the user. It's a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// An integer that shows how many budget name entries a paginated response
	// contains.
	MaxResults *int32

	// A generic string.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeBudgetNotificationsForAccountOutput struct {

	// A list of budget names and associated notifications for an account.
	BudgetNotificationsForAccount []types.BudgetNotificationsForAccount

	// A generic string.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBudgetNotificationsForAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBudgetNotificationsForAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBudgetNotificationsForAccount{}, middleware.After)
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
	if err = addOpDescribeBudgetNotificationsForAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBudgetNotificationsForAccount(options.Region), middleware.Before); err != nil {
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

// DescribeBudgetNotificationsForAccountAPIClient is a client that implements the
// DescribeBudgetNotificationsForAccount operation.
type DescribeBudgetNotificationsForAccountAPIClient interface {
	DescribeBudgetNotificationsForAccount(context.Context, *DescribeBudgetNotificationsForAccountInput, ...func(*Options)) (*DescribeBudgetNotificationsForAccountOutput, error)
}

var _ DescribeBudgetNotificationsForAccountAPIClient = (*Client)(nil)

// DescribeBudgetNotificationsForAccountPaginatorOptions is the paginator options
// for DescribeBudgetNotificationsForAccount
type DescribeBudgetNotificationsForAccountPaginatorOptions struct {
	// An integer that shows how many budget name entries a paginated response
	// contains.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBudgetNotificationsForAccountPaginator is a paginator for
// DescribeBudgetNotificationsForAccount
type DescribeBudgetNotificationsForAccountPaginator struct {
	options   DescribeBudgetNotificationsForAccountPaginatorOptions
	client    DescribeBudgetNotificationsForAccountAPIClient
	params    *DescribeBudgetNotificationsForAccountInput
	nextToken *string
	firstPage bool
}

// NewDescribeBudgetNotificationsForAccountPaginator returns a new
// DescribeBudgetNotificationsForAccountPaginator
func NewDescribeBudgetNotificationsForAccountPaginator(client DescribeBudgetNotificationsForAccountAPIClient, params *DescribeBudgetNotificationsForAccountInput, optFns ...func(*DescribeBudgetNotificationsForAccountPaginatorOptions)) *DescribeBudgetNotificationsForAccountPaginator {
	if params == nil {
		params = &DescribeBudgetNotificationsForAccountInput{}
	}

	options := DescribeBudgetNotificationsForAccountPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeBudgetNotificationsForAccountPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBudgetNotificationsForAccountPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeBudgetNotificationsForAccount page.
func (p *DescribeBudgetNotificationsForAccountPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBudgetNotificationsForAccountOutput, error) {
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

	result, err := p.client.DescribeBudgetNotificationsForAccount(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeBudgetNotificationsForAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DescribeBudgetNotificationsForAccount",
	}
}
