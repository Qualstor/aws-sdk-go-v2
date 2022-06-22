// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of slot types that match the specified criteria.
func (c *Client) ListSlotTypes(ctx context.Context, params *ListSlotTypesInput, optFns ...func(*Options)) (*ListSlotTypesOutput, error) {
	if params == nil {
		params = &ListSlotTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSlotTypes", params, optFns, c.addOperationListSlotTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSlotTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSlotTypesInput struct {

	// The unique identifier of the bot that contains the slot types.
	//
	// This member is required.
	BotId *string

	// The version of the bot that contains the slot type.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale of the slot types to list. The string
	// must match one of the supported locales. For more information, see Supported
	// languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html).
	//
	// This member is required.
	LocaleId *string

	// Provides the specification of a filter used to limit the slot types in the
	// response to only those that match the filter specification. You can only specify
	// one filter and only one string to filter on.
	Filters []types.SlotTypeFilter

	// The maximum number of slot types to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	MaxResults *int32

	// If the response from the ListSlotTypes operation contains more results than
	// specified in the maxResults parameter, a token is returned in the response. Use
	// that token in the nextToken parameter to return the next page of results.
	NextToken *string

	// Determines the sort order for the response from the ListSlotTypes operation. You
	// can choose to sort by the slot type name or last updated date in either
	// ascending or descending order.
	SortBy *types.SlotTypeSortBy

	noSmithyDocumentSerde
}

type ListSlotTypesOutput struct {

	// The identifier of the bot that contains the slot types.
	BotId *string

	// The version of the bot that contains the slot types.
	BotVersion *string

	// The language and local of the slot types in the list.
	LocaleId *string

	// A token that indicates whether there are more results to return in a response to
	// the ListSlotTypes operation. If the nextToken field is present, you send the
	// contents as the nextToken parameter of a ListSlotTypes operation request to get
	// the next page of results.
	NextToken *string

	// Summary information for the slot types that meet the filter criteria specified
	// in the request. The length of the list is specified in the maxResults parameter
	// of the request. If there are more slot types available, the nextToken field
	// contains a token to get the next page of results.
	SlotTypeSummaries []types.SlotTypeSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSlotTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSlotTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSlotTypes{}, middleware.After)
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
	if err = addOpListSlotTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSlotTypes(options.Region), middleware.Before); err != nil {
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

// ListSlotTypesAPIClient is a client that implements the ListSlotTypes operation.
type ListSlotTypesAPIClient interface {
	ListSlotTypes(context.Context, *ListSlotTypesInput, ...func(*Options)) (*ListSlotTypesOutput, error)
}

var _ ListSlotTypesAPIClient = (*Client)(nil)

// ListSlotTypesPaginatorOptions is the paginator options for ListSlotTypes
type ListSlotTypesPaginatorOptions struct {
	// The maximum number of slot types to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSlotTypesPaginator is a paginator for ListSlotTypes
type ListSlotTypesPaginator struct {
	options   ListSlotTypesPaginatorOptions
	client    ListSlotTypesAPIClient
	params    *ListSlotTypesInput
	nextToken *string
	firstPage bool
}

// NewListSlotTypesPaginator returns a new ListSlotTypesPaginator
func NewListSlotTypesPaginator(client ListSlotTypesAPIClient, params *ListSlotTypesInput, optFns ...func(*ListSlotTypesPaginatorOptions)) *ListSlotTypesPaginator {
	if params == nil {
		params = &ListSlotTypesInput{}
	}

	options := ListSlotTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSlotTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSlotTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSlotTypes page.
func (p *ListSlotTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSlotTypesOutput, error) {
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

	result, err := p.client.ListSlotTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSlotTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "ListSlotTypes",
	}
}
