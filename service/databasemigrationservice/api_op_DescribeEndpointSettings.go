// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the possible endpoint settings available when you
// create an endpoint for a specific database engine.
func (c *Client) DescribeEndpointSettings(ctx context.Context, params *DescribeEndpointSettingsInput, optFns ...func(*Options)) (*DescribeEndpointSettingsOutput, error) {
	if params == nil {
		params = &DescribeEndpointSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpointSettings", params, optFns, c.addOperationDescribeEndpointSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointSettingsInput struct {

	// The databse engine used for your source or target endpoint.
	//
	// This member is required.
	EngineName *string

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeEndpointSettingsOutput struct {

	// Descriptions of the endpoint settings available for your source or target
	// database engine.
	EndpointSettings []types.EndpointSetting

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEndpointSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEndpointSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEndpointSettings{}, middleware.After)
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
	if err = addOpDescribeEndpointSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpointSettings(options.Region), middleware.Before); err != nil {
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

// DescribeEndpointSettingsAPIClient is a client that implements the
// DescribeEndpointSettings operation.
type DescribeEndpointSettingsAPIClient interface {
	DescribeEndpointSettings(context.Context, *DescribeEndpointSettingsInput, ...func(*Options)) (*DescribeEndpointSettingsOutput, error)
}

var _ DescribeEndpointSettingsAPIClient = (*Client)(nil)

// DescribeEndpointSettingsPaginatorOptions is the paginator options for
// DescribeEndpointSettings
type DescribeEndpointSettingsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEndpointSettingsPaginator is a paginator for DescribeEndpointSettings
type DescribeEndpointSettingsPaginator struct {
	options   DescribeEndpointSettingsPaginatorOptions
	client    DescribeEndpointSettingsAPIClient
	params    *DescribeEndpointSettingsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEndpointSettingsPaginator returns a new
// DescribeEndpointSettingsPaginator
func NewDescribeEndpointSettingsPaginator(client DescribeEndpointSettingsAPIClient, params *DescribeEndpointSettingsInput, optFns ...func(*DescribeEndpointSettingsPaginatorOptions)) *DescribeEndpointSettingsPaginator {
	if params == nil {
		params = &DescribeEndpointSettingsInput{}
	}

	options := DescribeEndpointSettingsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEndpointSettingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEndpointSettingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEndpointSettings page.
func (p *DescribeEndpointSettingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEndpointSettingsOutput, error) {
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

	result, err := p.client.DescribeEndpointSettings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEndpointSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeEndpointSettings",
	}
}
