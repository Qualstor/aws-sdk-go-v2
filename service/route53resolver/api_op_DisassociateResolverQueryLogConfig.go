// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a VPC from a query logging configuration. Before you can delete a
// query logging configuration, you must first disassociate all VPCs from the
// configuration. If you used Resource Access Manager (RAM) to share a query
// logging configuration with other accounts, VPCs can be disassociated from the
// configuration in the following ways:
//
// * The accounts that you shared the
// configuration with can disassociate VPCs from the configuration.
//
// * You can stop
// sharing the configuration.
func (c *Client) DisassociateResolverQueryLogConfig(ctx context.Context, params *DisassociateResolverQueryLogConfigInput, optFns ...func(*Options)) (*DisassociateResolverQueryLogConfigOutput, error) {
	if params == nil {
		params = &DisassociateResolverQueryLogConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateResolverQueryLogConfig", params, optFns, c.addOperationDisassociateResolverQueryLogConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateResolverQueryLogConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateResolverQueryLogConfigInput struct {

	// The ID of the query logging configuration that you want to disassociate a
	// specified VPC from.
	//
	// This member is required.
	ResolverQueryLogConfigId *string

	// The ID of the Amazon VPC that you want to disassociate from a specified query
	// logging configuration.
	//
	// This member is required.
	ResourceId *string

	noSmithyDocumentSerde
}

type DisassociateResolverQueryLogConfigOutput struct {

	// A complex type that contains settings for the association that you deleted
	// between an Amazon VPC and a query logging configuration.
	ResolverQueryLogConfigAssociation *types.ResolverQueryLogConfigAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateResolverQueryLogConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateResolverQueryLogConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateResolverQueryLogConfig{}, middleware.After)
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
	if err = addOpDisassociateResolverQueryLogConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateResolverQueryLogConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateResolverQueryLogConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "DisassociateResolverQueryLogConfig",
	}
}
