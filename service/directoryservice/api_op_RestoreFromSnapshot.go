// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Restores a directory using an existing directory snapshot. When you restore a
// directory from a snapshot, any changes made to the directory after the snapshot
// date are overwritten. This action returns as soon as the restore operation is
// initiated. You can monitor the progress of the restore operation by calling the
// DescribeDirectories operation with the directory identifier. When the
// DirectoryDescription.Stage value changes to Active, the restore operation is
// complete.
func (c *Client) RestoreFromSnapshot(ctx context.Context, params *RestoreFromSnapshotInput, optFns ...func(*Options)) (*RestoreFromSnapshotOutput, error) {
	if params == nil {
		params = &RestoreFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreFromSnapshot", params, optFns, c.addOperationRestoreFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// An object representing the inputs for the RestoreFromSnapshot operation.
type RestoreFromSnapshotInput struct {

	// The identifier of the snapshot to restore from.
	//
	// This member is required.
	SnapshotId *string

	noSmithyDocumentSerde
}

// Contains the results of the RestoreFromSnapshot operation.
type RestoreFromSnapshotOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRestoreFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRestoreFromSnapshot{}, middleware.After)
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
	if err = addOpRestoreFromSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreFromSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreFromSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "RestoreFromSnapshot",
	}
}
