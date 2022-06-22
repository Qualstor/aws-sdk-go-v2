// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists deployment strategies.
func (c *Client) ListDeploymentStrategies(ctx context.Context, params *ListDeploymentStrategiesInput, optFns ...func(*Options)) (*ListDeploymentStrategiesOutput, error) {
	if params == nil {
		params = &ListDeploymentStrategiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeploymentStrategies", params, optFns, c.addOperationListDeploymentStrategiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeploymentStrategiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeploymentStrategiesInput struct {

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDeploymentStrategiesOutput struct {

	// The elements from this collection.
	Items []types.DeploymentStrategy

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDeploymentStrategiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDeploymentStrategies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeploymentStrategies{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentStrategies(options.Region), middleware.Before); err != nil {
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

// ListDeploymentStrategiesAPIClient is a client that implements the
// ListDeploymentStrategies operation.
type ListDeploymentStrategiesAPIClient interface {
	ListDeploymentStrategies(context.Context, *ListDeploymentStrategiesInput, ...func(*Options)) (*ListDeploymentStrategiesOutput, error)
}

var _ ListDeploymentStrategiesAPIClient = (*Client)(nil)

// ListDeploymentStrategiesPaginatorOptions is the paginator options for
// ListDeploymentStrategies
type ListDeploymentStrategiesPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeploymentStrategiesPaginator is a paginator for ListDeploymentStrategies
type ListDeploymentStrategiesPaginator struct {
	options   ListDeploymentStrategiesPaginatorOptions
	client    ListDeploymentStrategiesAPIClient
	params    *ListDeploymentStrategiesInput
	nextToken *string
	firstPage bool
}

// NewListDeploymentStrategiesPaginator returns a new
// ListDeploymentStrategiesPaginator
func NewListDeploymentStrategiesPaginator(client ListDeploymentStrategiesAPIClient, params *ListDeploymentStrategiesInput, optFns ...func(*ListDeploymentStrategiesPaginatorOptions)) *ListDeploymentStrategiesPaginator {
	if params == nil {
		params = &ListDeploymentStrategiesInput{}
	}

	options := ListDeploymentStrategiesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDeploymentStrategiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeploymentStrategiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDeploymentStrategies page.
func (p *ListDeploymentStrategiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeploymentStrategiesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDeploymentStrategies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDeploymentStrategies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "ListDeploymentStrategies",
	}
}
