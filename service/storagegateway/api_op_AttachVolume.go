// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Connects a volume to an iSCSI connection and then attaches the volume to the
// specified gateway. Detaching and attaching a volume enables you to recover your
// data from one gateway to a different gateway without creating a snapshot. It
// also makes it easier to move your volumes from an on-premises gateway to a
// gateway hosted on an Amazon EC2 instance.
func (c *Client) AttachVolume(ctx context.Context, params *AttachVolumeInput, optFns ...func(*Options)) (*AttachVolumeOutput, error) {
	if params == nil {
		params = &AttachVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachVolume", params, optFns, c.addOperationAttachVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// AttachVolumeInput
type AttachVolumeInput struct {

	// The Amazon Resource Name (ARN) of the gateway that you want to attach the volume
	// to.
	//
	// This member is required.
	GatewayARN *string

	// The network interface of the gateway on which to expose the iSCSI target. Only
	// IPv4 addresses are accepted. Use DescribeGatewayInformation to get a list of the
	// network interfaces available on a gateway. Valid Values: A valid IP address.
	//
	// This member is required.
	NetworkInterfaceId *string

	// The Amazon Resource Name (ARN) of the volume to attach to the specified gateway.
	//
	// This member is required.
	VolumeARN *string

	// The unique device ID or other distinguishing data that identifies the local disk
	// used to create the volume. This value is only required when you are attaching a
	// stored volume.
	DiskId *string

	// The name of the iSCSI target used by an initiator to connect to a volume and
	// used as a suffix for the target ARN. For example, specifying TargetName as
	// myvolume results in the target ARN of
	// arn:aws:storagegateway:us-east-2:111122223333:gateway/sgw-12A3456B/target/iqn.1997-05.com.amazon:myvolume.
	// The target name must be unique across all volumes on a gateway. If you don't
	// specify a value, Storage Gateway uses the value that was previously used for
	// this volume as the new target name.
	TargetName *string

	noSmithyDocumentSerde
}

// AttachVolumeOutput
type AttachVolumeOutput struct {

	// The Amazon Resource Name (ARN) of the volume target, which includes the iSCSI
	// name for the initiator that was used to connect to the target.
	TargetARN *string

	// The Amazon Resource Name (ARN) of the volume that was attached to the gateway.
	VolumeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAttachVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachVolume{}, middleware.After)
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
	if err = addOpAttachVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachVolume(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "AttachVolume",
	}
}
