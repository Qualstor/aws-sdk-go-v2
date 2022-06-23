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

// Returns information about the specified HSM client certificate. If no
// certificate ID is specified, returns information about all the HSM certificates
// owned by your Amazon Web Services account. If you specify both tag keys and tag
// values in the same request, Amazon Redshift returns all HSM client certificates
// that match any combination of the specified keys and values. For example, if you
// have owner and environment for tag keys, and admin and test for tag values, all
// HSM client certificates that have any combination of those values are returned.
// If both tag keys and values are omitted from the request, HSM client
// certificates are returned regardless of whether they have tag keys or values
// associated with them.
func (c *Client) DescribeHsmClientCertificates(ctx context.Context, params *DescribeHsmClientCertificatesInput, optFns ...func(*Options)) (*DescribeHsmClientCertificatesOutput, error) {
	if params == nil {
		params = &DescribeHsmClientCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHsmClientCertificates", params, optFns, c.addOperationDescribeHsmClientCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHsmClientCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeHsmClientCertificatesInput struct {

	// The identifier of a specific HSM client certificate for which you want
	// information. If no identifier is specified, information is returned for all HSM
	// client certificates owned by your Amazon Web Services account.
	HsmClientCertificateIdentifier *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeHsmClientCertificates request
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

	// A tag key or keys for which you want to return all matching HSM client
	// certificates that are associated with the specified key or keys. For example,
	// suppose that you have HSM client certificates that are tagged with keys called
	// owner and environment. If you specify both of these tag keys in the request,
	// Amazon Redshift returns a response with the HSM client certificates that have
	// either or both of these tag keys associated with them.
	TagKeys []string

	// A tag value or values for which you want to return all matching HSM client
	// certificates that are associated with the specified tag value or values. For
	// example, suppose that you have HSM client certificates that are tagged with
	// values called admin and test. If you specify both of these tag values in the
	// request, Amazon Redshift returns a response with the HSM client certificates
	// that have either or both of these tag values associated with them.
	TagValues []string

	noSmithyDocumentSerde
}

//
type DescribeHsmClientCertificatesOutput struct {

	// A list of the identifiers for one or more HSM client certificates used by Amazon
	// Redshift clusters to store and retrieve database encryption keys in an HSM.
	HsmClientCertificates []types.HsmClientCertificate

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHsmClientCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeHsmClientCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeHsmClientCertificates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHsmClientCertificates(options.Region), middleware.Before); err != nil {
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

// DescribeHsmClientCertificatesAPIClient is a client that implements the
// DescribeHsmClientCertificates operation.
type DescribeHsmClientCertificatesAPIClient interface {
	DescribeHsmClientCertificates(context.Context, *DescribeHsmClientCertificatesInput, ...func(*Options)) (*DescribeHsmClientCertificatesOutput, error)
}

var _ DescribeHsmClientCertificatesAPIClient = (*Client)(nil)

// DescribeHsmClientCertificatesPaginatorOptions is the paginator options for
// DescribeHsmClientCertificates
type DescribeHsmClientCertificatesPaginatorOptions struct {
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

// DescribeHsmClientCertificatesPaginator is a paginator for
// DescribeHsmClientCertificates
type DescribeHsmClientCertificatesPaginator struct {
	options   DescribeHsmClientCertificatesPaginatorOptions
	client    DescribeHsmClientCertificatesAPIClient
	params    *DescribeHsmClientCertificatesInput
	nextToken *string
	firstPage bool
}

// NewDescribeHsmClientCertificatesPaginator returns a new
// DescribeHsmClientCertificatesPaginator
func NewDescribeHsmClientCertificatesPaginator(client DescribeHsmClientCertificatesAPIClient, params *DescribeHsmClientCertificatesInput, optFns ...func(*DescribeHsmClientCertificatesPaginatorOptions)) *DescribeHsmClientCertificatesPaginator {
	if params == nil {
		params = &DescribeHsmClientCertificatesInput{}
	}

	options := DescribeHsmClientCertificatesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeHsmClientCertificatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeHsmClientCertificatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeHsmClientCertificates page.
func (p *DescribeHsmClientCertificatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeHsmClientCertificatesOutput, error) {
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

	result, err := p.client.DescribeHsmClientCertificates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeHsmClientCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeHsmClientCertificates",
	}
}
