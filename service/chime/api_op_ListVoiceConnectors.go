// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Amazon Chime Voice Connectors for the administrator's AWS account.
func (c *Client) ListVoiceConnectors(ctx context.Context, params *ListVoiceConnectorsInput, optFns ...func(*Options)) (*ListVoiceConnectorsOutput, error) {
	if params == nil {
		params = &ListVoiceConnectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVoiceConnectors", params, optFns, c.addOperationListVoiceConnectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVoiceConnectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVoiceConnectorsInput struct {

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVoiceConnectorsOutput struct {

	// The token to use to retrieve the next page of results.
	NextToken *string

	// The details of the Amazon Chime Voice Connectors.
	VoiceConnectors []types.VoiceConnector

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVoiceConnectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVoiceConnectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVoiceConnectors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVoiceConnectors(options.Region), middleware.Before); err != nil {
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

// ListVoiceConnectorsAPIClient is a client that implements the ListVoiceConnectors
// operation.
type ListVoiceConnectorsAPIClient interface {
	ListVoiceConnectors(context.Context, *ListVoiceConnectorsInput, ...func(*Options)) (*ListVoiceConnectorsOutput, error)
}

var _ ListVoiceConnectorsAPIClient = (*Client)(nil)

// ListVoiceConnectorsPaginatorOptions is the paginator options for
// ListVoiceConnectors
type ListVoiceConnectorsPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVoiceConnectorsPaginator is a paginator for ListVoiceConnectors
type ListVoiceConnectorsPaginator struct {
	options   ListVoiceConnectorsPaginatorOptions
	client    ListVoiceConnectorsAPIClient
	params    *ListVoiceConnectorsInput
	nextToken *string
	firstPage bool
}

// NewListVoiceConnectorsPaginator returns a new ListVoiceConnectorsPaginator
func NewListVoiceConnectorsPaginator(client ListVoiceConnectorsAPIClient, params *ListVoiceConnectorsInput, optFns ...func(*ListVoiceConnectorsPaginatorOptions)) *ListVoiceConnectorsPaginator {
	if params == nil {
		params = &ListVoiceConnectorsInput{}
	}

	options := ListVoiceConnectorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVoiceConnectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVoiceConnectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVoiceConnectors page.
func (p *ListVoiceConnectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVoiceConnectorsOutput, error) {
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

	result, err := p.client.ListVoiceConnectors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListVoiceConnectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListVoiceConnectors",
	}
}
