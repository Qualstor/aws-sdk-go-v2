// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the user pools associated with an Amazon Web Services account.
func (c *Client) ListUserPools(ctx context.Context, params *ListUserPoolsInput, optFns ...func(*Options)) (*ListUserPoolsOutput, error) {
	if params == nil {
		params = &ListUserPoolsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserPools", params, optFns, c.addOperationListUserPoolsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserPoolsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list user pools.
type ListUserPoolsInput struct {

	// The maximum number of results you want the request to return when listing the
	// user pools.
	//
	// This member is required.
	MaxResults int32

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	noSmithyDocumentSerde
}

// Represents the response to list user pools.
type ListUserPoolsOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// The user pools from the response to list users.
	UserPools []types.UserPoolDescriptionType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUserPoolsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUserPools{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUserPools{}, middleware.After)
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
	if err = addOpListUserPoolsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUserPools(options.Region), middleware.Before); err != nil {
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

// ListUserPoolsAPIClient is a client that implements the ListUserPools operation.
type ListUserPoolsAPIClient interface {
	ListUserPools(context.Context, *ListUserPoolsInput, ...func(*Options)) (*ListUserPoolsOutput, error)
}

var _ ListUserPoolsAPIClient = (*Client)(nil)

// ListUserPoolsPaginatorOptions is the paginator options for ListUserPools
type ListUserPoolsPaginatorOptions struct {
	// The maximum number of results you want the request to return when listing the
	// user pools.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUserPoolsPaginator is a paginator for ListUserPools
type ListUserPoolsPaginator struct {
	options   ListUserPoolsPaginatorOptions
	client    ListUserPoolsAPIClient
	params    *ListUserPoolsInput
	nextToken *string
	firstPage bool
}

// NewListUserPoolsPaginator returns a new ListUserPoolsPaginator
func NewListUserPoolsPaginator(client ListUserPoolsAPIClient, params *ListUserPoolsInput, optFns ...func(*ListUserPoolsPaginatorOptions)) *ListUserPoolsPaginator {
	if params == nil {
		params = &ListUserPoolsInput{}
	}

	options := ListUserPoolsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUserPoolsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUserPoolsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUserPools page.
func (p *ListUserPoolsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUserPoolsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListUserPools(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUserPools(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "ListUserPools",
	}
}
