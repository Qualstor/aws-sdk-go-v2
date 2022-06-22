// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns metadata about the images in a repository. Beginning with Docker version
// 1.9, the Docker client compresses image layers before pushing them to a V2
// Docker registry. The output of the docker images command shows the uncompressed
// image size, so it may return a larger image size than the image sizes returned
// by DescribeImages.
func (c *Client) DescribeImages(ctx context.Context, params *DescribeImagesInput, optFns ...func(*Options)) (*DescribeImagesOutput, error) {
	if params == nil {
		params = &DescribeImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImages", params, optFns, c.addOperationDescribeImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImagesInput struct {

	// The repository that contains the images to describe.
	//
	// This member is required.
	RepositoryName *string

	// The filter key and value with which to filter your DescribeImages results.
	Filter *types.DescribeImagesFilter

	// The list of image IDs for the requested repository.
	ImageIds []types.ImageIdentifier

	// The maximum number of repository results returned by DescribeImages in paginated
	// output. When this parameter is used, DescribeImages only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeImages
	// request with the returned nextToken value. This value can be between 1 and 1000.
	// If this parameter is not used, then DescribeImages returns up to 100 results and
	// a nextToken value, if applicable. This option cannot be used when you specify
	// images with imageIds.
	MaxResults *int32

	// The nextToken value returned from a previous paginated DescribeImages request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	// This option cannot be used when you specify images with imageIds.
	NextToken *string

	// The Amazon Web Services account ID associated with the registry that contains
	// the repository in which to describe images. If you do not specify a registry,
	// the default registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type DescribeImagesOutput struct {

	// A list of ImageDetail objects that contain data about the image.
	ImageDetails []types.ImageDetail

	// The nextToken value to include in a future DescribeImages request. When the
	// results of a DescribeImages request exceed maxResults, this value can be used to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImages{}, middleware.After)
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
	if err = addOpDescribeImagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImages(options.Region), middleware.Before); err != nil {
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

// DescribeImagesAPIClient is a client that implements the DescribeImages
// operation.
type DescribeImagesAPIClient interface {
	DescribeImages(context.Context, *DescribeImagesInput, ...func(*Options)) (*DescribeImagesOutput, error)
}

var _ DescribeImagesAPIClient = (*Client)(nil)

// DescribeImagesPaginatorOptions is the paginator options for DescribeImages
type DescribeImagesPaginatorOptions struct {
	// The maximum number of repository results returned by DescribeImages in paginated
	// output. When this parameter is used, DescribeImages only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeImages
	// request with the returned nextToken value. This value can be between 1 and 1000.
	// If this parameter is not used, then DescribeImages returns up to 100 results and
	// a nextToken value, if applicable. This option cannot be used when you specify
	// images with imageIds.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeImagesPaginator is a paginator for DescribeImages
type DescribeImagesPaginator struct {
	options   DescribeImagesPaginatorOptions
	client    DescribeImagesAPIClient
	params    *DescribeImagesInput
	nextToken *string
	firstPage bool
}

// NewDescribeImagesPaginator returns a new DescribeImagesPaginator
func NewDescribeImagesPaginator(client DescribeImagesAPIClient, params *DescribeImagesInput, optFns ...func(*DescribeImagesPaginatorOptions)) *DescribeImagesPaginator {
	if params == nil {
		params = &DescribeImagesInput{}
	}

	options := DescribeImagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeImagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeImagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeImages page.
func (p *DescribeImagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeImagesOutput, error) {
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

	result, err := p.client.DescribeImages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "DescribeImages",
	}
}
