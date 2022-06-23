// Code generated by smithy-go-codegen DO NOT EDIT.

package amplifyuibuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/amplifyuibuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports component configurations to code that is ready to integrate into an
// Amplify app.
func (c *Client) ExportComponents(ctx context.Context, params *ExportComponentsInput, optFns ...func(*Options)) (*ExportComponentsOutput, error) {
	if params == nil {
		params = &ExportComponentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportComponents", params, optFns, c.addOperationExportComponentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportComponentsInput struct {

	// The unique ID of the Amplify app to export components to.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment that is a part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ExportComponentsOutput struct {

	// Represents the configuration of the exported components.
	//
	// This member is required.
	Entities []types.Component

	// The pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportComponentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportComponents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportComponents{}, middleware.After)
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
	if err = addOpExportComponentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportComponents(options.Region), middleware.Before); err != nil {
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

// ExportComponentsAPIClient is a client that implements the ExportComponents
// operation.
type ExportComponentsAPIClient interface {
	ExportComponents(context.Context, *ExportComponentsInput, ...func(*Options)) (*ExportComponentsOutput, error)
}

var _ ExportComponentsAPIClient = (*Client)(nil)

// ExportComponentsPaginatorOptions is the paginator options for ExportComponents
type ExportComponentsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ExportComponentsPaginator is a paginator for ExportComponents
type ExportComponentsPaginator struct {
	options   ExportComponentsPaginatorOptions
	client    ExportComponentsAPIClient
	params    *ExportComponentsInput
	nextToken *string
	firstPage bool
}

// NewExportComponentsPaginator returns a new ExportComponentsPaginator
func NewExportComponentsPaginator(client ExportComponentsAPIClient, params *ExportComponentsInput, optFns ...func(*ExportComponentsPaginatorOptions)) *ExportComponentsPaginator {
	if params == nil {
		params = &ExportComponentsInput{}
	}

	options := ExportComponentsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ExportComponentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ExportComponentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ExportComponents page.
func (p *ExportComponentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ExportComponentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ExportComponents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opExportComponents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplifyuibuilder",
		OperationName: "ExportComponents",
	}
}
