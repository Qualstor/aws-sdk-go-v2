// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/redshiftdata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the detailed information about a table from metadata in the cluster.
// The information includes its columns. A token is returned to page through the
// column list. Depending on the authorization method, use one of the following
// combinations of request parameters:
//
// * Secrets Manager - when connecting to a
// cluster, specify the Amazon Resource Name (ARN) of the secret, the database
// name, and the cluster identifier that matches the cluster in the secret. When
// connecting to a serverless endpoint, specify the Amazon Resource Name (ARN) of
// the secret and the database name.
//
// * Temporary credentials - when connecting to
// a cluster, specify the cluster identifier, the database name, and the database
// user name. Also, permission to call the redshift:GetClusterCredentials operation
// is required. When connecting to a serverless endpoint, specify the database
// name.
func (c *Client) DescribeTable(ctx context.Context, params *DescribeTableInput, optFns ...func(*Options)) (*DescribeTableOutput, error) {
	if params == nil {
		params = &DescribeTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTable", params, optFns, c.addOperationDescribeTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTableInput struct {

	// The name of the database that contains the tables to be described. If
	// ConnectedDatabase is not specified, this is also the database to connect to with
	// your authentication credentials.
	//
	// This member is required.
	Database *string

	// The cluster identifier. This parameter is required when connecting to a cluster
	// and authenticating using either Secrets Manager or temporary credentials.
	ClusterIdentifier *string

	// A database name. The connected database is specified when you connect with your
	// authentication credentials.
	ConnectedDatabase *string

	// The database user name. This parameter is required when connecting to a cluster
	// and authenticating using temporary credentials.
	DbUser *string

	// The maximum number of tables to return in the response. If more tables exist
	// than fit in one response, then NextToken is returned to page through the
	// results.
	MaxResults int32

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// The schema that contains the table. If no schema is specified, then matching
	// tables for all schemas are returned.
	Schema *string

	// The name or ARN of the secret that enables access to the database. This
	// parameter is required when authenticating using Secrets Manager.
	SecretArn *string

	// The table name. If no table is specified, then all tables for all matching
	// schemas are returned. If no table and no schema is specified, then all tables
	// for all schemas in the database are returned
	Table *string

	noSmithyDocumentSerde
}

type DescribeTableOutput struct {

	// A list of columns in the table.
	ColumnList []types.ColumnMetadata

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// The table name.
	TableName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTable{}, middleware.After)
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
	if err = addOpDescribeTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTable(options.Region), middleware.Before); err != nil {
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

// DescribeTableAPIClient is a client that implements the DescribeTable operation.
type DescribeTableAPIClient interface {
	DescribeTable(context.Context, *DescribeTableInput, ...func(*Options)) (*DescribeTableOutput, error)
}

var _ DescribeTableAPIClient = (*Client)(nil)

// DescribeTablePaginatorOptions is the paginator options for DescribeTable
type DescribeTablePaginatorOptions struct {
	// The maximum number of tables to return in the response. If more tables exist
	// than fit in one response, then NextToken is returned to page through the
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTablePaginator is a paginator for DescribeTable
type DescribeTablePaginator struct {
	options   DescribeTablePaginatorOptions
	client    DescribeTableAPIClient
	params    *DescribeTableInput
	nextToken *string
	firstPage bool
}

// NewDescribeTablePaginator returns a new DescribeTablePaginator
func NewDescribeTablePaginator(client DescribeTableAPIClient, params *DescribeTableInput, optFns ...func(*DescribeTablePaginatorOptions)) *DescribeTablePaginator {
	if params == nil {
		params = &DescribeTableInput{}
	}

	options := DescribeTablePaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTablePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTablePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTable page.
func (p *DescribeTablePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTableOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeTable(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-data",
		OperationName: "DescribeTable",
	}
}
