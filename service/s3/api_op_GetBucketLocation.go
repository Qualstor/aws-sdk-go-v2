// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/Qualstor/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/Qualstor/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/aws/smithy-go"
	smithyxml "github.com/aws/smithy-go/encoding/xml"
	smithyio "github.com/aws/smithy-go/io"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Returns the Region the bucket resides in. You set the bucket's Region using the
// LocationConstraint request parameter in a CreateBucket request. For more
// information, see CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html). To use
// this implementation of the operation, you must be the bucket owner. To use this
// API against an access point, provide the alias of the access point in place of
// the bucket name. The following operations are related to GetBucketLocation:
//
// *
// GetObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
//
// *
// CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
func (c *Client) GetBucketLocation(ctx context.Context, params *GetBucketLocationInput, optFns ...func(*Options)) (*GetBucketLocationOutput, error) {
	if params == nil {
		params = &GetBucketLocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketLocation", params, optFns, c.addOperationGetBucketLocationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketLocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketLocationInput struct {

	// The name of the bucket for which to get the location.
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

type GetBucketLocationOutput struct {

	// Specifies the Region where the bucket resides. For a list of all the Amazon S3
	// supported location constraints by Region, see Regions and Endpoints
	// (https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region). Buckets in
	// Region us-east-1 have a LocationConstraint of null.
	LocationConstraint types.BucketLocationConstraint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketLocationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketLocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketLocation{}, middleware.After)
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
	if err = swapDeserializerHelper(stack); err != nil {
		return err
	}
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetBucketLocationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketLocation(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addGetBucketLocationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type awsRestxml_deserializeOpGetBucketLocation_custom struct {
}

func (*awsRestxml_deserializeOpGetBucketLocation_custom) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestxml_deserializeOpGetBucketLocation_custom) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestxml_deserializeOpErrorGetBucketLocation(response, &metadata)
	}
	output := &GetBucketLocationOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	body := io.TeeReader(response.Body, ringBuffer)
	rootDecoder := xml.NewDecoder(body)
	decoder := smithyxml.WrapNodeDecoder(rootDecoder, xml.StartElement{})
	err = awsRestxml_deserializeOpDocumentGetBucketLocationOutput(&output, decoder)
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

// Helper to swap in a custom deserializer
func swapDeserializerHelper(stack *middleware.Stack) error {
	_, err := stack.Deserialize.Swap("OperationDeserializer", &awsRestxml_deserializeOpGetBucketLocation_custom{})
	if err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetBucketLocation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketLocation",
	}
}

// getGetBucketLocationBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getGetBucketLocationBucketMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketLocationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addGetBucketLocationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getGetBucketLocationBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
