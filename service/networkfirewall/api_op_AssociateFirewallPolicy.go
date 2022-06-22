// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a FirewallPolicy to a Firewall. A firewall policy defines how to
// monitor and manage your VPC network traffic, using a collection of inspection
// rule groups and other settings. Each firewall requires one firewall policy
// association, and you can use the same firewall policy for multiple firewalls.
func (c *Client) AssociateFirewallPolicy(ctx context.Context, params *AssociateFirewallPolicyInput, optFns ...func(*Options)) (*AssociateFirewallPolicyOutput, error) {
	if params == nil {
		params = &AssociateFirewallPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateFirewallPolicy", params, optFns, c.addOperationAssociateFirewallPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateFirewallPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateFirewallPolicyInput struct {

	// The Amazon Resource Name (ARN) of the firewall policy.
	//
	// This member is required.
	FirewallPolicyArn *string

	// The Amazon Resource Name (ARN) of the firewall. You must specify the ARN or the
	// name, and you can specify both.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it. You must specify the ARN or the name, and you can specify
	// both.
	FirewallName *string

	// An optional token that you can use for optimistic locking. Network Firewall
	// returns a token to your requests that access the firewall. The token marks the
	// state of the firewall resource at the time of the request. To make an
	// unconditional change to the firewall, omit the token in your update request.
	// Without the token, Network Firewall performs your updates regardless of whether
	// the firewall has changed since you last retrieved it. To make a conditional
	// change to the firewall, provide the token in your update request. Network
	// Firewall uses the token to ensure that the firewall hasn't changed since you
	// last retrieved it. If it has changed, the operation fails with an
	// InvalidTokenException. If this happens, retrieve the firewall again to get a
	// current copy of it with a new token. Reapply your changes as needed, then try
	// the operation again using the new token.
	UpdateToken *string

	noSmithyDocumentSerde
}

type AssociateFirewallPolicyOutput struct {

	// The Amazon Resource Name (ARN) of the firewall.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it.
	FirewallName *string

	// The Amazon Resource Name (ARN) of the firewall policy.
	FirewallPolicyArn *string

	// An optional token that you can use for optimistic locking. Network Firewall
	// returns a token to your requests that access the firewall. The token marks the
	// state of the firewall resource at the time of the request. To make an
	// unconditional change to the firewall, omit the token in your update request.
	// Without the token, Network Firewall performs your updates regardless of whether
	// the firewall has changed since you last retrieved it. To make a conditional
	// change to the firewall, provide the token in your update request. Network
	// Firewall uses the token to ensure that the firewall hasn't changed since you
	// last retrieved it. If it has changed, the operation fails with an
	// InvalidTokenException. If this happens, retrieve the firewall again to get a
	// current copy of it with a new token. Reapply your changes as needed, then try
	// the operation again using the new token.
	UpdateToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateFirewallPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpAssociateFirewallPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpAssociateFirewallPolicy{}, middleware.After)
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
	if err = addOpAssociateFirewallPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateFirewallPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateFirewallPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "AssociateFirewallPolicy",
	}
}
