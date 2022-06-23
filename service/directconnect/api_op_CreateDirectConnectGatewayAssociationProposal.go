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

// Creates a proposal to associate the specified virtual private gateway or transit
// gateway with the specified Direct Connect gateway. You can associate a Direct
// Connect gateway and virtual private gateway or transit gateway that is owned by
// any Amazon Web Services account.
func (c *Client) CreateDirectConnectGatewayAssociationProposal(ctx context.Context, params *CreateDirectConnectGatewayAssociationProposalInput, optFns ...func(*Options)) (*CreateDirectConnectGatewayAssociationProposalOutput, error) {
	if params == nil {
		params = &CreateDirectConnectGatewayAssociationProposalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDirectConnectGatewayAssociationProposal", params, optFns, c.addOperationCreateDirectConnectGatewayAssociationProposalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDirectConnectGatewayAssociationProposalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDirectConnectGatewayAssociationProposalInput struct {

	// The ID of the Direct Connect gateway.
	//
	// This member is required.
	DirectConnectGatewayId *string

	// The ID of the Amazon Web Services account that owns the Direct Connect gateway.
	//
	// This member is required.
	DirectConnectGatewayOwnerAccount *string

	// The ID of the virtual private gateway or transit gateway.
	//
	// This member is required.
	GatewayId *string

	// The Amazon VPC prefixes to advertise to the Direct Connect gateway.
	AddAllowedPrefixesToDirectConnectGateway []types.RouteFilterPrefix

	// The Amazon VPC prefixes to no longer advertise to the Direct Connect gateway.
	RemoveAllowedPrefixesToDirectConnectGateway []types.RouteFilterPrefix

	noSmithyDocumentSerde
}

type CreateDirectConnectGatewayAssociationProposalOutput struct {

	// Information about the Direct Connect gateway proposal.
	DirectConnectGatewayAssociationProposal *types.DirectConnectGatewayAssociationProposal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDirectConnectGatewayAssociationProposalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDirectConnectGatewayAssociationProposal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDirectConnectGatewayAssociationProposal{}, middleware.After)
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
	if err = addOpCreateDirectConnectGatewayAssociationProposalValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDirectConnectGatewayAssociationProposal(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDirectConnectGatewayAssociationProposal(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "CreateDirectConnectGatewayAssociationProposal",
	}
}
