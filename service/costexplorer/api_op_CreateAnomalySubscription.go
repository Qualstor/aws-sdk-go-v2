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

// Adds a subscription to a cost anomaly detection monitor. You can use each
// subscription to define subscribers with email or SNS notifications. Email
// subscribers can set a dollar threshold and a time frequency for receiving
// notifications.
func (c *Client) CreateAnomalySubscription(ctx context.Context, params *CreateAnomalySubscriptionInput, optFns ...func(*Options)) (*CreateAnomalySubscriptionOutput, error) {
	if params == nil {
		params = &CreateAnomalySubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAnomalySubscription", params, optFns, c.addOperationCreateAnomalySubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAnomalySubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAnomalySubscriptionInput struct {

	// The cost anomaly subscription object that you want to create.
	//
	// This member is required.
	AnomalySubscription *types.AnomalySubscription

	// An optional list of tags to associate with the specified AnomalySubscription
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_AnomalySubscription.html).
	// You can use resource tags to control access to your subscription using IAM
	// policies. Each tag consists of a key and a value, and each key must be unique
	// for the resource. The following restrictions apply to resource tags:
	//
	// * Although
	// the maximum number of array members is 200, you can assign a maximum of 50
	// user-tags to one resource. The remaining are reserved for Amazon Web Services
	// use
	//
	// * The maximum length of a key is 128 characters
	//
	// * The maximum length of a
	// value is 256 characters
	//
	// * Keys and values can only contain alphanumeric
	// characters, spaces, and any of the following: _.:/=+@-
	//
	// * Keys and values are
	// case sensitive
	//
	// * Keys and values are trimmed for any leading or trailing
	// whitespaces
	//
	// * Don’t use aws: as a prefix for your keys. This prefix is reserved
	// for Amazon Web Services use
	ResourceTags []types.ResourceTag

	noSmithyDocumentSerde
}

type CreateAnomalySubscriptionOutput struct {

	// The unique identifier of your newly created cost anomaly subscription.
	//
	// This member is required.
	SubscriptionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAnomalySubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAnomalySubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAnomalySubscription{}, middleware.After)
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
	if err = addOpCreateAnomalySubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAnomalySubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAnomalySubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "CreateAnomalySubscription",
	}
}
