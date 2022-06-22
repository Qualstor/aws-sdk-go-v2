// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the staging labels attached to a version of a secret. Secrets Manager
// uses staging labels to track a version as it progresses through the secret
// rotation process. Each staging label can be attached to only one version at a
// time. To add a staging label to a version when it is already attached to another
// version, Secrets Manager first removes it from the other version first and then
// attaches it to this one. For more information about versions and staging labels,
// see Concepts: Version
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/getting-started.html#term_version).
// The staging labels that you specify in the VersionStage parameter are added to
// the existing list of staging labels for the version. You can move the AWSCURRENT
// staging label to this version by including it in this call. Whenever you move
// AWSCURRENT, Secrets Manager automatically moves the label AWSPREVIOUS to the
// version that AWSCURRENT was removed from. If this action results in the last
// label being removed from a version, then the version is considered to be
// 'deprecated' and can be deleted by Secrets Manager. Required permissions:
// secretsmanager:UpdateSecretVersionStage. For more information, see  IAM policy
// actions for Secrets Manager
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions)
// and Authentication and access control in Secrets Manager
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html).
func (c *Client) UpdateSecretVersionStage(ctx context.Context, params *UpdateSecretVersionStageInput, optFns ...func(*Options)) (*UpdateSecretVersionStageOutput, error) {
	if params == nil {
		params = &UpdateSecretVersionStageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSecretVersionStage", params, optFns, c.addOperationUpdateSecretVersionStageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSecretVersionStageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSecretVersionStageInput struct {

	// The ARN or the name of the secret with the version and staging labelsto modify.
	// For an ARN, we recommend that you specify a complete ARN rather than a partial
	// ARN. See Finding a secret from a partial ARN
	// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/troubleshoot.html#ARN_secretnamehyphen).
	//
	// This member is required.
	SecretId *string

	// The staging label to add to this version.
	//
	// This member is required.
	VersionStage *string

	// The ID of the version to add the staging label to. To remove a label from a
	// version, then do not specify this parameter. If the staging label is already
	// attached to a different version of the secret, then you must also specify the
	// RemoveFromVersionId parameter.
	MoveToVersionId *string

	// The ID of the version that the staging label is to be removed from. If the
	// staging label you are trying to attach to one version is already attached to a
	// different version, then you must include this parameter and specify the version
	// that the label is to be removed from. If the label is attached and you either do
	// not specify this parameter, or the version ID does not match, then the operation
	// fails.
	RemoveFromVersionId *string

	noSmithyDocumentSerde
}

type UpdateSecretVersionStageOutput struct {

	// The ARN of the secret that was updated.
	ARN *string

	// The name of the secret that was updated.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSecretVersionStageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSecretVersionStage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSecretVersionStage{}, middleware.After)
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
	if err = addOpUpdateSecretVersionStageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSecretVersionStage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSecretVersionStage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "UpdateSecretVersionStage",
	}
}
