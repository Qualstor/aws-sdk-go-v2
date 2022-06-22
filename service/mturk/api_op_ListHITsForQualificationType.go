// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListHITsForQualificationType operation returns the HITs that use the given
// Qualification type for a Qualification requirement. The operation returns HITs
// of any status, except for HITs that have been deleted with the DeleteHIT
// operation or that have been auto-deleted.
func (c *Client) ListHITsForQualificationType(ctx context.Context, params *ListHITsForQualificationTypeInput, optFns ...func(*Options)) (*ListHITsForQualificationTypeOutput, error) {
	if params == nil {
		params = &ListHITsForQualificationTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHITsForQualificationType", params, optFns, c.addOperationListHITsForQualificationTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHITsForQualificationTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHITsForQualificationTypeInput struct {

	// The ID of the Qualification type to use when querying HITs.
	//
	// This member is required.
	QualificationTypeId *string

	// Limit the number of results returned.
	MaxResults *int32

	// Pagination Token
	NextToken *string

	noSmithyDocumentSerde
}

type ListHITsForQualificationTypeOutput struct {

	// The list of HIT elements returned by the query.
	HITs []types.HIT

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of HITs on this page in the filtered results list, equivalent to the
	// number of HITs being returned by this call.
	NumResults *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHITsForQualificationTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListHITsForQualificationType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHITsForQualificationType{}, middleware.After)
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
	if err = addOpListHITsForQualificationTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHITsForQualificationType(options.Region), middleware.Before); err != nil {
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

// ListHITsForQualificationTypeAPIClient is a client that implements the
// ListHITsForQualificationType operation.
type ListHITsForQualificationTypeAPIClient interface {
	ListHITsForQualificationType(context.Context, *ListHITsForQualificationTypeInput, ...func(*Options)) (*ListHITsForQualificationTypeOutput, error)
}

var _ ListHITsForQualificationTypeAPIClient = (*Client)(nil)

// ListHITsForQualificationTypePaginatorOptions is the paginator options for
// ListHITsForQualificationType
type ListHITsForQualificationTypePaginatorOptions struct {
	// Limit the number of results returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHITsForQualificationTypePaginator is a paginator for
// ListHITsForQualificationType
type ListHITsForQualificationTypePaginator struct {
	options   ListHITsForQualificationTypePaginatorOptions
	client    ListHITsForQualificationTypeAPIClient
	params    *ListHITsForQualificationTypeInput
	nextToken *string
	firstPage bool
}

// NewListHITsForQualificationTypePaginator returns a new
// ListHITsForQualificationTypePaginator
func NewListHITsForQualificationTypePaginator(client ListHITsForQualificationTypeAPIClient, params *ListHITsForQualificationTypeInput, optFns ...func(*ListHITsForQualificationTypePaginatorOptions)) *ListHITsForQualificationTypePaginator {
	if params == nil {
		params = &ListHITsForQualificationTypeInput{}
	}

	options := ListHITsForQualificationTypePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHITsForQualificationTypePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHITsForQualificationTypePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListHITsForQualificationType page.
func (p *ListHITsForQualificationTypePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHITsForQualificationTypeOutput, error) {
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

	result, err := p.client.ListHITsForQualificationType(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListHITsForQualificationType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListHITsForQualificationType",
	}
}
