// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new text message and sends it to a recipient's phone number. SMS
// throughput limits are measured in Message Parts per Second (MPS). Your MPS limit
// depends on the destination country of your messages, as well as the type of
// phone number (origination number) that you use to send the message. For more
// information, see Message Parts per Second (MPS) limits
// (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-limitations-mps.html)
// in the Amazon Pinpoint User Guide.
func (c *Client) SendTextMessage(ctx context.Context, params *SendTextMessageInput, optFns ...func(*Options)) (*SendTextMessageOutput, error) {
	if params == nil {
		params = &SendTextMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendTextMessage", params, optFns, c.addOperationSendTextMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendTextMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendTextMessageInput struct {

	// The destination phone number in E.164 format.
	//
	// This member is required.
	DestinationPhoneNumber *string

	// The name of the configuration set to use. This can be either the
	// ConfigurationSetName or ConfigurationSetArn.
	ConfigurationSetName *string

	// You can specify custom data in this field. If you do, that data is logged to the
	// event destination.
	Context map[string]string

	// This field is used for any country-specific registration requirements.
	// Currently, this setting is only used when you send messages to recipients in
	// India using a sender ID. For more information see Special requirements for
	// sending SMS messages to recipients in India
	// (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-senderid-india.html).
	DestinationCountryParameters map[string]string

	// When set to true, the message is checked and validated, but isn't sent to the
	// end recipient.
	DryRun bool

	// When you register a short code in the US, you must specify a program name. If
	// you don’t have a US short code, omit this attribute.
	Keyword *string

	// The maximum amount that you want to spend, in US dollars, per each text message
	// part. A text message can contain multiple parts.
	MaxPrice *string

	// The body of the text message.
	MessageBody *string

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	MessageType types.MessageType

	// The origination identity of the message. This can be either the PhoneNumber,
	// PhoneNumberId, PhoneNumberArn, SenderId, SenderIdArn, PoolId, or PoolArn.
	OriginationIdentity *string

	// How long the text message is valid for. By default this is 72 hours.
	TimeToLive *int32

	noSmithyDocumentSerde
}

type SendTextMessageOutput struct {

	// The unique identifier for the message.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendTextMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSendTextMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSendTextMessage{}, middleware.After)
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
	if err = addOpSendTextMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendTextMessage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendTextMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "SendTextMessage",
	}
}
