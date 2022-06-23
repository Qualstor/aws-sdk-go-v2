// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/migrationhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the created artifacts attached to a given migration task in an update
// stream. This API has the following traits:
//
// * Gets the list of the created
// artifacts while migration is taking place.
//
// * Shows the artifacts created by the
// migration tool that was associated by the AssociateCreatedArtifact API.
//
// * Lists
// created artifacts in a paginated interface.
func (c *Client) ListCreatedArtifacts(ctx context.Context, params *ListCreatedArtifactsInput, optFns ...func(*Options)) (*ListCreatedArtifactsOutput, error) {
	if params == nil {
		params = &ListCreatedArtifactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCreatedArtifacts", params, optFns, c.addOperationListCreatedArtifactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCreatedArtifactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCreatedArtifactsInput struct {

	// Unique identifier that references the migration task. Do not store personal data
	// in this field.
	//
	// This member is required.
	MigrationTaskName *string

	// The name of the ProgressUpdateStream.
	//
	// This member is required.
	ProgressUpdateStream *string

	// Maximum number of results to be returned per page.
	MaxResults *int32

	// If a NextToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in NextToken.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCreatedArtifactsOutput struct {

	// List of created artifacts up to the maximum number of results specified in the
	// request.
	CreatedArtifactList []types.CreatedArtifact

	// If there are more created artifacts than the max result, return the next token
	// to be passed to the next call as a bookmark of where to start from.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCreatedArtifactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCreatedArtifacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCreatedArtifacts{}, middleware.After)
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
	if err = addOpListCreatedArtifactsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCreatedArtifacts(options.Region), middleware.Before); err != nil {
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

// ListCreatedArtifactsAPIClient is a client that implements the
// ListCreatedArtifacts operation.
type ListCreatedArtifactsAPIClient interface {
	ListCreatedArtifacts(context.Context, *ListCreatedArtifactsInput, ...func(*Options)) (*ListCreatedArtifactsOutput, error)
}

var _ ListCreatedArtifactsAPIClient = (*Client)(nil)

// ListCreatedArtifactsPaginatorOptions is the paginator options for
// ListCreatedArtifacts
type ListCreatedArtifactsPaginatorOptions struct {
	// Maximum number of results to be returned per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCreatedArtifactsPaginator is a paginator for ListCreatedArtifacts
type ListCreatedArtifactsPaginator struct {
	options   ListCreatedArtifactsPaginatorOptions
	client    ListCreatedArtifactsAPIClient
	params    *ListCreatedArtifactsInput
	nextToken *string
	firstPage bool
}

// NewListCreatedArtifactsPaginator returns a new ListCreatedArtifactsPaginator
func NewListCreatedArtifactsPaginator(client ListCreatedArtifactsAPIClient, params *ListCreatedArtifactsInput, optFns ...func(*ListCreatedArtifactsPaginatorOptions)) *ListCreatedArtifactsPaginator {
	if params == nil {
		params = &ListCreatedArtifactsInput{}
	}

	options := ListCreatedArtifactsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCreatedArtifactsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCreatedArtifactsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCreatedArtifacts page.
func (p *ListCreatedArtifactsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCreatedArtifactsOutput, error) {
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

	result, err := p.client.ListCreatedArtifacts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCreatedArtifacts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "ListCreatedArtifacts",
	}
}
