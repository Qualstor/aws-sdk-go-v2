// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates some of the parameters of a previously created location for Network File
// System (NFS) access. For information about creating an NFS location, see
// Creating a location for NFS
// (https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html).
func (c *Client) UpdateLocationNfs(ctx context.Context, params *UpdateLocationNfsInput, optFns ...func(*Options)) (*UpdateLocationNfsOutput, error) {
	if params == nil {
		params = &UpdateLocationNfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationNfs", params, optFns, c.addOperationUpdateLocationNfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationNfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationNfsInput struct {

	// The Amazon Resource Name (ARN) of the NFS location to update.
	//
	// This member is required.
	LocationArn *string

	// Represents the mount options that are available for DataSync to access an NFS
	// location.
	MountOptions *types.NfsMountOptions

	// A list of Amazon Resource Names (ARNs) of agents to use for a Network File
	// System (NFS) location.
	OnPremConfig *types.OnPremConfig

	// The subdirectory in the NFS file system that is used to read data from the NFS
	// source location or write data to the NFS destination. The NFS path should be a
	// path that's exported by the NFS server, or a subdirectory of that path. The path
	// should be such that it can be mounted by other NFS clients in your network. To
	// see all the paths exported by your NFS server, run "showmount -e
	// nfs-server-name" from an NFS client that has access to your server. You can
	// specify any directory that appears in the results, and any subdirectory of that
	// directory. Ensure that the NFS export is accessible without Kerberos
	// authentication. To transfer all the data in the folder that you specified,
	// DataSync must have permissions to read all the data. To ensure this, either
	// configure the NFS export with no_root_squash, or ensure that the files you want
	// DataSync to access have permissions that allow read access for all users. Doing
	// either option enables the agent to read the files. For the agent to access
	// directories, you must additionally enable all execute access. If you are copying
	// data to or from your Snowcone device, see NFS Server on Snowcone
	// (https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone)
	// for more information. For information about NFS export configuration, see 18.7.
	// The /etc/exports Configuration File in the Red Hat Enterprise Linux
	// documentation.
	Subdirectory *string

	noSmithyDocumentSerde
}

type UpdateLocationNfsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationNfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationNfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationNfs{}, middleware.After)
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
	if err = addOpUpdateLocationNfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationNfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLocationNfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "UpdateLocationNfs",
	}
}
