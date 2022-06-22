// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a subscription definition.
func (c *Client) UpdateSubscriptionDefinition(ctx context.Context, params *UpdateSubscriptionDefinitionInput, optFns ...func(*Options)) (*UpdateSubscriptionDefinitionOutput, error) {
	if params == nil {
		params = &UpdateSubscriptionDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSubscriptionDefinition", params, optFns, c.addOperationUpdateSubscriptionDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSubscriptionDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSubscriptionDefinitionInput struct {

	// The ID of the subscription definition.
	//
	// This member is required.
	SubscriptionDefinitionId *string

	// The name of the definition.
	Name *string

	noSmithyDocumentSerde
}

type UpdateSubscriptionDefinitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSubscriptionDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSubscriptionDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSubscriptionDefinition{}, middleware.After)
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
	if err = addOpUpdateSubscriptionDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSubscriptionDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSubscriptionDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "UpdateSubscriptionDefinition",
	}
}
