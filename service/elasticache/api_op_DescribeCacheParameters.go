// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the detailed parameter list for a particular cache parameter group.
func (c *Client) DescribeCacheParameters(ctx context.Context, params *DescribeCacheParametersInput, optFns ...func(*Options)) (*DescribeCacheParametersOutput, error) {
	if params == nil {
		params = &DescribeCacheParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCacheParameters", params, optFns, c.addOperationDescribeCacheParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCacheParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeCacheParameters operation.
type DescribeCacheParametersInput struct {

	// The name of a specific cache parameter group to return details for.
	//
	// This member is required.
	CacheParameterGroupName *string

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	MaxRecords *int32

	// The parameter types to return. Valid values: user | system | engine-default
	Source *string

	noSmithyDocumentSerde
}

// Represents the output of a DescribeCacheParameters operation.
type DescribeCacheParametersOutput struct {

	// A list of parameters specific to a particular cache node type. Each element in
	// the list contains detailed information about one parameter.
	CacheNodeTypeSpecificParameters []types.CacheNodeTypeSpecificParameter

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string

	// A list of Parameter instances.
	Parameters []types.Parameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCacheParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCacheParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCacheParameters{}, middleware.After)
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
	if err = addOpDescribeCacheParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCacheParameters(options.Region), middleware.Before); err != nil {
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

// DescribeCacheParametersAPIClient is a client that implements the
// DescribeCacheParameters operation.
type DescribeCacheParametersAPIClient interface {
	DescribeCacheParameters(context.Context, *DescribeCacheParametersInput, ...func(*Options)) (*DescribeCacheParametersOutput, error)
}

var _ DescribeCacheParametersAPIClient = (*Client)(nil)

// DescribeCacheParametersPaginatorOptions is the paginator options for
// DescribeCacheParameters
type DescribeCacheParametersPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCacheParametersPaginator is a paginator for DescribeCacheParameters
type DescribeCacheParametersPaginator struct {
	options   DescribeCacheParametersPaginatorOptions
	client    DescribeCacheParametersAPIClient
	params    *DescribeCacheParametersInput
	nextToken *string
	firstPage bool
}

// NewDescribeCacheParametersPaginator returns a new
// DescribeCacheParametersPaginator
func NewDescribeCacheParametersPaginator(client DescribeCacheParametersAPIClient, params *DescribeCacheParametersInput, optFns ...func(*DescribeCacheParametersPaginatorOptions)) *DescribeCacheParametersPaginator {
	if params == nil {
		params = &DescribeCacheParametersInput{}
	}

	options := DescribeCacheParametersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCacheParametersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCacheParametersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeCacheParameters page.
func (p *DescribeCacheParametersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCacheParametersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeCacheParameters(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeCacheParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeCacheParameters",
	}
}
