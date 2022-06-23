// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a channel to which you can add users and send messages. Restriction: You
// can't change a channel's privacy. The x-amz-chime-bearer request header is
// mandatory. Use the AppInstanceUserArn of the user that makes the API call as the
// value in the header.
func (c *Client) CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) {
	if params == nil {
		params = &CreateChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChannel", params, optFns, c.addOperationCreateChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChannelInput struct {

	// The ARN of the channel request.
	//
	// This member is required.
	AppInstanceArn *string

	// The client token for the request. An Idempotency token.
	//
	// This member is required.
	ClientRequestToken *string

	// The name of the channel.
	//
	// This member is required.
	Name *string

	// The AppInstanceUserArn of the user that makes the API call.
	ChimeBearer *string

	// The metadata of the creation request. Limited to 1KB and UTF-8.
	Metadata *string

	// The channel mode: UNRESTRICTED or RESTRICTED. Administrators, moderators, and
	// channel members can add themselves and other members to unrestricted channels.
	// Only administrators and moderators can add members to restricted channels.
	Mode types.ChannelMode

	// The channel's privacy level: PUBLIC or PRIVATE. Private channels aren't
	// discoverable by users outside the channel. Public channels are discoverable by
	// anyone in the AppInstance.
	Privacy types.ChannelPrivacy

	// The tags for the creation request.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateChannelOutput struct {

	// The ARN of the channel.
	ChannelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChannel{}, middleware.After)
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
	if err = addEndpointPrefix_opCreateChannelMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateChannelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChannel(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateChannelMiddleware struct {
}

func (*endpointPrefix_opCreateChannelMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateChannelMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "messaging-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateChannelMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateChannelMiddleware{}, `OperationSerializer`, middleware.After)
}

type idempotencyToken_initializeOpCreateChannel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateChannel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateChannelInput ")
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
func addIdempotencyToken_opCreateChannelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateChannel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateChannel",
	}
}
