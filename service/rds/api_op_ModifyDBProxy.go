// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the settings for an existing DB proxy.
func (c *Client) ModifyDBProxy(ctx context.Context, params *ModifyDBProxyInput, optFns ...func(*Options)) (*ModifyDBProxyOutput, error) {
	if params == nil {
		params = &ModifyDBProxyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBProxy", params, optFns, c.addOperationModifyDBProxyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBProxyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBProxyInput struct {

	// The identifier for the DBProxy to modify.
	//
	// This member is required.
	DBProxyName *string

	// The new authentication settings for the DBProxy.
	Auth []types.UserAuthConfig

	// Whether the proxy includes detailed information about SQL statements in its
	// logs. This information helps you to debug issues involving SQL behavior or the
	// performance and scalability of the proxy connections. The debug information
	// includes the text of SQL statements that you submit through the proxy. Thus,
	// only enable this setting when needed for debugging, and only when you have
	// security measures in place to safeguard any sensitive information that appears
	// in the logs.
	DebugLogging *bool

	// The number of seconds that a connection to the proxy can be inactive before the
	// proxy disconnects it. You can set this value higher or lower than the connection
	// timeout limit for the associated database.
	IdleClientTimeout *int32

	// The new identifier for the DBProxy. An identifier must begin with a letter and
	// must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen
	// or contain two consecutive hyphens.
	NewDBProxyName *string

	// Whether Transport Layer Security (TLS) encryption is required for connections to
	// the proxy. By enabling this setting, you can enforce encrypted TLS connections
	// to the proxy, even if the associated database doesn't use TLS.
	RequireTLS *bool

	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access
	// secrets in Amazon Web Services Secrets Manager.
	RoleArn *string

	// The new list of security groups for the DBProxy.
	SecurityGroups []string

	noSmithyDocumentSerde
}

type ModifyDBProxyOutput struct {

	// The DBProxy object representing the new settings for the proxy.
	DBProxy *types.DBProxy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBProxyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBProxy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBProxy{}, middleware.After)
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
	if err = addOpModifyDBProxyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBProxy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBProxy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBProxy",
	}
}
