// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a CloudFormation template that streamlines and automates the
// integration of VPC flow logs with Amazon Athena. This make it easier for you to
// query and gain insights from VPC flow logs data. Based on the information that
// you provide, we configure resources in the template to do the following:
//
// *
// Create a table in Athena that maps fields to a custom log format
//
// * Create a
// Lambda function that updates the table with new partitions on a daily, weekly,
// or monthly basis
//
// * Create a table partitioned between two timestamps in the
// past
//
// * Create a set of named queries in Athena that you can use to get started
// quickly
func (c *Client) GetFlowLogsIntegrationTemplate(ctx context.Context, params *GetFlowLogsIntegrationTemplateInput, optFns ...func(*Options)) (*GetFlowLogsIntegrationTemplateOutput, error) {
	if params == nil {
		params = &GetFlowLogsIntegrationTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFlowLogsIntegrationTemplate", params, optFns, c.addOperationGetFlowLogsIntegrationTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFlowLogsIntegrationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFlowLogsIntegrationTemplateInput struct {

	// To store the CloudFormation template in Amazon S3, specify the location in
	// Amazon S3.
	//
	// This member is required.
	ConfigDeliveryS3DestinationArn *string

	// The ID of the flow log.
	//
	// This member is required.
	FlowLogId *string

	// Information about the service integration.
	//
	// This member is required.
	IntegrateServices *types.IntegrateServices

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type GetFlowLogsIntegrationTemplateOutput struct {

	// The generated CloudFormation template.
	Result *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFlowLogsIntegrationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetFlowLogsIntegrationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetFlowLogsIntegrationTemplate{}, middleware.After)
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
	if err = addOpGetFlowLogsIntegrationTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFlowLogsIntegrationTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFlowLogsIntegrationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetFlowLogsIntegrationTemplate",
	}
}
