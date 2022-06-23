// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the specified pools or all pools associated with your Amazon Web
// Services account. If you specify pool IDs, the output includes information for
// only the specified pools. If you specify filters, the output includes
// information for only those pools that meet the filter criteria. If you don't
// specify pool IDs or filters, the output includes information for all pools. If
// you specify a pool ID that isn't valid, an Error is returned. A pool is a
// collection of phone numbers and SenderIds. A pool can include one or more phone
// numbers and SenderIds that are associated with your Amazon Web Services account.
func (c *Client) DescribePools(ctx context.Context, params *DescribePoolsInput, optFns ...func(*Options)) (*DescribePoolsOutput, error) {
	if params == nil {
		params = &DescribePoolsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePools", params, optFns, c.addOperationDescribePoolsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePoolsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePoolsInput struct {

	// An array of PoolFilter objects to filter the results.
	Filters []types.PoolFilter

	// The maximum number of results to return per each request.
	MaxResults *int32

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	// The unique identifier of pools to find. This is an array of strings that can be
	// either the PoolId or PoolArn.
	PoolIds []string

	noSmithyDocumentSerde
}

type DescribePoolsOutput struct {

	// The token to be used for the next set of paginated results. If this field is
	// empty then there are no more results.
	NextToken *string

	// An array of PoolInformation objects that contain the details for the requested
	// pools.
	Pools []types.PoolInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePoolsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribePools{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribePools{}, middleware.After)
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
	if err = addOpDescribePoolsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePools(options.Region), middleware.Before); err != nil {
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

// DescribePoolsAPIClient is a client that implements the DescribePools operation.
type DescribePoolsAPIClient interface {
	DescribePools(context.Context, *DescribePoolsInput, ...func(*Options)) (*DescribePoolsOutput, error)
}

var _ DescribePoolsAPIClient = (*Client)(nil)

// DescribePoolsPaginatorOptions is the paginator options for DescribePools
type DescribePoolsPaginatorOptions struct {
	// The maximum number of results to return per each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePoolsPaginator is a paginator for DescribePools
type DescribePoolsPaginator struct {
	options   DescribePoolsPaginatorOptions
	client    DescribePoolsAPIClient
	params    *DescribePoolsInput
	nextToken *string
	firstPage bool
}

// NewDescribePoolsPaginator returns a new DescribePoolsPaginator
func NewDescribePoolsPaginator(client DescribePoolsAPIClient, params *DescribePoolsInput, optFns ...func(*DescribePoolsPaginatorOptions)) *DescribePoolsPaginator {
	if params == nil {
		params = &DescribePoolsInput{}
	}

	options := DescribePoolsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePoolsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePoolsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribePools page.
func (p *DescribePoolsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePoolsOutput, error) {
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

	result, err := p.client.DescribePools(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribePools(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DescribePools",
	}
}
