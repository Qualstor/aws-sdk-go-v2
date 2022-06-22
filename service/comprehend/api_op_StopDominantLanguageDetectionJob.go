// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops a dominant language detection job in progress. If the job state is
// IN_PROGRESS the job is marked for termination and put into the STOP_REQUESTED
// state. If the job completes before it can be stopped, it is put into the
// COMPLETED state; otherwise the job is stopped and put into the STOPPED state. If
// the job is in the COMPLETED or FAILED state when you call the
// StopDominantLanguageDetectionJob operation, the operation returns a 400 Internal
// Request Exception. When a job is stopped, any documents already processed are
// written to the output location.
func (c *Client) StopDominantLanguageDetectionJob(ctx context.Context, params *StopDominantLanguageDetectionJobInput, optFns ...func(*Options)) (*StopDominantLanguageDetectionJobOutput, error) {
	if params == nil {
		params = &StopDominantLanguageDetectionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDominantLanguageDetectionJob", params, optFns, c.addOperationStopDominantLanguageDetectionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDominantLanguageDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDominantLanguageDetectionJobInput struct {

	// The identifier of the dominant language detection job to stop.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type StopDominantLanguageDetectionJobOutput struct {

	// The identifier of the dominant language detection job to stop.
	JobId *string

	// Either STOP_REQUESTED if the job is currently running, or STOPPED if the job was
	// previously stopped with the StopDominantLanguageDetectionJob operation.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopDominantLanguageDetectionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopDominantLanguageDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopDominantLanguageDetectionJob{}, middleware.After)
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
	if err = addOpStopDominantLanguageDetectionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopDominantLanguageDetectionJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopDominantLanguageDetectionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "StopDominantLanguageDetectionJob",
	}
}
