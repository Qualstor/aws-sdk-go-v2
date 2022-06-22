// Code generated by smithy-go-codegen DO NOT EDIT.

package iotevents

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/iotevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an alarm model to monitor an AWS IoT Events input attribute. You can use
// the alarm to get notified when the value is outside a specified range. For more
// information, see Create an alarm model
// (https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html)
// in the AWS IoT Events Developer Guide.
func (c *Client) CreateAlarmModel(ctx context.Context, params *CreateAlarmModelInput, optFns ...func(*Options)) (*CreateAlarmModelOutput, error) {
	if params == nil {
		params = &CreateAlarmModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAlarmModel", params, optFns, c.addOperationCreateAlarmModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAlarmModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAlarmModelInput struct {

	// A unique name that helps you identify the alarm model. You can't change this
	// name after you create the alarm model.
	//
	// This member is required.
	AlarmModelName *string

	// Defines when your alarm is invoked.
	//
	// This member is required.
	AlarmRule *types.AlarmRule

	// The ARN of the IAM role that allows the alarm to perform actions and access AWS
	// resources. For more information, see Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	RoleArn *string

	// Contains the configuration information of alarm state changes.
	AlarmCapabilities *types.AlarmCapabilities

	// Contains information about one or more alarm actions.
	AlarmEventActions *types.AlarmEventActions

	// A description that tells you what the alarm model detects.
	AlarmModelDescription *string

	// Contains information about one or more notification actions.
	AlarmNotification *types.AlarmNotification

	// An input attribute used as a key to create an alarm. AWS IoT Events routes
	// inputs
	// (https://docs.aws.amazon.com/iotevents/latest/apireference/API_Input.html)
	// associated with this key to the alarm.
	Key *string

	// A non-negative integer that reflects the severity level of the alarm.
	Severity *int32

	// A list of key-value pairs that contain metadata for the alarm model. The tags
	// help you manage the alarm model. For more information, see Tagging your AWS IoT
	// Events resources
	// (https://docs.aws.amazon.com/iotevents/latest/developerguide/tagging-iotevents.html)
	// in the AWS IoT Events Developer Guide. You can create up to 50 tags for one
	// alarm model.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAlarmModelOutput struct {

	// The ARN of the alarm model. For more information, see Amazon Resource Names
	// (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	AlarmModelArn *string

	// The version of the alarm model.
	AlarmModelVersion *string

	// The time the alarm model was created, in the Unix epoch format.
	CreationTime *time.Time

	// The time the alarm model was last updated, in the Unix epoch format.
	LastUpdateTime *time.Time

	// The status of the alarm model. The status can be one of the following values:
	//
	// *
	// ACTIVE - The alarm model is active and it's ready to evaluate data.
	//
	// *
	// ACTIVATING - AWS IoT Events is activating your alarm model. Activating an alarm
	// model can take up to a few minutes.
	//
	// * INACTIVE - The alarm model is inactive,
	// so it isn't ready to evaluate data. Check your alarm model information and
	// update the alarm model.
	//
	// * FAILED - You couldn't create or update the alarm
	// model. Check your alarm model information and try again.
	Status types.AlarmModelVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAlarmModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAlarmModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAlarmModel{}, middleware.After)
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
	if err = addOpCreateAlarmModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAlarmModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAlarmModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotevents",
		OperationName: "CreateAlarmModel",
	}
}
