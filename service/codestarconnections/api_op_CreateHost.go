// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarconnections

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/codestarconnections/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a resource that represents the infrastructure where a third-party
// provider is installed. The host is used when you create connections to an
// installed third-party provider type, such as GitHub Enterprise Server. You
// create one host for all connections to that provider. A host created through the
// CLI or the SDK is in `PENDING` status by default. You can make its status
// `AVAILABLE` by setting up the host in the console.
func (c *Client) CreateHost(ctx context.Context, params *CreateHostInput, optFns ...func(*Options)) (*CreateHostOutput, error) {
	if params == nil {
		params = &CreateHostInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHost", params, optFns, c.addOperationCreateHostMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHostOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHostInput struct {

	// The name of the host to be created. The name must be unique in the calling AWS
	// account.
	//
	// This member is required.
	Name *string

	// The endpoint of the infrastructure to be represented by the host after it is
	// created.
	//
	// This member is required.
	ProviderEndpoint *string

	// The name of the installed provider to be associated with your connection. The
	// host resource represents the infrastructure where your provider type is
	// installed. The valid provider type is GitHub Enterprise Server.
	//
	// This member is required.
	ProviderType types.ProviderType

	Tags []types.Tag

	// The VPC configuration to be provisioned for the host. A VPC must be configured
	// and the infrastructure to be represented by the host must already be connected
	// to the VPC.
	VpcConfiguration *types.VpcConfiguration

	noSmithyDocumentSerde
}

type CreateHostOutput struct {

	// The Amazon Resource Name (ARN) of the host to be created.
	HostArn *string

	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHostMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateHost{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateHost{}, middleware.After)
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
	if err = addOpCreateHostValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHost(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHost(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-connections",
		OperationName: "CreateHost",
	}
}
