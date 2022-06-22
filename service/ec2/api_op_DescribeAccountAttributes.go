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

// Describes attributes of your Amazon Web Services account. The following are the
// supported account attributes:
//
// * supported-platforms: Indicates whether your
// account can launch instances into EC2-Classic and EC2-VPC, or only into
// EC2-VPC.
//
// * default-vpc: The ID of the default VPC for your account, or none.
//
// *
// max-instances: This attribute is no longer supported. The returned value does
// not reflect your actual vCPU limit for running On-Demand Instances. For more
// information, see On-Demand Instance Limits
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-on-demand-instances.html#ec2-on-demand-instances-limits)
// in the Amazon Elastic Compute Cloud User Guide.
//
// *
// vpc-max-security-groups-per-interface: The maximum number of security groups
// that you can assign to a network interface.
//
// * max-elastic-ips: The maximum
// number of Elastic IP addresses that you can allocate for use with
// EC2-Classic.
//
// * vpc-max-elastic-ips: The maximum number of Elastic IP addresses
// that you can allocate for use with EC2-VPC.
func (c *Client) DescribeAccountAttributes(ctx context.Context, params *DescribeAccountAttributesInput, optFns ...func(*Options)) (*DescribeAccountAttributesOutput, error) {
	if params == nil {
		params = &DescribeAccountAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountAttributes", params, optFns, c.addOperationDescribeAccountAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountAttributesInput struct {

	// The account attribute names.
	AttributeNames []types.AccountAttributeName

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type DescribeAccountAttributesOutput struct {

	// Information about the account attributes.
	AccountAttributes []types.AccountAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccountAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeAccountAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeAccountAttributes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccountAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeAccountAttributes",
	}
}
