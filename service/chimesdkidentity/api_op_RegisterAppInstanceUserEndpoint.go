// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers an endpoint under an Amazon Chime AppInstanceUser. The endpoint
// receives messages for a user. For push notifications, the endpoint is a mobile
// device used to receive mobile push notifications for a user.
func (c *Client) RegisterAppInstanceUserEndpoint(ctx context.Context, params *RegisterAppInstanceUserEndpointInput, optFns ...func(*Options)) (*RegisterAppInstanceUserEndpointOutput, error) {
	if params == nil {
		params = &RegisterAppInstanceUserEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterAppInstanceUserEndpoint", params, optFns, c.addOperationRegisterAppInstanceUserEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterAppInstanceUserEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterAppInstanceUserEndpointInput struct {

	// The ARN of the AppInstanceUser.
	//
	// This member is required.
	AppInstanceUserArn *string

	// The idempotency token for each client request.
	//
	// This member is required.
	ClientRequestToken *string

	// The attributes of an Endpoint.
	//
	// This member is required.
	EndpointAttributes *types.EndpointAttributes

	// The ARN of the resource to which the endpoint belongs.
	//
	// This member is required.
	ResourceArn *string

	// The type of the AppInstanceUserEndpoint. Supported types:
	//
	// * APNS: The mobile
	// notification service for an Apple device.
	//
	// * APNS_SANDBOX: The sandbox
	// environment of the mobile notification service for an Apple device.
	//
	// * GCM: The
	// mobile notification service for an Android device.
	//
	// Populate the ResourceArn
	// value of each type as PinpointAppArn.
	//
	// This member is required.
	Type types.AppInstanceUserEndpointType

	// Boolean that controls whether the AppInstanceUserEndpoint is opted in to receive
	// messages. ALL indicates the endpoint receives all messages. NONE indicates the
	// endpoint receives no messages.
	AllowMessages types.AllowMessages

	// The name of the AppInstanceUserEndpoint.
	Name *string

	noSmithyDocumentSerde
}

type RegisterAppInstanceUserEndpointOutput struct {

	// The ARN of the AppInstanceUser.
	AppInstanceUserArn *string

	// The unique identifier of the AppInstanceUserEndpoint.
	EndpointId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterAppInstanceUserEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterAppInstanceUserEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterAppInstanceUserEndpoint{}, middleware.After)
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
	if err = addIdempotencyToken_opRegisterAppInstanceUserEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRegisterAppInstanceUserEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterAppInstanceUserEndpoint(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpRegisterAppInstanceUserEndpoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRegisterAppInstanceUserEndpoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRegisterAppInstanceUserEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RegisterAppInstanceUserEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RegisterAppInstanceUserEndpointInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opRegisterAppInstanceUserEndpointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRegisterAppInstanceUserEndpoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRegisterAppInstanceUserEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "RegisterAppInstanceUserEndpoint",
	}
}
