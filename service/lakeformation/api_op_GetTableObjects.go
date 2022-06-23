// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the set of Amazon S3 objects that make up the specified governed table.
// A transaction ID or timestamp can be specified for time-travel queries.
func (c *Client) GetTableObjects(ctx context.Context, params *GetTableObjectsInput, optFns ...func(*Options)) (*GetTableObjectsOutput, error) {
	if params == nil {
		params = &GetTableObjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTableObjects", params, optFns, c.addOperationGetTableObjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTableObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTableObjectsInput struct {

	// The database containing the governed table.
	//
	// This member is required.
	DatabaseName *string

	// The governed table for which to retrieve objects.
	//
	// This member is required.
	TableName *string

	// The catalog containing the governed table. Defaults to the caller’s account.
	CatalogId *string

	// Specifies how many values to return in a page.
	MaxResults *int32

	// A continuation token if this is not the first call to retrieve these objects.
	NextToken *string

	// A predicate to filter the objects returned based on the partition keys defined
	// in the governed table.
	//
	// * The comparison operators supported are: =, >, <, >=,
	// <=
	//
	// * The logical operators supported are: AND
	//
	// * The data types supported are
	// integer, long, date(yyyy-MM-dd), timestamp(yyyy-MM-dd HH:mm:ssXXX or yyyy-MM-dd
	// HH:mm:ss"), string and decimal.
	PartitionPredicate *string

	// The time as of when to read the governed table contents. If not set, the most
	// recent transaction commit time is used. Cannot be specified along with
	// TransactionId.
	QueryAsOfTime *time.Time

	// The transaction ID at which to read the governed table contents. If this
	// transaction has aborted, an error is returned. If not set, defaults to the most
	// recent committed transaction. Cannot be specified along with QueryAsOfTime.
	TransactionId *string

	noSmithyDocumentSerde
}

type GetTableObjectsOutput struct {

	// A continuation token indicating whether additional data is available.
	NextToken *string

	// A list of objects organized by partition keys.
	Objects []types.PartitionObjects

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTableObjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTableObjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTableObjects{}, middleware.After)
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
	if err = addOpGetTableObjectsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTableObjects(options.Region), middleware.Before); err != nil {
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

// GetTableObjectsAPIClient is a client that implements the GetTableObjects
// operation.
type GetTableObjectsAPIClient interface {
	GetTableObjects(context.Context, *GetTableObjectsInput, ...func(*Options)) (*GetTableObjectsOutput, error)
}

var _ GetTableObjectsAPIClient = (*Client)(nil)

// GetTableObjectsPaginatorOptions is the paginator options for GetTableObjects
type GetTableObjectsPaginatorOptions struct {
	// Specifies how many values to return in a page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTableObjectsPaginator is a paginator for GetTableObjects
type GetTableObjectsPaginator struct {
	options   GetTableObjectsPaginatorOptions
	client    GetTableObjectsAPIClient
	params    *GetTableObjectsInput
	nextToken *string
	firstPage bool
}

// NewGetTableObjectsPaginator returns a new GetTableObjectsPaginator
func NewGetTableObjectsPaginator(client GetTableObjectsAPIClient, params *GetTableObjectsInput, optFns ...func(*GetTableObjectsPaginatorOptions)) *GetTableObjectsPaginator {
	if params == nil {
		params = &GetTableObjectsInput{}
	}

	options := GetTableObjectsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTableObjectsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTableObjectsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTableObjects page.
func (p *GetTableObjectsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTableObjectsOutput, error) {
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

	result, err := p.client.GetTableObjects(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTableObjects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "GetTableObjects",
	}
}
