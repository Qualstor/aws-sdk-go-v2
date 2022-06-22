// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the TLS certificates that are associated with the
// specified Lightsail load balancer. TLS is just an updated, more secure version
// of Secure Socket Layer (SSL). You can have a maximum of 2 certificates
// associated with a Lightsail load balancer. One is active and the other is
// inactive.
func (c *Client) GetLoadBalancerTlsCertificates(ctx context.Context, params *GetLoadBalancerTlsCertificatesInput, optFns ...func(*Options)) (*GetLoadBalancerTlsCertificatesOutput, error) {
	if params == nil {
		params = &GetLoadBalancerTlsCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLoadBalancerTlsCertificates", params, optFns, c.addOperationGetLoadBalancerTlsCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLoadBalancerTlsCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLoadBalancerTlsCertificatesInput struct {

	// The name of the load balancer you associated with your SSL/TLS certificate.
	//
	// This member is required.
	LoadBalancerName *string

	noSmithyDocumentSerde
}

type GetLoadBalancerTlsCertificatesOutput struct {

	// An array of LoadBalancerTlsCertificate objects describing your SSL/TLS
	// certificates.
	TlsCertificates []types.LoadBalancerTlsCertificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLoadBalancerTlsCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLoadBalancerTlsCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLoadBalancerTlsCertificates{}, middleware.After)
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
	if err = addOpGetLoadBalancerTlsCertificatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLoadBalancerTlsCertificates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLoadBalancerTlsCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetLoadBalancerTlsCertificates",
	}
}
