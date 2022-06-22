// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the namespaces for the specified Amazon Web Services account.
func (c *Client) ListNamespaces(ctx context.Context, params *ListNamespacesInput, optFns ...func(*Options)) (*ListNamespacesOutput, error) {
	if params == nil {
		params = &ListNamespacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNamespaces", params, optFns, c.addOperationListNamespacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNamespacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNamespacesInput struct {

	// The ID for the Amazon Web Services account that contains the Amazon QuickSight
	// namespaces that you want to list.
	//
	// This member is required.
	AwsAccountId *string

	// The maximum number of results to return.
	MaxResults int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListNamespacesOutput struct {

	// The information about the namespaces in this Amazon Web Services account. The
	// response includes the namespace ARN, name, Amazon Web Services Region,
	// notification email address, creation status, and identity store.
	Namespaces []types.NamespaceInfoV2

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNamespacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNamespaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNamespaces{}, middleware.After)
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
	if err = addOpListNamespacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNamespaces(options.Region), middleware.Before); err != nil {
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

// ListNamespacesAPIClient is a client that implements the ListNamespaces
// operation.
type ListNamespacesAPIClient interface {
	ListNamespaces(context.Context, *ListNamespacesInput, ...func(*Options)) (*ListNamespacesOutput, error)
}

var _ ListNamespacesAPIClient = (*Client)(nil)

// ListNamespacesPaginatorOptions is the paginator options for ListNamespaces
type ListNamespacesPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNamespacesPaginator is a paginator for ListNamespaces
type ListNamespacesPaginator struct {
	options   ListNamespacesPaginatorOptions
	client    ListNamespacesAPIClient
	params    *ListNamespacesInput
	nextToken *string
	firstPage bool
}

// NewListNamespacesPaginator returns a new ListNamespacesPaginator
func NewListNamespacesPaginator(client ListNamespacesAPIClient, params *ListNamespacesInput, optFns ...func(*ListNamespacesPaginatorOptions)) *ListNamespacesPaginator {
	if params == nil {
		params = &ListNamespacesInput{}
	}

	options := ListNamespacesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNamespacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNamespacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNamespaces page.
func (p *ListNamespacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNamespacesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListNamespaces(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNamespaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListNamespaces",
	}
}
