// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Submits a request to change the health status of a custom health check to
// healthy or unhealthy. You can use UpdateInstanceCustomHealthStatus to change the
// status only for custom health checks, which you define using
// HealthCheckCustomConfig when you create a service. You can't use it to change
// the status for Route 53 health checks, which you define using HealthCheckConfig.
// For more information, see HealthCheckCustomConfig
// (https://docs.aws.amazon.com/cloud-map/latest/api/API_HealthCheckCustomConfig.html).
func (c *Client) UpdateInstanceCustomHealthStatus(ctx context.Context, params *UpdateInstanceCustomHealthStatusInput, optFns ...func(*Options)) (*UpdateInstanceCustomHealthStatusOutput, error) {
	if params == nil {
		params = &UpdateInstanceCustomHealthStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInstanceCustomHealthStatus", params, optFns, c.addOperationUpdateInstanceCustomHealthStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInstanceCustomHealthStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateInstanceCustomHealthStatusInput struct {

	// The ID of the instance that you want to change the health status for.
	//
	// This member is required.
	InstanceId *string

	// The ID of the service that includes the configuration for the custom health
	// check that you want to change the status for.
	//
	// This member is required.
	ServiceId *string

	// The new status of the instance, HEALTHY or UNHEALTHY.
	//
	// This member is required.
	Status types.CustomHealthStatus

	noSmithyDocumentSerde
}

type UpdateInstanceCustomHealthStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateInstanceCustomHealthStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateInstanceCustomHealthStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateInstanceCustomHealthStatus{}, middleware.After)
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
	if err = addOpUpdateInstanceCustomHealthStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInstanceCustomHealthStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateInstanceCustomHealthStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "UpdateInstanceCustomHealthStatus",
	}
}
