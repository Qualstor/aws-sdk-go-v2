// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes a connection.
func (c *Client) DeleteConnection(ctx context.Context, params *DeleteConnectionInput, optFns ...func(*Options)) (*DeleteConnectionOutput, error) {
	if params == nil {
		params = &DeleteConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConnection", params, optFns, c.addOperationDeleteConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConnectionInput struct {

	// The name of the connection to delete.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DeleteConnectionOutput struct {

	// The ARN of the connection that was deleted.
	ConnectionArn *string

	// The state of the connection before it was deleted.
	ConnectionState types.ConnectionState

	// A time stamp for the time that the connection was created.
	CreationTime *time.Time

	// A time stamp for the time that the connection was last authorized before it wa
	// deleted.
	LastAuthorizedTime *time.Time

	// A time stamp for the time that the connection was last modified before it was
	// deleted.
	LastModifiedTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteConnection{}, middleware.After)
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
	if err = addOpDeleteConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DeleteConnection",
	}
}
