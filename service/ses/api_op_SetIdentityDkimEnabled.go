// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables or disables Easy DKIM signing of email sent from an identity. If Easy
// DKIM signing is enabled for a domain, then Amazon SES uses DKIM to sign all
// email that it sends from addresses on that domain. If Easy DKIM signing is
// enabled for an email address, then Amazon SES uses DKIM to sign all email it
// sends from that address. For email addresses (for example, user@example.com),
// you can only enable DKIM signing if the corresponding domain (in this case,
// example.com) has been set up to use Easy DKIM. You can enable DKIM signing for
// an identity at any time after you start the verification process for the
// identity, even if the verification process isn't complete. You can execute this
// operation no more than once per second. For more information about Easy DKIM
// signing, go to the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).
func (c *Client) SetIdentityDkimEnabled(ctx context.Context, params *SetIdentityDkimEnabledInput, optFns ...func(*Options)) (*SetIdentityDkimEnabledOutput, error) {
	if params == nil {
		params = &SetIdentityDkimEnabledInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetIdentityDkimEnabled", params, optFns, c.addOperationSetIdentityDkimEnabledMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetIdentityDkimEnabledOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to enable or disable Amazon SES Easy DKIM signing for an
// identity. For more information about setting up Easy DKIM, see the Amazon SES
// Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).
type SetIdentityDkimEnabledInput struct {

	// Sets whether DKIM signing is enabled for an identity. Set to true to enable DKIM
	// signing for this identity; false to disable it.
	//
	// This member is required.
	DkimEnabled bool

	// The identity for which DKIM signing should be enabled or disabled.
	//
	// This member is required.
	Identity *string

	noSmithyDocumentSerde
}

// An empty element returned on a successful request.
type SetIdentityDkimEnabledOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetIdentityDkimEnabledMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetIdentityDkimEnabled{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetIdentityDkimEnabled{}, middleware.After)
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
	if err = addOpSetIdentityDkimEnabledValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetIdentityDkimEnabled(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetIdentityDkimEnabled(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SetIdentityDkimEnabled",
	}
}
