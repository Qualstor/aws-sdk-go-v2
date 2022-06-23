// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows an AppInstanceUser to search the channels that they belong to. The
// AppInstanceUser can search by membership or external ID. An AppInstanceAdmin can
// search across all channels within the AppInstance.
func (c *Client) SearchChannels(ctx context.Context, params *SearchChannelsInput, optFns ...func(*Options)) (*SearchChannelsOutput, error) {
	if params == nil {
		params = &SearchChannelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchChannels", params, optFns, c.addOperationSearchChannelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchChannelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchChannelsInput struct {

	// A list of the Field objects in the channel being searched.
	//
	// This member is required.
	Fields []types.SearchField

	// The AppInstanceUserArn of the user making the API call.
	ChimeBearer *string

	// The maximum number of channels that you want returned.
	MaxResults *int32

	// The token returned from previous API requests until the number of channels is
	// reached.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchChannelsOutput struct {

	// A list of the channels in the request.
	Channels []types.ChannelSummary

	// The token returned from previous API responses until the number of channels is
	// reached.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchChannelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchChannels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchChannels{}, middleware.After)
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
	if err = addOpSearchChannelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchChannels(options.Region), middleware.Before); err != nil {
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

// SearchChannelsAPIClient is a client that implements the SearchChannels
// operation.
type SearchChannelsAPIClient interface {
	SearchChannels(context.Context, *SearchChannelsInput, ...func(*Options)) (*SearchChannelsOutput, error)
}

var _ SearchChannelsAPIClient = (*Client)(nil)

// SearchChannelsPaginatorOptions is the paginator options for SearchChannels
type SearchChannelsPaginatorOptions struct {
	// The maximum number of channels that you want returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchChannelsPaginator is a paginator for SearchChannels
type SearchChannelsPaginator struct {
	options   SearchChannelsPaginatorOptions
	client    SearchChannelsAPIClient
	params    *SearchChannelsInput
	nextToken *string
	firstPage bool
}

// NewSearchChannelsPaginator returns a new SearchChannelsPaginator
func NewSearchChannelsPaginator(client SearchChannelsAPIClient, params *SearchChannelsInput, optFns ...func(*SearchChannelsPaginatorOptions)) *SearchChannelsPaginator {
	if params == nil {
		params = &SearchChannelsInput{}
	}

	options := SearchChannelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchChannelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchChannelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchChannels page.
func (p *SearchChannelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchChannelsOutput, error) {
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

	result, err := p.client.SearchChannels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchChannels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "SearchChannels",
	}
}
