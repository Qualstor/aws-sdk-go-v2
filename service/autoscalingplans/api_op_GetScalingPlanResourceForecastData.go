// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscalingplans

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/autoscalingplans/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the forecast data for a scalable resource. Capacity forecasts are
// represented as predicted values, or data points, that are calculated using
// historical data points from a specified CloudWatch load metric. Data points are
// available for up to 56 days.
func (c *Client) GetScalingPlanResourceForecastData(ctx context.Context, params *GetScalingPlanResourceForecastDataInput, optFns ...func(*Options)) (*GetScalingPlanResourceForecastDataOutput, error) {
	if params == nil {
		params = &GetScalingPlanResourceForecastDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetScalingPlanResourceForecastData", params, optFns, c.addOperationGetScalingPlanResourceForecastDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetScalingPlanResourceForecastDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetScalingPlanResourceForecastDataInput struct {

	// The exclusive end time of the time range for the forecast data to get. The
	// maximum time duration between the start and end time is seven days. Although
	// this parameter can accept a date and time that is more than two days in the
	// future, the availability of forecast data has limits. AWS Auto Scaling only
	// issues forecasts for periods of two days in advance.
	//
	// This member is required.
	EndTime *time.Time

	// The type of forecast data to get.
	//
	// * LoadForecast: The load metric forecast.
	//
	// *
	// CapacityForecast: The capacity forecast.
	//
	// * ScheduledActionMinCapacity: The
	// minimum capacity for each scheduled scaling action. This data is calculated as
	// the larger of two values: the capacity forecast or the minimum capacity in the
	// scaling instruction.
	//
	// * ScheduledActionMaxCapacity: The maximum capacity for
	// each scheduled scaling action. The calculation used is determined by the
	// predictive scaling maximum capacity behavior setting in the scaling instruction.
	//
	// This member is required.
	ForecastDataType types.ForecastDataType

	// The ID of the resource. This string consists of a prefix (autoScalingGroup)
	// followed by the name of a specified Auto Scaling group (my-asg). Example:
	// autoScalingGroup/my-asg.
	//
	// This member is required.
	ResourceId *string

	// The scalable dimension for the resource. The only valid value is
	// autoscaling:autoScalingGroup:DesiredCapacity.
	//
	// This member is required.
	ScalableDimension types.ScalableDimension

	// The name of the scaling plan.
	//
	// This member is required.
	ScalingPlanName *string

	// The version number of the scaling plan. Currently, the only valid value is 1.
	//
	// This member is required.
	ScalingPlanVersion *int64

	// The namespace of the AWS service. The only valid value is autoscaling.
	//
	// This member is required.
	ServiceNamespace types.ServiceNamespace

	// The inclusive start time of the time range for the forecast data to get. The
	// date and time can be at most 56 days before the current date and time.
	//
	// This member is required.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type GetScalingPlanResourceForecastDataOutput struct {

	// The data points to return.
	//
	// This member is required.
	Datapoints []types.Datapoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetScalingPlanResourceForecastDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetScalingPlanResourceForecastData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetScalingPlanResourceForecastData{}, middleware.After)
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
	if err = addOpGetScalingPlanResourceForecastDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetScalingPlanResourceForecastData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetScalingPlanResourceForecastData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling-plans",
		OperationName: "GetScalingPlanResourceForecastData",
	}
}
