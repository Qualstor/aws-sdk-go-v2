// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a block storage disk that can be attached to an Amazon Lightsail
// instance in the same Availability Zone (e.g., us-east-2a). The create disk
// operation supports tag-based access control via request tags. For more
// information, see the Amazon Lightsail Developer Guide
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateDisk(ctx context.Context, params *CreateDiskInput, optFns ...func(*Options)) (*CreateDiskOutput, error) {
	if params == nil {
		params = &CreateDiskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDisk", params, optFns, c.addOperationCreateDiskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDiskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDiskInput struct {

	// The Availability Zone where you want to create the disk (e.g., us-east-2a). Use
	// the same Availability Zone as the Lightsail instance to which you want to attach
	// the disk. Use the get regions operation to list the Availability Zones where
	// Lightsail is currently available.
	//
	// This member is required.
	AvailabilityZone *string

	// The unique Lightsail disk name (e.g., my-disk).
	//
	// This member is required.
	DiskName *string

	// The size of the disk in GB (e.g., 32).
	//
	// This member is required.
	SizeInGb *int32

	// An array of objects that represent the add-ons to enable for the new disk.
	AddOns []types.AddOnRequest

	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDiskOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDiskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDisk{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDisk{}, middleware.After)
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
	if err = addOpCreateDiskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDisk(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDisk(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateDisk",
	}
}
