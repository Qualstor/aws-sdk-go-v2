// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified static route associated with a VPN connection between an
// existing virtual private gateway and a VPN customer gateway. The static route
// allows traffic to be routed from the virtual private gateway to the VPN customer
// gateway.
func (c *Client) DeleteVpnConnectionRoute(ctx context.Context, params *DeleteVpnConnectionRouteInput, optFns ...func(*Options)) (*DeleteVpnConnectionRouteOutput, error) {
	if params == nil {
		params = &DeleteVpnConnectionRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVpnConnectionRoute", params, optFns, c.addOperationDeleteVpnConnectionRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVpnConnectionRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DeleteVpnConnectionRoute.
type DeleteVpnConnectionRouteInput struct {

	// The CIDR block associated with the local subnet of the customer network.
	//
	// This member is required.
	DestinationCidrBlock *string

	// The ID of the VPN connection.
	//
	// This member is required.
	VpnConnectionId *string

	noSmithyDocumentSerde
}

type DeleteVpnConnectionRouteOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteVpnConnectionRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteVpnConnectionRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteVpnConnectionRoute{}, middleware.After)
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
	if err = addOpDeleteVpnConnectionRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVpnConnectionRoute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVpnConnectionRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteVpnConnectionRoute",
	}
}
