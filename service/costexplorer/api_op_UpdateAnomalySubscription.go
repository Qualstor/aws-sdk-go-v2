// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing cost anomaly monitor subscription.
func (c *Client) UpdateAnomalySubscription(ctx context.Context, params *UpdateAnomalySubscriptionInput, optFns ...func(*Options)) (*UpdateAnomalySubscriptionOutput, error) {
	if params == nil {
		params = &UpdateAnomalySubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAnomalySubscription", params, optFns, c.addOperationUpdateAnomalySubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAnomalySubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAnomalySubscriptionInput struct {

	// A cost anomaly subscription Amazon Resource Name (ARN).
	//
	// This member is required.
	SubscriptionArn *string

	// The update to the frequency value that subscribers receive notifications.
	Frequency types.AnomalySubscriptionFrequency

	// A list of cost anomaly monitor ARNs.
	MonitorArnList []string

	// The update to the subscriber list.
	Subscribers []types.Subscriber

	// The new name of the subscription.
	SubscriptionName *string

	// The update to the threshold value for receiving notifications.
	Threshold *float64

	noSmithyDocumentSerde
}

type UpdateAnomalySubscriptionOutput struct {

	// A cost anomaly subscription ARN.
	//
	// This member is required.
	SubscriptionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAnomalySubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAnomalySubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAnomalySubscription{}, middleware.After)
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
	if err = addOpUpdateAnomalySubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAnomalySubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAnomalySubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "UpdateAnomalySubscription",
	}
}
