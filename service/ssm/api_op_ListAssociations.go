// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all State Manager associations in the current Amazon Web Services
// account and Amazon Web Services Region. You can limit the results to a specific
// State Manager association document or managed node by specifying a filter. State
// Manager is a capability of Amazon Web Services Systems Manager.
func (c *Client) ListAssociations(ctx context.Context, params *ListAssociationsInput, optFns ...func(*Options)) (*ListAssociationsOutput, error) {
	if params == nil {
		params = &ListAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssociations", params, optFns, c.addOperationListAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssociationsInput struct {

	// One or more filters. Use a filter to return a more specific list of results.
	// Filtering associations using the InstanceID attribute only returns legacy
	// associations created using the InstanceID attribute. Associations targeting the
	// managed node that are part of the Target Attributes ResourceGroup or Tags aren't
	// returned.
	AssociationFilterList []types.AssociationFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssociationsOutput struct {

	// The associations.
	Associations []types.Association

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAssociations{}, middleware.After)
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
	if err = addOpListAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssociations(options.Region), middleware.Before); err != nil {
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

// ListAssociationsAPIClient is a client that implements the ListAssociations
// operation.
type ListAssociationsAPIClient interface {
	ListAssociations(context.Context, *ListAssociationsInput, ...func(*Options)) (*ListAssociationsOutput, error)
}

var _ ListAssociationsAPIClient = (*Client)(nil)

// ListAssociationsPaginatorOptions is the paginator options for ListAssociations
type ListAssociationsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssociationsPaginator is a paginator for ListAssociations
type ListAssociationsPaginator struct {
	options   ListAssociationsPaginatorOptions
	client    ListAssociationsAPIClient
	params    *ListAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListAssociationsPaginator returns a new ListAssociationsPaginator
func NewListAssociationsPaginator(client ListAssociationsAPIClient, params *ListAssociationsInput, optFns ...func(*ListAssociationsPaginatorOptions)) *ListAssociationsPaginator {
	if params == nil {
		params = &ListAssociationsInput{}
	}

	options := ListAssociationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssociations page.
func (p *ListAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssociationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListAssociations",
	}
}
