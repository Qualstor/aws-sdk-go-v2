// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the customer gateway or the target gateway of an Amazon Web Services
// Site-to-Site VPN connection. To modify the target gateway, the following
// migration options are available:
//
// * An existing virtual private gateway to a new
// virtual private gateway
//
// * An existing virtual private gateway to a transit
// gateway
//
// * An existing transit gateway to a new transit gateway
//
// * An existing
// transit gateway to a virtual private gateway
//
// Before you perform the migration
// to the new gateway, you must configure the new gateway. Use CreateVpnGateway to
// create a virtual private gateway, or CreateTransitGateway to create a transit
// gateway. This step is required when you migrate from a virtual private gateway
// with static routes to a transit gateway. You must delete the static routes
// before you migrate to the new gateway. Keep a copy of the static route before
// you delete it. You will need to add back these routes to the transit gateway
// after the VPN connection migration is complete. After you migrate to the new
// gateway, you might need to modify your VPC route table. Use CreateRoute and
// DeleteRoute to make the changes described in Update VPC route tables
// (https://docs.aws.amazon.com/vpn/latest/s2svpn/modify-vpn-target.html#step-update-routing)
// in the Amazon Web Services Site-to-Site VPN User Guide. When the new gateway is
// a transit gateway, modify the transit gateway route table to allow traffic
// between the VPC and the Amazon Web Services Site-to-Site VPN connection. Use
// CreateTransitGatewayRoute to add the routes. If you deleted VPN static routes,
// you must add the static routes to the transit gateway route table. After you
// perform this operation, the VPN endpoint's IP addresses on the Amazon Web
// Services side and the tunnel options remain intact. Your Amazon Web Services
// Site-to-Site VPN connection will be temporarily unavailable for a brief period
// while we provision the new endpoints.
func (c *Client) ModifyVpnConnection(ctx context.Context, params *ModifyVpnConnectionInput, optFns ...func(*Options)) (*ModifyVpnConnectionOutput, error) {
	if params == nil {
		params = &ModifyVpnConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpnConnection", params, optFns, c.addOperationModifyVpnConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpnConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpnConnectionInput struct {

	// The ID of the VPN connection.
	//
	// This member is required.
	VpnConnectionId *string

	// The ID of the customer gateway at your end of the VPN connection.
	CustomerGatewayId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the transit gateway.
	TransitGatewayId *string

	// The ID of the virtual private gateway at the Amazon Web Services side of the VPN
	// connection.
	VpnGatewayId *string

	noSmithyDocumentSerde
}

type ModifyVpnConnectionOutput struct {

	// Describes a VPN connection.
	VpnConnection *types.VpnConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVpnConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpnConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpnConnection{}, middleware.After)
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
	if err = addOpModifyVpnConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpnConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyVpnConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpnConnection",
	}
}
