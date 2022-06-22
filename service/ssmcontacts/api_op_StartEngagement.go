// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an engagement to a contact or escalation plan. The engagement engages
// each contact specified in the incident.
func (c *Client) StartEngagement(ctx context.Context, params *StartEngagementInput, optFns ...func(*Options)) (*StartEngagementOutput, error) {
	if params == nil {
		params = &StartEngagementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartEngagement", params, optFns, c.addOperationStartEngagementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartEngagementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartEngagementInput struct {

	// The Amazon Resource Name (ARN) of the contact being engaged.
	//
	// This member is required.
	ContactId *string

	// The secure content of the message that was sent to the contact. Use this field
	// for engagements to VOICE or EMAIL.
	//
	// This member is required.
	Content *string

	// The user that started the engagement.
	//
	// This member is required.
	Sender *string

	// The secure subject of the message that was sent to the contact. Use this field
	// for engagements to VOICE or EMAIL.
	//
	// This member is required.
	Subject *string

	// A token ensuring that the operation is called only once with the specified
	// details.
	IdempotencyToken *string

	// The ARN of the incident that the engagement is part of.
	IncidentId *string

	// The insecure content of the message that was sent to the contact. Use this field
	// for engagements to SMS.
	PublicContent *string

	// The insecure subject of the message that was sent to the contact. Use this field
	// for engagements to SMS.
	PublicSubject *string

	noSmithyDocumentSerde
}

type StartEngagementOutput struct {

	// The ARN of the engagement.
	//
	// This member is required.
	EngagementArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartEngagementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartEngagement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartEngagement{}, middleware.After)
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
	if err = addIdempotencyToken_opStartEngagementMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartEngagementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartEngagement(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartEngagement struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartEngagement) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartEngagement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartEngagementInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartEngagementInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartEngagementMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartEngagement{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartEngagement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-contacts",
		OperationName: "StartEngagement",
	}
}
