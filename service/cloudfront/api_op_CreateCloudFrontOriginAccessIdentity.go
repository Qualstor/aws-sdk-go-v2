// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new origin access identity. If you're using Amazon S3 for your origin,
// you can use an origin access identity to require users to access your content
// using a CloudFront URL instead of the Amazon S3 URL. For more information about
// how to use origin access identities, see Serving Private Content through
// CloudFront
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html)
// in the Amazon CloudFront Developer Guide.
func (c *Client) CreateCloudFrontOriginAccessIdentity(ctx context.Context, params *CreateCloudFrontOriginAccessIdentityInput, optFns ...func(*Options)) (*CreateCloudFrontOriginAccessIdentityOutput, error) {
	if params == nil {
		params = &CreateCloudFrontOriginAccessIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCloudFrontOriginAccessIdentity", params, optFns, c.addOperationCreateCloudFrontOriginAccessIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCloudFrontOriginAccessIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to create a new origin access identity (OAI). An origin access
// identity is a special CloudFront user that you can associate with Amazon S3
// origins, so that you can secure all or just some of your Amazon S3 content. For
// more information, see  Restricting Access to Amazon S3 Content by Using an
// Origin Access Identity
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html)
// in the Amazon CloudFront Developer Guide.
type CreateCloudFrontOriginAccessIdentityInput struct {

	// The current configuration information for the identity.
	//
	// This member is required.
	CloudFrontOriginAccessIdentityConfig *types.CloudFrontOriginAccessIdentityConfig

	noSmithyDocumentSerde
}

// The returned result of the corresponding request.
type CreateCloudFrontOriginAccessIdentityOutput struct {

	// The origin access identity's information.
	CloudFrontOriginAccessIdentity *types.CloudFrontOriginAccessIdentity

	// The current version of the origin access identity created.
	ETag *string

	// The fully qualified URI of the new origin access identity just created.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCloudFrontOriginAccessIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateCloudFrontOriginAccessIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateCloudFrontOriginAccessIdentity{}, middleware.After)
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
	if err = addOpCreateCloudFrontOriginAccessIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCloudFrontOriginAccessIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCloudFrontOriginAccessIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateCloudFrontOriginAccessIdentity",
	}
}
