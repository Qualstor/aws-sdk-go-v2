// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes tags from the specified delivery stream. Removed tags are deleted, and
// you can't recover them after this operation successfully completes. If you
// specify a tag that doesn't exist, the operation ignores it. This operation has a
// limit of five transactions per second per account.
func (c *Client) UntagDeliveryStream(ctx context.Context, params *UntagDeliveryStreamInput, optFns ...func(*Options)) (*UntagDeliveryStreamOutput, error) {
	if params == nil {
		params = &UntagDeliveryStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UntagDeliveryStream", params, optFns, c.addOperationUntagDeliveryStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UntagDeliveryStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagDeliveryStreamInput struct {

	// The name of the delivery stream.
	//
	// This member is required.
	DeliveryStreamName *string

	// A list of tag keys. Each corresponding tag is removed from the delivery stream.
	//
	// This member is required.
	TagKeys []string

	noSmithyDocumentSerde
}

type UntagDeliveryStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUntagDeliveryStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUntagDeliveryStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUntagDeliveryStream{}, middleware.After)
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
	if err = addOpUntagDeliveryStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUntagDeliveryStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUntagDeliveryStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "UntagDeliveryStream",
	}
}
