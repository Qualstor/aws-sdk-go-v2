// Code generated by smithy-go-codegen DO NOT EDIT.

package voiceid

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/voiceid/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the speaker enrollment jobs in the domain with the specified
// JobStatus. If JobStatus is not provided, this lists all jobs with all possible
// speaker enrollment job statuses.
func (c *Client) ListSpeakerEnrollmentJobs(ctx context.Context, params *ListSpeakerEnrollmentJobsInput, optFns ...func(*Options)) (*ListSpeakerEnrollmentJobsOutput, error) {
	if params == nil {
		params = &ListSpeakerEnrollmentJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSpeakerEnrollmentJobs", params, optFns, c.addOperationListSpeakerEnrollmentJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSpeakerEnrollmentJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSpeakerEnrollmentJobsInput struct {

	// The identifier of the domain containing the speaker enrollment jobs.
	//
	// This member is required.
	DomainId *string

	// Provides the status of your speaker enrollment Job.
	JobStatus types.SpeakerEnrollmentJobStatus

	// The maximum number of results that are returned per call. You can use NextToken
	// to obtain further pages of results. The default is 100; the maximum allowed page
	// size is also 100.
	MaxResults *int32

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSpeakerEnrollmentJobsOutput struct {

	// A list containing details about each specified speaker enrollment job.
	JobSummaries []types.SpeakerEnrollmentJobSummary

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSpeakerEnrollmentJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListSpeakerEnrollmentJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListSpeakerEnrollmentJobs{}, middleware.After)
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
	if err = addOpListSpeakerEnrollmentJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSpeakerEnrollmentJobs(options.Region), middleware.Before); err != nil {
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

// ListSpeakerEnrollmentJobsAPIClient is a client that implements the
// ListSpeakerEnrollmentJobs operation.
type ListSpeakerEnrollmentJobsAPIClient interface {
	ListSpeakerEnrollmentJobs(context.Context, *ListSpeakerEnrollmentJobsInput, ...func(*Options)) (*ListSpeakerEnrollmentJobsOutput, error)
}

var _ ListSpeakerEnrollmentJobsAPIClient = (*Client)(nil)

// ListSpeakerEnrollmentJobsPaginatorOptions is the paginator options for
// ListSpeakerEnrollmentJobs
type ListSpeakerEnrollmentJobsPaginatorOptions struct {
	// The maximum number of results that are returned per call. You can use NextToken
	// to obtain further pages of results. The default is 100; the maximum allowed page
	// size is also 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSpeakerEnrollmentJobsPaginator is a paginator for ListSpeakerEnrollmentJobs
type ListSpeakerEnrollmentJobsPaginator struct {
	options   ListSpeakerEnrollmentJobsPaginatorOptions
	client    ListSpeakerEnrollmentJobsAPIClient
	params    *ListSpeakerEnrollmentJobsInput
	nextToken *string
	firstPage bool
}

// NewListSpeakerEnrollmentJobsPaginator returns a new
// ListSpeakerEnrollmentJobsPaginator
func NewListSpeakerEnrollmentJobsPaginator(client ListSpeakerEnrollmentJobsAPIClient, params *ListSpeakerEnrollmentJobsInput, optFns ...func(*ListSpeakerEnrollmentJobsPaginatorOptions)) *ListSpeakerEnrollmentJobsPaginator {
	if params == nil {
		params = &ListSpeakerEnrollmentJobsInput{}
	}

	options := ListSpeakerEnrollmentJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSpeakerEnrollmentJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSpeakerEnrollmentJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSpeakerEnrollmentJobs page.
func (p *ListSpeakerEnrollmentJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSpeakerEnrollmentJobsOutput, error) {
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

	result, err := p.client.ListSpeakerEnrollmentJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSpeakerEnrollmentJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "voiceid",
		OperationName: "ListSpeakerEnrollmentJobs",
	}
}
