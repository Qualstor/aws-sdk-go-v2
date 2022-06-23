// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an on-demand replication run for the specified replication job. This
// replication run starts immediately. This replication run is in addition to the
// ones already scheduled. There is a limit on the number of on-demand replications
// runs that you can request in a 24-hour period.
func (c *Client) StartOnDemandReplicationRun(ctx context.Context, params *StartOnDemandReplicationRunInput, optFns ...func(*Options)) (*StartOnDemandReplicationRunOutput, error) {
	if params == nil {
		params = &StartOnDemandReplicationRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartOnDemandReplicationRun", params, optFns, c.addOperationStartOnDemandReplicationRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartOnDemandReplicationRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartOnDemandReplicationRunInput struct {

	// The ID of the replication job.
	//
	// This member is required.
	ReplicationJobId *string

	// The description of the replication run.
	Description *string

	noSmithyDocumentSerde
}

type StartOnDemandReplicationRunOutput struct {

	// The ID of the replication run.
	ReplicationRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartOnDemandReplicationRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartOnDemandReplicationRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartOnDemandReplicationRun{}, middleware.After)
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
	if err = addOpStartOnDemandReplicationRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartOnDemandReplicationRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartOnDemandReplicationRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "StartOnDemandReplicationRun",
	}
}
