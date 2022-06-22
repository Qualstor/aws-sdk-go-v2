// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more attachments between resources and transit gateways. By
// default, all attachments are described. Alternatively, you can filter the
// results by attachment ID, attachment state, resource ID, or resource owner.
func (c *Client) DescribeTransitGatewayAttachments(ctx context.Context, params *DescribeTransitGatewayAttachmentsInput, optFns ...func(*Options)) (*DescribeTransitGatewayAttachmentsOutput, error) {
	if params == nil {
		params = &DescribeTransitGatewayAttachmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGatewayAttachments", params, optFns, c.addOperationDescribeTransitGatewayAttachmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransitGatewayAttachmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewayAttachmentsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	// * association.state - The state
	// of the association (associating | associated | disassociating).
	//
	// *
	// association.transit-gateway-route-table-id - The ID of the route table for the
	// transit gateway.
	//
	// * resource-id - The ID of the resource.
	//
	// * resource-owner-id -
	// The ID of the Amazon Web Services account that owns the resource.
	//
	// *
	// resource-type - The resource type. Valid values are vpc | vpn |
	// direct-connect-gateway | peering | connect.
	//
	// * state - The state of the
	// attachment. Valid values are available | deleted | deleting | failed | failing |
	// initiatingRequest | modifying | pendingAcceptance | pending | rollingBack |
	// rejected | rejecting.
	//
	// * transit-gateway-attachment-id - The ID of the
	// attachment.
	//
	// * transit-gateway-id - The ID of the transit gateway.
	//
	// *
	// transit-gateway-owner-id - The ID of the Amazon Web Services account that owns
	// the transit gateway.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The IDs of the attachments.
	TransitGatewayAttachmentIds []string

	noSmithyDocumentSerde
}

type DescribeTransitGatewayAttachmentsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the attachments.
	TransitGatewayAttachments []types.TransitGatewayAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTransitGatewayAttachmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGatewayAttachments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGatewayAttachments{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGatewayAttachments(options.Region), middleware.Before); err != nil {
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

// DescribeTransitGatewayAttachmentsAPIClient is a client that implements the
// DescribeTransitGatewayAttachments operation.
type DescribeTransitGatewayAttachmentsAPIClient interface {
	DescribeTransitGatewayAttachments(context.Context, *DescribeTransitGatewayAttachmentsInput, ...func(*Options)) (*DescribeTransitGatewayAttachmentsOutput, error)
}

var _ DescribeTransitGatewayAttachmentsAPIClient = (*Client)(nil)

// DescribeTransitGatewayAttachmentsPaginatorOptions is the paginator options for
// DescribeTransitGatewayAttachments
type DescribeTransitGatewayAttachmentsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTransitGatewayAttachmentsPaginator is a paginator for
// DescribeTransitGatewayAttachments
type DescribeTransitGatewayAttachmentsPaginator struct {
	options   DescribeTransitGatewayAttachmentsPaginatorOptions
	client    DescribeTransitGatewayAttachmentsAPIClient
	params    *DescribeTransitGatewayAttachmentsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTransitGatewayAttachmentsPaginator returns a new
// DescribeTransitGatewayAttachmentsPaginator
func NewDescribeTransitGatewayAttachmentsPaginator(client DescribeTransitGatewayAttachmentsAPIClient, params *DescribeTransitGatewayAttachmentsInput, optFns ...func(*DescribeTransitGatewayAttachmentsPaginatorOptions)) *DescribeTransitGatewayAttachmentsPaginator {
	if params == nil {
		params = &DescribeTransitGatewayAttachmentsInput{}
	}

	options := DescribeTransitGatewayAttachmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTransitGatewayAttachmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTransitGatewayAttachmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTransitGatewayAttachments page.
func (p *DescribeTransitGatewayAttachmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTransitGatewayAttachmentsOutput, error) {
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

	result, err := p.client.DescribeTransitGatewayAttachments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTransitGatewayAttachments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTransitGatewayAttachments",
	}
}
