// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Add cost allocation tags to the specified Amazon SQS queue. For an overview, see
// Tagging Your Amazon SQS Queues
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html)
// in the Amazon SQS Developer Guide. When you use queue tags, keep the following
// guidelines in mind:
//
// * Adding more than 50 tags to a queue isn't recommended.
//
// *
// Tags don't have any semantic meaning. Amazon SQS interprets tags as character
// strings.
//
// * Tags are case-sensitive.
//
// * A new tag with a key identical to that
// of an existing tag overwrites the existing tag.
//
// For a full list of tag
// restrictions, see Quotas related to queues
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-limits.html#limits-queues)
// in the Amazon SQS Developer Guide. Cross-account permissions don't apply to this
// action. For more information, see Grant cross-account permissions to a role and
// a user name
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon SQS Developer Guide.
func (c *Client) TagQueue(ctx context.Context, params *TagQueueInput, optFns ...func(*Options)) (*TagQueueOutput, error) {
	if params == nil {
		params = &TagQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagQueue", params, optFns, c.addOperationTagQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagQueueInput struct {

	// The URL of the queue.
	//
	// This member is required.
	QueueUrl *string

	// The list of tags to be added to the specified queue.
	//
	// This member is required.
	Tags map[string]string

	noSmithyDocumentSerde
}

type TagQueueOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTagQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpTagQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpTagQueue{}, middleware.After)
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
	if err = addOpTagQueueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagQueue(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTagQueue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "TagQueue",
	}
}
