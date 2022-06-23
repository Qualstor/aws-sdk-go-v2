// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The DeleteHIT operation is used to delete HIT that is no longer needed. Only the
// Requester who created the HIT can delete it. You can only dispose of HITs that
// are in the Reviewable state, with all of their submitted assignments already
// either approved or rejected. If you call the DeleteHIT operation on a HIT that
// is not in the Reviewable state (for example, that has not expired, or still has
// active assignments), or on a HIT that is Reviewable but without all of its
// submitted assignments already approved or rejected, the service will return an
// error.
//
// * HITs are automatically disposed of after 120 days.
//
// * After you
// dispose of a HIT, you can no longer approve the HIT's rejected assignments.
//
// *
// Disposed HITs are not returned in results for the ListHITs operation.
//
// *
// Disposing HITs can improve the performance of operations such as
// ListReviewableHITs and ListHITs.
func (c *Client) DeleteHIT(ctx context.Context, params *DeleteHITInput, optFns ...func(*Options)) (*DeleteHITOutput, error) {
	if params == nil {
		params = &DeleteHITInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteHIT", params, optFns, c.addOperationDeleteHITMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteHITInput struct {

	// The ID of the HIT to be deleted.
	//
	// This member is required.
	HITId *string

	noSmithyDocumentSerde
}

type DeleteHITOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteHITMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteHIT{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteHIT{}, middleware.After)
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
	if err = addOpDeleteHITValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHIT(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteHIT(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "DeleteHIT",
	}
}
