// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/grafana/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the users and groups who have the Grafana Admin and Editor roles in this
// workspace. If you use this operation without specifying userId or groupId, the
// operation returns the roles of all users and groups. If you specify a userId or
// a groupId, only the roles for that user or group are returned. If you do this,
// you can specify only one userId or one groupId.
func (c *Client) ListPermissions(ctx context.Context, params *ListPermissionsInput, optFns ...func(*Options)) (*ListPermissionsOutput, error) {
	if params == nil {
		params = &ListPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPermissions", params, optFns, c.addOperationListPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionsInput struct {

	// The ID of the workspace to list permissions for. This parameter is required.
	//
	// This member is required.
	WorkspaceId *string

	// (Optional) Limits the results to only the group that matches this ID.
	GroupId *string

	// The maximum number of results to include in the response.
	MaxResults *int32

	// The token to use when requesting the next set of results. You received this
	// token from a previous ListPermissions operation.
	NextToken *string

	// (Optional) Limits the results to only the user that matches this ID.
	UserId *string

	// (Optional) If you specify SSO_USER, then only the permissions of Amazon Web
	// Services SSO users are returned. If you specify SSO_GROUP, only the permissions
	// of Amazon Web Services SSO groups are returned.
	UserType types.UserType

	noSmithyDocumentSerde
}

type ListPermissionsOutput struct {

	// The permissions returned by the operation.
	//
	// This member is required.
	Permissions []types.PermissionEntry

	// The token to use in a subsequent ListPermissions operation to return the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPermissions{}, middleware.After)
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
	if err = addOpListPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissions(options.Region), middleware.Before); err != nil {
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

// ListPermissionsAPIClient is a client that implements the ListPermissions
// operation.
type ListPermissionsAPIClient interface {
	ListPermissions(context.Context, *ListPermissionsInput, ...func(*Options)) (*ListPermissionsOutput, error)
}

var _ ListPermissionsAPIClient = (*Client)(nil)

// ListPermissionsPaginatorOptions is the paginator options for ListPermissions
type ListPermissionsPaginatorOptions struct {
	// The maximum number of results to include in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPermissionsPaginator is a paginator for ListPermissions
type ListPermissionsPaginator struct {
	options   ListPermissionsPaginatorOptions
	client    ListPermissionsAPIClient
	params    *ListPermissionsInput
	nextToken *string
	firstPage bool
}

// NewListPermissionsPaginator returns a new ListPermissionsPaginator
func NewListPermissionsPaginator(client ListPermissionsAPIClient, params *ListPermissionsInput, optFns ...func(*ListPermissionsPaginatorOptions)) *ListPermissionsPaginator {
	if params == nil {
		params = &ListPermissionsInput{}
	}

	options := ListPermissionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPermissionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPermissionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPermissions page.
func (p *ListPermissionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPermissionsOutput, error) {
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

	result, err := p.client.ListPermissions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "grafana",
		OperationName: "ListPermissions",
	}
}
