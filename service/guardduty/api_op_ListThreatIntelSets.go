// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the ThreatIntelSets of the GuardDuty service specified by the detector ID.
// If you use this operation from a member account, the ThreatIntelSets associated
// with the administrator account are returned.
func (c *Client) ListThreatIntelSets(ctx context.Context, params *ListThreatIntelSetsInput, optFns ...func(*Options)) (*ListThreatIntelSetsOutput, error) {
	if params == nil {
		params = &ListThreatIntelSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThreatIntelSets", params, optFns, c.addOperationListThreatIntelSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThreatIntelSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThreatIntelSetsInput struct {

	// The unique ID of the detector that the threatIntelSet is associated with.
	//
	// This member is required.
	DetectorId *string

	// You can use this parameter to indicate the maximum number of items that you want
	// in the response. The default value is 50. The maximum value is 50.
	MaxResults int32

	// You can use this parameter to paginate results in the response. Set the value of
	// this parameter to null on your first call to the list action. For subsequent
	// calls to the action, fill nextToken in the request with the value of NextToken
	// from the previous response to continue listing data.
	NextToken *string

	noSmithyDocumentSerde
}

type ListThreatIntelSetsOutput struct {

	// The IDs of the ThreatIntelSet resources.
	//
	// This member is required.
	ThreatIntelSetIds []string

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThreatIntelSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThreatIntelSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThreatIntelSets{}, middleware.After)
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
	if err = addOpListThreatIntelSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThreatIntelSets(options.Region), middleware.Before); err != nil {
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

// ListThreatIntelSetsAPIClient is a client that implements the ListThreatIntelSets
// operation.
type ListThreatIntelSetsAPIClient interface {
	ListThreatIntelSets(context.Context, *ListThreatIntelSetsInput, ...func(*Options)) (*ListThreatIntelSetsOutput, error)
}

var _ ListThreatIntelSetsAPIClient = (*Client)(nil)

// ListThreatIntelSetsPaginatorOptions is the paginator options for
// ListThreatIntelSets
type ListThreatIntelSetsPaginatorOptions struct {
	// You can use this parameter to indicate the maximum number of items that you want
	// in the response. The default value is 50. The maximum value is 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThreatIntelSetsPaginator is a paginator for ListThreatIntelSets
type ListThreatIntelSetsPaginator struct {
	options   ListThreatIntelSetsPaginatorOptions
	client    ListThreatIntelSetsAPIClient
	params    *ListThreatIntelSetsInput
	nextToken *string
	firstPage bool
}

// NewListThreatIntelSetsPaginator returns a new ListThreatIntelSetsPaginator
func NewListThreatIntelSetsPaginator(client ListThreatIntelSetsAPIClient, params *ListThreatIntelSetsInput, optFns ...func(*ListThreatIntelSetsPaginatorOptions)) *ListThreatIntelSetsPaginator {
	if params == nil {
		params = &ListThreatIntelSetsInput{}
	}

	options := ListThreatIntelSetsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThreatIntelSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThreatIntelSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListThreatIntelSets page.
func (p *ListThreatIntelSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThreatIntelSetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListThreatIntelSets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListThreatIntelSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ListThreatIntelSets",
	}
}
