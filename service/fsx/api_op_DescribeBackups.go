// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the description of a specific Amazon FSx backup, if a BackupIds value is
// provided for that backup. Otherwise, it returns all backups owned by your Amazon
// Web Services account in the Amazon Web Services Region of the endpoint that
// you're calling. When retrieving all backups, you can optionally specify the
// MaxResults parameter to limit the number of backups in a response. If more
// backups remain, Amazon FSx returns a NextToken value in the response. In this
// case, send a later request with the NextToken request parameter set to the value
// of the NextToken value from the last response. This operation is used in an
// iterative process to retrieve a list of your backups. DescribeBackups is called
// first without a NextToken value. Then the operation continues to be called with
// the NextToken parameter set to the value of the last NextToken value until a
// response has no NextToken value. When using this operation, keep the following
// in mind:
//
// * The operation might return fewer than the MaxResults value of backup
// descriptions while still including a NextToken value.
//
// * The order of the
// backups returned in the response of one DescribeBackups call and the order of
// the backups returned across the responses of a multi-call iteration is
// unspecified.
func (c *Client) DescribeBackups(ctx context.Context, params *DescribeBackupsInput, optFns ...func(*Options)) (*DescribeBackupsOutput, error) {
	if params == nil {
		params = &DescribeBackupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBackups", params, optFns, c.addOperationDescribeBackupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for the DescribeBackups operation.
type DescribeBackupsInput struct {

	// The IDs of the backups that you want to retrieve. This parameter value overrides
	// any filters. If any IDs aren't found, a BackupNotFound error occurs.
	BackupIds []string

	// The filters structure. The supported names are file-system-id, backup-type,
	// file-system-type, and volume-id.
	Filters []types.Filter

	// Maximum number of backups to return in the response. This parameter value must
	// be greater than 0. The number of items that Amazon FSx returns is the minimum of
	// the MaxResults parameter specified in the request and the service's internal
	// maximum number of items per page.
	MaxResults *int32

	// An opaque pagination token returned from a previous DescribeBackups operation.
	// If a token is present, the operation continues the list from where the returning
	// call left off.
	NextToken *string

	noSmithyDocumentSerde
}

// Response object for the DescribeBackups operation.
type DescribeBackupsOutput struct {

	// An array of backups.
	Backups []types.Backup

	// A NextToken value is present if there are more backups than returned in the
	// response. You can use the NextToken value in the subsequent request to fetch the
	// backups.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBackupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBackups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBackups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBackups(options.Region), middleware.Before); err != nil {
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

// DescribeBackupsAPIClient is a client that implements the DescribeBackups
// operation.
type DescribeBackupsAPIClient interface {
	DescribeBackups(context.Context, *DescribeBackupsInput, ...func(*Options)) (*DescribeBackupsOutput, error)
}

var _ DescribeBackupsAPIClient = (*Client)(nil)

// DescribeBackupsPaginatorOptions is the paginator options for DescribeBackups
type DescribeBackupsPaginatorOptions struct {
	// Maximum number of backups to return in the response. This parameter value must
	// be greater than 0. The number of items that Amazon FSx returns is the minimum of
	// the MaxResults parameter specified in the request and the service's internal
	// maximum number of items per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBackupsPaginator is a paginator for DescribeBackups
type DescribeBackupsPaginator struct {
	options   DescribeBackupsPaginatorOptions
	client    DescribeBackupsAPIClient
	params    *DescribeBackupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeBackupsPaginator returns a new DescribeBackupsPaginator
func NewDescribeBackupsPaginator(client DescribeBackupsAPIClient, params *DescribeBackupsInput, optFns ...func(*DescribeBackupsPaginatorOptions)) *DescribeBackupsPaginator {
	if params == nil {
		params = &DescribeBackupsInput{}
	}

	options := DescribeBackupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeBackupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBackupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeBackups page.
func (p *DescribeBackupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBackupsOutput, error) {
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

	result, err := p.client.DescribeBackups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeBackups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DescribeBackups",
	}
}
