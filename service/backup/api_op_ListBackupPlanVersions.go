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
)

// Returns version metadata of your backup plans, including Amazon Resource Names
// (ARNs), backup plan IDs, creation and deletion dates, plan names, and version
// IDs.
func (c *Client) ListBackupPlanVersions(ctx context.Context, params *ListBackupPlanVersionsInput, optFns ...func(*Options)) (*ListBackupPlanVersionsOutput, error) {
	if params == nil {
		params = &ListBackupPlanVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackupPlanVersions", params, optFns, c.addOperationListBackupPlanVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackupPlanVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBackupPlanVersionsInput struct {

	// Uniquely identifies a backup plan.
	//
	// This member is required.
	BackupPlanId *string

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBackupPlanVersionsOutput struct {

	// An array of version list items containing metadata about your backup plans.
	BackupPlanVersionsList []types.BackupPlansListMember

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBackupPlanVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackupPlanVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackupPlanVersions{}, middleware.After)
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
	if err = addOpListBackupPlanVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBackupPlanVersions(options.Region), middleware.Before); err != nil {
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

// ListBackupPlanVersionsAPIClient is a client that implements the
// ListBackupPlanVersions operation.
type ListBackupPlanVersionsAPIClient interface {
	ListBackupPlanVersions(context.Context, *ListBackupPlanVersionsInput, ...func(*Options)) (*ListBackupPlanVersionsOutput, error)
}

var _ ListBackupPlanVersionsAPIClient = (*Client)(nil)

// ListBackupPlanVersionsPaginatorOptions is the paginator options for
// ListBackupPlanVersions
type ListBackupPlanVersionsPaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBackupPlanVersionsPaginator is a paginator for ListBackupPlanVersions
type ListBackupPlanVersionsPaginator struct {
	options   ListBackupPlanVersionsPaginatorOptions
	client    ListBackupPlanVersionsAPIClient
	params    *ListBackupPlanVersionsInput
	nextToken *string
	firstPage bool
}

// NewListBackupPlanVersionsPaginator returns a new ListBackupPlanVersionsPaginator
func NewListBackupPlanVersionsPaginator(client ListBackupPlanVersionsAPIClient, params *ListBackupPlanVersionsInput, optFns ...func(*ListBackupPlanVersionsPaginatorOptions)) *ListBackupPlanVersionsPaginator {
	if params == nil {
		params = &ListBackupPlanVersionsInput{}
	}

	options := ListBackupPlanVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBackupPlanVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBackupPlanVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBackupPlanVersions page.
func (p *ListBackupPlanVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBackupPlanVersionsOutput, error) {
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

	result, err := p.client.ListBackupPlanVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBackupPlanVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListBackupPlanVersions",
	}
}
