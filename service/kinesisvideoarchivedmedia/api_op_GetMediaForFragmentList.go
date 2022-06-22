// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Gets media for a list of fragments (specified by fragment number) from the
// archived data in an Amazon Kinesis video stream. You must first call the
// GetDataEndpoint API to get an endpoint. Then send the GetMediaForFragmentList
// requests to this endpoint using the --endpoint-url parameter
// (https://docs.aws.amazon.com/cli/latest/reference/). For limits, see Kinesis
// Video Streams Limits
// (http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/limits.html). If an
// error is thrown after invoking a Kinesis Video Streams archived media API, in
// addition to the HTTP status code and the response body, it includes the
// following pieces of information:
//
// * x-amz-ErrorType HTTP header – contains a
// more specific error type in addition to what the HTTP status code provides.
//
// *
// x-amz-RequestId HTTP header – if you want to report an issue to AWS, the support
// team can better diagnose the problem if given the Request Id.
//
// Both the HTTP
// status code and the ErrorType header can be utilized to make programmatic
// decisions about whether errors are retry-able and under what conditions, as well
// as provide information on what actions the client programmer might need to take
// in order to successfully try again. For more information, see the Errors section
// at the bottom of this topic, as well as Common Errors
// (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/CommonErrors.html).
func (c *Client) GetMediaForFragmentList(ctx context.Context, params *GetMediaForFragmentListInput, optFns ...func(*Options)) (*GetMediaForFragmentListOutput, error) {
	if params == nil {
		params = &GetMediaForFragmentListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMediaForFragmentList", params, optFns, c.addOperationGetMediaForFragmentListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMediaForFragmentListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMediaForFragmentListInput struct {

	// A list of the numbers of fragments for which to retrieve media. You retrieve
	// these values with ListFragments.
	//
	// This member is required.
	Fragments []string

	// The Amazon Resource Name (ARN) of the stream from which to retrieve fragment
	// media. Specify either this parameter or the StreamName parameter.
	StreamARN *string

	// The name of the stream from which to retrieve fragment media. Specify either
	// this parameter or the StreamARN parameter.
	StreamName *string

	noSmithyDocumentSerde
}

type GetMediaForFragmentListOutput struct {

	// The content type of the requested media.
	ContentType *string

	// The payload that Kinesis Video Streams returns is a sequence of chunks from the
	// specified stream. For information about the chunks, see PutMedia
	// (http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/API_dataplane_PutMedia.html).
	// The chunks that Kinesis Video Streams returns in the GetMediaForFragmentList
	// call also include the following additional Matroska (MKV) tags:
	//
	// *
	// AWS_KINESISVIDEO_FRAGMENT_NUMBER - Fragment number returned in the chunk.
	//
	// *
	// AWS_KINESISVIDEO_SERVER_SIDE_TIMESTAMP - Server-side timestamp of the
	// fragment.
	//
	// * AWS_KINESISVIDEO_PRODUCER_SIDE_TIMESTAMP - Producer-side timestamp
	// of the fragment.
	//
	// The following tags will be included if an exception occurs:
	//
	// *
	// AWS_KINESISVIDEO_FRAGMENT_NUMBER - The number of the fragment that threw the
	// exception
	//
	// * AWS_KINESISVIDEO_EXCEPTION_ERROR_CODE - The integer code of the
	// exception
	//
	// * AWS_KINESISVIDEO_EXCEPTION_MESSAGE - A text description of the
	// exception
	Payload io.ReadCloser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMediaForFragmentListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMediaForFragmentList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMediaForFragmentList{}, middleware.After)
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
	if err = addOpGetMediaForFragmentListValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMediaForFragmentList(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMediaForFragmentList(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "GetMediaForFragmentList",
	}
}
