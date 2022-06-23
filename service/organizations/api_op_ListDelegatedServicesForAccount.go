// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the Amazon Web Services services for which the specified account is a
// delegated administrator. This operation can be called only from the
// organization's management account or by a member account that is a delegated
// administrator for an Amazon Web Services service.
func (c *Client) ListDelegatedServicesForAccount(ctx context.Context, params *ListDelegatedServicesForAccountInput, optFns ...func(*Options)) (*ListDelegatedServicesForAccountOutput, error) {
	if params == nil {
		params = &ListDelegatedServicesForAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDelegatedServicesForAccount", params, optFns, c.addOperationListDelegatedServicesForAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDelegatedServicesForAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDelegatedServicesForAccountInput struct {

	// The account ID number of a delegated administrator account in the organization.
	//
	// This member is required.
	AccountId *string

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDelegatedServicesForAccountOutput struct {

	// The services for which the account is a delegated administrator.
	DelegatedServices []types.DelegatedService

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDelegatedServicesForAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDelegatedServicesForAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDelegatedServicesForAccount{}, middleware.After)
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
	if err = addOpListDelegatedServicesForAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDelegatedServicesForAccount(options.Region), middleware.Before); err != nil {
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

// ListDelegatedServicesForAccountAPIClient is a client that implements the
// ListDelegatedServicesForAccount operation.
type ListDelegatedServicesForAccountAPIClient interface {
	ListDelegatedServicesForAccount(context.Context, *ListDelegatedServicesForAccountInput, ...func(*Options)) (*ListDelegatedServicesForAccountOutput, error)
}

var _ ListDelegatedServicesForAccountAPIClient = (*Client)(nil)

// ListDelegatedServicesForAccountPaginatorOptions is the paginator options for
// ListDelegatedServicesForAccount
type ListDelegatedServicesForAccountPaginatorOptions struct {
	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDelegatedServicesForAccountPaginator is a paginator for
// ListDelegatedServicesForAccount
type ListDelegatedServicesForAccountPaginator struct {
	options   ListDelegatedServicesForAccountPaginatorOptions
	client    ListDelegatedServicesForAccountAPIClient
	params    *ListDelegatedServicesForAccountInput
	nextToken *string
	firstPage bool
}

// NewListDelegatedServicesForAccountPaginator returns a new
// ListDelegatedServicesForAccountPaginator
func NewListDelegatedServicesForAccountPaginator(client ListDelegatedServicesForAccountAPIClient, params *ListDelegatedServicesForAccountInput, optFns ...func(*ListDelegatedServicesForAccountPaginatorOptions)) *ListDelegatedServicesForAccountPaginator {
	if params == nil {
		params = &ListDelegatedServicesForAccountInput{}
	}

	options := ListDelegatedServicesForAccountPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDelegatedServicesForAccountPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDelegatedServicesForAccountPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDelegatedServicesForAccount page.
func (p *ListDelegatedServicesForAccountPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDelegatedServicesForAccountOutput, error) {
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

	result, err := p.client.ListDelegatedServicesForAccount(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDelegatedServicesForAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListDelegatedServicesForAccount",
	}
}
