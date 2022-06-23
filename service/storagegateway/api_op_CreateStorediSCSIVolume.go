// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a volume on a specified gateway. This operation is only supported in the
// stored volume gateway type. The size of the volume to create is inferred from
// the disk size. You can choose to preserve existing data on the disk, create
// volume from an existing snapshot, or create an empty volume. If you choose to
// create an empty gateway volume, then any existing data on the disk is erased. In
// the request, you must specify the gateway and the disk information on which you
// are creating the volume. In response, the gateway creates the volume and returns
// volume information such as the volume Amazon Resource Name (ARN), its size, and
// the iSCSI target ARN that initiators can use to connect to the volume target.
func (c *Client) CreateStorediSCSIVolume(ctx context.Context, params *CreateStorediSCSIVolumeInput, optFns ...func(*Options)) (*CreateStorediSCSIVolumeOutput, error) {
	if params == nil {
		params = &CreateStorediSCSIVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStorediSCSIVolume", params, optFns, c.addOperationCreateStorediSCSIVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStorediSCSIVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing one or more of the following fields:
//
// *
// CreateStorediSCSIVolumeInput$DiskId
//
// *
// CreateStorediSCSIVolumeInput$NetworkInterfaceId
//
// *
// CreateStorediSCSIVolumeInput$PreserveExistingData
//
// *
// CreateStorediSCSIVolumeInput$SnapshotId
//
// *
// CreateStorediSCSIVolumeInput$TargetName
type CreateStorediSCSIVolumeInput struct {

	// The unique identifier for the gateway local disk that is configured as a stored
	// volume. Use ListLocalDisks
	// (https://docs.aws.amazon.com/storagegateway/latest/userguide/API_ListLocalDisks.html)
	// to list disk IDs for a gateway.
	//
	// This member is required.
	DiskId *string

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// The network interface of the gateway on which to expose the iSCSI target. Only
	// IPv4 addresses are accepted. Use DescribeGatewayInformation to get a list of the
	// network interfaces available on a gateway. Valid Values: A valid IP address.
	//
	// This member is required.
	NetworkInterfaceId *string

	// Set to true if you want to preserve the data on the local disk. Otherwise, set
	// to false to create an empty volume. Valid Values: true | false
	//
	// This member is required.
	PreserveExistingData bool

	// The name of the iSCSI target used by an initiator to connect to a volume and
	// used as a suffix for the target ARN. For example, specifying TargetName as
	// myvolume results in the target ARN of
	// arn:aws:storagegateway:us-east-2:111122223333:gateway/sgw-12A3456B/target/iqn.1997-05.com.amazon:myvolume.
	// The target name must be unique across all volumes on a gateway. If you don't
	// specify a value, Storage Gateway uses the value that was previously used for
	// this volume as the new target name.
	//
	// This member is required.
	TargetName *string

	// Set to true to use Amazon S3 server-side encryption with your own KMS key, or
	// false to use a key managed by Amazon S3. Optional. Valid Values: true | false
	KMSEncrypted *bool

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for
	// Amazon S3 server-side encryption. Storage Gateway does not support asymmetric
	// CMKs. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string

	// The snapshot ID (e.g., "snap-1122aabb") of the snapshot to restore as the new
	// stored volume. Specify this field if you want to create the iSCSI storage volume
	// from a snapshot; otherwise, do not include this field. To list snapshots for
	// your account use DescribeSnapshots
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshots.html)
	// in the Amazon Elastic Compute Cloud API Reference.
	SnapshotId *string

	// A list of up to 50 tags that can be assigned to a stored volume. Each tag is a
	// key-value pair. Valid characters for key and value are letters, spaces, and
	// numbers representable in UTF-8 format, and the following special characters: + -
	// = . _ : / @. The maximum length of a tag's key is 128 characters, and the
	// maximum length for a tag's value is 256.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// A JSON object containing the following fields:
type CreateStorediSCSIVolumeOutput struct {

	// The Amazon Resource Name (ARN) of the volume target, which includes the iSCSI
	// name that initiators can use to connect to the target.
	TargetARN *string

	// The Amazon Resource Name (ARN) of the configured volume.
	VolumeARN *string

	// The size of the volume in bytes.
	VolumeSizeInBytes int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStorediSCSIVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStorediSCSIVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStorediSCSIVolume{}, middleware.After)
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
	if err = addOpCreateStorediSCSIVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStorediSCSIVolume(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStorediSCSIVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateStorediSCSIVolume",
	}
}
