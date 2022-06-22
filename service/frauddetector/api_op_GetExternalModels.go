// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the details for one or more Amazon SageMaker models that have been imported
// into the service. This is a paginated API. If you provide a null maxResults,
// this actions retrieves a maximum of 10 records per page. If you provide a
// maxResults, the value must be between 5 and 10. To get the next page results,
// provide the pagination token from the GetExternalModelsResult as part of your
// request. A null pagination token fetches the records from the beginning.
func (c *Client) GetExternalModels(ctx context.Context, params *GetExternalModelsInput, optFns ...func(*Options)) (*GetExternalModelsOutput, error) {
	if params == nil {
		params = &GetExternalModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetExternalModels", params, optFns, c.addOperationGetExternalModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetExternalModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExternalModelsInput struct {

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The Amazon SageMaker model endpoint.
	ModelEndpoint *string

	// The next page token for the request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetExternalModelsOutput struct {

	// Gets the Amazon SageMaker models.
	ExternalModels []types.ExternalModel

	// The next page token to be used in subsequent requests.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetExternalModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetExternalModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetExternalModels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetExternalModels(options.Region), middleware.Before); err != nil {
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

// GetExternalModelsAPIClient is a client that implements the GetExternalModels
// operation.
type GetExternalModelsAPIClient interface {
	GetExternalModels(context.Context, *GetExternalModelsInput, ...func(*Options)) (*GetExternalModelsOutput, error)
}

var _ GetExternalModelsAPIClient = (*Client)(nil)

// GetExternalModelsPaginatorOptions is the paginator options for GetExternalModels
type GetExternalModelsPaginatorOptions struct {
	// The maximum number of objects to return for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetExternalModelsPaginator is a paginator for GetExternalModels
type GetExternalModelsPaginator struct {
	options   GetExternalModelsPaginatorOptions
	client    GetExternalModelsAPIClient
	params    *GetExternalModelsInput
	nextToken *string
	firstPage bool
}

// NewGetExternalModelsPaginator returns a new GetExternalModelsPaginator
func NewGetExternalModelsPaginator(client GetExternalModelsAPIClient, params *GetExternalModelsInput, optFns ...func(*GetExternalModelsPaginatorOptions)) *GetExternalModelsPaginator {
	if params == nil {
		params = &GetExternalModelsInput{}
	}

	options := GetExternalModelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetExternalModelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetExternalModelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetExternalModels page.
func (p *GetExternalModelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetExternalModelsOutput, error) {
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

	result, err := p.client.GetExternalModels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetExternalModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetExternalModels",
	}
}
