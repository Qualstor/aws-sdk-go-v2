// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an entity.
func (c *Client) DeleteEntity(ctx context.Context, params *DeleteEntityInput, optFns ...func(*Options)) (*DeleteEntityOutput, error) {
	if params == nil {
		params = &DeleteEntityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEntity", params, optFns, c.addOperationDeleteEntityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEntityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEntityInput struct {

	// The ID of the entity to delete.
	//
	// This member is required.
	EntityId *string

	// The ID of the workspace that contains the entity to delete.
	//
	// This member is required.
	WorkspaceId *string

	// A Boolean value that specifies whether the operation deletes child entities.
	IsRecursive *bool

	noSmithyDocumentSerde
}

type DeleteEntityOutput struct {

	// The current state of the deleted entity.
	//
	// This member is required.
	State types.State

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteEntityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteEntity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteEntity{}, middleware.After)
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
	if err = addEndpointPrefix_opDeleteEntityMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteEntityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEntity(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDeleteEntityMiddleware struct {
}

func (*endpointPrefix_opDeleteEntityMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteEntityMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDeleteEntityMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDeleteEntityMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteEntity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iottwinmaker",
		OperationName: "DeleteEntity",
	}
}
