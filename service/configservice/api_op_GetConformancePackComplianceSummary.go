// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns compliance details for the conformance pack based on the cumulative
// compliance results of all the rules in that conformance pack.
func (c *Client) GetConformancePackComplianceSummary(ctx context.Context, params *GetConformancePackComplianceSummaryInput, optFns ...func(*Options)) (*GetConformancePackComplianceSummaryOutput, error) {
	if params == nil {
		params = &GetConformancePackComplianceSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConformancePackComplianceSummary", params, optFns, c.addOperationGetConformancePackComplianceSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConformancePackComplianceSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConformancePackComplianceSummaryInput struct {

	// Names of conformance packs.
	//
	// This member is required.
	ConformancePackNames []string

	// The maximum number of conformance packs returned on each page.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type GetConformancePackComplianceSummaryOutput struct {

	// A list of ConformancePackComplianceSummary objects.
	ConformancePackComplianceSummaryList []types.ConformancePackComplianceSummary

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConformancePackComplianceSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetConformancePackComplianceSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetConformancePackComplianceSummary{}, middleware.After)
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
	if err = addOpGetConformancePackComplianceSummaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConformancePackComplianceSummary(options.Region), middleware.Before); err != nil {
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

// GetConformancePackComplianceSummaryAPIClient is a client that implements the
// GetConformancePackComplianceSummary operation.
type GetConformancePackComplianceSummaryAPIClient interface {
	GetConformancePackComplianceSummary(context.Context, *GetConformancePackComplianceSummaryInput, ...func(*Options)) (*GetConformancePackComplianceSummaryOutput, error)
}

var _ GetConformancePackComplianceSummaryAPIClient = (*Client)(nil)

// GetConformancePackComplianceSummaryPaginatorOptions is the paginator options for
// GetConformancePackComplianceSummary
type GetConformancePackComplianceSummaryPaginatorOptions struct {
	// The maximum number of conformance packs returned on each page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetConformancePackComplianceSummaryPaginator is a paginator for
// GetConformancePackComplianceSummary
type GetConformancePackComplianceSummaryPaginator struct {
	options   GetConformancePackComplianceSummaryPaginatorOptions
	client    GetConformancePackComplianceSummaryAPIClient
	params    *GetConformancePackComplianceSummaryInput
	nextToken *string
	firstPage bool
}

// NewGetConformancePackComplianceSummaryPaginator returns a new
// GetConformancePackComplianceSummaryPaginator
func NewGetConformancePackComplianceSummaryPaginator(client GetConformancePackComplianceSummaryAPIClient, params *GetConformancePackComplianceSummaryInput, optFns ...func(*GetConformancePackComplianceSummaryPaginatorOptions)) *GetConformancePackComplianceSummaryPaginator {
	if params == nil {
		params = &GetConformancePackComplianceSummaryInput{}
	}

	options := GetConformancePackComplianceSummaryPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetConformancePackComplianceSummaryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetConformancePackComplianceSummaryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetConformancePackComplianceSummary page.
func (p *GetConformancePackComplianceSummaryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetConformancePackComplianceSummaryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.GetConformancePackComplianceSummary(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetConformancePackComplianceSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetConformancePackComplianceSummary",
	}
}
