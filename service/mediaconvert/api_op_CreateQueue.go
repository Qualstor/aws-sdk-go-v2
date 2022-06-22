// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new transcoding queue. For information about queues, see Working With
// Queues in the User Guide at
// https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-queues.html
func (c *Client) CreateQueue(ctx context.Context, params *CreateQueueInput, optFns ...func(*Options)) (*CreateQueueOutput, error) {
	if params == nil {
		params = &CreateQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateQueue", params, optFns, c.addOperationCreateQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateQueueInput struct {

	// The name of the queue that you are creating.
	//
	// This member is required.
	Name *string

	// Optional. A description of the queue that you are creating.
	Description *string

	// Specifies whether the pricing plan for the queue is on-demand or reserved. For
	// on-demand, you pay per minute, billed in increments of .01 minute. For reserved,
	// you pay for the transcoding capacity of the entire queue, regardless of how much
	// or how little you use it. Reserved pricing requires a 12-month commitment. When
	// you use the API to create a queue, the default is on-demand.
	PricingPlan types.PricingPlan

	// Details about the pricing plan for your reserved queue. Required for reserved
	// queues and not applicable to on-demand queues.
	ReservationPlanSettings *types.ReservationPlanSettings

	// Initial state of the queue. If you create a paused queue, then jobs in that
	// queue won't begin.
	Status types.QueueStatus

	// The tags that you want to add to the resource. You can tag resources with a
	// key-value pair or with only a key.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateQueueOutput struct {

	// You can use queues to manage the resources that are available to your AWS
	// account for running multiple transcoding jobs at the same time. If you don't
	// specify a queue, the service sends all jobs through the default queue. For more
	// information, see
	// https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-queues.html.
	Queue *types.Queue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateQueue{}, middleware.After)
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
	if err = addOpCreateQueueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateQueue(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateQueue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "CreateQueue",
	}
}
