// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the executions of a state machine that meet the filtering criteria.
// Results are sorted by time, with the most recent execution first. If nextToken
// is returned, there are more results available. The value of nextToken is a
// unique pagination token for each page. Make the call again using the returned
// token to retrieve the next page. Keep all other arguments unchanged. Each
// pagination token expires after 24 hours. Using an expired pagination token will
// return an HTTP 400 InvalidToken error. This operation is eventually consistent.
// The results are best effort and may not reflect very recent updates and changes.
// This API action is not supported by EXPRESS state machines.
func (c *Client) ListExecutions(ctx context.Context, params *ListExecutionsInput, optFns ...func(*Options)) (*ListExecutionsOutput, error) {
	if params == nil {
		params = &ListExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExecutions", params, optFns, c.addOperationListExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExecutionsInput struct {

	// The Amazon Resource Name (ARN) of the state machine whose executions is listed.
	//
	// This member is required.
	StateMachineArn *string

	// The maximum number of results that are returned per call. You can use nextToken
	// to obtain further pages of results. The default is 100 and the maximum allowed
	// page size is 1000. A value of 0 uses the default. This is only an upper limit.
	// The actual number of results returned per call might be fewer than the specified
	// maximum.
	MaxResults int32

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string

	// If specified, only list the executions whose current execution status matches
	// the given filter.
	StatusFilter types.ExecutionStatus

	noSmithyDocumentSerde
}

type ListExecutionsOutput struct {

	// The list of matching executions.
	//
	// This member is required.
	Executions []types.ExecutionListItem

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListExecutions{}, middleware.After)
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
	if err = addOpListExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExecutions(options.Region), middleware.Before); err != nil {
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

// ListExecutionsAPIClient is a client that implements the ListExecutions
// operation.
type ListExecutionsAPIClient interface {
	ListExecutions(context.Context, *ListExecutionsInput, ...func(*Options)) (*ListExecutionsOutput, error)
}

var _ ListExecutionsAPIClient = (*Client)(nil)

// ListExecutionsPaginatorOptions is the paginator options for ListExecutions
type ListExecutionsPaginatorOptions struct {
	// The maximum number of results that are returned per call. You can use nextToken
	// to obtain further pages of results. The default is 100 and the maximum allowed
	// page size is 1000. A value of 0 uses the default. This is only an upper limit.
	// The actual number of results returned per call might be fewer than the specified
	// maximum.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExecutionsPaginator is a paginator for ListExecutions
type ListExecutionsPaginator struct {
	options   ListExecutionsPaginatorOptions
	client    ListExecutionsAPIClient
	params    *ListExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListExecutionsPaginator returns a new ListExecutionsPaginator
func NewListExecutionsPaginator(client ListExecutionsAPIClient, params *ListExecutionsInput, optFns ...func(*ListExecutionsPaginatorOptions)) *ListExecutionsPaginator {
	if params == nil {
		params = &ListExecutionsInput{}
	}

	options := ListExecutionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExecutions page.
func (p *ListExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExecutionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "ListExecutions",
	}
}
