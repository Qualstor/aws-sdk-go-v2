// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the scaling activities in the account and Region. When
// scaling events occur, you see a record of the scaling activity in the scaling
// activities. For more information, see Verifying a scaling activity for an Auto
// Scaling group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-verify-scaling-activity.html)
// in the Amazon EC2 Auto Scaling User Guide. If the scaling event succeeds, the
// value of the StatusCode element in the response is Successful. If an attempt to
// launch instances failed, the StatusCode value is Failed or Cancelled and the
// StatusMessage element in the response indicates the cause of the failure. For
// help interpreting the StatusMessage, see Troubleshooting Amazon EC2 Auto Scaling
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/CHAP_Troubleshooting.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DescribeScalingActivities(ctx context.Context, params *DescribeScalingActivitiesInput, optFns ...func(*Options)) (*DescribeScalingActivitiesOutput, error) {
	if params == nil {
		params = &DescribeScalingActivitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScalingActivities", params, optFns, c.addOperationDescribeScalingActivitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScalingActivitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScalingActivitiesInput struct {

	// The activity IDs of the desired scaling activities. If you omit this parameter,
	// all activities for the past six weeks are described. If unknown activities are
	// requested, they are ignored with no error. If you specify an Auto Scaling group,
	// the results are limited to that group. Array Members: Maximum number of 50 IDs.
	ActivityIds []string

	// The name of the Auto Scaling group.
	AutoScalingGroupName *string

	// Indicates whether to include scaling activity from deleted Auto Scaling groups.
	IncludeDeletedGroups *bool

	// The maximum number of items to return with this call. The default value is 100
	// and the maximum value is 100.
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeScalingActivitiesOutput struct {

	// The scaling activities. Activities are sorted by start time. Activities still in
	// progress are described first.
	//
	// This member is required.
	Activities []types.Activity

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScalingActivitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeScalingActivities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeScalingActivities{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScalingActivities(options.Region), middleware.Before); err != nil {
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

// DescribeScalingActivitiesAPIClient is a client that implements the
// DescribeScalingActivities operation.
type DescribeScalingActivitiesAPIClient interface {
	DescribeScalingActivities(context.Context, *DescribeScalingActivitiesInput, ...func(*Options)) (*DescribeScalingActivitiesOutput, error)
}

var _ DescribeScalingActivitiesAPIClient = (*Client)(nil)

// DescribeScalingActivitiesPaginatorOptions is the paginator options for
// DescribeScalingActivities
type DescribeScalingActivitiesPaginatorOptions struct {
	// The maximum number of items to return with this call. The default value is 100
	// and the maximum value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeScalingActivitiesPaginator is a paginator for DescribeScalingActivities
type DescribeScalingActivitiesPaginator struct {
	options   DescribeScalingActivitiesPaginatorOptions
	client    DescribeScalingActivitiesAPIClient
	params    *DescribeScalingActivitiesInput
	nextToken *string
	firstPage bool
}

// NewDescribeScalingActivitiesPaginator returns a new
// DescribeScalingActivitiesPaginator
func NewDescribeScalingActivitiesPaginator(client DescribeScalingActivitiesAPIClient, params *DescribeScalingActivitiesInput, optFns ...func(*DescribeScalingActivitiesPaginatorOptions)) *DescribeScalingActivitiesPaginator {
	if params == nil {
		params = &DescribeScalingActivitiesInput{}
	}

	options := DescribeScalingActivitiesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeScalingActivitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeScalingActivitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeScalingActivities page.
func (p *DescribeScalingActivitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeScalingActivitiesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeScalingActivities(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeScalingActivities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeScalingActivities",
	}
}
