// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a file share from an S3 File Gateway. This operation is only supported
// for S3 File Gateways.
func (c *Client) DeleteFileShare(ctx context.Context, params *DeleteFileShareInput, optFns ...func(*Options)) (*DeleteFileShareOutput, error) {
	if params == nil {
		params = &DeleteFileShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFileShare", params, optFns, c.addOperationDeleteFileShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFileShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DeleteFileShareInput
type DeleteFileShareInput struct {

	// The Amazon Resource Name (ARN) of the file share to be deleted.
	//
	// This member is required.
	FileShareARN *string

	// If this value is set to true, the operation deletes a file share immediately and
	// aborts all data uploads to Amazon Web Services. Otherwise, the file share is not
	// deleted until all data is uploaded to Amazon Web Services. This process aborts
	// the data upload process, and the file share enters the FORCE_DELETING status.
	// Valid Values: true | false
	ForceDelete bool

	noSmithyDocumentSerde
}

// DeleteFileShareOutput
type DeleteFileShareOutput struct {

	// The Amazon Resource Name (ARN) of the deleted file share.
	FileShareARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFileShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteFileShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteFileShare{}, middleware.After)
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
	if err = addOpDeleteFileShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFileShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFileShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DeleteFileShare",
	}
}
