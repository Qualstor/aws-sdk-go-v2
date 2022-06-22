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

// Represents a collection of BasePathMapping resources.
func (c *Client) GetBasePathMappings(ctx context.Context, params *GetBasePathMappingsInput, optFns ...func(*Options)) (*GetBasePathMappingsOutput, error) {
	if params == nil {
		params = &GetBasePathMappingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBasePathMappings", params, optFns, c.addOperationGetBasePathMappingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBasePathMappingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get information about a collection of BasePathMapping resources.
type GetBasePathMappingsInput struct {

	// The domain name of a BasePathMapping resource.
	//
	// This member is required.
	DomainName *string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string

	noSmithyDocumentSerde
}

// Represents a collection of BasePathMapping resources.
type GetBasePathMappingsOutput struct {

	// The current page of elements from this collection.
	Items []types.BasePathMapping

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBasePathMappingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBasePathMappings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBasePathMappings{}, middleware.After)
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
	if err = addOpGetBasePathMappingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBasePathMappings(options.Region), middleware.Before); err != nil {
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

// GetBasePathMappingsAPIClient is a client that implements the GetBasePathMappings
// operation.
type GetBasePathMappingsAPIClient interface {
	GetBasePathMappings(context.Context, *GetBasePathMappingsInput, ...func(*Options)) (*GetBasePathMappingsOutput, error)
}

var _ GetBasePathMappingsAPIClient = (*Client)(nil)

// GetBasePathMappingsPaginatorOptions is the paginator options for
// GetBasePathMappings
type GetBasePathMappingsPaginatorOptions struct {
	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetBasePathMappingsPaginator is a paginator for GetBasePathMappings
type GetBasePathMappingsPaginator struct {
	options   GetBasePathMappingsPaginatorOptions
	client    GetBasePathMappingsAPIClient
	params    *GetBasePathMappingsInput
	nextToken *string
	firstPage bool
}

// NewGetBasePathMappingsPaginator returns a new GetBasePathMappingsPaginator
func NewGetBasePathMappingsPaginator(client GetBasePathMappingsAPIClient, params *GetBasePathMappingsInput, optFns ...func(*GetBasePathMappingsPaginatorOptions)) *GetBasePathMappingsPaginator {
	if params == nil {
		params = &GetBasePathMappingsInput{}
	}

	options := GetBasePathMappingsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetBasePathMappingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Position,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetBasePathMappingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetBasePathMappings page.
func (p *GetBasePathMappingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetBasePathMappingsOutput, error) {
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

	result, err := p.client.GetBasePathMappings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetBasePathMappings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetBasePathMappings",
	}
}
