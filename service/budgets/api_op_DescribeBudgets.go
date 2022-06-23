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

// Lists the budgets that are associated with an account. The Request Syntax
// section shows the BudgetLimit syntax. For PlannedBudgetLimits, see the Examples
// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_DescribeBudgets.html#API_DescribeBudgets_Examples)
// section.
func (c *Client) DescribeBudgets(ctx context.Context, params *DescribeBudgetsInput, optFns ...func(*Options)) (*DescribeBudgetsOutput, error) {
	if params == nil {
		params = &DescribeBudgetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBudgets", params, optFns, c.addOperationDescribeBudgetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBudgetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of DescribeBudgets
type DescribeBudgetsInput struct {

	// The accountId that is associated with the budgets that you want descriptions of.
	//
	// This member is required.
	AccountId *string

	// An optional integer that represents how many entries a paginated response
	// contains. The maximum is 100.
	MaxResults *int32

	// The pagination token that you include in your request to indicate the next set
	// of results that you want to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

// Response of DescribeBudgets
type DescribeBudgetsOutput struct {

	// A list of budgets.
	Budgets []types.Budget

	// The pagination token in the service response that indicates the next set of
	// results that you can retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBudgetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBudgets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBudgets{}, middleware.After)
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
	if err = addOpDescribeBudgetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBudgets(options.Region), middleware.Before); err != nil {
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

// DescribeBudgetsAPIClient is a client that implements the DescribeBudgets
// operation.
type DescribeBudgetsAPIClient interface {
	DescribeBudgets(context.Context, *DescribeBudgetsInput, ...func(*Options)) (*DescribeBudgetsOutput, error)
}

var _ DescribeBudgetsAPIClient = (*Client)(nil)

// DescribeBudgetsPaginatorOptions is the paginator options for DescribeBudgets
type DescribeBudgetsPaginatorOptions struct {
	// An optional integer that represents how many entries a paginated response
	// contains. The maximum is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBudgetsPaginator is a paginator for DescribeBudgets
type DescribeBudgetsPaginator struct {
	options   DescribeBudgetsPaginatorOptions
	client    DescribeBudgetsAPIClient
	params    *DescribeBudgetsInput
	nextToken *string
	firstPage bool
}

// NewDescribeBudgetsPaginator returns a new DescribeBudgetsPaginator
func NewDescribeBudgetsPaginator(client DescribeBudgetsAPIClient, params *DescribeBudgetsInput, optFns ...func(*DescribeBudgetsPaginatorOptions)) *DescribeBudgetsPaginator {
	if params == nil {
		params = &DescribeBudgetsInput{}
	}

	options := DescribeBudgetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeBudgetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBudgetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeBudgets page.
func (p *DescribeBudgetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBudgetsOutput, error) {
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

	result, err := p.client.DescribeBudgets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeBudgets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DescribeBudgets",
	}
}
