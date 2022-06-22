// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates an attribute group with an application to augment the application's
// metadata with the group's attributes. This feature enables applications to be
// described with user-defined details that are machine-readable, such as
// third-party integrations.
func (c *Client) AssociateAttributeGroup(ctx context.Context, params *AssociateAttributeGroupInput, optFns ...func(*Options)) (*AssociateAttributeGroupOutput, error) {
	if params == nil {
		params = &AssociateAttributeGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateAttributeGroup", params, optFns, c.addOperationAssociateAttributeGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateAttributeGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateAttributeGroupInput struct {

	// The name or ID of the application.
	//
	// This member is required.
	Application *string

	// The name or ID of the attribute group that holds the attributes to describe the
	// application.
	//
	// This member is required.
	AttributeGroup *string

	noSmithyDocumentSerde
}

type AssociateAttributeGroupOutput struct {

	// The Amazon resource name (ARN) of the application that was augmented with
	// attributes.
	ApplicationArn *string

	// The Amazon resource name (ARN) of the attribute group that contains the
	// application's new attributes.
	AttributeGroupArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateAttributeGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateAttributeGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateAttributeGroup{}, middleware.After)
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
	if err = addOpAssociateAttributeGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateAttributeGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateAttributeGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "AssociateAttributeGroup",
	}
}
