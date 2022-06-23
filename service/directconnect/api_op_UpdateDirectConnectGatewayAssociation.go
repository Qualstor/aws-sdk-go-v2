// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified attributes of the Direct Connect gateway association. Add
// or remove prefixes from the association.
func (c *Client) UpdateDirectConnectGatewayAssociation(ctx context.Context, params *UpdateDirectConnectGatewayAssociationInput, optFns ...func(*Options)) (*UpdateDirectConnectGatewayAssociationOutput, error) {
	if params == nil {
		params = &UpdateDirectConnectGatewayAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDirectConnectGatewayAssociation", params, optFns, c.addOperationUpdateDirectConnectGatewayAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDirectConnectGatewayAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDirectConnectGatewayAssociationInput struct {

	// The Amazon VPC prefixes to advertise to the Direct Connect gateway.
	AddAllowedPrefixesToDirectConnectGateway []types.RouteFilterPrefix

	// The ID of the Direct Connect gateway association.
	AssociationId *string

	// The Amazon VPC prefixes to no longer advertise to the Direct Connect gateway.
	RemoveAllowedPrefixesToDirectConnectGateway []types.RouteFilterPrefix

	noSmithyDocumentSerde
}

type UpdateDirectConnectGatewayAssociationOutput struct {

	// Information about an association between a Direct Connect gateway and a virtual
	// private gateway or transit gateway.
	DirectConnectGatewayAssociation *types.DirectConnectGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDirectConnectGatewayAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDirectConnectGatewayAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDirectConnectGatewayAssociation{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDirectConnectGatewayAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDirectConnectGatewayAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "UpdateDirectConnectGatewayAssociation",
	}
}
