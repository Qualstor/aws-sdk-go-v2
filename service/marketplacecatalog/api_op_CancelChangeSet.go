// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecatalog

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used to cancel an open change request. Must be sent before the status of the
// request changes to APPLYING, the final stage of completing your change request.
// You can describe a change during the 60-day request history retention period for
// API calls.
func (c *Client) CancelChangeSet(ctx context.Context, params *CancelChangeSetInput, optFns ...func(*Options)) (*CancelChangeSetOutput, error) {
	if params == nil {
		params = &CancelChangeSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelChangeSet", params, optFns, c.addOperationCancelChangeSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelChangeSetInput struct {

	// Required. The catalog related to the request. Fixed value: AWSMarketplace.
	//
	// This member is required.
	Catalog *string

	// Required. The unique identifier of the StartChangeSet request that you want to
	// cancel.
	//
	// This member is required.
	ChangeSetId *string

	noSmithyDocumentSerde
}

type CancelChangeSetOutput struct {

	// The ARN associated with the change set referenced in this request.
	ChangeSetArn *string

	// The unique identifier for the change set referenced in this request.
	ChangeSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelChangeSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelChangeSet{}, middleware.After)
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
	if err = addOpCancelChangeSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelChangeSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelChangeSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "CancelChangeSet",
	}
}
