// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels a matchmaking ticket or match backfill ticket that is currently being
// processed. To stop the matchmaking operation, specify the ticket ID. If
// successful, work on the ticket is stopped, and the ticket status is changed to
// CANCELLED. This call is also used to turn off automatic backfill for an
// individual game session. This is for game sessions that are created with a
// matchmaking configuration that has automatic backfill enabled. The ticket ID is
// included in the MatchmakerData of an updated game session object, which is
// provided to the game server. If the operation is successful, the service sends
// back an empty JSON struct with the HTTP 200 response (not an empty HTTP body).
// Learn more  Add FlexMatch to a game client
// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-client.html)
// Related actions StartMatchmaking | DescribeMatchmaking | StopMatchmaking |
// AcceptMatch | StartMatchBackfill | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) StopMatchmaking(ctx context.Context, params *StopMatchmakingInput, optFns ...func(*Options)) (*StopMatchmakingOutput, error) {
	if params == nil {
		params = &StopMatchmakingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopMatchmaking", params, optFns, c.addOperationStopMatchmakingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopMatchmakingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type StopMatchmakingInput struct {

	// A unique identifier for a matchmaking ticket.
	//
	// This member is required.
	TicketId *string

	noSmithyDocumentSerde
}

type StopMatchmakingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopMatchmakingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopMatchmaking{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopMatchmaking{}, middleware.After)
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
	if err = addOpStopMatchmakingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopMatchmaking(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopMatchmaking(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "StopMatchmaking",
	}
}
