// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/Qualstor/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/Qualstor/aws-sdk-go-v2/service/glacier/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation lists all vaults owned by the calling user's account. The list
// returned in the response is ASCII-sorted by vault name. By default, this
// operation returns up to 10 items. If there are more vaults to list, the response
// marker field contains the vault Amazon Resource Name (ARN) at which to continue
// the list with a new List Vaults request; otherwise, the marker field is null. To
// return a list of vaults that begins at a specific vault, set the marker request
// parameter to the vault ARN you obtained from a previous List Vaults request. You
// can also limit the number of vaults returned in the response by specifying the
// limit parameter in the request. An AWS account has full permission to perform
// all operations (actions). However, AWS Identity and Access Management (IAM)
// users don't have any permissions by default. You must grant them explicit
// permission to perform specific actions. For more information, see Access Control
// Using AWS Identity and Access Management (IAM)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
// For conceptual information and underlying REST API, see Retrieving Vault
// Metadata in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/retrieving-vault-info.html)
// and List Vaults
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vaults-get.html) in
// the Amazon Glacier Developer Guide.
func (c *Client) ListVaults(ctx context.Context, params *ListVaultsInput, optFns ...func(*Options)) (*ListVaultsOutput, error) {
	if params == nil {
		params = &ListVaultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVaults", params, optFns, c.addOperationListVaultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVaultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to retrieve the vault list owned by the calling user's account.
// The list provides metadata information for each vault.
type ListVaultsInput struct {

	// The AccountId value is the AWS account ID. This value must match the AWS account
	// ID associated with the credentials used to sign the request. You can either
	// specify an AWS account ID or optionally a single '-' (hyphen), in which case
	// Amazon Glacier uses the AWS account ID associated with the credentials used to
	// sign the request. If you specify your account ID, do not include any hyphens
	// ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The maximum number of vaults to be returned. The default limit is 10. The number
	// of vaults returned might be fewer than the specified limit, but the number of
	// returned vaults never exceeds the limit.
	Limit *int32

	// A string used for pagination. The marker specifies the vault ARN after which the
	// listing of vaults should begin.
	Marker *string

	noSmithyDocumentSerde
}

// Contains the Amazon S3 Glacier response to your request.
type ListVaultsOutput struct {

	// The vault ARN at which to continue pagination of the results. You use the marker
	// in another List Vaults request to obtain more vaults in the list.
	Marker *string

	// List of vaults.
	VaultList []types.DescribeVaultOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVaultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVaults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVaults{}, middleware.After)
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
	if err = addOpListVaultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVaults(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListVaultsAPIClient is a client that implements the ListVaults operation.
type ListVaultsAPIClient interface {
	ListVaults(context.Context, *ListVaultsInput, ...func(*Options)) (*ListVaultsOutput, error)
}

var _ ListVaultsAPIClient = (*Client)(nil)

// ListVaultsPaginatorOptions is the paginator options for ListVaults
type ListVaultsPaginatorOptions struct {
	// The maximum number of vaults to be returned. The default limit is 10. The number
	// of vaults returned might be fewer than the specified limit, but the number of
	// returned vaults never exceeds the limit.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVaultsPaginator is a paginator for ListVaults
type ListVaultsPaginator struct {
	options   ListVaultsPaginatorOptions
	client    ListVaultsAPIClient
	params    *ListVaultsInput
	nextToken *string
	firstPage bool
}

// NewListVaultsPaginator returns a new ListVaultsPaginator
func NewListVaultsPaginator(client ListVaultsAPIClient, params *ListVaultsInput, optFns ...func(*ListVaultsPaginatorOptions)) *ListVaultsPaginator {
	if params == nil {
		params = &ListVaultsInput{}
	}

	options := ListVaultsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVaultsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVaultsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVaults page.
func (p *ListVaultsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVaultsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListVaults(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListVaults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "ListVaults",
	}
}
