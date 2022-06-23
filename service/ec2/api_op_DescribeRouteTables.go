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

// Describes one or more of your route tables. Each subnet in your VPC must be
// associated with a route table. If a subnet is not explicitly associated with any
// route table, it is implicitly associated with the main route table. This command
// does not return the subnet ID for implicit associations. For more information,
// see Route tables
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) DescribeRouteTables(ctx context.Context, params *DescribeRouteTablesInput, optFns ...func(*Options)) (*DescribeRouteTablesOutput, error) {
	if params == nil {
		params = &DescribeRouteTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRouteTables", params, optFns, c.addOperationDescribeRouteTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRouteTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRouteTablesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * association.route-table-association-id - The ID of an
	// association ID for the route table.
	//
	// * association.route-table-id - The ID of
	// the route table involved in the association.
	//
	// * association.subnet-id - The ID
	// of the subnet involved in the association.
	//
	// * association.main - Indicates
	// whether the route table is the main route table for the VPC (true | false).
	// Route tables that do not have an association ID are not returned in the
	// response.
	//
	// * owner-id - The ID of the Amazon Web Services account that owns the
	// route table.
	//
	// * route-table-id - The ID of the route table.
	//
	// *
	// route.destination-cidr-block - The IPv4 CIDR range specified in a route in the
	// table.
	//
	// * route.destination-ipv6-cidr-block - The IPv6 CIDR range specified in a
	// route in the route table.
	//
	// * route.destination-prefix-list-id - The ID (prefix)
	// of the Amazon Web Service specified in a route in the table.
	//
	// *
	// route.egress-only-internet-gateway-id - The ID of an egress-only Internet
	// gateway specified in a route in the route table.
	//
	// * route.gateway-id - The ID of
	// a gateway specified in a route in the table.
	//
	// * route.instance-id - The ID of an
	// instance specified in a route in the table.
	//
	// * route.nat-gateway-id - The ID of
	// a NAT gateway.
	//
	// * route.transit-gateway-id - The ID of a transit gateway.
	//
	// *
	// route.origin - Describes how the route was created. CreateRouteTable indicates
	// that the route was automatically created when the route table was created;
	// CreateRoute indicates that the route was manually added to the route table;
	// EnableVgwRoutePropagation indicates that the route was propagated by route
	// propagation.
	//
	// * route.state - The state of a route in the route table (active |
	// blackhole). The blackhole state indicates that the route's target isn't
	// available (for example, the specified gateway isn't attached to the VPC, the
	// specified NAT instance has been terminated, and so on).
	//
	// *
	// route.vpc-peering-connection-id - The ID of a VPC peering connection specified
	// in a route in the table.
	//
	// * tag: - The key/value combination of a tag assigned
	// to the resource. Use the tag key in the filter name and the tag value as the
	// filter value. For example, to find all resources that have a tag with the key
	// Owner and the value TeamA, specify tag:Owner for the filter name and TeamA for
	// the filter value.
	//
	// * tag-key - The key of a tag assigned to the resource. Use
	// this filter to find all resources assigned a tag with a specific key, regardless
	// of the tag value.
	//
	// * vpc-id - The ID of the VPC for the route table.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more route table IDs. Default: Describes all your route tables.
	RouteTableIds []string

	noSmithyDocumentSerde
}

// Contains the output of DescribeRouteTables.
type DescribeRouteTablesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about one or more route tables.
	RouteTables []types.RouteTable

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRouteTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeRouteTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeRouteTables{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRouteTables(options.Region), middleware.Before); err != nil {
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

// DescribeRouteTablesAPIClient is a client that implements the DescribeRouteTables
// operation.
type DescribeRouteTablesAPIClient interface {
	DescribeRouteTables(context.Context, *DescribeRouteTablesInput, ...func(*Options)) (*DescribeRouteTablesOutput, error)
}

var _ DescribeRouteTablesAPIClient = (*Client)(nil)

// DescribeRouteTablesPaginatorOptions is the paginator options for
// DescribeRouteTables
type DescribeRouteTablesPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRouteTablesPaginator is a paginator for DescribeRouteTables
type DescribeRouteTablesPaginator struct {
	options   DescribeRouteTablesPaginatorOptions
	client    DescribeRouteTablesAPIClient
	params    *DescribeRouteTablesInput
	nextToken *string
	firstPage bool
}

// NewDescribeRouteTablesPaginator returns a new DescribeRouteTablesPaginator
func NewDescribeRouteTablesPaginator(client DescribeRouteTablesAPIClient, params *DescribeRouteTablesInput, optFns ...func(*DescribeRouteTablesPaginatorOptions)) *DescribeRouteTablesPaginator {
	if params == nil {
		params = &DescribeRouteTablesInput{}
	}

	options := DescribeRouteTablesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRouteTablesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRouteTablesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRouteTables page.
func (p *DescribeRouteTablesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRouteTablesOutput, error) {
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

	result, err := p.client.DescribeRouteTables(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRouteTables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeRouteTables",
	}
}
