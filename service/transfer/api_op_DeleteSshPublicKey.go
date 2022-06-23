// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a user's Secure Shell (SSH) public key.
func (c *Client) DeleteSshPublicKey(ctx context.Context, params *DeleteSshPublicKeyInput, optFns ...func(*Options)) (*DeleteSshPublicKeyOutput, error) {
	if params == nil {
		params = &DeleteSshPublicKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSshPublicKey", params, optFns, c.addOperationDeleteSshPublicKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSshPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSshPublicKeyInput struct {

	// A system-assigned unique identifier for a file transfer protocol-enabled server
	// instance that has the user assigned to it.
	//
	// This member is required.
	ServerId *string

	// A unique identifier used to reference your user's specific SSH key.
	//
	// This member is required.
	SshPublicKeyId *string

	// A unique string that identifies a user whose public key is being deleted.
	//
	// This member is required.
	UserName *string

	noSmithyDocumentSerde
}

type DeleteSshPublicKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSshPublicKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSshPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSshPublicKey{}, middleware.After)
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
	if err = addOpDeleteSshPublicKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSshPublicKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSshPublicKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "DeleteSshPublicKey",
	}
}
