// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Redacts message content, but not metadata. The message exists in the back end,
// but the action returns null content, and the state shows as redacted. The
// x-amz-chime-bearer request header is mandatory. Use the AppInstanceUserArn of
// the user that makes the API call as the value in the header.
func (c *Client) RedactChannelMessage(ctx context.Context, params *RedactChannelMessageInput, optFns ...func(*Options)) (*RedactChannelMessageOutput, error) {
	if params == nil {
		params = &RedactChannelMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RedactChannelMessage", params, optFns, c.addOperationRedactChannelMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RedactChannelMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RedactChannelMessageInput struct {

	// The ARN of the channel containing the messages that you want to redact.
	//
	// This member is required.
	ChannelArn *string

	// The ID of the message being redacted.
	//
	// This member is required.
	MessageId *string

	// The AppInstanceUserArn of the user that makes the API call.
	ChimeBearer *string

	noSmithyDocumentSerde
}

type RedactChannelMessageOutput struct {

	// The ARN of the channel containing the messages that you want to redact.
	ChannelArn *string

	// The ID of the message being redacted.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRedactChannelMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRedactChannelMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRedactChannelMessage{}, middleware.After)
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
	if err = addEndpointPrefix_opRedactChannelMessageMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRedactChannelMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRedactChannelMessage(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opRedactChannelMessageMiddleware struct {
}

func (*endpointPrefix_opRedactChannelMessageMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opRedactChannelMessageMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
func addEndpointPrefix_opRedactChannelMessageMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opRedactChannelMessageMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opRedactChannelMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "RedactChannelMessage",
	}
}
