// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts the specified replay. Events are not necessarily replayed in the exact
// same order that they were added to the archive. A replay processes events to
// replay based on the time in the event, and replays them using 1 minute
// intervals. If you specify an EventStartTime and an EventEndTime that covers a 20
// minute time range, the events are replayed from the first minute of that 20
// minute range first. Then the events from the second minute are replayed. You can
// use DescribeReplay to determine the progress of a replay. The value returned for
// EventLastReplayedTime indicates the time within the specified time range
// associated with the last event replayed.
func (c *Client) StartReplay(ctx context.Context, params *StartReplayInput, optFns ...func(*Options)) (*StartReplayOutput, error) {
	if params == nil {
		params = &StartReplayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartReplay", params, optFns, c.addOperationStartReplayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartReplayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartReplayInput struct {

	// A ReplayDestination object that includes details about the destination for the
	// replay.
	//
	// This member is required.
	Destination *types.ReplayDestination

	// A time stamp for the time to stop replaying events. Only events that occurred
	// between the EventStartTime and EventEndTime are replayed.
	//
	// This member is required.
	EventEndTime *time.Time

	// The ARN of the archive to replay events from.
	//
	// This member is required.
	EventSourceArn *string

	// A time stamp for the time to start replaying events. Only events that occurred
	// between the EventStartTime and EventEndTime are replayed.
	//
	// This member is required.
	EventStartTime *time.Time

	// The name of the replay to start.
	//
	// This member is required.
	ReplayName *string

	// A description for the replay to start.
	Description *string

	noSmithyDocumentSerde
}

type StartReplayOutput struct {

	// The ARN of the replay.
	ReplayArn *string

	// The time at which the replay started.
	ReplayStartTime *time.Time

	// The state of the replay.
	State types.ReplayState

	// The reason that the replay is in the state.
	StateReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartReplayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartReplay{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartReplay{}, middleware.After)
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
	if err = addOpStartReplayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartReplay(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartReplay(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "StartReplay",
	}
}
