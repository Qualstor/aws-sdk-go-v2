// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified target groups or all of your target groups. By default,
// all target groups are described. Alternatively, you can specify one of the
// following to filter the results: the ARN of the load balancer, the names of one
// or more target groups, or the ARNs of one or more target groups.
func (c *Client) DescribeTargetGroups(ctx context.Context, params *DescribeTargetGroupsInput, optFns ...func(*Options)) (*DescribeTargetGroupsOutput, error) {
	if params == nil {
		params = &DescribeTargetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTargetGroups", params, optFns, c.addOperationDescribeTargetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTargetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTargetGroupsInput struct {

	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn *string

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	// The names of the target groups.
	Names []string

	// The maximum number of results to return with this call.
	PageSize *int32

	// The Amazon Resource Names (ARN) of the target groups.
	TargetGroupArns []string

	noSmithyDocumentSerde
}

type DescribeTargetGroupsOutput struct {

	// If there are additional results, this is the marker for the next set of results.
	// Otherwise, this is null.
	NextMarker *string

	// Information about the target groups.
	TargetGroups []types.TargetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTargetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeTargetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeTargetGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTargetGroups(options.Region), middleware.Before); err != nil {
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

// DescribeTargetGroupsAPIClient is a client that implements the
// DescribeTargetGroups operation.
type DescribeTargetGroupsAPIClient interface {
	DescribeTargetGroups(context.Context, *DescribeTargetGroupsInput, ...func(*Options)) (*DescribeTargetGroupsOutput, error)
}

var _ DescribeTargetGroupsAPIClient = (*Client)(nil)

// DescribeTargetGroupsPaginatorOptions is the paginator options for
// DescribeTargetGroups
type DescribeTargetGroupsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTargetGroupsPaginator is a paginator for DescribeTargetGroups
type DescribeTargetGroupsPaginator struct {
	options   DescribeTargetGroupsPaginatorOptions
	client    DescribeTargetGroupsAPIClient
	params    *DescribeTargetGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTargetGroupsPaginator returns a new DescribeTargetGroupsPaginator
func NewDescribeTargetGroupsPaginator(client DescribeTargetGroupsAPIClient, params *DescribeTargetGroupsInput, optFns ...func(*DescribeTargetGroupsPaginatorOptions)) *DescribeTargetGroupsPaginator {
	if params == nil {
		params = &DescribeTargetGroupsInput{}
	}

	options := DescribeTargetGroupsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTargetGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTargetGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTargetGroups page.
func (p *DescribeTargetGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTargetGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.DescribeTargetGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeTargetGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeTargetGroups",
	}
}
