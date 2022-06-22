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

// Returns the history of the specified execution as a list of events. By default,
// the results are returned in ascending order of the timeStamp of the events. Use
// the reverseOrder parameter to get the latest events first. If nextToken is
// returned, there are more results available. The value of nextToken is a unique
// pagination token for each page. Make the call again using the returned token to
// retrieve the next page. Keep all other arguments unchanged. Each pagination
// token expires after 24 hours. Using an expired pagination token will return an
// HTTP 400 InvalidToken error. This API action is not supported by EXPRESS state
// machines.
func (c *Client) GetExecutionHistory(ctx context.Context, params *GetExecutionHistoryInput, optFns ...func(*Options)) (*GetExecutionHistoryOutput, error) {
	if params == nil {
		params = &GetExecutionHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetExecutionHistory", params, optFns, c.addOperationGetExecutionHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetExecutionHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExecutionHistoryInput struct {

	// The Amazon Resource Name (ARN) of the execution.
	//
	// This member is required.
	ExecutionArn *string

	// You can select whether execution data (input or output of a history event) is
	// returned. The default is true.
	IncludeExecutionData *bool

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

	// Lists events in descending order of their timeStamp.
	ReverseOrder bool

	noSmithyDocumentSerde
}

type GetExecutionHistoryOutput struct {

	// The list of events that occurred in the execution.
	//
	// This member is required.
	Events []types.HistoryEvent

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

func (c *Client) addOperationGetExecutionHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetExecutionHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetExecutionHistory{}, middleware.After)
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
	if err = addOpGetExecutionHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetExecutionHistory(options.Region), middleware.Before); err != nil {
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

// GetExecutionHistoryAPIClient is a client that implements the GetExecutionHistory
// operation.
type GetExecutionHistoryAPIClient interface {
	GetExecutionHistory(context.Context, *GetExecutionHistoryInput, ...func(*Options)) (*GetExecutionHistoryOutput, error)
}

var _ GetExecutionHistoryAPIClient = (*Client)(nil)

// GetExecutionHistoryPaginatorOptions is the paginator options for
// GetExecutionHistory
type GetExecutionHistoryPaginatorOptions struct {
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

// GetExecutionHistoryPaginator is a paginator for GetExecutionHistory
type GetExecutionHistoryPaginator struct {
	options   GetExecutionHistoryPaginatorOptions
	client    GetExecutionHistoryAPIClient
	params    *GetExecutionHistoryInput
	nextToken *string
	firstPage bool
}

// NewGetExecutionHistoryPaginator returns a new GetExecutionHistoryPaginator
func NewGetExecutionHistoryPaginator(client GetExecutionHistoryAPIClient, params *GetExecutionHistoryInput, optFns ...func(*GetExecutionHistoryPaginatorOptions)) *GetExecutionHistoryPaginator {
	if params == nil {
		params = &GetExecutionHistoryInput{}
	}

	options := GetExecutionHistoryPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetExecutionHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetExecutionHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetExecutionHistory page.
func (p *GetExecutionHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetExecutionHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetExecutionHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetExecutionHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "GetExecutionHistory",
	}
}
