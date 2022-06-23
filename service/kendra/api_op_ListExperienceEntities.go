// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists users or groups in your Amazon Web Services SSO identity source that are
// granted access to your Amazon Kendra experience. You can create an Amazon Kendra
// experience such as a search application. For more information on creating a
// search application experience, see Building a search experience with no code
// (https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html).
func (c *Client) ListExperienceEntities(ctx context.Context, params *ListExperienceEntitiesInput, optFns ...func(*Options)) (*ListExperienceEntitiesOutput, error) {
	if params == nil {
		params = &ListExperienceEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExperienceEntities", params, optFns, c.addOperationListExperienceEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExperienceEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExperienceEntitiesInput struct {

	// The identifier of your Amazon Kendra experience.
	//
	// This member is required.
	Id *string

	// The identifier of the index for your Amazon Kendra experience.
	//
	// This member is required.
	IndexId *string

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Kendra returns a pagination token in the response. You can use
	// this pagination token to retrieve the next set of users or groups.
	NextToken *string

	noSmithyDocumentSerde
}

type ListExperienceEntitiesOutput struct {

	// If the response is truncated, Amazon Kendra returns this token, which you can
	// use in a later request to retrieve the next set of users or groups.
	NextToken *string

	// An array of summary information for one or more users or groups.
	SummaryItems []types.ExperienceEntitiesSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExperienceEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListExperienceEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListExperienceEntities{}, middleware.After)
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
	if err = addOpListExperienceEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExperienceEntities(options.Region), middleware.Before); err != nil {
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

// ListExperienceEntitiesAPIClient is a client that implements the
// ListExperienceEntities operation.
type ListExperienceEntitiesAPIClient interface {
	ListExperienceEntities(context.Context, *ListExperienceEntitiesInput, ...func(*Options)) (*ListExperienceEntitiesOutput, error)
}

var _ ListExperienceEntitiesAPIClient = (*Client)(nil)

// ListExperienceEntitiesPaginatorOptions is the paginator options for
// ListExperienceEntities
type ListExperienceEntitiesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExperienceEntitiesPaginator is a paginator for ListExperienceEntities
type ListExperienceEntitiesPaginator struct {
	options   ListExperienceEntitiesPaginatorOptions
	client    ListExperienceEntitiesAPIClient
	params    *ListExperienceEntitiesInput
	nextToken *string
	firstPage bool
}

// NewListExperienceEntitiesPaginator returns a new ListExperienceEntitiesPaginator
func NewListExperienceEntitiesPaginator(client ListExperienceEntitiesAPIClient, params *ListExperienceEntitiesInput, optFns ...func(*ListExperienceEntitiesPaginatorOptions)) *ListExperienceEntitiesPaginator {
	if params == nil {
		params = &ListExperienceEntitiesInput{}
	}

	options := ListExperienceEntitiesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExperienceEntitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExperienceEntitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExperienceEntities page.
func (p *ListExperienceEntitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExperienceEntitiesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListExperienceEntities(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExperienceEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "ListExperienceEntities",
	}
}
