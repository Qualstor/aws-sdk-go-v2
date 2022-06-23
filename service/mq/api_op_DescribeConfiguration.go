// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/mq/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the specified configuration.
func (c *Client) DescribeConfiguration(ctx context.Context, params *DescribeConfigurationInput, optFns ...func(*Options)) (*DescribeConfigurationOutput, error) {
	if params == nil {
		params = &DescribeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfiguration", params, optFns, c.addOperationDescribeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationInput struct {

	// The unique ID that Amazon MQ generates for the configuration.
	//
	// This member is required.
	ConfigurationId *string

	noSmithyDocumentSerde
}

type DescribeConfigurationOutput struct {

	// Required. The ARN of the configuration.
	Arn *string

	// Optional. The authentication strategy associated with the configuration. The
	// default is SIMPLE.
	AuthenticationStrategy types.AuthenticationStrategy

	// Required. The date and time of the configuration revision.
	Created *time.Time

	// Required. The description of the configuration.
	Description *string

	// Required. The type of broker engine. Currently, Amazon MQ supports ACTIVEMQ and
	// RABBITMQ.
	EngineType types.EngineType

	// Required. The broker engine's version. For a list of supported engine versions,
	// see, Supported engines
	// (https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker-engine.html).
	EngineVersion *string

	// Required. The unique ID that Amazon MQ generates for the configuration.
	Id *string

	// Required. The latest revision of the configuration.
	LatestRevision *types.ConfigurationRevision

	// Required. The name of the configuration. This value can contain only
	// alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~).
	// This value must be 1-150 characters long.
	Name *string

	// The list of all tags associated with this configuration.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConfiguration{}, middleware.After)
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
	if err = addOpDescribeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "DescribeConfiguration",
	}
}
