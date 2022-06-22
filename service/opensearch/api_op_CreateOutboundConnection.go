// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new cross-cluster connection from a local OpenSearch domain to a
// remote OpenSearch domain.
func (c *Client) CreateOutboundConnection(ctx context.Context, params *CreateOutboundConnectionInput, optFns ...func(*Options)) (*CreateOutboundConnectionOutput, error) {
	if params == nil {
		params = &CreateOutboundConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOutboundConnection", params, optFns, c.addOperationCreateOutboundConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOutboundConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the CreateOutboundConnection operation.
type CreateOutboundConnectionInput struct {

	// The connection alias used used by the customer for this cross-cluster
	// connection.
	//
	// This member is required.
	ConnectionAlias *string

	// The AWSDomainInformation for the local OpenSearch domain.
	//
	// This member is required.
	LocalDomainInfo *types.DomainInformationContainer

	// The AWSDomainInformation for the remote OpenSearch domain.
	//
	// This member is required.
	RemoteDomainInfo *types.DomainInformationContainer

	noSmithyDocumentSerde
}

// The result of a CreateOutboundConnection request. Contains the details about the
// newly created cross-cluster connection.
type CreateOutboundConnectionOutput struct {

	// The connection alias provided during the create connection request.
	ConnectionAlias *string

	// The unique ID for the created outbound connection, which is used for subsequent
	// operations on the connection.
	ConnectionId *string

	// The OutboundConnectionStatus for the newly created connection.
	ConnectionStatus *types.OutboundConnectionStatus

	// The AWSDomainInformation for the local OpenSearch domain.
	LocalDomainInfo *types.DomainInformationContainer

	// The AWSDomainInformation for the remote OpenSearch domain.
	RemoteDomainInfo *types.DomainInformationContainer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOutboundConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateOutboundConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateOutboundConnection{}, middleware.After)
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
	if err = addOpCreateOutboundConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOutboundConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateOutboundConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "CreateOutboundConnection",
	}
}
