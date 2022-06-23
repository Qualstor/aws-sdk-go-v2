// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the Firewall Manager administrator account. The account must be a member of
// the organization in Organizations whose resources you want to protect. Firewall
// Manager sets the permissions that allow the account to administer your Firewall
// Manager policies. The account that you associate with Firewall Manager is called
// the Firewall Manager administrator account.
func (c *Client) AssociateAdminAccount(ctx context.Context, params *AssociateAdminAccountInput, optFns ...func(*Options)) (*AssociateAdminAccountOutput, error) {
	if params == nil {
		params = &AssociateAdminAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateAdminAccount", params, optFns, c.addOperationAssociateAdminAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateAdminAccountInput struct {

	// The Amazon Web Services account ID to associate with Firewall Manager as the
	// Firewall Manager administrator account. This must be an Organizations member
	// account. For more information about Organizations, see Managing the Amazon Web
	// Services Accounts in Your Organization
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts.html).
	//
	// This member is required.
	AdminAccount *string

	noSmithyDocumentSerde
}

type AssociateAdminAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateAdminAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateAdminAccount{}, middleware.After)
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
	if err = addOpAssociateAdminAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateAdminAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateAdminAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "AssociateAdminAccount",
	}
}
