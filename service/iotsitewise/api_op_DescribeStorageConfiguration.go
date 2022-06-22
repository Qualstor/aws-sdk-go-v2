// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about the storage configuration for IoT SiteWise.
func (c *Client) DescribeStorageConfiguration(ctx context.Context, params *DescribeStorageConfigurationInput, optFns ...func(*Options)) (*DescribeStorageConfigurationOutput, error) {
	if params == nil {
		params = &DescribeStorageConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStorageConfiguration", params, optFns, c.addOperationDescribeStorageConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStorageConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStorageConfigurationInput struct {
	noSmithyDocumentSerde
}

type DescribeStorageConfigurationOutput struct {

	// Contains current status information for the configuration.
	//
	// This member is required.
	ConfigurationStatus *types.ConfigurationStatus

	// The storage tier that you specified for your data. The storageType parameter can
	// be one of the following values:
	//
	// * SITEWISE_DEFAULT_STORAGE – IoT SiteWise saves
	// your data into the hot tier. The hot tier is a service-managed database.
	//
	// *
	// MULTI_LAYER_STORAGE – IoT SiteWise saves your data in both the cold tier and the
	// hot tier. The cold tier is a customer-managed Amazon S3 bucket.
	//
	// This member is required.
	StorageType types.StorageType

	// Contains the storage configuration for time series (data streams) that aren't
	// associated with asset properties. The disassociatedDataStorage can be one of the
	// following values:
	//
	// * ENABLED – IoT SiteWise accepts time series that aren't
	// associated with asset properties. After the disassociatedDataStorage is enabled,
	// you can't disable it.
	//
	// * DISABLED – IoT SiteWise doesn't accept time series
	// (data streams) that aren't associated with asset properties.
	//
	// For more
	// information, see Data streams
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/data-streams.html) in
	// the IoT SiteWise User Guide.
	DisassociatedDataStorage types.DisassociatedDataStorageState

	// The date the storage configuration was last updated, in Unix epoch time.
	LastUpdateDate *time.Time

	// Contains information about the storage destination.
	MultiLayerStorage *types.MultiLayerStorage

	// How many days your data is kept in the hot tier. By default, your data is kept
	// indefinitely in the hot tier.
	RetentionPeriod *types.RetentionPeriod

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStorageConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeStorageConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeStorageConfiguration{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribeStorageConfigurationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStorageConfiguration(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeStorageConfigurationMiddleware struct {
}

func (*endpointPrefix_opDescribeStorageConfigurationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeStorageConfigurationMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeStorageConfigurationMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeStorageConfigurationMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStorageConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeStorageConfiguration",
	}
}
