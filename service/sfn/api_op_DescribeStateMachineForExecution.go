// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the state machine associated with a specific execution. This operation
// is eventually consistent. The results are best effort and may not reflect very
// recent updates and changes. This API action is not supported by EXPRESS state
// machines.
func (c *Client) DescribeStateMachineForExecution(ctx context.Context, params *DescribeStateMachineForExecutionInput, optFns ...func(*Options)) (*DescribeStateMachineForExecutionOutput, error) {
	if params == nil {
		params = &DescribeStateMachineForExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStateMachineForExecution", params, optFns, c.addOperationDescribeStateMachineForExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStateMachineForExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStateMachineForExecutionInput struct {

	// The Amazon Resource Name (ARN) of the execution you want state machine
	// information for.
	//
	// This member is required.
	ExecutionArn *string

	noSmithyDocumentSerde
}

type DescribeStateMachineForExecutionOutput struct {

	// The Amazon States Language definition of the state machine. See Amazon States
	// Language
	// (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).
	//
	// This member is required.
	Definition *string

	// The name of the state machine associated with the execution.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM role of the State Machine for the
	// execution.
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the state machine associated with the
	// execution.
	//
	// This member is required.
	StateMachineArn *string

	// The date and time the state machine associated with an execution was updated.
	// For a newly created state machine, this is the creation date.
	//
	// This member is required.
	UpdateDate *time.Time

	// The LoggingConfiguration data type is used to set CloudWatch Logs options.
	LoggingConfiguration *types.LoggingConfiguration

	// Selects whether AWS X-Ray tracing is enabled.
	TracingConfiguration *types.TracingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStateMachineForExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeStateMachineForExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeStateMachineForExecution{}, middleware.After)
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
	if err = addOpDescribeStateMachineForExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStateMachineForExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStateMachineForExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "DescribeStateMachineForExecution",
	}
}
