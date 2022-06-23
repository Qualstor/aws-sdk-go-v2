// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is used with the GameLift FleetIQ solution and game server
// groups. Reinstates activity on a game server group after it has been suspended.
// A game server group might be suspended by theSuspendGameServerGroup operation,
// or it might be suspended involuntarily due to a configuration problem. In the
// second case, you can manually resume activity on the group once the
// configuration problem has been resolved. Refer to the game server group status
// and status reason for more information on why group activity is suspended. To
// resume activity, specify a game server group ARN and the type of activity to be
// resumed. If successful, a GameServerGroup object is returned showing that the
// resumed activity is no longer listed in SuspendedActions. Learn more GameLift
// FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
// Related actions CreateGameServerGroup | ListGameServerGroups |
// DescribeGameServerGroup | UpdateGameServerGroup | DeleteGameServerGroup |
// ResumeGameServerGroup | SuspendGameServerGroup | DescribeGameServerInstances |
// All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/reference-awssdk-fleetiq.html)
func (c *Client) ResumeGameServerGroup(ctx context.Context, params *ResumeGameServerGroupInput, optFns ...func(*Options)) (*ResumeGameServerGroupOutput, error) {
	if params == nil {
		params = &ResumeGameServerGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResumeGameServerGroup", params, optFns, c.addOperationResumeGameServerGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResumeGameServerGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResumeGameServerGroupInput struct {

	// A unique identifier for the game server group. Use either the GameServerGroup
	// name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The activity to resume for this game server group.
	//
	// This member is required.
	ResumeActions []types.GameServerGroupAction

	noSmithyDocumentSerde
}

type ResumeGameServerGroupOutput struct {

	// An object that describes the game server group resource, with the
	// SuspendedActions property updated to reflect the resumed activity.
	GameServerGroup *types.GameServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResumeGameServerGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResumeGameServerGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResumeGameServerGroup{}, middleware.After)
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
	if err = addOpResumeGameServerGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResumeGameServerGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResumeGameServerGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ResumeGameServerGroup",
	}
}
