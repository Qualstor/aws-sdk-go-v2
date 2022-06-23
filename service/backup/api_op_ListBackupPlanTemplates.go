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

// Returns metadata of your saved backup plan templates, including the template ID,
// name, and the creation and deletion dates.
func (c *Client) ListBackupPlanTemplates(ctx context.Context, params *ListBackupPlanTemplatesInput, optFns ...func(*Options)) (*ListBackupPlanTemplatesOutput, error) {
	if params == nil {
		params = &ListBackupPlanTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackupPlanTemplates", params, optFns, c.addOperationListBackupPlanTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackupPlanTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBackupPlanTemplatesInput struct {

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBackupPlanTemplatesOutput struct {

	// An array of template list items containing metadata about your saved templates.
	BackupPlanTemplatesList []types.BackupPlanTemplatesListMember

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBackupPlanTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackupPlanTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackupPlanTemplates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBackupPlanTemplates(options.Region), middleware.Before); err != nil {
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

// ListBackupPlanTemplatesAPIClient is a client that implements the
// ListBackupPlanTemplates operation.
type ListBackupPlanTemplatesAPIClient interface {
	ListBackupPlanTemplates(context.Context, *ListBackupPlanTemplatesInput, ...func(*Options)) (*ListBackupPlanTemplatesOutput, error)
}

var _ ListBackupPlanTemplatesAPIClient = (*Client)(nil)

// ListBackupPlanTemplatesPaginatorOptions is the paginator options for
// ListBackupPlanTemplates
type ListBackupPlanTemplatesPaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBackupPlanTemplatesPaginator is a paginator for ListBackupPlanTemplates
type ListBackupPlanTemplatesPaginator struct {
	options   ListBackupPlanTemplatesPaginatorOptions
	client    ListBackupPlanTemplatesAPIClient
	params    *ListBackupPlanTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListBackupPlanTemplatesPaginator returns a new
// ListBackupPlanTemplatesPaginator
func NewListBackupPlanTemplatesPaginator(client ListBackupPlanTemplatesAPIClient, params *ListBackupPlanTemplatesInput, optFns ...func(*ListBackupPlanTemplatesPaginatorOptions)) *ListBackupPlanTemplatesPaginator {
	if params == nil {
		params = &ListBackupPlanTemplatesInput{}
	}

	options := ListBackupPlanTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBackupPlanTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBackupPlanTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBackupPlanTemplates page.
func (p *ListBackupPlanTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBackupPlanTemplatesOutput, error) {
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

	result, err := p.client.ListBackupPlanTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBackupPlanTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListBackupPlanTemplates",
	}
}
