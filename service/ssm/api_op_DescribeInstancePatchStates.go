// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the high-level patch state of one or more managed nodes.
func (c *Client) DescribeInstancePatchStates(ctx context.Context, params *DescribeInstancePatchStatesInput, optFns ...func(*Options)) (*DescribeInstancePatchStatesOutput, error) {
	if params == nil {
		params = &DescribeInstancePatchStatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstancePatchStates", params, optFns, c.addOperationDescribeInstancePatchStatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstancePatchStatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstancePatchStatesInput struct {

	// The ID of the managed node for which patch state information should be
	// retrieved.
	//
	// This member is required.
	InstanceIds []string

	// The maximum number of managed nodes to return (per page).
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeInstancePatchStatesOutput struct {

	// The high-level patch state for the requested managed nodes.
	InstancePatchStates []types.InstancePatchState

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstancePatchStatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInstancePatchStates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInstancePatchStates{}, middleware.After)
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
	if err = addOpDescribeInstancePatchStatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstancePatchStates(options.Region), middleware.Before); err != nil {
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

// DescribeInstancePatchStatesAPIClient is a client that implements the
// DescribeInstancePatchStates operation.
type DescribeInstancePatchStatesAPIClient interface {
	DescribeInstancePatchStates(context.Context, *DescribeInstancePatchStatesInput, ...func(*Options)) (*DescribeInstancePatchStatesOutput, error)
}

var _ DescribeInstancePatchStatesAPIClient = (*Client)(nil)

// DescribeInstancePatchStatesPaginatorOptions is the paginator options for
// DescribeInstancePatchStates
type DescribeInstancePatchStatesPaginatorOptions struct {
	// The maximum number of managed nodes to return (per page).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstancePatchStatesPaginator is a paginator for
// DescribeInstancePatchStates
type DescribeInstancePatchStatesPaginator struct {
	options   DescribeInstancePatchStatesPaginatorOptions
	client    DescribeInstancePatchStatesAPIClient
	params    *DescribeInstancePatchStatesInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstancePatchStatesPaginator returns a new
// DescribeInstancePatchStatesPaginator
func NewDescribeInstancePatchStatesPaginator(client DescribeInstancePatchStatesAPIClient, params *DescribeInstancePatchStatesInput, optFns ...func(*DescribeInstancePatchStatesPaginatorOptions)) *DescribeInstancePatchStatesPaginator {
	if params == nil {
		params = &DescribeInstancePatchStatesInput{}
	}

	options := DescribeInstancePatchStatesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstancePatchStatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstancePatchStatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInstancePatchStates page.
func (p *DescribeInstancePatchStatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstancePatchStatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeInstancePatchStates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeInstancePatchStates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeInstancePatchStates",
	}
}
