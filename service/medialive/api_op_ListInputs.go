// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Produces list of inputs that have been created
func (c *Client) ListInputs(ctx context.Context, params *ListInputsInput, optFns ...func(*Options)) (*ListInputsOutput, error) {
	if params == nil {
		params = &ListInputsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInputs", params, optFns, c.addOperationListInputsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInputsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListInputsRequest
type ListInputsInput struct {

	// Placeholder documentation for MaxResults
	MaxResults int32

	// Placeholder documentation for __string
	NextToken *string

	noSmithyDocumentSerde
}

// Placeholder documentation for ListInputsResponse
type ListInputsOutput struct {

	// Placeholder documentation for __listOfInput
	Inputs []types.Input

	// Placeholder documentation for __string
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInputsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInputs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInputs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInputs(options.Region), middleware.Before); err != nil {
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

// ListInputsAPIClient is a client that implements the ListInputs operation.
type ListInputsAPIClient interface {
	ListInputs(context.Context, *ListInputsInput, ...func(*Options)) (*ListInputsOutput, error)
}

var _ ListInputsAPIClient = (*Client)(nil)

// ListInputsPaginatorOptions is the paginator options for ListInputs
type ListInputsPaginatorOptions struct {
	// Placeholder documentation for MaxResults
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInputsPaginator is a paginator for ListInputs
type ListInputsPaginator struct {
	options   ListInputsPaginatorOptions
	client    ListInputsAPIClient
	params    *ListInputsInput
	nextToken *string
	firstPage bool
}

// NewListInputsPaginator returns a new ListInputsPaginator
func NewListInputsPaginator(client ListInputsAPIClient, params *ListInputsInput, optFns ...func(*ListInputsPaginatorOptions)) *ListInputsPaginator {
	if params == nil {
		params = &ListInputsInput{}
	}

	options := ListInputsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInputsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInputsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInputs page.
func (p *ListInputsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInputsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListInputs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInputs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "ListInputs",
	}
}
