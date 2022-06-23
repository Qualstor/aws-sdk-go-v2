// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/polly/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the content of the specified pronunciation lexicon stored in an Amazon
// Web Services Region. For more information, see Managing Lexicons
// (https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
func (c *Client) GetLexicon(ctx context.Context, params *GetLexiconInput, optFns ...func(*Options)) (*GetLexiconOutput, error) {
	if params == nil {
		params = &GetLexiconInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLexicon", params, optFns, c.addOperationGetLexiconMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLexiconOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLexiconInput struct {

	// Name of the lexicon.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetLexiconOutput struct {

	// Lexicon object that provides name and the string content of the lexicon.
	Lexicon *types.Lexicon

	// Metadata of the lexicon, including phonetic alphabetic used, language code,
	// lexicon ARN, number of lexemes defined in the lexicon, and size of lexicon in
	// bytes.
	LexiconAttributes *types.LexiconAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLexiconMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLexicon{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLexicon{}, middleware.After)
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
	if err = addOpGetLexiconValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLexicon(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLexicon(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "GetLexicon",
	}
}
