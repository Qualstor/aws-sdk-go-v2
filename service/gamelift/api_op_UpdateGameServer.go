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
// groups. Updates information about a registered game server to help GameLift
// FleetIQ to track game server availability. This operation is called by a game
// server process that is running on an instance in a game server group. Use this
// operation to update the following types of game server information. You can make
// all three types of updates in the same request:
//
// * To update the game server's
// utilization status, identify the game server and game server group and specify
// the current utilization status. Use this status to identify when game servers
// are currently hosting games and when they are available to be claimed.
//
// * To
// report health status, identify the game server and game server group and set
// health check to HEALTHY. If a game server does not report health status for a
// certain length of time, the game server is no longer considered healthy. As a
// result, it will be eventually deregistered from the game server group to avoid
// affecting utilization metrics. The best practice is to report health every 60
// seconds.
//
// * To change game server metadata, provide updated game server
// data.
//
// Once a game server is successfully updated, the relevant statuses and
// timestamps are updated. Learn more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
// Related actions RegisterGameServer | ListGameServers | ClaimGameServer |
// DescribeGameServer | UpdateGameServer | DeregisterGameServer | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/reference-awssdk-fleetiq.html)
func (c *Client) UpdateGameServer(ctx context.Context, params *UpdateGameServerInput, optFns ...func(*Options)) (*UpdateGameServerOutput, error) {
	if params == nil {
		params = &UpdateGameServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGameServer", params, optFns, c.addOperationUpdateGameServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGameServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGameServerInput struct {

	// A unique identifier for the game server group where the game server is running.
	// Use either the GameServerGroup name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// A custom string that uniquely identifies the game server to update.
	//
	// This member is required.
	GameServerId *string

	// A set of custom game server properties, formatted as a single string value. This
	// data is passed to a game client or service when it requests information on game
	// servers using ListGameServers or ClaimGameServer.
	GameServerData *string

	// Indicates health status of the game server. A request that includes this
	// parameter updates the game server's LastHealthCheckTime timestamp.
	HealthCheck types.GameServerHealthCheck

	// Indicates whether the game server is available or is currently hosting gameplay.
	UtilizationStatus types.GameServerUtilizationStatus

	noSmithyDocumentSerde
}

type UpdateGameServerOutput struct {

	// Object that describes the newly updated game server.
	GameServer *types.GameServer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGameServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateGameServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateGameServer{}, middleware.After)
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
	if err = addOpUpdateGameServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGameServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateGameServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateGameServer",
	}
}
