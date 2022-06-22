// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the data storage options for this project. If you store evaluation
// events, you an keep them and analyze them on your own. If you choose not to
// store evaluation events, Evidently deletes them after using them to produce
// metrics and other experiment results that you can view. You can't specify both
// cloudWatchLogs and s3Destination in the same operation.
func (c *Client) UpdateProjectDataDelivery(ctx context.Context, params *UpdateProjectDataDeliveryInput, optFns ...func(*Options)) (*UpdateProjectDataDeliveryOutput, error) {
	if params == nil {
		params = &UpdateProjectDataDeliveryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProjectDataDelivery", params, optFns, c.addOperationUpdateProjectDataDeliveryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProjectDataDeliveryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProjectDataDeliveryInput struct {

	// The name or ARN of the project that you want to modify the data storage options
	// for.
	//
	// This member is required.
	Project *string

	// A structure containing the CloudWatch Logs log group where you want to store
	// evaluation events.
	CloudWatchLogs *types.CloudWatchLogsDestinationConfig

	// A structure containing the S3 bucket name and bucket prefix where you want to
	// store evaluation events.
	S3Destination *types.S3DestinationConfig

	noSmithyDocumentSerde
}

type UpdateProjectDataDeliveryOutput struct {

	// A structure containing details about the project that you updated.
	//
	// This member is required.
	Project *types.Project

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProjectDataDeliveryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateProjectDataDelivery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateProjectDataDelivery{}, middleware.After)
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
	if err = addOpUpdateProjectDataDeliveryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProjectDataDelivery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProjectDataDelivery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "UpdateProjectDataDelivery",
	}
}
