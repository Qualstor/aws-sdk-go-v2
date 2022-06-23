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

// Creates a task. A task includes a source location and a destination location,
// and a configuration that specifies how data is transferred. A task always
// transfers data from the source location to the destination location. The
// configuration specifies options such as task scheduling, bandwidth limits, etc.
// A task is the complete definition of a data transfer. When you create a task
// that transfers data between Amazon Web Services services in different Amazon Web
// Services Regions, one of the two locations that you specify must reside in the
// Region where DataSync is being used. The other location must be specified in a
// different Region. You can transfer data between commercial Amazon Web Services
// Regions except for China, or between Amazon Web Services GovCloud (US) Regions.
// When you use DataSync to copy files or objects between Amazon Web Services
// Regions, you pay for data transfer between Regions. This is billed as data
// transfer OUT from your source Region to your destination Region. For more
// information, see Data Transfer pricing
// (http://aws.amazon.com/ec2/pricing/on-demand/#Data_Transfer).
func (c *Client) CreateTask(ctx context.Context, params *CreateTaskInput, optFns ...func(*Options)) (*CreateTaskOutput, error) {
	if params == nil {
		params = &CreateTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTask", params, optFns, c.addOperationCreateTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateTaskRequest
type CreateTaskInput struct {

	// The Amazon Resource Name (ARN) of an Amazon Web Services storage resource's
	// location.
	//
	// This member is required.
	DestinationLocationArn *string

	// The Amazon Resource Name (ARN) of the source location for the task.
	//
	// This member is required.
	SourceLocationArn *string

	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that is used
	// to monitor and log events in the task.
	CloudWatchLogGroupArn *string

	// A list of filter rules that determines which files to exclude from a task. The
	// list should contain a single filter string that consists of the patterns to
	// exclude. The patterns are delimited by "|" (that is, a pipe), for example,
	// "/folder1|/folder2".
	Excludes []types.FilterRule

	// A list of filter rules that determines which files to include when running a
	// task. The pattern contains a single filter string that consists of the patterns
	// to include. The patterns are delimited by "|" (that is, a pipe), for example,
	// "/folder1|/folder2".
	Includes []types.FilterRule

	// The name of a task. This value is a text reference that is used to identify the
	// task in the console.
	Name *string

	// The set of configuration options that control the behavior of a single execution
	// of the task that occurs when you call StartTaskExecution. You can configure
	// these options to preserve metadata such as user ID (UID) and group ID (GID),
	// file permissions, data integrity verification, and so on. For each individual
	// task execution, you can override these options by specifying the OverrideOptions
	// before starting the task execution. For more information, see the
	// StartTaskExecution
	// (https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html)
	// operation.
	Options *types.Options

	// Specifies a schedule used to periodically transfer files from a source to a
	// destination location. The schedule should be specified in UTC time. For more
	// information, see Scheduling your task
	// (https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html).
	Schedule *types.TaskSchedule

	// The key-value pair that represents the tag that you want to add to the resource.
	// The value can be an empty string.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

// CreateTaskResponse
type CreateTaskOutput struct {

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTask{}, middleware.After)
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
	if err = addOpCreateTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateTask",
	}
}
