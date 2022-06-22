// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns configuration details about all the features in the specified project.
func (c *Client) ListFeatures(ctx context.Context, params *ListFeaturesInput, optFns ...func(*Options)) (*ListFeaturesOutput, error) {
	if params == nil {
		params = &ListFeaturesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFeatures", params, optFns, c.addOperationListFeaturesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFeaturesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFeaturesInput struct {

	// The name or ARN of the project to return the feature list from.
	//
	// This member is required.
	Project *string

	// The maximum number of results to include in the response.
	MaxResults *int32

	// The token to use when requesting the next set of results. You received this
	// token from a previous ListFeatures operation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFeaturesOutput struct {

	// An array of structures that contain the configuration details of the features in
	// the specified project.
	Features []types.FeatureSummary

	// The token to use in a subsequent ListFeatures operation to return the next set
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFeaturesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFeatures{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFeatures{}, middleware.After)
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
	if err = addOpListFeaturesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFeatures(options.Region), middleware.Before); err != nil {
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

// ListFeaturesAPIClient is a client that implements the ListFeatures operation.
type ListFeaturesAPIClient interface {
	ListFeatures(context.Context, *ListFeaturesInput, ...func(*Options)) (*ListFeaturesOutput, error)
}

var _ ListFeaturesAPIClient = (*Client)(nil)

// ListFeaturesPaginatorOptions is the paginator options for ListFeatures
type ListFeaturesPaginatorOptions struct {
	// The maximum number of results to include in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFeaturesPaginator is a paginator for ListFeatures
type ListFeaturesPaginator struct {
	options   ListFeaturesPaginatorOptions
	client    ListFeaturesAPIClient
	params    *ListFeaturesInput
	nextToken *string
	firstPage bool
}

// NewListFeaturesPaginator returns a new ListFeaturesPaginator
func NewListFeaturesPaginator(client ListFeaturesAPIClient, params *ListFeaturesInput, optFns ...func(*ListFeaturesPaginatorOptions)) *ListFeaturesPaginator {
	if params == nil {
		params = &ListFeaturesInput{}
	}

	options := ListFeaturesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFeaturesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFeaturesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFeatures page.
func (p *ListFeaturesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFeaturesOutput, error) {
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

	result, err := p.client.ListFeatures(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFeatures(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "ListFeatures",
	}
}
