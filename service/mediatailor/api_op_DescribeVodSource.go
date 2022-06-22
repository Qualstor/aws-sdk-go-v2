// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides details about a specific VOD source in a specific source location.
func (c *Client) DescribeVodSource(ctx context.Context, params *DescribeVodSourceInput, optFns ...func(*Options)) (*DescribeVodSourceOutput, error) {
	if params == nil {
		params = &DescribeVodSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVodSource", params, optFns, c.addOperationDescribeVodSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVodSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVodSourceInput struct {

	// The identifier for the source location you are working on.
	//
	// This member is required.
	SourceLocationName *string

	// The identifier for the VOD source you are working on.
	//
	// This member is required.
	VodSourceName *string

	noSmithyDocumentSerde
}

type DescribeVodSourceOutput struct {

	// The ARN of the VOD source.
	Arn *string

	// The timestamp that indicates when the VOD source was created.
	CreationTime *time.Time

	// The HTTP package configurations.
	HttpPackageConfigurations []types.HttpPackageConfiguration

	// The last modified time of the VOD source.
	LastModifiedTime *time.Time

	// The name of the source location associated with the VOD source.
	SourceLocationName *string

	// The tags assigned to the VOD source.
	Tags map[string]string

	// The name of the VOD source.
	VodSourceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVodSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeVodSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeVodSource{}, middleware.After)
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
	if err = addOpDescribeVodSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVodSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVodSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "DescribeVodSource",
	}
}
