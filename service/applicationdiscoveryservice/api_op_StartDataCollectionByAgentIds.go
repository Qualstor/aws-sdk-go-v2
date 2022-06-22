// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Instructs the specified agents or connectors to start collecting data.
func (c *Client) StartDataCollectionByAgentIds(ctx context.Context, params *StartDataCollectionByAgentIdsInput, optFns ...func(*Options)) (*StartDataCollectionByAgentIdsOutput, error) {
	if params == nil {
		params = &StartDataCollectionByAgentIdsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDataCollectionByAgentIds", params, optFns, c.addOperationStartDataCollectionByAgentIdsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDataCollectionByAgentIdsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDataCollectionByAgentIdsInput struct {

	// The IDs of the agents or connectors from which to start collecting data. If you
	// send a request to an agent/connector ID that you do not have permission to
	// contact, according to your Amazon Web Services account, the service does not
	// throw an exception. Instead, it returns the error in the Description field. If
	// you send a request to multiple agents/connectors and you do not have permission
	// to contact some of those agents/connectors, the system does not throw an
	// exception. Instead, the system shows Failed in the Description field.
	//
	// This member is required.
	AgentIds []string

	noSmithyDocumentSerde
}

type StartDataCollectionByAgentIdsOutput struct {

	// Information about agents or the connector that were instructed to start
	// collecting data. Information includes the agent/connector ID, a description of
	// the operation performed, and whether the agent/connector configuration was
	// updated.
	AgentsConfigurationStatus []types.AgentConfigurationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDataCollectionByAgentIdsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartDataCollectionByAgentIds{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartDataCollectionByAgentIds{}, middleware.After)
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
	if err = addOpStartDataCollectionByAgentIdsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDataCollectionByAgentIds(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDataCollectionByAgentIds(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "StartDataCollectionByAgentIds",
	}
}
