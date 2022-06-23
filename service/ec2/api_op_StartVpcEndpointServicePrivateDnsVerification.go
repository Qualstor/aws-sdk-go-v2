// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates the verification process to prove that the service provider owns the
// private DNS name domain for the endpoint service. The service provider must
// successfully perform the verification before the consumer can use the name to
// access the service. Before the service provider runs this command, they must add
// a record to the DNS server.
func (c *Client) StartVpcEndpointServicePrivateDnsVerification(ctx context.Context, params *StartVpcEndpointServicePrivateDnsVerificationInput, optFns ...func(*Options)) (*StartVpcEndpointServicePrivateDnsVerificationOutput, error) {
	if params == nil {
		params = &StartVpcEndpointServicePrivateDnsVerificationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartVpcEndpointServicePrivateDnsVerification", params, optFns, c.addOperationStartVpcEndpointServicePrivateDnsVerificationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartVpcEndpointServicePrivateDnsVerificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartVpcEndpointServicePrivateDnsVerificationInput struct {

	// The ID of the endpoint service.
	//
	// This member is required.
	ServiceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type StartVpcEndpointServicePrivateDnsVerificationOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartVpcEndpointServicePrivateDnsVerificationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpStartVpcEndpointServicePrivateDnsVerification{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpStartVpcEndpointServicePrivateDnsVerification{}, middleware.After)
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
	if err = addOpStartVpcEndpointServicePrivateDnsVerificationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartVpcEndpointServicePrivateDnsVerification(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartVpcEndpointServicePrivateDnsVerification(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "StartVpcEndpointServicePrivateDnsVerification",
	}
}
