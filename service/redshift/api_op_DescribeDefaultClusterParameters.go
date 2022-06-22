// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of parameter settings for the specified parameter group family.
// For more information about parameters and parameter groups, go to Amazon
// Redshift Parameter Groups
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) DescribeDefaultClusterParameters(ctx context.Context, params *DescribeDefaultClusterParametersInput, optFns ...func(*Options)) (*DescribeDefaultClusterParametersOutput, error) {
	if params == nil {
		params = &DescribeDefaultClusterParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDefaultClusterParameters", params, optFns, c.addOperationDescribeDefaultClusterParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDefaultClusterParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDefaultClusterParametersInput struct {

	// The name of the cluster parameter group family.
	//
	// This member is required.
	ParameterGroupFamily *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeDefaultClusterParameters request
	// exceed the value specified in MaxRecords, Amazon Web Services returns a value in
	// the Marker field of the response. You can retrieve the next set of response
	// records by providing the returned marker value in the Marker parameter and
	// retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeDefaultClusterParametersOutput struct {

	// Describes the default cluster parameters for a parameter group family.
	DefaultClusterParameters *types.DefaultClusterParameters

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDefaultClusterParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDefaultClusterParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDefaultClusterParameters{}, middleware.After)
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
	if err = addOpDescribeDefaultClusterParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDefaultClusterParameters(options.Region), middleware.Before); err != nil {
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

// DescribeDefaultClusterParametersAPIClient is a client that implements the
// DescribeDefaultClusterParameters operation.
type DescribeDefaultClusterParametersAPIClient interface {
	DescribeDefaultClusterParameters(context.Context, *DescribeDefaultClusterParametersInput, ...func(*Options)) (*DescribeDefaultClusterParametersOutput, error)
}

var _ DescribeDefaultClusterParametersAPIClient = (*Client)(nil)

// DescribeDefaultClusterParametersPaginatorOptions is the paginator options for
// DescribeDefaultClusterParameters
type DescribeDefaultClusterParametersPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDefaultClusterParametersPaginator is a paginator for
// DescribeDefaultClusterParameters
type DescribeDefaultClusterParametersPaginator struct {
	options   DescribeDefaultClusterParametersPaginatorOptions
	client    DescribeDefaultClusterParametersAPIClient
	params    *DescribeDefaultClusterParametersInput
	nextToken *string
	firstPage bool
}

// NewDescribeDefaultClusterParametersPaginator returns a new
// DescribeDefaultClusterParametersPaginator
func NewDescribeDefaultClusterParametersPaginator(client DescribeDefaultClusterParametersAPIClient, params *DescribeDefaultClusterParametersInput, optFns ...func(*DescribeDefaultClusterParametersPaginatorOptions)) *DescribeDefaultClusterParametersPaginator {
	if params == nil {
		params = &DescribeDefaultClusterParametersInput{}
	}

	options := DescribeDefaultClusterParametersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDefaultClusterParametersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDefaultClusterParametersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDefaultClusterParameters page.
func (p *DescribeDefaultClusterParametersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDefaultClusterParametersOutput, error) {
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

	result, err := p.client.DescribeDefaultClusterParameters(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = nil
	if result.DefaultClusterParameters != nil {
		p.nextToken = result.DefaultClusterParameters.Marker
	}

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDefaultClusterParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeDefaultClusterParameters",
	}
}
