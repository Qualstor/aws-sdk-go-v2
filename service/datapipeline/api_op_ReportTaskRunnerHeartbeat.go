// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Task runners call ReportTaskRunnerHeartbeat every 15 minutes to indicate that
// they are operational. If the AWS Data Pipeline Task Runner is launched on a
// resource managed by AWS Data Pipeline, the web service can use this call to
// detect when the task runner application has failed and restart a new instance.
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.ReportTaskRunnerHeartbeat Content-Length: 84 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams {"taskrunnerId": "1234567890", "workerGroup":
// "wg-12345", "hostname": "example.com"} Status: x-amzn-RequestId:
// b3104dc5-0734-11e2-af6f-6bc7a6be60d9 Content-Type: application/x-amz-json-1.1
// Content-Length: 20 Date: Mon, 12 Nov 2012 17:50:53 GMT {"terminate": false}
func (c *Client) ReportTaskRunnerHeartbeat(ctx context.Context, params *ReportTaskRunnerHeartbeatInput, optFns ...func(*Options)) (*ReportTaskRunnerHeartbeatOutput, error) {
	if params == nil {
		params = &ReportTaskRunnerHeartbeatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReportTaskRunnerHeartbeat", params, optFns, c.addOperationReportTaskRunnerHeartbeatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReportTaskRunnerHeartbeatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ReportTaskRunnerHeartbeat.
type ReportTaskRunnerHeartbeatInput struct {

	// The ID of the task runner. This value should be unique across your AWS account.
	// In the case of AWS Data Pipeline Task Runner launched on a resource managed by
	// AWS Data Pipeline, the web service provides a unique identifier when it launches
	// the application. If you have written a custom task runner, you should assign a
	// unique identifier for the task runner.
	//
	// This member is required.
	TaskrunnerId *string

	// The public DNS name of the task runner.
	Hostname *string

	// The type of task the task runner is configured to accept and process. The worker
	// group is set as a field on objects in the pipeline when they are created. You
	// can only specify a single value for workerGroup. There are no wildcard values
	// permitted in workerGroup; the string must be an exact, case-sensitive, match.
	WorkerGroup *string

	noSmithyDocumentSerde
}

// Contains the output of ReportTaskRunnerHeartbeat.
type ReportTaskRunnerHeartbeatOutput struct {

	// Indicates whether the calling task runner should terminate.
	//
	// This member is required.
	Terminate bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReportTaskRunnerHeartbeatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpReportTaskRunnerHeartbeat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpReportTaskRunnerHeartbeat{}, middleware.After)
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
	if err = addOpReportTaskRunnerHeartbeatValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReportTaskRunnerHeartbeat(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReportTaskRunnerHeartbeat(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "ReportTaskRunnerHeartbeat",
	}
}
