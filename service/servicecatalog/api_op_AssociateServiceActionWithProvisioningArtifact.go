// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a self-service action with a provisioning artifact.
func (c *Client) AssociateServiceActionWithProvisioningArtifact(ctx context.Context, params *AssociateServiceActionWithProvisioningArtifactInput, optFns ...func(*Options)) (*AssociateServiceActionWithProvisioningArtifactOutput, error) {
	if params == nil {
		params = &AssociateServiceActionWithProvisioningArtifactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateServiceActionWithProvisioningArtifact", params, optFns, c.addOperationAssociateServiceActionWithProvisioningArtifactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateServiceActionWithProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateServiceActionWithProvisioningArtifactInput struct {

	// The product identifier. For example, prod-abcdzk7xy33qa.
	//
	// This member is required.
	ProductId *string

	// The identifier of the provisioning artifact. For example, pa-4abcdjnxjj6ne.
	//
	// This member is required.
	ProvisioningArtifactId *string

	// The self-service action identifier. For example, act-fs7abcd89wxyz.
	//
	// This member is required.
	ServiceActionId *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	noSmithyDocumentSerde
}

type AssociateServiceActionWithProvisioningArtifactOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateServiceActionWithProvisioningArtifactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateServiceActionWithProvisioningArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateServiceActionWithProvisioningArtifact{}, middleware.After)
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
	if err = addOpAssociateServiceActionWithProvisioningArtifactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateServiceActionWithProvisioningArtifact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateServiceActionWithProvisioningArtifact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "AssociateServiceActionWithProvisioningArtifact",
	}
}
