// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an Amazon FSx for OpenZFS snapshot. After deletion, the snapshot no
// longer exists, and its data is gone. Deleting a snapshot doesn't affect
// snapshots stored in a file system backup. The DeleteSnapshot operation returns
// instantly. The snapshot appears with the lifecycle status of DELETING until the
// deletion is complete.
func (c *Client) DeleteSnapshot(ctx context.Context, params *DeleteSnapshotInput, optFns ...func(*Options)) (*DeleteSnapshotOutput, error) {
	if params == nil {
		params = &DeleteSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSnapshot", params, optFns, c.addOperationDeleteSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSnapshotInput struct {

	// The ID of the snapshot that you want to delete.
	//
	// This member is required.
	SnapshotId *string

	// (Optional) An idempotency token for resource creation, in a string of up to 64
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type DeleteSnapshotOutput struct {

	// The lifecycle status of the snapshot. If the DeleteSnapshot operation is
	// successful, this status is DELETING.
	Lifecycle types.SnapshotLifecycle

	// The ID of the deleted snapshot.
	SnapshotId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSnapshot{}, middleware.After)
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
	if err = addIdempotencyToken_opDeleteSnapshotMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSnapshot(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteSnapshot struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteSnapshot) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteSnapshotInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteSnapshotMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteSnapshot{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DeleteSnapshot",
	}
}
