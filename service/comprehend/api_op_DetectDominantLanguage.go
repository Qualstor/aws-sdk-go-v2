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

// Determines the dominant language of the input text. For a list of languages that
// Amazon Comprehend can detect, see Amazon Comprehend Supported Languages
// (https://docs.aws.amazon.com/comprehend/latest/dg/how-languages.html).
func (c *Client) DetectDominantLanguage(ctx context.Context, params *DetectDominantLanguageInput, optFns ...func(*Options)) (*DetectDominantLanguageOutput, error) {
	if params == nil {
		params = &DetectDominantLanguageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectDominantLanguage", params, optFns, c.addOperationDetectDominantLanguageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectDominantLanguageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectDominantLanguageInput struct {

	// A UTF-8 text string. Each string should contain at least 20 characters and must
	// contain fewer that 5,000 bytes of UTF-8 encoded characters.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

type DetectDominantLanguageOutput struct {

	// The languages that Amazon Comprehend detected in the input text. For each
	// language, the response returns the RFC 5646 language code and the level of
	// confidence that Amazon Comprehend has in the accuracy of its inference. For more
	// information about RFC 5646, see Tags for Identifying Languages
	// (https://tools.ietf.org/html/rfc5646) on the IETF Tools web site.
	Languages []types.DominantLanguage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectDominantLanguageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectDominantLanguage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectDominantLanguage{}, middleware.After)
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
	if err = addOpDetectDominantLanguageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectDominantLanguage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectDominantLanguage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DetectDominantLanguage",
	}
}
