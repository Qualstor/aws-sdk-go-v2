// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a subset of information about all the custom data identifiers for an
// account.
func (c *Client) ListCustomDataIdentifiers(ctx context.Context, params *ListCustomDataIdentifiersInput, optFns ...func(*Options)) (*ListCustomDataIdentifiersOutput, error) {
	if params == nil {
		params = &ListCustomDataIdentifiersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomDataIdentifiers", params, optFns, c.addOperationListCustomDataIdentifiersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomDataIdentifiersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomDataIdentifiersInput struct {

	// The maximum number of items to include in each page of the response.
	MaxResults int32

	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCustomDataIdentifiersOutput struct {

	// An array of objects, one for each custom data identifier.
	Items []types.CustomDataIdentifierSummary

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomDataIdentifiersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCustomDataIdentifiers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCustomDataIdentifiers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomDataIdentifiers(options.Region), middleware.Before); err != nil {
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

// ListCustomDataIdentifiersAPIClient is a client that implements the
// ListCustomDataIdentifiers operation.
type ListCustomDataIdentifiersAPIClient interface {
	ListCustomDataIdentifiers(context.Context, *ListCustomDataIdentifiersInput, ...func(*Options)) (*ListCustomDataIdentifiersOutput, error)
}

var _ ListCustomDataIdentifiersAPIClient = (*Client)(nil)

// ListCustomDataIdentifiersPaginatorOptions is the paginator options for
// ListCustomDataIdentifiers
type ListCustomDataIdentifiersPaginatorOptions struct {
	// The maximum number of items to include in each page of the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomDataIdentifiersPaginator is a paginator for ListCustomDataIdentifiers
type ListCustomDataIdentifiersPaginator struct {
	options   ListCustomDataIdentifiersPaginatorOptions
	client    ListCustomDataIdentifiersAPIClient
	params    *ListCustomDataIdentifiersInput
	nextToken *string
	firstPage bool
}

// NewListCustomDataIdentifiersPaginator returns a new
// ListCustomDataIdentifiersPaginator
func NewListCustomDataIdentifiersPaginator(client ListCustomDataIdentifiersAPIClient, params *ListCustomDataIdentifiersInput, optFns ...func(*ListCustomDataIdentifiersPaginatorOptions)) *ListCustomDataIdentifiersPaginator {
	if params == nil {
		params = &ListCustomDataIdentifiersInput{}
	}

	options := ListCustomDataIdentifiersPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomDataIdentifiersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomDataIdentifiersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomDataIdentifiers page.
func (p *ListCustomDataIdentifiersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomDataIdentifiersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListCustomDataIdentifiers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomDataIdentifiers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "ListCustomDataIdentifiers",
	}
}
