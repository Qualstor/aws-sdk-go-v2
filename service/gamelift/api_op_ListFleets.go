// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a collection of fleet resources in an Amazon Web Services Region. You
// can call this operation to get fleets in a previously selected default Region
// (see
// https://docs.aws.amazon.com/credref/latest/refdocs/setting-global-region.html
// (https://docs.aws.amazon.com/credref/latest/refdocs/setting-global-region.html)or
// specify a Region in your request. You can filter the result set to find only
// those fleets that are deployed with a specific build or script. For fleets that
// have multiple locations, this operation retrieves fleets based on their home
// Region only. This operation can be used in the following ways:
//
// * To get a list
// of all fleets in a Region, don't provide a build or script identifier.
//
// * To get
// a list of all fleets where a specific custom game build is deployed, provide the
// build ID.
//
// * To get a list of all Realtime Servers fleets with a specific
// configuration script, provide the script ID.
//
// Use the pagination parameters to
// retrieve results as a set of sequential pages. If successful, a list of fleet
// IDs that match the request parameters is returned. A NextToken value is also
// returned if there are more result pages to retrieve. Fleet resources are not
// listed in a particular order. Learn more Setting up GameLift fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related actions CreateFleet | UpdateFleetCapacity | PutScalingPolicy |
// DescribeEC2InstanceLimits | DescribeFleetAttributes |
// DescribeFleetLocationAttributes | UpdateFleetAttributes | StopFleetActions |
// DeleteFleet | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) ListFleets(ctx context.Context, params *ListFleetsInput, optFns ...func(*Options)) (*ListFleetsOutput, error) {
	if params == nil {
		params = &ListFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFleets", params, optFns, c.addOperationListFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type ListFleetsInput struct {

	// A unique identifier for the build to request fleets for. Use this parameter to
	// return only fleets using a specified build. Use either the build ID or ARN
	// value.
	BuildId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// A token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// A unique identifier for the Realtime script to request fleets for. Use this
	// parameter to return only fleets using a specified script. Use either the script
	// ID or ARN value.
	ScriptId *string

	noSmithyDocumentSerde
}

// Represents the returned data in response to a request operation.
type ListFleetsOutput struct {

	// A set of fleet IDs that match the list request. You can retrieve additional
	// information about all returned fleets by passing this result set to a
	// DescribeFleetAttributes, DescribeFleetCapacity, or DescribeFleetUtilization
	// call.
	FleetIds []string

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFleets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFleets(options.Region), middleware.Before); err != nil {
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

// ListFleetsAPIClient is a client that implements the ListFleets operation.
type ListFleetsAPIClient interface {
	ListFleets(context.Context, *ListFleetsInput, ...func(*Options)) (*ListFleetsOutput, error)
}

var _ ListFleetsAPIClient = (*Client)(nil)

// ListFleetsPaginatorOptions is the paginator options for ListFleets
type ListFleetsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFleetsPaginator is a paginator for ListFleets
type ListFleetsPaginator struct {
	options   ListFleetsPaginatorOptions
	client    ListFleetsAPIClient
	params    *ListFleetsInput
	nextToken *string
	firstPage bool
}

// NewListFleetsPaginator returns a new ListFleetsPaginator
func NewListFleetsPaginator(client ListFleetsAPIClient, params *ListFleetsInput, optFns ...func(*ListFleetsPaginatorOptions)) *ListFleetsPaginator {
	if params == nil {
		params = &ListFleetsInput{}
	}

	options := ListFleetsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFleetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFleetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFleets page.
func (p *ListFleetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFleetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListFleets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFleets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ListFleets",
	}
}
