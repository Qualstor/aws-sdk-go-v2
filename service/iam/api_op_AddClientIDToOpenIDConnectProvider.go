// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a new client ID (also known as audience) to the list of client IDs already
// registered for the specified IAM OpenID Connect (OIDC) provider resource. This
// operation is idempotent; it does not fail or return an error if you add an
// existing client ID to the provider.
func (c *Client) AddClientIDToOpenIDConnectProvider(ctx context.Context, params *AddClientIDToOpenIDConnectProviderInput, optFns ...func(*Options)) (*AddClientIDToOpenIDConnectProviderOutput, error) {
	if params == nil {
		params = &AddClientIDToOpenIDConnectProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddClientIDToOpenIDConnectProvider", params, optFns, c.addOperationAddClientIDToOpenIDConnectProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddClientIDToOpenIDConnectProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddClientIDToOpenIDConnectProviderInput struct {

	// The client ID (also known as audience) to add to the IAM OpenID Connect provider
	// resource.
	//
	// This member is required.
	ClientID *string

	// The Amazon Resource Name (ARN) of the IAM OpenID Connect (OIDC) provider
	// resource to add the client ID to. You can get a list of OIDC provider ARNs by
	// using the ListOpenIDConnectProviders operation.
	//
	// This member is required.
	OpenIDConnectProviderArn *string

	noSmithyDocumentSerde
}

type AddClientIDToOpenIDConnectProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddClientIDToOpenIDConnectProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddClientIDToOpenIDConnectProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddClientIDToOpenIDConnectProvider{}, middleware.After)
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
	if err = addOpAddClientIDToOpenIDConnectProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddClientIDToOpenIDConnectProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddClientIDToOpenIDConnectProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "AddClientIDToOpenIDConnectProvider",
	}
}
