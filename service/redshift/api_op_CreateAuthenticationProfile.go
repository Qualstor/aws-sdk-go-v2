// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an authentication profile with the specified parameters.
func (c *Client) CreateAuthenticationProfile(ctx context.Context, params *CreateAuthenticationProfileInput, optFns ...func(*Options)) (*CreateAuthenticationProfileOutput, error) {
	if params == nil {
		params = &CreateAuthenticationProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAuthenticationProfile", params, optFns, c.addOperationCreateAuthenticationProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAuthenticationProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAuthenticationProfileInput struct {

	// The content of the authentication profile in JSON format. The maximum length of
	// the JSON string is determined by a quota for your account.
	//
	// This member is required.
	AuthenticationProfileContent *string

	// The name of the authentication profile to be created.
	//
	// This member is required.
	AuthenticationProfileName *string

	noSmithyDocumentSerde
}

type CreateAuthenticationProfileOutput struct {

	// The content of the authentication profile in JSON format.
	AuthenticationProfileContent *string

	// The name of the authentication profile that was created.
	AuthenticationProfileName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAuthenticationProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateAuthenticationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateAuthenticationProfile{}, middleware.After)
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
	if err = addOpCreateAuthenticationProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAuthenticationProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAuthenticationProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateAuthenticationProfile",
	}
}
