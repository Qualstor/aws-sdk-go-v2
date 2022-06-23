// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Represents a collection of DomainName resources.
func (c *Client) GetDomainNames(ctx context.Context, params *GetDomainNamesInput, optFns ...func(*Options)) (*GetDomainNamesOutput, error) {
	if params == nil {
		params = &GetDomainNamesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomainNames", params, optFns, c.addOperationGetDomainNamesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainNamesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe a collection of DomainName resources.
type GetDomainNamesInput struct {

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string

	noSmithyDocumentSerde
}

// Represents a collection of DomainName resources.
type GetDomainNamesOutput struct {

	// The current page of elements from this collection.
	Items []types.DomainName

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDomainNamesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDomainNames{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomainNames{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainNames(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// GetDomainNamesAPIClient is a client that implements the GetDomainNames
// operation.
type GetDomainNamesAPIClient interface {
	GetDomainNames(context.Context, *GetDomainNamesInput, ...func(*Options)) (*GetDomainNamesOutput, error)
}

var _ GetDomainNamesAPIClient = (*Client)(nil)

// GetDomainNamesPaginatorOptions is the paginator options for GetDomainNames
type GetDomainNamesPaginatorOptions struct {
	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetDomainNamesPaginator is a paginator for GetDomainNames
type GetDomainNamesPaginator struct {
	options   GetDomainNamesPaginatorOptions
	client    GetDomainNamesAPIClient
	params    *GetDomainNamesInput
	nextToken *string
	firstPage bool
}

// NewGetDomainNamesPaginator returns a new GetDomainNamesPaginator
func NewGetDomainNamesPaginator(client GetDomainNamesAPIClient, params *GetDomainNamesInput, optFns ...func(*GetDomainNamesPaginatorOptions)) *GetDomainNamesPaginator {
	if params == nil {
		params = &GetDomainNamesInput{}
	}

	options := GetDomainNamesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetDomainNamesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Position,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetDomainNamesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetDomainNames page.
func (p *GetDomainNamesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetDomainNamesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Position = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.GetDomainNames(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Position

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetDomainNames(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetDomainNames",
	}
}
