// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Associates a LoggingConfiguration with a specified web ACL. You can
// access information about all traffic that AWS WAF inspects using the following
// steps:
//
// * Create an Amazon Kinesis Data Firehose. Create the data firehose with
// a PUT source and in the region that you are operating. However, if you are
// capturing logs for Amazon CloudFront, always create the firehose in US East (N.
// Virginia). Do not create the data firehose using a Kinesis stream as your
// source.
//
// * Associate that firehose to your web ACL using a
// PutLoggingConfiguration request.
//
// When you successfully enable logging using a
// PutLoggingConfiguration request, AWS WAF will create a service linked role with
// the necessary permissions to write logs to the Amazon Kinesis Data Firehose. For
// more information, see Logging Web ACL Traffic Information
// (https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) in the AWS
// WAF Developer Guide.
func (c *Client) PutLoggingConfiguration(ctx context.Context, params *PutLoggingConfigurationInput, optFns ...func(*Options)) (*PutLoggingConfigurationOutput, error) {
	if params == nil {
		params = &PutLoggingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLoggingConfiguration", params, optFns, c.addOperationPutLoggingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLoggingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLoggingConfigurationInput struct {

	// The Amazon Kinesis Data Firehose that contains the inspected traffic
	// information, the redacted fields details, and the Amazon Resource Name (ARN) of
	// the web ACL to monitor. When specifying Type in RedactedFields, you must use one
	// of the following values: URI, QUERY_STRING, HEADER, or METHOD.
	//
	// This member is required.
	LoggingConfiguration *types.LoggingConfiguration

	noSmithyDocumentSerde
}

type PutLoggingConfigurationOutput struct {

	// The LoggingConfiguration that you submitted in the request.
	LoggingConfiguration *types.LoggingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutLoggingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutLoggingConfiguration{}, middleware.After)
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
	if err = addOpPutLoggingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLoggingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutLoggingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "PutLoggingConfiguration",
	}
}
