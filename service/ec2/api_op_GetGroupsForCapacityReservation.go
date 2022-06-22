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

// Lists the resource groups to which a Capacity Reservation has been added.
func (c *Client) GetGroupsForCapacityReservation(ctx context.Context, params *GetGroupsForCapacityReservationInput, optFns ...func(*Options)) (*GetGroupsForCapacityReservationOutput, error) {
	if params == nil {
		params = &GetGroupsForCapacityReservationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGroupsForCapacityReservation", params, optFns, c.addOperationGetGroupsForCapacityReservationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGroupsForCapacityReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGroupsForCapacityReservationInput struct {

	// The ID of the Capacity Reservation.
	//
	// This member is required.
	CapacityReservationId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given a
	// larger value than 500, you receive an error.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetGroupsForCapacityReservationOutput struct {

	// Information about the resource groups to which the Capacity Reservation has been
	// added.
	CapacityReservationGroups []types.CapacityReservationGroup

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGroupsForCapacityReservationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetGroupsForCapacityReservation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetGroupsForCapacityReservation{}, middleware.After)
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
	if err = addOpGetGroupsForCapacityReservationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGroupsForCapacityReservation(options.Region), middleware.Before); err != nil {
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

// GetGroupsForCapacityReservationAPIClient is a client that implements the
// GetGroupsForCapacityReservation operation.
type GetGroupsForCapacityReservationAPIClient interface {
	GetGroupsForCapacityReservation(context.Context, *GetGroupsForCapacityReservationInput, ...func(*Options)) (*GetGroupsForCapacityReservationOutput, error)
}

var _ GetGroupsForCapacityReservationAPIClient = (*Client)(nil)

// GetGroupsForCapacityReservationPaginatorOptions is the paginator options for
// GetGroupsForCapacityReservation
type GetGroupsForCapacityReservationPaginatorOptions struct {
	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given a
	// larger value than 500, you receive an error.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetGroupsForCapacityReservationPaginator is a paginator for
// GetGroupsForCapacityReservation
type GetGroupsForCapacityReservationPaginator struct {
	options   GetGroupsForCapacityReservationPaginatorOptions
	client    GetGroupsForCapacityReservationAPIClient
	params    *GetGroupsForCapacityReservationInput
	nextToken *string
	firstPage bool
}

// NewGetGroupsForCapacityReservationPaginator returns a new
// GetGroupsForCapacityReservationPaginator
func NewGetGroupsForCapacityReservationPaginator(client GetGroupsForCapacityReservationAPIClient, params *GetGroupsForCapacityReservationInput, optFns ...func(*GetGroupsForCapacityReservationPaginatorOptions)) *GetGroupsForCapacityReservationPaginator {
	if params == nil {
		params = &GetGroupsForCapacityReservationInput{}
	}

	options := GetGroupsForCapacityReservationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetGroupsForCapacityReservationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetGroupsForCapacityReservationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetGroupsForCapacityReservation page.
func (p *GetGroupsForCapacityReservationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetGroupsForCapacityReservationOutput, error) {
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

	result, err := p.client.GetGroupsForCapacityReservation(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetGroupsForCapacityReservation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetGroupsForCapacityReservation",
	}
}
