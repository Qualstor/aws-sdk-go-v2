// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves attribute data along with aggregate utilization and savings data for a
// given time period. This doesn't support granular or grouped data (daily/monthly)
// in response. You can't retrieve data by dates in a single response similar to
// GetSavingsPlanUtilization, but you have the option to make multiple calls to
// GetSavingsPlanUtilizationDetails by providing individual dates. You can use
// GetDimensionValues in SAVINGS_PLANS to determine the possible dimension values.
// GetSavingsPlanUtilizationDetails internally groups data by SavingsPlansArn.
func (c *Client) GetSavingsPlansUtilizationDetails(ctx context.Context, params *GetSavingsPlansUtilizationDetailsInput, optFns ...func(*Options)) (*GetSavingsPlansUtilizationDetailsOutput, error) {
	if params == nil {
		params = &GetSavingsPlansUtilizationDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSavingsPlansUtilizationDetails", params, optFns, c.addOperationGetSavingsPlansUtilizationDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSavingsPlansUtilizationDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSavingsPlansUtilizationDetailsInput struct {

	// The time period that you want the usage and costs for. The Start date must be
	// within 13 months. The End date must be after the Start date, and before the
	// current date. Future dates can't be used as an End date.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// The data type.
	DataType []types.SavingsPlansDataType

	// Filters Savings Plans utilization coverage data for active Savings Plans
	// dimensions. You can filter data with the following dimensions:
	//
	// *
	// LINKED_ACCOUNT
	//
	// * SAVINGS_PLAN_ARN
	//
	// * REGION
	//
	// * PAYMENT_OPTION
	//
	// *
	// INSTANCE_TYPE_FAMILY
	//
	// GetSavingsPlansUtilizationDetails uses the same Expression
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension.
	Filter *types.Expression

	// The number of items to be returned in a response. The default is 20, with a
	// minimum value of 1.
	MaxResults int32

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string

	// The value that you want to sort the data by. The following values are supported
	// for Key:
	//
	// * UtilizationPercentage
	//
	// * TotalCommitment
	//
	// * UsedCommitment
	//
	// *
	// UnusedCommitment
	//
	// * NetSavings
	//
	// * AmortizedRecurringCommitment
	//
	// *
	// AmortizedUpfrontCommitment
	//
	// The supported values for SortOrder are ASCENDING and
	// DESCENDING.
	SortBy *types.SortDefinition

	noSmithyDocumentSerde
}

type GetSavingsPlansUtilizationDetailsOutput struct {

	// Retrieves a single daily or monthly Savings Plans utilization rate and details
	// for your account.
	//
	// This member is required.
	SavingsPlansUtilizationDetails []types.SavingsPlansUtilizationDetail

	// The time period of the request.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string

	// The total Savings Plans utilization, regardless of time period.
	Total *types.SavingsPlansUtilizationAggregates

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSavingsPlansUtilizationDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSavingsPlansUtilizationDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSavingsPlansUtilizationDetails{}, middleware.After)
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
	if err = addOpGetSavingsPlansUtilizationDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSavingsPlansUtilizationDetails(options.Region), middleware.Before); err != nil {
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

// GetSavingsPlansUtilizationDetailsAPIClient is a client that implements the
// GetSavingsPlansUtilizationDetails operation.
type GetSavingsPlansUtilizationDetailsAPIClient interface {
	GetSavingsPlansUtilizationDetails(context.Context, *GetSavingsPlansUtilizationDetailsInput, ...func(*Options)) (*GetSavingsPlansUtilizationDetailsOutput, error)
}

var _ GetSavingsPlansUtilizationDetailsAPIClient = (*Client)(nil)

// GetSavingsPlansUtilizationDetailsPaginatorOptions is the paginator options for
// GetSavingsPlansUtilizationDetails
type GetSavingsPlansUtilizationDetailsPaginatorOptions struct {
	// The number of items to be returned in a response. The default is 20, with a
	// minimum value of 1.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetSavingsPlansUtilizationDetailsPaginator is a paginator for
// GetSavingsPlansUtilizationDetails
type GetSavingsPlansUtilizationDetailsPaginator struct {
	options   GetSavingsPlansUtilizationDetailsPaginatorOptions
	client    GetSavingsPlansUtilizationDetailsAPIClient
	params    *GetSavingsPlansUtilizationDetailsInput
	nextToken *string
	firstPage bool
}

// NewGetSavingsPlansUtilizationDetailsPaginator returns a new
// GetSavingsPlansUtilizationDetailsPaginator
func NewGetSavingsPlansUtilizationDetailsPaginator(client GetSavingsPlansUtilizationDetailsAPIClient, params *GetSavingsPlansUtilizationDetailsInput, optFns ...func(*GetSavingsPlansUtilizationDetailsPaginatorOptions)) *GetSavingsPlansUtilizationDetailsPaginator {
	if params == nil {
		params = &GetSavingsPlansUtilizationDetailsInput{}
	}

	options := GetSavingsPlansUtilizationDetailsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetSavingsPlansUtilizationDetailsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetSavingsPlansUtilizationDetailsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetSavingsPlansUtilizationDetails page.
func (p *GetSavingsPlansUtilizationDetailsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetSavingsPlansUtilizationDetailsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetSavingsPlansUtilizationDetails(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetSavingsPlansUtilizationDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetSavingsPlansUtilizationDetails",
	}
}
