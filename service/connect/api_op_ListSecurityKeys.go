// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Returns a paginated list of all security keys associated with the instance.
func (c *Client) ListSecurityKeys(ctx context.Context, params *ListSecurityKeysInput, optFns ...func(*Options)) (*ListSecurityKeysOutput, error) {
	if params == nil {
		params = &ListSecurityKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecurityKeys", params, optFns, c.addOperationListSecurityKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecurityKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityKeysInput struct {

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSecurityKeysOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The security keys.
	SecurityKeys []types.SecurityKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSecurityKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSecurityKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSecurityKeys{}, middleware.After)
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
	if err = addOpListSecurityKeysValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityKeys(options.Region), middleware.Before); err != nil {
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

// ListSecurityKeysAPIClient is a client that implements the ListSecurityKeys
// operation.
type ListSecurityKeysAPIClient interface {
	ListSecurityKeys(context.Context, *ListSecurityKeysInput, ...func(*Options)) (*ListSecurityKeysOutput, error)
}

var _ ListSecurityKeysAPIClient = (*Client)(nil)

// ListSecurityKeysPaginatorOptions is the paginator options for ListSecurityKeys
type ListSecurityKeysPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSecurityKeysPaginator is a paginator for ListSecurityKeys
type ListSecurityKeysPaginator struct {
	options   ListSecurityKeysPaginatorOptions
	client    ListSecurityKeysAPIClient
	params    *ListSecurityKeysInput
	nextToken *string
	firstPage bool
}

// NewListSecurityKeysPaginator returns a new ListSecurityKeysPaginator
func NewListSecurityKeysPaginator(client ListSecurityKeysAPIClient, params *ListSecurityKeysInput, optFns ...func(*ListSecurityKeysPaginatorOptions)) *ListSecurityKeysPaginator {
	if params == nil {
		params = &ListSecurityKeysInput{}
	}

	options := ListSecurityKeysPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSecurityKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSecurityKeysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSecurityKeys page.
func (p *ListSecurityKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSecurityKeysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListSecurityKeys(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSecurityKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListSecurityKeys",
	}
}
