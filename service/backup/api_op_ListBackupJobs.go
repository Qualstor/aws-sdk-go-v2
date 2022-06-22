// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of existing backup jobs for an authenticated account for the last
// 30 days. For a longer period of time, consider using these monitoring tools
// (https://docs.aws.amazon.com/aws-backup/latest/devguide/monitoring.html).
func (c *Client) ListBackupJobs(ctx context.Context, params *ListBackupJobsInput, optFns ...func(*Options)) (*ListBackupJobsOutput, error) {
	if params == nil {
		params = &ListBackupJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackupJobs", params, optFns, c.addOperationListBackupJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackupJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBackupJobsInput struct {

	// The account ID to list the jobs from. Returns only backup jobs associated with
	// the specified account ID. If used from an Organizations management account,
	// passing * returns all jobs across the organization.
	ByAccountId *string

	// Returns only backup jobs that will be stored in the specified backup vault.
	// Backup vaults are identified by names that are unique to the account used to
	// create them and the Amazon Web Services Region where they are created. They
	// consist of lowercase letters, numbers, and hyphens.
	ByBackupVaultName *string

	// Returns only backup jobs completed after a date expressed in Unix format and
	// Coordinated Universal Time (UTC).
	ByCompleteAfter *time.Time

	// Returns only backup jobs completed before a date expressed in Unix format and
	// Coordinated Universal Time (UTC).
	ByCompleteBefore *time.Time

	// Returns only backup jobs that were created after the specified date.
	ByCreatedAfter *time.Time

	// Returns only backup jobs that were created before the specified date.
	ByCreatedBefore *time.Time

	// Returns only backup jobs that match the specified resource Amazon Resource Name
	// (ARN).
	ByResourceArn *string

	// Returns only backup jobs for the specified resources:
	//
	// * Aurora for Amazon
	// Aurora
	//
	// * DocumentDB for Amazon DocumentDB (with MongoDB compatibility)
	//
	// *
	// DynamoDB for Amazon DynamoDB
	//
	// * EBS for Amazon Elastic Block Store
	//
	// * EC2 for
	// Amazon Elastic Compute Cloud
	//
	// * EFS for Amazon Elastic File System
	//
	// * FSx for
	// Amazon FSx
	//
	// * Neptune for Amazon Neptune
	//
	// * RDS for Amazon Relational Database
	// Service
	//
	// * Storage Gateway for Storage Gateway
	//
	// * S3 for Amazon S3
	//
	// *
	// VirtualMachine for virtual machines
	ByResourceType *string

	// Returns only backup jobs that are in the specified state.
	ByState types.BackupJobState

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBackupJobsOutput struct {

	// An array of structures containing metadata about your backup jobs returned in
	// JSON format.
	BackupJobs []types.BackupJob

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBackupJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackupJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackupJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBackupJobs(options.Region), middleware.Before); err != nil {
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

// ListBackupJobsAPIClient is a client that implements the ListBackupJobs
// operation.
type ListBackupJobsAPIClient interface {
	ListBackupJobs(context.Context, *ListBackupJobsInput, ...func(*Options)) (*ListBackupJobsOutput, error)
}

var _ ListBackupJobsAPIClient = (*Client)(nil)

// ListBackupJobsPaginatorOptions is the paginator options for ListBackupJobs
type ListBackupJobsPaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBackupJobsPaginator is a paginator for ListBackupJobs
type ListBackupJobsPaginator struct {
	options   ListBackupJobsPaginatorOptions
	client    ListBackupJobsAPIClient
	params    *ListBackupJobsInput
	nextToken *string
	firstPage bool
}

// NewListBackupJobsPaginator returns a new ListBackupJobsPaginator
func NewListBackupJobsPaginator(client ListBackupJobsAPIClient, params *ListBackupJobsInput, optFns ...func(*ListBackupJobsPaginatorOptions)) *ListBackupJobsPaginator {
	if params == nil {
		params = &ListBackupJobsInput{}
	}

	options := ListBackupJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBackupJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBackupJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBackupJobs page.
func (p *ListBackupJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBackupJobsOutput, error) {
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

	result, err := p.client.ListBackupJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBackupJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListBackupJobs",
	}
}
