// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete a user from launch profile membership.
func (c *Client) DeleteLaunchProfileMember(ctx context.Context, params *DeleteLaunchProfileMemberInput, optFns ...func(*Options)) (*DeleteLaunchProfileMemberOutput, error) {
	if params == nil {
		params = &DeleteLaunchProfileMemberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteLaunchProfileMember", params, optFns, c.addOperationDeleteLaunchProfileMemberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteLaunchProfileMemberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLaunchProfileMemberInput struct {

	// The Launch Profile ID.
	//
	// This member is required.
	LaunchProfileId *string

	// The principal ID. This currently supports a Amazon Web Services SSO UserId.
	//
	// This member is required.
	PrincipalId *string

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don’t specify a client token, the AWS SDK automatically
	// generates a client token and uses it for the request to ensure idempotency.
	ClientToken *string

	noSmithyDocumentSerde
}

type DeleteLaunchProfileMemberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteLaunchProfileMemberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteLaunchProfileMember{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteLaunchProfileMember{}, middleware.After)
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
	if err = addIdempotencyToken_opDeleteLaunchProfileMemberMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteLaunchProfileMemberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLaunchProfileMember(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteLaunchProfileMember struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteLaunchProfileMember) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteLaunchProfileMember) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteLaunchProfileMemberInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteLaunchProfileMemberInput ")
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
func addIdempotencyToken_opDeleteLaunchProfileMemberMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteLaunchProfileMember{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteLaunchProfileMember(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "DeleteLaunchProfileMember",
	}
}
