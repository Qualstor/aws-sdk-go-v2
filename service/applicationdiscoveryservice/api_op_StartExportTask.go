// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Begins the export of discovered data to an S3 bucket. If you specify agentIds in
// a filter, the task exports up to 72 hours of detailed data collected by the
// identified Application Discovery Agent, including network, process, and
// performance details. A time range for exported agent data may be set by using
// startTime and endTime. Export of detailed agent data is limited to five
// concurrently running exports. If you do not include an agentIds filter, summary
// data is exported that includes both Amazon Web Services Agentless Discovery
// Connector data and summary data from Amazon Web Services Discovery Agents.
// Export of summary data is limited to two exports per day.
func (c *Client) StartExportTask(ctx context.Context, params *StartExportTaskInput, optFns ...func(*Options)) (*StartExportTaskOutput, error) {
	if params == nil {
		params = &StartExportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartExportTask", params, optFns, c.addOperationStartExportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartExportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartExportTaskInput struct {

	// The end timestamp for exported data from the single Application Discovery Agent
	// selected in the filters. If no value is specified, exported data includes the
	// most recent data collected by the agent.
	EndTime *time.Time

	// The file format for the returned export data. Default value is CSV. Note: The
	// GRAPHML option has been deprecated.
	ExportDataFormat []types.ExportDataFormat

	// If a filter is present, it selects the single agentId of the Application
	// Discovery Agent for which data is exported. The agentId can be found in the
	// results of the DescribeAgents API or CLI. If no filter is present, startTime and
	// endTime are ignored and exported data includes both Agentless Discovery
	// Connector data and summary data from Application Discovery agents.
	Filters []types.ExportFilter

	// The start timestamp for exported data from the single Application Discovery
	// Agent selected in the filters. If no value is specified, data is exported
	// starting from the first data collected by the agent.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type StartExportTaskOutput struct {

	// A unique identifier used to query the status of an export request.
	ExportId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartExportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartExportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartExportTask{}, middleware.After)
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
	if err = addOpStartExportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartExportTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartExportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "StartExportTask",
	}
}
