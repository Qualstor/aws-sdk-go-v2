// Code generated by smithy-go-codegen DO NOT EDIT.

package ssooidc

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates and returns an access token for the authorized client. The access token
// issued will be used to fetch short-term credentials for the assigned roles in
// the AWS account.
func (c *Client) CreateToken(ctx context.Context, params *CreateTokenInput, optFns ...func(*Options)) (*CreateTokenOutput, error) {
	if params == nil {
		params = &CreateTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateToken", params, optFns, c.addOperationCreateTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTokenInput struct {

	// The unique identifier string for each client. This value should come from the
	// persisted result of the RegisterClient API.
	//
	// This member is required.
	ClientId *string

	// A secret string generated for the client. This value should come from the
	// persisted result of the RegisterClient API.
	//
	// This member is required.
	ClientSecret *string

	// Used only when calling this API for the device code grant type. This short-term
	// code is used to identify this authentication attempt. This should come from an
	// in-memory reference to the result of the StartDeviceAuthorization API.
	//
	// This member is required.
	DeviceCode *string

	// Supports grant types for authorization code, refresh token, and device code
	// request.
	//
	// This member is required.
	GrantType *string

	// The authorization code received from the authorization service. This parameter
	// is required to perform an authorization grant request to get access to a token.
	Code *string

	// The location of the application that will receive the authorization code. Users
	// authorize the service to send the request to this location.
	RedirectUri *string

	// The token used to obtain an access token in the event that the access token is
	// invalid or expired. This token is not issued by the service.
	RefreshToken *string

	// The list of scopes that is defined by the client. Upon authorization, this list
	// is used to restrict permissions when granting an access token.
	Scope []string

	noSmithyDocumentSerde
}

type CreateTokenOutput struct {

	// An opaque token to access AWS SSO resources assigned to a user.
	AccessToken *string

	// Indicates the time in seconds when an access token will expire.
	ExpiresIn int32

	// The identifier of the user that associated with the access token, if present.
	IdToken *string

	// A token that, if present, can be used to refresh a previously issued access
	// token that might have expired.
	RefreshToken *string

	// Used to notify the client that the returned token is an access token. The
	// supported type is BearerToken.
	TokenType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateToken{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpCreateTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateToken",
	}
}
