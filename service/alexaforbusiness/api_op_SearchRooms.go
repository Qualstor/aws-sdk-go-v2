// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches rooms and lists the ones that meet a set of filter and sort criteria.
func (c *Client) SearchRooms(ctx context.Context, params *SearchRoomsInput, optFns ...func(*Options)) (*SearchRoomsOutput, error) {
	if params == nil {
		params = &SearchRoomsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchRooms", params, optFns, c.addOperationSearchRoomsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchRoomsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchRoomsInput struct {

	// The filters to use to list a specified set of rooms. The supported filter keys
	// are RoomName and ProfileName.
	Filters []types.Filter

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// The sort order to use in listing the specified set of rooms. The supported sort
	// keys are RoomName and ProfileName.
	SortCriteria []types.Sort

	noSmithyDocumentSerde
}

type SearchRoomsOutput struct {

	// The token returned to indicate that there is more data available.
	NextToken *string

	// The rooms that meet the specified set of filter criteria, in sort order.
	Rooms []types.RoomData

	// The total number of rooms returned.
	TotalCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchRoomsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchRooms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchRooms{}, middleware.After)
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
	if err = addOpSearchRoomsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchRooms(options.Region), middleware.Before); err != nil {
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

// SearchRoomsAPIClient is a client that implements the SearchRooms operation.
type SearchRoomsAPIClient interface {
	SearchRooms(context.Context, *SearchRoomsInput, ...func(*Options)) (*SearchRoomsOutput, error)
}

var _ SearchRoomsAPIClient = (*Client)(nil)

// SearchRoomsPaginatorOptions is the paginator options for SearchRooms
type SearchRoomsPaginatorOptions struct {
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchRoomsPaginator is a paginator for SearchRooms
type SearchRoomsPaginator struct {
	options   SearchRoomsPaginatorOptions
	client    SearchRoomsAPIClient
	params    *SearchRoomsInput
	nextToken *string
	firstPage bool
}

// NewSearchRoomsPaginator returns a new SearchRoomsPaginator
func NewSearchRoomsPaginator(client SearchRoomsAPIClient, params *SearchRoomsInput, optFns ...func(*SearchRoomsPaginatorOptions)) *SearchRoomsPaginator {
	if params == nil {
		params = &SearchRoomsInput{}
	}

	options := SearchRoomsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchRoomsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchRoomsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchRooms page.
func (p *SearchRoomsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchRoomsOutput, error) {
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

	result, err := p.client.SearchRooms(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchRooms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchRooms",
	}
}
