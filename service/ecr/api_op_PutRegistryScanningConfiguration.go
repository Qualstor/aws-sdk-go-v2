// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates the scanning configuration for your private registry.
func (c *Client) PutRegistryScanningConfiguration(ctx context.Context, params *PutRegistryScanningConfigurationInput, optFns ...func(*Options)) (*PutRegistryScanningConfigurationOutput, error) {
	if params == nil {
		params = &PutRegistryScanningConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRegistryScanningConfiguration", params, optFns, c.addOperationPutRegistryScanningConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRegistryScanningConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRegistryScanningConfigurationInput struct {

	// The scanning rules to use for the registry. A scanning rule is used to determine
	// which repository filters are used and at what frequency scanning will occur.
	Rules []types.RegistryScanningRule

	// The scanning type to set for the registry. When a registry scanning
	// configuration is not defined, by default the BASIC scan type is used. When basic
	// scanning is used, you may specify filters to determine which individual
	// repositories, or all repositories, are scanned when new images are pushed to
	// those repositories. Alternatively, you can do manual scans of images with basic
	// scanning. When the ENHANCED scan type is set, Amazon Inspector provides
	// automated vulnerability scanning. You may choose between continuous scanning or
	// scan on push and you may specify filters to determine which individual
	// repositories, or all repositories, are scanned.
	ScanType types.ScanType

	noSmithyDocumentSerde
}

type PutRegistryScanningConfigurationOutput struct {

	// The scanning configuration for your registry.
	RegistryScanningConfiguration *types.RegistryScanningConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRegistryScanningConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRegistryScanningConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRegistryScanningConfiguration{}, middleware.After)
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
	if err = addOpPutRegistryScanningConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRegistryScanningConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRegistryScanningConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "PutRegistryScanningConfiguration",
	}
}
