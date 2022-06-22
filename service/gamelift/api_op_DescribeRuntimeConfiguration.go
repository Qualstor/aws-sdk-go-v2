// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a fleet's runtime configuration settings. The runtime configuration
// tells GameLift which server processes to run (and how) on each instance in the
// fleet. To get the runtime configuration that is currently in forces for a fleet,
// provide the fleet ID. If successful, a RuntimeConfiguration object is returned
// for the requested fleet. If the requested fleet has been deleted, the result set
// is empty. Learn more Setting up GameLift fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)Running
// multiple processes on a fleet
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-multiprocess.html)
// Related actions ListFleets | DescribeEC2InstanceLimits | DescribeFleetAttributes
// | DescribeFleetCapacity | DescribeFleetEvents | DescribeFleetLocationAttributes
// | DescribeFleetPortSettings | DescribeFleetUtilization |
// DescribeRuntimeConfiguration | DescribeScalingPolicies | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DescribeRuntimeConfiguration(ctx context.Context, params *DescribeRuntimeConfigurationInput, optFns ...func(*Options)) (*DescribeRuntimeConfigurationOutput, error) {
	if params == nil {
		params = &DescribeRuntimeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRuntimeConfiguration", params, optFns, c.addOperationDescribeRuntimeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRuntimeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeRuntimeConfigurationInput struct {

	// A unique identifier for the fleet to get the runtime configuration for. You can
	// use either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	noSmithyDocumentSerde
}

// Represents the returned data in response to a request operation.
type DescribeRuntimeConfigurationOutput struct {

	// Instructions that describe how server processes should be launched and
	// maintained on each instance in the fleet.
	RuntimeConfiguration *types.RuntimeConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRuntimeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRuntimeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRuntimeConfiguration{}, middleware.After)
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
	if err = addOpDescribeRuntimeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRuntimeConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRuntimeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeRuntimeConfiguration",
	}
}
