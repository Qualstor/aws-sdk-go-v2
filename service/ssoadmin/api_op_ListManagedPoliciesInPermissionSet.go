// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the IAM managed policy that is attached to a specified permission set.
func (c *Client) ListManagedPoliciesInPermissionSet(ctx context.Context, params *ListManagedPoliciesInPermissionSetInput, optFns ...func(*Options)) (*ListManagedPoliciesInPermissionSetOutput, error) {
	if params == nil {
		params = &ListManagedPoliciesInPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListManagedPoliciesInPermissionSet", params, optFns, c.addOperationListManagedPoliciesInPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListManagedPoliciesInPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListManagedPoliciesInPermissionSetInput struct {

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and Amazon Web Services
	// Service Namespaces in the Amazon Web Services General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the PermissionSet whose managed policies will be listed.
	//
	// This member is required.
	PermissionSetArn *string

	// The maximum number of results to display for the PermissionSet.
	MaxResults *int32

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	noSmithyDocumentSerde
}

type ListManagedPoliciesInPermissionSetOutput struct {

	// The array of the AttachedManagedPolicy data type object.
	AttachedManagedPolicies []types.AttachedManagedPolicy

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListManagedPoliciesInPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListManagedPoliciesInPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListManagedPoliciesInPermissionSet{}, middleware.After)
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
	if err = addOpListManagedPoliciesInPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListManagedPoliciesInPermissionSet(options.Region), middleware.Before); err != nil {
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

// ListManagedPoliciesInPermissionSetAPIClient is a client that implements the
// ListManagedPoliciesInPermissionSet operation.
type ListManagedPoliciesInPermissionSetAPIClient interface {
	ListManagedPoliciesInPermissionSet(context.Context, *ListManagedPoliciesInPermissionSetInput, ...func(*Options)) (*ListManagedPoliciesInPermissionSetOutput, error)
}

var _ ListManagedPoliciesInPermissionSetAPIClient = (*Client)(nil)

// ListManagedPoliciesInPermissionSetPaginatorOptions is the paginator options for
// ListManagedPoliciesInPermissionSet
type ListManagedPoliciesInPermissionSetPaginatorOptions struct {
	// The maximum number of results to display for the PermissionSet.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListManagedPoliciesInPermissionSetPaginator is a paginator for
// ListManagedPoliciesInPermissionSet
type ListManagedPoliciesInPermissionSetPaginator struct {
	options   ListManagedPoliciesInPermissionSetPaginatorOptions
	client    ListManagedPoliciesInPermissionSetAPIClient
	params    *ListManagedPoliciesInPermissionSetInput
	nextToken *string
	firstPage bool
}

// NewListManagedPoliciesInPermissionSetPaginator returns a new
// ListManagedPoliciesInPermissionSetPaginator
func NewListManagedPoliciesInPermissionSetPaginator(client ListManagedPoliciesInPermissionSetAPIClient, params *ListManagedPoliciesInPermissionSetInput, optFns ...func(*ListManagedPoliciesInPermissionSetPaginatorOptions)) *ListManagedPoliciesInPermissionSetPaginator {
	if params == nil {
		params = &ListManagedPoliciesInPermissionSetInput{}
	}

	options := ListManagedPoliciesInPermissionSetPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListManagedPoliciesInPermissionSetPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListManagedPoliciesInPermissionSetPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListManagedPoliciesInPermissionSet page.
func (p *ListManagedPoliciesInPermissionSetPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListManagedPoliciesInPermissionSetOutput, error) {
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

	result, err := p.client.ListManagedPoliciesInPermissionSet(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListManagedPoliciesInPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "ListManagedPoliciesInPermissionSet",
	}
}
