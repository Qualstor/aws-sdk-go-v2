// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all pending aggregation requests.
func (c *Client) DescribePendingAggregationRequests(ctx context.Context, params *DescribePendingAggregationRequestsInput, optFns ...func(*Options)) (*DescribePendingAggregationRequestsOutput, error) {
	if params == nil {
		params = &DescribePendingAggregationRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePendingAggregationRequests", params, optFns, c.addOperationDescribePendingAggregationRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePendingAggregationRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePendingAggregationRequestsInput struct {

	// The maximum number of evaluation results returned on each page. The default is
	// maximum. If you specify 0, Config uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribePendingAggregationRequestsOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Returns a PendingAggregationRequests object.
	PendingAggregationRequests []types.PendingAggregationRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePendingAggregationRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePendingAggregationRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePendingAggregationRequests{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePendingAggregationRequests(options.Region), middleware.Before); err != nil {
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

// DescribePendingAggregationRequestsAPIClient is a client that implements the
// DescribePendingAggregationRequests operation.
type DescribePendingAggregationRequestsAPIClient interface {
	DescribePendingAggregationRequests(context.Context, *DescribePendingAggregationRequestsInput, ...func(*Options)) (*DescribePendingAggregationRequestsOutput, error)
}

var _ DescribePendingAggregationRequestsAPIClient = (*Client)(nil)

// DescribePendingAggregationRequestsPaginatorOptions is the paginator options for
// DescribePendingAggregationRequests
type DescribePendingAggregationRequestsPaginatorOptions struct {
	// The maximum number of evaluation results returned on each page. The default is
	// maximum. If you specify 0, Config uses the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePendingAggregationRequestsPaginator is a paginator for
// DescribePendingAggregationRequests
type DescribePendingAggregationRequestsPaginator struct {
	options   DescribePendingAggregationRequestsPaginatorOptions
	client    DescribePendingAggregationRequestsAPIClient
	params    *DescribePendingAggregationRequestsInput
	nextToken *string
	firstPage bool
}

// NewDescribePendingAggregationRequestsPaginator returns a new
// DescribePendingAggregationRequestsPaginator
func NewDescribePendingAggregationRequestsPaginator(client DescribePendingAggregationRequestsAPIClient, params *DescribePendingAggregationRequestsInput, optFns ...func(*DescribePendingAggregationRequestsPaginatorOptions)) *DescribePendingAggregationRequestsPaginator {
	if params == nil {
		params = &DescribePendingAggregationRequestsInput{}
	}

	options := DescribePendingAggregationRequestsPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePendingAggregationRequestsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePendingAggregationRequestsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribePendingAggregationRequests page.
func (p *DescribePendingAggregationRequestsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePendingAggregationRequestsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribePendingAggregationRequests(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribePendingAggregationRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribePendingAggregationRequests",
	}
}
