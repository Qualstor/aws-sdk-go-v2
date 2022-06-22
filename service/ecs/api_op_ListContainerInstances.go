// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of container instances in a specified cluster. You can filter the
// results of a ListContainerInstances operation with cluster query language
// statements inside the filter parameter. For more information, see Cluster Query
// Language
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) ListContainerInstances(ctx context.Context, params *ListContainerInstancesInput, optFns ...func(*Options)) (*ListContainerInstancesOutput, error) {
	if params == nil {
		params = &ListContainerInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContainerInstances", params, optFns, c.addOperationListContainerInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContainerInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContainerInstancesInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// container instances to list. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string

	// You can filter the results of a ListContainerInstances operation with cluster
	// query language statements. For more information, see Cluster Query Language
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html)
	// in the Amazon Elastic Container Service Developer Guide.
	Filter *string

	// The maximum number of container instance results that ListContainerInstances
	// returned in paginated output. When this parameter is used,
	// ListContainerInstances only returns maxResults results in a single page along
	// with a nextToken response element. The remaining results of the initial request
	// can be seen by sending another ListContainerInstances request with the returned
	// nextToken value. This value can be between 1 and 100. If this parameter isn't
	// used, then ListContainerInstances returns up to 100 results and a nextToken
	// value if applicable.
	MaxResults *int32

	// The nextToken value returned from a ListContainerInstances request indicating
	// that more results are available to fulfill the request and further calls are
	// needed. If maxResults was provided, it's possible the number of results to be
	// fewer than maxResults. This token should be treated as an opaque identifier that
	// is only used to retrieve the next items in a list and not for other programmatic
	// purposes.
	NextToken *string

	// Filters the container instances by status. For example, if you specify the
	// DRAINING status, the results include only container instances that have been set
	// to DRAINING using UpdateContainerInstancesState. If you don't specify this
	// parameter, the default is to include container instances set to all states other
	// than INACTIVE.
	Status types.ContainerInstanceStatus

	noSmithyDocumentSerde
}

type ListContainerInstancesOutput struct {

	// The list of container instances with full ARN entries for each container
	// instance associated with the specified cluster.
	ContainerInstanceArns []string

	// The nextToken value to include in a future ListContainerInstances request. When
	// the results of a ListContainerInstances request exceed maxResults, this value
	// can be used to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListContainerInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListContainerInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListContainerInstances{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContainerInstances(options.Region), middleware.Before); err != nil {
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

// ListContainerInstancesAPIClient is a client that implements the
// ListContainerInstances operation.
type ListContainerInstancesAPIClient interface {
	ListContainerInstances(context.Context, *ListContainerInstancesInput, ...func(*Options)) (*ListContainerInstancesOutput, error)
}

var _ ListContainerInstancesAPIClient = (*Client)(nil)

// ListContainerInstancesPaginatorOptions is the paginator options for
// ListContainerInstances
type ListContainerInstancesPaginatorOptions struct {
	// The maximum number of container instance results that ListContainerInstances
	// returned in paginated output. When this parameter is used,
	// ListContainerInstances only returns maxResults results in a single page along
	// with a nextToken response element. The remaining results of the initial request
	// can be seen by sending another ListContainerInstances request with the returned
	// nextToken value. This value can be between 1 and 100. If this parameter isn't
	// used, then ListContainerInstances returns up to 100 results and a nextToken
	// value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListContainerInstancesPaginator is a paginator for ListContainerInstances
type ListContainerInstancesPaginator struct {
	options   ListContainerInstancesPaginatorOptions
	client    ListContainerInstancesAPIClient
	params    *ListContainerInstancesInput
	nextToken *string
	firstPage bool
}

// NewListContainerInstancesPaginator returns a new ListContainerInstancesPaginator
func NewListContainerInstancesPaginator(client ListContainerInstancesAPIClient, params *ListContainerInstancesInput, optFns ...func(*ListContainerInstancesPaginatorOptions)) *ListContainerInstancesPaginator {
	if params == nil {
		params = &ListContainerInstancesInput{}
	}

	options := ListContainerInstancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListContainerInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListContainerInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListContainerInstances page.
func (p *ListContainerInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListContainerInstancesOutput, error) {
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

	result, err := p.client.ListContainerInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListContainerInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "ListContainerInstances",
	}
}
