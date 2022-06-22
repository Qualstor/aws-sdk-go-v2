// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides metrics on the accuracy of the models that were trained by the
// CreatePredictor operation. Use metrics to see how well the model performed and
// to decide whether to use the predictor to generate a forecast. For more
// information, see Predictor Metrics
// (https://docs.aws.amazon.com/forecast/latest/dg/metrics.html). This operation
// generates metrics for each backtest window that was evaluated. The number of
// backtest windows (NumberOfBacktestWindows) is specified using the
// EvaluationParameters object, which is optionally included in the CreatePredictor
// request. If NumberOfBacktestWindows isn't specified, the number defaults to one.
// The parameters of the filling method determine which items contribute to the
// metrics. If you want all items to contribute, specify zero. If you want only
// those items that have complete data in the range being evaluated to contribute,
// specify nan. For more information, see FeaturizationMethod. Before you can get
// accuracy metrics, the Status of the predictor must be ACTIVE, signifying that
// training has completed. To get the status, use the DescribePredictor operation.
func (c *Client) GetAccuracyMetrics(ctx context.Context, params *GetAccuracyMetricsInput, optFns ...func(*Options)) (*GetAccuracyMetricsOutput, error) {
	if params == nil {
		params = &GetAccuracyMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccuracyMetrics", params, optFns, c.addOperationGetAccuracyMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccuracyMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccuracyMetricsInput struct {

	// The Amazon Resource Name (ARN) of the predictor to get metrics for.
	//
	// This member is required.
	PredictorArn *string

	noSmithyDocumentSerde
}

type GetAccuracyMetricsOutput struct {

	// The LatencyOptimized AutoML override strategy is only available in private beta.
	// Contact AWS Support or your account manager to learn more about access
	// privileges. The AutoML strategy used to train the predictor. Unless
	// LatencyOptimized is specified, the AutoML strategy optimizes predictor accuracy.
	// This parameter is only valid for predictors trained using AutoML.
	AutoMLOverrideStrategy types.AutoMLOverrideStrategy

	// Whether the predictor was created with CreateAutoPredictor.
	IsAutoPredictor *bool

	// The accuracy metric used to optimize the predictor.
	OptimizationMetric types.OptimizationMetric

	// An array of results from evaluating the predictor.
	PredictorEvaluationResults []types.EvaluationResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccuracyMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAccuracyMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAccuracyMetrics{}, middleware.After)
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
	if err = addOpGetAccuracyMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccuracyMetrics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAccuracyMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "GetAccuracyMetrics",
	}
}
