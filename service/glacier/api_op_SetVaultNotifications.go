// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/Qualstor/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/Qualstor/aws-sdk-go-v2/service/glacier/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation configures notifications that will be sent when specific events
// happen to a vault. By default, you don't get any notifications. To configure
// vault notifications, send a PUT request to the notification-configuration
// subresource of the vault. The request should include a JSON document that
// provides an Amazon SNS topic and specific events for which you want Amazon S3
// Glacier to send notifications to the topic. Amazon SNS topics must grant
// permission to the vault to be allowed to publish notifications to the topic. You
// can configure a vault to publish a notification for the following vault
// events:
//
// * ArchiveRetrievalCompleted This event occurs when a job that was
// initiated for an archive retrieval is completed (InitiateJob). The status of the
// completed job can be "Succeeded" or "Failed". The notification sent to the SNS
// topic is the same output as returned from DescribeJob.
//
// *
// InventoryRetrievalCompleted This event occurs when a job that was initiated for
// an inventory retrieval is completed (InitiateJob). The status of the completed
// job can be "Succeeded" or "Failed". The notification sent to the SNS topic is
// the same output as returned from DescribeJob.
//
// An AWS account has full
// permission to perform all operations (actions). However, AWS Identity and Access
// Management (IAM) users don't have any permissions by default. You must grant
// them explicit permission to perform specific actions. For more information, see
// Access Control Using AWS Identity and Access Management (IAM)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
// For conceptual information and underlying REST API, see Configuring Vault
// Notifications in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/configuring-notifications.html)
// and Set Vault Notification Configuration
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-put.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) SetVaultNotifications(ctx context.Context, params *SetVaultNotificationsInput, optFns ...func(*Options)) (*SetVaultNotificationsOutput, error) {
	if params == nil {
		params = &SetVaultNotificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetVaultNotifications", params, optFns, c.addOperationSetVaultNotificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetVaultNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to configure notifications that will be sent when specific
// events happen to a vault.
type SetVaultNotificationsInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// Provides options for specifying notification configuration.
	VaultNotificationConfig *types.VaultNotificationConfig

	noSmithyDocumentSerde
}

type SetVaultNotificationsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetVaultNotificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSetVaultNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSetVaultNotifications{}, middleware.After)
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
	if err = addOpSetVaultNotificationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetVaultNotifications(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSetVaultNotifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "SetVaultNotifications",
	}
}
