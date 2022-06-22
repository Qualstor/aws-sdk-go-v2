// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Disassociates a connection from a link aggregation group (LAG). The connection
// is interrupted and re-established as a standalone connection (the connection is
// not deleted; to delete the connection, use the DeleteConnection request). If the
// LAG has associated virtual interfaces or hosted connections, they remain
// associated with the LAG. A disassociated connection owned by an Direct Connect
// Partner is automatically converted to an interconnect. If disassociating the
// connection would cause the LAG to fall below its setting for minimum number of
// operational connections, the request fails, except when it's the last member of
// the LAG. If all connections are disassociated, the LAG continues to exist as an
// empty LAG with no physical connections.
func (c *Client) DisassociateConnectionFromLag(ctx context.Context, params *DisassociateConnectionFromLagInput, optFns ...func(*Options)) (*DisassociateConnectionFromLagOutput, error) {
	if params == nil {
		params = &DisassociateConnectionFromLagInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateConnectionFromLag", params, optFns, c.addOperationDisassociateConnectionFromLagMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateConnectionFromLagOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateConnectionFromLagInput struct {

	// The ID of the connection.
	//
	// This member is required.
	ConnectionId *string

	// The ID of the LAG.
	//
	// This member is required.
	LagId *string

	noSmithyDocumentSerde
}

// Information about an Direct Connect connection.
type DisassociateConnectionFromLagOutput struct {

	// The Direct Connect endpoint on which the physical connection terminates.
	//
	// Deprecated: This member has been deprecated.
	AwsDevice *string

	// The Direct Connect endpoint that terminates the physical connection.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string

	// The bandwidth of the connection.
	Bandwidth *string

	// The ID of the connection.
	ConnectionId *string

	// The name of the connection.
	ConnectionName *string

	// The state of the connection. The following are the possible values:
	//
	// * ordering:
	// The initial state of a hosted connection provisioned on an interconnect. The
	// connection stays in the ordering state until the owner of the hosted connection
	// confirms or declines the connection order.
	//
	// * requested: The initial state of a
	// standard connection. The connection stays in the requested state until the
	// Letter of Authorization (LOA) is sent to the customer.
	//
	// * pending: The
	// connection has been approved and is being initialized.
	//
	// * available: The network
	// link is up and the connection is ready for use.
	//
	// * down: The network link is
	// down.
	//
	// * deleting: The connection is being deleted.
	//
	// * deleted: The connection
	// has been deleted.
	//
	// * rejected: A hosted connection in the ordering state enters
	// the rejected state if it is deleted by the customer.
	//
	// * unknown: The state of
	// the connection is not available.
	ConnectionState types.ConnectionState

	// The MAC Security (MACsec) connection encryption mode. The valid values are
	// no_encrypt, should_encrypt, and must_encrypt.
	EncryptionMode *string

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time

	// The location of the connection.
	Location *string

	// Indicates whether the connection supports MAC Security (MACsec).
	MacSecCapable *bool

	// The MAC Security (MACsec) security keys associated with the connection.
	MacSecKeys []types.MacSecKey

	// The ID of the Amazon Web Services account that owns the connection.
	OwnerAccount *string

	// The name of the Direct Connect service provider associated with the connection.
	PartnerName *string

	// The MAC Security (MACsec) port link status of the connection. The valid values
	// are Encryption Up, which means that there is an active Connection Key Name, or
	// Encryption Down.
	PortEncryptionStatus *string

	// The name of the service provider associated with the connection.
	ProviderName *string

	// The Amazon Web Services Region where the connection is located.
	Region *string

	// The tags associated with the connection.
	Tags []types.Tag

	// The ID of the VLAN.
	Vlan int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateConnectionFromLagMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateConnectionFromLag{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateConnectionFromLag{}, middleware.After)
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
	if err = addOpDisassociateConnectionFromLagValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateConnectionFromLag(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateConnectionFromLag(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DisassociateConnectionFromLag",
	}
}
