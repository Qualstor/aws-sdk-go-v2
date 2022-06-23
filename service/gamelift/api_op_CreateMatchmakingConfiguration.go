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

// Defines a new matchmaking configuration for use with FlexMatch. Whether your are
// using FlexMatch with GameLift hosting or as a standalone matchmaking service,
// the matchmaking configuration sets out rules for matching players and forming
// teams. If you're also using GameLift hosting, it defines how to start game
// sessions for each match. Your matchmaking system can use multiple configurations
// to handle different game scenarios. All matchmaking requests (StartMatchmaking
// or StartMatchBackfill) identify the matchmaking configuration to use and provide
// player attributes consistent with that configuration. To create a matchmaking
// configuration, you must provide the following: configuration name and FlexMatch
// mode (with or without GameLift hosting); a rule set that specifies how to
// evaluate players and find acceptable matches; whether player acceptance is
// required; and the maximum time allowed for a matchmaking attempt. When using
// FlexMatch with GameLift hosting, you also need to identify the game session
// queue to use when starting a game session for the match. In addition, you must
// set up an Amazon Simple Notification Service topic to receive matchmaking
// notifications. Provide the topic ARN in the matchmaking configuration. An
// alternative method, continuously polling ticket status with DescribeMatchmaking,
// is only suitable for games in development with low matchmaking usage. Learn more
// Design a FlexMatch matchmaker
// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-configuration.html)
// Set up FlexMatch event notification
// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html)
// Related actions CreateMatchmakingConfiguration |
// DescribeMatchmakingConfigurations | UpdateMatchmakingConfiguration |
// DeleteMatchmakingConfiguration | CreateMatchmakingRuleSet |
// DescribeMatchmakingRuleSets | ValidateMatchmakingRuleSet |
// DeleteMatchmakingRuleSet | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) CreateMatchmakingConfiguration(ctx context.Context, params *CreateMatchmakingConfigurationInput, optFns ...func(*Options)) (*CreateMatchmakingConfigurationOutput, error) {
	if params == nil {
		params = &CreateMatchmakingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMatchmakingConfiguration", params, optFns, c.addOperationCreateMatchmakingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMatchmakingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type CreateMatchmakingConfigurationInput struct {

	// A flag that determines whether a match that was created with this configuration
	// must be accepted by the matched players. To require acceptance, set to TRUE.
	// With this option enabled, matchmaking tickets use the status REQUIRES_ACCEPTANCE
	// to indicate when a completed potential match is waiting for player acceptance.
	//
	// This member is required.
	AcceptanceRequired *bool

	// A unique identifier for the matchmaking configuration. This name is used to
	// identify the configuration associated with a matchmaking request or ticket.
	//
	// This member is required.
	Name *string

	// The maximum duration, in seconds, that a matchmaking ticket can remain in
	// process before timing out. Requests that fail due to timing out can be
	// resubmitted as needed.
	//
	// This member is required.
	RequestTimeoutSeconds *int32

	// A unique identifier for the matchmaking rule set to use with this configuration.
	// You can use either the rule set name or ARN value. A matchmaking configuration
	// can only use rule sets that are defined in the same Region.
	//
	// This member is required.
	RuleSetName *string

	// The length of time (in seconds) to wait for players to accept a proposed match,
	// if acceptance is required.
	AcceptanceTimeoutSeconds *int32

	// The number of player slots in a match to keep open for future players. For
	// example, if the configuration's rule set specifies a match for a single
	// 12-person team, and the additional player count is set to 2, only 10 players are
	// selected for the match. This parameter is not used if FlexMatchMode is set to
	// STANDALONE.
	AdditionalPlayerCount *int32

	// The method used to backfill game sessions that are created with this matchmaking
	// configuration. Specify MANUAL when your game manages backfill requests manually
	// or does not use the match backfill feature. Specify AUTOMATIC to have GameLift
	// create a StartMatchBackfill request whenever a game session has one or more open
	// slots. Learn more about manual and automatic backfill in  Backfill Existing
	// Games with FlexMatch
	// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-backfill.html).
	// Automatic backfill is not available when FlexMatchMode is set to STANDALONE.
	BackfillMode types.BackfillMode

	// Information to be added to all events related to this matchmaking configuration.
	CustomEventData *string

	// A human-readable description of the matchmaking configuration.
	Description *string

	// Indicates whether this matchmaking configuration is being used with GameLift
	// hosting or as a standalone matchmaking solution.
	//
	// * STANDALONE - FlexMatch forms
	// matches and returns match information, including players and team assignments,
	// in a  MatchmakingSucceeded
	// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-events.html#match-events-matchmakingsucceeded)
	// event.
	//
	// * WITH_QUEUE - FlexMatch forms matches and uses the specified GameLift
	// queue to start a game session for the match.
	FlexMatchMode types.FlexMatchMode

	// A set of custom properties for a game session, formatted as key:value pairs.
	// These properties are passed to a game server process in the GameSession object
	// with a request to start a new game session (see Start a Game Session
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	// This information is added to the new GameSession object that is created for a
	// successful match. This parameter is not used if FlexMatchMode is set to
	// STANDALONE.
	GameProperties []types.GameProperty

	// A set of custom game session properties, formatted as a single string value.
	// This data is passed to a game server process in the GameSession object with a
	// request to start a new game session (see Start a Game Session
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	// This information is added to the new GameSession object that is created for a
	// successful match. This parameter is not used if FlexMatchMode is set to
	// STANDALONE.
	GameSessionData *string

	// The Amazon Resource Name (ARN
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is
	// assigned to a GameLift game session queue resource and uniquely identifies it.
	// ARNs are unique across all Regions. Format is
	// arn:aws:gamelift:::gamesessionqueue/. Queues can be located in any Region.
	// Queues are used to start new GameLift-hosted game sessions for matches that are
	// created with this matchmaking configuration. If FlexMatchMode is set to
	// STANDALONE, do not set this parameter.
	GameSessionQueueArns []string

	// An SNS topic ARN that is set up to receive matchmaking notifications. See
	// Setting up notifications for matchmaking
	// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html)
	// for more information.
	NotificationTarget *string

	// A list of labels to assign to the new matchmaking configuration resource. Tags
	// are developer-defined key-value pairs. Tagging Amazon Web Services resources are
	// useful for resource management, access management and cost allocation. For more
	// information, see  Tagging Amazon Web Services Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the Amazon
	// Web Services General Reference. Once the resource is created, you can use
	// TagResource, UntagResource, and ListTagsForResource to add, remove, and view
	// tags. The maximum tag limit may be lower than stated. See the Amazon Web
	// Services General Reference for actual tagging limits.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Represents the returned data in response to a request operation.
type CreateMatchmakingConfigurationOutput struct {

	// Object that describes the newly created matchmaking configuration.
	Configuration *types.MatchmakingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMatchmakingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMatchmakingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMatchmakingConfiguration{}, middleware.After)
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
	if err = addOpCreateMatchmakingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMatchmakingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMatchmakingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateMatchmakingConfiguration",
	}
}
