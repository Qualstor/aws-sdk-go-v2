// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an Glue machine learning transform. Machine learning transforms are a
// special type of transform that use machine learning to learn the details of the
// transformation to be performed by learning from examples provided by humans.
// These transformations are then saved by Glue. If you no longer need a transform,
// you can delete it by calling DeleteMLTransforms. However, any Glue jobs that
// still reference the deleted transform will no longer succeed.
func (c *Client) DeleteMLTransform(ctx context.Context, params *DeleteMLTransformInput, optFns ...func(*Options)) (*DeleteMLTransformOutput, error) {
	if params == nil {
		params = &DeleteMLTransformInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMLTransform", params, optFns, c.addOperationDeleteMLTransformMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMLTransformOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMLTransformInput struct {

	// The unique identifier of the transform to delete.
	//
	// This member is required.
	TransformId *string

	noSmithyDocumentSerde
}

type DeleteMLTransformOutput struct {

	// The unique identifier of the transform that was deleted.
	TransformId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMLTransformMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMLTransform{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMLTransform{}, middleware.After)
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
	if err = addOpDeleteMLTransformValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMLTransform(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMLTransform(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteMLTransform",
	}
}
