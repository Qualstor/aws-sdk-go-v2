// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about AWS CloudHSM clusters. This is a paginated operation,
// which means that each response might contain only a subset of all the clusters.
// When the response contains only a subset of clusters, it includes a NextToken
// value. Use this value in a subsequent DescribeClusters request to get more
// clusters. When you receive a response with no NextToken (or an empty or null
// value), that means there are no more clusters to get.
func (c *Client) DescribeClusters(ctx context.Context, params *DescribeClustersInput, optFns ...func(*Options)) (*DescribeClustersOutput, error) {
	if params == nil {
		params = &DescribeClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusters", params, optFns, c.addOperationDescribeClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClustersInput struct {

	// One or more filters to limit the items returned in the response. Use the
	// clusterIds filter to return only the specified clusters. Specify clusters by
	// their cluster identifier (ID). Use the vpcIds filter to return only the clusters
	// in the specified virtual private clouds (VPCs). Specify VPCs by their VPC
	// identifier (ID). Use the states filter to return only clusters that match the
	// specified state.
	Filters map[string][]string

	// The maximum number of clusters to return in the response. When there are more
	// clusters than the number you specify, the response contains a NextToken value.
	MaxResults *int32

	// The NextToken value that you received in the previous response. Use this value
	// to get more clusters.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeClustersOutput struct {

	// A list of clusters.
	Clusters []types.Cluster

	// An opaque string that indicates that the response contains only a subset of
	// clusters. Use this value in a subsequent DescribeClusters request to get more
	// clusters.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeClusters{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusters(options.Region), middleware.Before); err != nil {
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

// DescribeClustersAPIClient is a client that implements the DescribeClusters
// operation.
type DescribeClustersAPIClient interface {
	DescribeClusters(context.Context, *DescribeClustersInput, ...func(*Options)) (*DescribeClustersOutput, error)
}

var _ DescribeClustersAPIClient = (*Client)(nil)

// DescribeClustersPaginatorOptions is the paginator options for DescribeClusters
type DescribeClustersPaginatorOptions struct {
	// The maximum number of clusters to return in the response. When there are more
	// clusters than the number you specify, the response contains a NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClustersPaginator is a paginator for DescribeClusters
type DescribeClustersPaginator struct {
	options   DescribeClustersPaginatorOptions
	client    DescribeClustersAPIClient
	params    *DescribeClustersInput
	nextToken *string
	firstPage bool
}

// NewDescribeClustersPaginator returns a new DescribeClustersPaginator
func NewDescribeClustersPaginator(client DescribeClustersAPIClient, params *DescribeClustersInput, optFns ...func(*DescribeClustersPaginatorOptions)) *DescribeClustersPaginator {
	if params == nil {
		params = &DescribeClustersInput{}
	}

	options := DescribeClustersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeClustersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClustersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeClusters page.
func (p *DescribeClustersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClustersOutput, error) {
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

	result, err := p.client.DescribeClusters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "DescribeClusters",
	}
}
