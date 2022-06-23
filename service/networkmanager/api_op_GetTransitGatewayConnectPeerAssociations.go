// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about one or more of your transit gateway Connect peer
// associations in a global network.
func (c *Client) GetTransitGatewayConnectPeerAssociations(ctx context.Context, params *GetTransitGatewayConnectPeerAssociationsInput, optFns ...func(*Options)) (*GetTransitGatewayConnectPeerAssociationsOutput, error) {
	if params == nil {
		params = &GetTransitGatewayConnectPeerAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransitGatewayConnectPeerAssociations", params, optFns, c.addOperationGetTransitGatewayConnectPeerAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransitGatewayConnectPeerAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayConnectPeerAssociationsInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more transit gateway Connect peer Amazon Resource Names (ARNs).
	TransitGatewayConnectPeerArns []string

	noSmithyDocumentSerde
}

type GetTransitGatewayConnectPeerAssociationsOutput struct {

	// The token to use for the next page of results.
	NextToken *string

	// Information about the transit gateway Connect peer associations.
	TransitGatewayConnectPeerAssociations []types.TransitGatewayConnectPeerAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTransitGatewayConnectPeerAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTransitGatewayConnectPeerAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTransitGatewayConnectPeerAssociations{}, middleware.After)
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
	if err = addOpGetTransitGatewayConnectPeerAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayConnectPeerAssociations(options.Region), middleware.Before); err != nil {
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

// GetTransitGatewayConnectPeerAssociationsAPIClient is a client that implements
// the GetTransitGatewayConnectPeerAssociations operation.
type GetTransitGatewayConnectPeerAssociationsAPIClient interface {
	GetTransitGatewayConnectPeerAssociations(context.Context, *GetTransitGatewayConnectPeerAssociationsInput, ...func(*Options)) (*GetTransitGatewayConnectPeerAssociationsOutput, error)
}

var _ GetTransitGatewayConnectPeerAssociationsAPIClient = (*Client)(nil)

// GetTransitGatewayConnectPeerAssociationsPaginatorOptions is the paginator
// options for GetTransitGatewayConnectPeerAssociations
type GetTransitGatewayConnectPeerAssociationsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTransitGatewayConnectPeerAssociationsPaginator is a paginator for
// GetTransitGatewayConnectPeerAssociations
type GetTransitGatewayConnectPeerAssociationsPaginator struct {
	options   GetTransitGatewayConnectPeerAssociationsPaginatorOptions
	client    GetTransitGatewayConnectPeerAssociationsAPIClient
	params    *GetTransitGatewayConnectPeerAssociationsInput
	nextToken *string
	firstPage bool
}

// NewGetTransitGatewayConnectPeerAssociationsPaginator returns a new
// GetTransitGatewayConnectPeerAssociationsPaginator
func NewGetTransitGatewayConnectPeerAssociationsPaginator(client GetTransitGatewayConnectPeerAssociationsAPIClient, params *GetTransitGatewayConnectPeerAssociationsInput, optFns ...func(*GetTransitGatewayConnectPeerAssociationsPaginatorOptions)) *GetTransitGatewayConnectPeerAssociationsPaginator {
	if params == nil {
		params = &GetTransitGatewayConnectPeerAssociationsInput{}
	}

	options := GetTransitGatewayConnectPeerAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTransitGatewayConnectPeerAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTransitGatewayConnectPeerAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTransitGatewayConnectPeerAssociations page.
func (p *GetTransitGatewayConnectPeerAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTransitGatewayConnectPeerAssociationsOutput, error) {
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

	result, err := p.client.GetTransitGatewayConnectPeerAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTransitGatewayConnectPeerAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "GetTransitGatewayConnectPeerAssociations",
	}
}
