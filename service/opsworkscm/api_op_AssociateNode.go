// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a new node with the server. For more information about how to
// disassociate a node, see DisassociateNode. On a Chef server: This command is an
// alternative to knife bootstrap. Example (Chef): aws opsworks-cm associate-node
// --server-name MyServer --node-name MyManagedNode --engine-attributes
// "Name=CHEF_ORGANIZATION,Value=default"
// "Name=CHEF_NODE_PUBLIC_KEY,Value=public-key-pem" On a Puppet server, this
// command is an alternative to the puppet cert sign command that signs a Puppet
// node CSR. Example (Puppet): aws opsworks-cm associate-node --server-name
// MyServer --node-name MyManagedNode --engine-attributes
// "Name=PUPPET_NODE_CSR,Value=csr-pem" A node can can only be associated with
// servers that are in a HEALTHY state. Otherwise, an InvalidStateException is
// thrown. A ResourceNotFoundException is thrown when the server does not exist. A
// ValidationException is raised when parameters of the request are not valid. The
// AssociateNode API call can be integrated into Auto Scaling configurations, AWS
// Cloudformation templates, or the user data of a server's instance.
func (c *Client) AssociateNode(ctx context.Context, params *AssociateNodeInput, optFns ...func(*Options)) (*AssociateNodeOutput, error) {
	if params == nil {
		params = &AssociateNodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateNode", params, optFns, c.addOperationAssociateNodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateNodeInput struct {

	// Engine attributes used for associating the node. Attributes accepted in a
	// AssociateNode request for Chef
	//
	// * CHEF_ORGANIZATION: The Chef organization with
	// which the node is associated. By default only one organization named default can
	// exist.
	//
	// * CHEF_NODE_PUBLIC_KEY: A PEM-formatted public key. This key is required
	// for the chef-client agent to access the Chef API.
	//
	// Attributes accepted in a
	// AssociateNode request for Puppet
	//
	// * PUPPET_NODE_CSR: A PEM-formatted
	// certificate-signing request (CSR) that is created by the node.
	//
	// This member is required.
	EngineAttributes []types.EngineAttribute

	// The name of the node.
	//
	// This member is required.
	NodeName *string

	// The name of the server with which to associate the node.
	//
	// This member is required.
	ServerName *string

	noSmithyDocumentSerde
}

type AssociateNodeOutput struct {

	// Contains a token which can be passed to the DescribeNodeAssociationStatus API
	// call to get the status of the association request.
	NodeAssociationStatusToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateNodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateNode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateNode{}, middleware.After)
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
	if err = addOpAssociateNodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateNode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateNode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "AssociateNode",
	}
}
