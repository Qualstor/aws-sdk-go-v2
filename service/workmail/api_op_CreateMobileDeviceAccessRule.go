// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new mobile device access rule for the specified Amazon WorkMail
// organization.
func (c *Client) CreateMobileDeviceAccessRule(ctx context.Context, params *CreateMobileDeviceAccessRuleInput, optFns ...func(*Options)) (*CreateMobileDeviceAccessRuleOutput, error) {
	if params == nil {
		params = &CreateMobileDeviceAccessRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMobileDeviceAccessRule", params, optFns, c.addOperationCreateMobileDeviceAccessRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMobileDeviceAccessRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMobileDeviceAccessRuleInput struct {

	// The effect of the rule when it matches. Allowed values are ALLOW or DENY.
	//
	// This member is required.
	Effect types.MobileDeviceAccessRuleEffect

	// The rule name.
	//
	// This member is required.
	Name *string

	// The Amazon WorkMail organization under which the rule will be created.
	//
	// This member is required.
	OrganizationId *string

	// The idempotency token for the client request.
	ClientToken *string

	// The rule description.
	Description *string

	// Device models that the rule will match.
	DeviceModels []string

	// Device operating systems that the rule will match.
	DeviceOperatingSystems []string

	// Device types that the rule will match.
	DeviceTypes []string

	// Device user agents that the rule will match.
	DeviceUserAgents []string

	// Device models that the rule will not match. All other device models will match.
	NotDeviceModels []string

	// Device operating systems that the rule will not match. All other device
	// operating systems will match.
	NotDeviceOperatingSystems []string

	// Device types that the rule will not match. All other device types will match.
	NotDeviceTypes []string

	// Device user agents that the rule will not match. All other device user agents
	// will match.
	NotDeviceUserAgents []string

	noSmithyDocumentSerde
}

type CreateMobileDeviceAccessRuleOutput struct {

	// The identifier for the newly created mobile device access rule.
	MobileDeviceAccessRuleId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMobileDeviceAccessRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMobileDeviceAccessRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMobileDeviceAccessRule{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateMobileDeviceAccessRuleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMobileDeviceAccessRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMobileDeviceAccessRule(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateMobileDeviceAccessRule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMobileDeviceAccessRule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMobileDeviceAccessRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMobileDeviceAccessRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMobileDeviceAccessRuleInput ")
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
func addIdempotencyToken_opCreateMobileDeviceAccessRuleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMobileDeviceAccessRule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMobileDeviceAccessRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "CreateMobileDeviceAccessRule",
	}
}
