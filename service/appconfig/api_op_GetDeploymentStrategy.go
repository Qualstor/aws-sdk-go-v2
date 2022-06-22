// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about a deployment strategy. A deployment strategy defines
// important criteria for rolling out your configuration to the designated targets.
// A deployment strategy includes the overall duration required, a percentage of
// targets to receive the deployment during each interval, an algorithm that
// defines how percentage grows, and bake time.
func (c *Client) GetDeploymentStrategy(ctx context.Context, params *GetDeploymentStrategyInput, optFns ...func(*Options)) (*GetDeploymentStrategyOutput, error) {
	if params == nil {
		params = &GetDeploymentStrategyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeploymentStrategy", params, optFns, c.addOperationGetDeploymentStrategyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeploymentStrategyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeploymentStrategyInput struct {

	// The ID of the deployment strategy to get.
	//
	// This member is required.
	DeploymentStrategyId *string

	noSmithyDocumentSerde
}

type GetDeploymentStrategyOutput struct {

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes int32

	// The description of the deployment strategy.
	Description *string

	// The amount of time that AppConfig monitored for alarms before considering the
	// deployment to be complete and no longer eligible for automatic rollback.
	FinalBakeTimeInMinutes int32

	// The percentage of targets that received a deployed configuration during each
	// interval.
	GrowthFactor float32

	// The algorithm used to define how percentage grew over time.
	GrowthType types.GrowthType

	// The deployment strategy ID.
	Id *string

	// The name of the deployment strategy.
	Name *string

	// Save the deployment strategy to a Systems Manager (SSM) document.
	ReplicateTo types.ReplicateTo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDeploymentStrategyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDeploymentStrategy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDeploymentStrategy{}, middleware.After)
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
	if err = addOpGetDeploymentStrategyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeploymentStrategy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDeploymentStrategy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "GetDeploymentStrategy",
	}
}
