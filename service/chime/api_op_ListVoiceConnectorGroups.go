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

// Lists the Amazon Chime Voice Connector groups for the administrator's AWS
// account.
func (c *Client) ListVoiceConnectorGroups(ctx context.Context, params *ListVoiceConnectorGroupsInput, optFns ...func(*Options)) (*ListVoiceConnectorGroupsOutput, error) {
	if params == nil {
		params = &ListVoiceConnectorGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVoiceConnectorGroups", params, optFns, c.addOperationListVoiceConnectorGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVoiceConnectorGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVoiceConnectorGroupsInput struct {

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVoiceConnectorGroupsOutput struct {

	// The token to use to retrieve the next page of results.
	NextToken *string

	// The details of the Amazon Chime Voice Connector groups.
	VoiceConnectorGroups []types.VoiceConnectorGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVoiceConnectorGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVoiceConnectorGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVoiceConnectorGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVoiceConnectorGroups(options.Region), middleware.Before); err != nil {
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

// ListVoiceConnectorGroupsAPIClient is a client that implements the
// ListVoiceConnectorGroups operation.
type ListVoiceConnectorGroupsAPIClient interface {
	ListVoiceConnectorGroups(context.Context, *ListVoiceConnectorGroupsInput, ...func(*Options)) (*ListVoiceConnectorGroupsOutput, error)
}

var _ ListVoiceConnectorGroupsAPIClient = (*Client)(nil)

// ListVoiceConnectorGroupsPaginatorOptions is the paginator options for
// ListVoiceConnectorGroups
type ListVoiceConnectorGroupsPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVoiceConnectorGroupsPaginator is a paginator for ListVoiceConnectorGroups
type ListVoiceConnectorGroupsPaginator struct {
	options   ListVoiceConnectorGroupsPaginatorOptions
	client    ListVoiceConnectorGroupsAPIClient
	params    *ListVoiceConnectorGroupsInput
	nextToken *string
	firstPage bool
}

// NewListVoiceConnectorGroupsPaginator returns a new
// ListVoiceConnectorGroupsPaginator
func NewListVoiceConnectorGroupsPaginator(client ListVoiceConnectorGroupsAPIClient, params *ListVoiceConnectorGroupsInput, optFns ...func(*ListVoiceConnectorGroupsPaginatorOptions)) *ListVoiceConnectorGroupsPaginator {
	if params == nil {
		params = &ListVoiceConnectorGroupsInput{}
	}

	options := ListVoiceConnectorGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVoiceConnectorGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVoiceConnectorGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVoiceConnectorGroups page.
func (p *ListVoiceConnectorGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVoiceConnectorGroupsOutput, error) {
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

	result, err := p.client.ListVoiceConnectorGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListVoiceConnectorGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListVoiceConnectorGroups",
	}
}
