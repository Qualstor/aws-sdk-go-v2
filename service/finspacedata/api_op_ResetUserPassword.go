// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Resets the password for a specified user ID and generates a temporary one. Only
// a superuser can reset password for other users. Resetting the password
// immediately invalidates the previous password associated with the user.
func (c *Client) ResetUserPassword(ctx context.Context, params *ResetUserPasswordInput, optFns ...func(*Options)) (*ResetUserPasswordOutput, error) {
	if params == nil {
		params = &ResetUserPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetUserPassword", params, optFns, c.addOperationResetUserPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetUserPasswordInput struct {

	// The unique identifier of the user that a temporary password is requested for.
	//
	// This member is required.
	UserId *string

	// A token that ensures idempotency. This token expires in 10 minutes.
	ClientToken *string

	noSmithyDocumentSerde
}

type ResetUserPasswordOutput struct {

	// A randomly generated temporary password for the requested user account. This
	// password expires in 7 days.
	TemporaryPassword *string

	// The unique identifier of the user that a new password is generated for.
	UserId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResetUserPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpResetUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpResetUserPassword{}, middleware.After)
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opResetUserPasswordMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpResetUserPasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetUserPassword(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpResetUserPassword struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpResetUserPassword) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpResetUserPassword) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ResetUserPasswordInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ResetUserPasswordInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opResetUserPasswordMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpResetUserPassword{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opResetUserPassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "ResetUserPassword",
	}
}
