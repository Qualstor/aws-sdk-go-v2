// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Records a heartbeat for the lifecycle action associated with the specified token
// or instance. This extends the timeout by the length of time defined using the
// PutLifecycleHook API call. This step is a part of the procedure for adding a
// lifecycle hook to an Auto Scaling group:
//
// * (Optional) Create a launch template
// or launch configuration with a user data script that runs while an instance is
// in a wait state due to a lifecycle hook.
//
// * (Optional) Create a Lambda function
// and a rule that allows Amazon EventBridge to invoke your Lambda function when an
// instance is put into a wait state due to a lifecycle hook.
//
// * (Optional) Create
// a notification target and an IAM role. The target can be either an Amazon SQS
// queue or an Amazon SNS topic. The role allows Amazon EC2 Auto Scaling to publish
// lifecycle notifications to the target.
//
// * Create the lifecycle hook. Specify
// whether the hook is used when the instances launch or terminate.
//
// * If you need
// more time, record the lifecycle action heartbeat to keep the instance in a wait
// state.
//
// * If you finish before the timeout period ends, send a callback by using
// the CompleteLifecycleAction API call.
//
// For more information, see Amazon EC2 Auto
// Scaling lifecycle hooks
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html) in
// the Amazon EC2 Auto Scaling User Guide.
func (c *Client) RecordLifecycleActionHeartbeat(ctx context.Context, params *RecordLifecycleActionHeartbeatInput, optFns ...func(*Options)) (*RecordLifecycleActionHeartbeatOutput, error) {
	if params == nil {
		params = &RecordLifecycleActionHeartbeatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RecordLifecycleActionHeartbeat", params, optFns, c.addOperationRecordLifecycleActionHeartbeatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RecordLifecycleActionHeartbeatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RecordLifecycleActionHeartbeatInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The name of the lifecycle hook.
	//
	// This member is required.
	LifecycleHookName *string

	// The ID of the instance.
	InstanceId *string

	// A token that uniquely identifies a specific lifecycle action associated with an
	// instance. Amazon EC2 Auto Scaling sends this token to the notification target
	// that you specified when you created the lifecycle hook.
	LifecycleActionToken *string

	noSmithyDocumentSerde
}

type RecordLifecycleActionHeartbeatOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRecordLifecycleActionHeartbeatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRecordLifecycleActionHeartbeat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRecordLifecycleActionHeartbeat{}, middleware.After)
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
	if err = addOpRecordLifecycleActionHeartbeatValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRecordLifecycleActionHeartbeat(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRecordLifecycleActionHeartbeat(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "RecordLifecycleActionHeartbeat",
	}
}
