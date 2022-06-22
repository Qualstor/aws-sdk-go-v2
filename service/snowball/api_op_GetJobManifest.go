// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a link to an Amazon S3 presigned URL for the manifest file associated
// with the specified JobId value. You can access the manifest file for up to 60
// minutes after this request has been made. To access the manifest file after 60
// minutes have passed, you'll have to make another call to the GetJobManifest
// action. The manifest is an encrypted file that you can download after your job
// enters the WithCustomer status. The manifest is decrypted by using the
// UnlockCode code value, when you pass both values to the Snow device through the
// Snowball client when the client is started for the first time. As a best
// practice, we recommend that you don't save a copy of an UnlockCode value in the
// same location as the manifest file for that job. Saving these separately helps
// prevent unauthorized parties from gaining access to the Snow device associated
// with that job. The credentials of a given job, including its manifest file and
// unlock code, expire 360 days after the job is created.
func (c *Client) GetJobManifest(ctx context.Context, params *GetJobManifestInput, optFns ...func(*Options)) (*GetJobManifestOutput, error) {
	if params == nil {
		params = &GetJobManifestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetJobManifest", params, optFns, c.addOperationGetJobManifestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetJobManifestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJobManifestInput struct {

	// The ID for a job that you want to get the manifest file for, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type GetJobManifestOutput struct {

	// The Amazon S3 presigned URL for the manifest file associated with the specified
	// JobId value.
	ManifestURI *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetJobManifestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetJobManifest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetJobManifest{}, middleware.After)
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
	if err = addOpGetJobManifestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetJobManifest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetJobManifest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "GetJobManifest",
	}
}
