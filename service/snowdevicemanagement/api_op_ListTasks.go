// Code generated by smithy-go-codegen DO NOT EDIT.

package snowdevicemanagement

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/snowdevicemanagement/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of tasks that can be filtered by state.
func (c *Client) ListTasks(ctx context.Context, params *ListTasksInput, optFns ...func(*Options)) (*ListTasksOutput, error) {
	if params == nil {
		params = &ListTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTasks", params, optFns, c.addOperationListTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTasksInput struct {

	// The maximum number of tasks per page.
	MaxResults *int32

	// A pagination token to continue to the next page of tasks.
	NextToken *string

	// A structure used to filter the list of tasks.
	State types.TaskState

	noSmithyDocumentSerde
}

type ListTasksOutput struct {

	// A pagination token to continue to the next page of tasks.
	NextToken *string

	// A list of task structures containing details about each task.
	Tasks []types.TaskSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTasks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTasks(options.Region), middleware.Before); err != nil {
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

// ListTasksAPIClient is a client that implements the ListTasks operation.
type ListTasksAPIClient interface {
	ListTasks(context.Context, *ListTasksInput, ...func(*Options)) (*ListTasksOutput, error)
}

var _ ListTasksAPIClient = (*Client)(nil)

// ListTasksPaginatorOptions is the paginator options for ListTasks
type ListTasksPaginatorOptions struct {
	// The maximum number of tasks per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTasksPaginator is a paginator for ListTasks
type ListTasksPaginator struct {
	options   ListTasksPaginatorOptions
	client    ListTasksAPIClient
	params    *ListTasksInput
	nextToken *string
	firstPage bool
}

// NewListTasksPaginator returns a new ListTasksPaginator
func NewListTasksPaginator(client ListTasksAPIClient, params *ListTasksInput, optFns ...func(*ListTasksPaginatorOptions)) *ListTasksPaginator {
	if params == nil {
		params = &ListTasksInput{}
	}

	options := ListTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTasks page.
func (p *ListTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTasksOutput, error) {
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

	result, err := p.client.ListTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snow-device-management",
		OperationName: "ListTasks",
	}
}
