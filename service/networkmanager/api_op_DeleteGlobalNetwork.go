// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an existing global network. You must first delete all global network
// objects (devices, links, and sites), deregister all transit gateways, and delete
// any core networks.
func (c *Client) DeleteGlobalNetwork(ctx context.Context, params *DeleteGlobalNetworkInput, optFns ...func(*Options)) (*DeleteGlobalNetworkOutput, error) {
	if params == nil {
		params = &DeleteGlobalNetworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGlobalNetwork", params, optFns, c.addOperationDeleteGlobalNetworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGlobalNetworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGlobalNetworkInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	noSmithyDocumentSerde
}

type DeleteGlobalNetworkOutput struct {

	// Information about the global network.
	GlobalNetwork *types.GlobalNetwork

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteGlobalNetworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteGlobalNetwork{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteGlobalNetwork{}, middleware.After)
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
	if err = addOpDeleteGlobalNetworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGlobalNetwork(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteGlobalNetwork(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "DeleteGlobalNetwork",
	}
}
