// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes cross-account permissions from a signing profile.
func (c *Client) RemoveProfilePermission(ctx context.Context, params *RemoveProfilePermissionInput, optFns ...func(*Options)) (*RemoveProfilePermissionOutput, error) {
	if params == nil {
		params = &RemoveProfilePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveProfilePermission", params, optFns, c.addOperationRemoveProfilePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveProfilePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveProfilePermissionInput struct {

	// A human-readable name for the signing profile with permissions to be removed.
	//
	// This member is required.
	ProfileName *string

	// An identifier for the current revision of the signing profile permissions.
	//
	// This member is required.
	RevisionId *string

	// A unique identifier for the cross-account permissions statement.
	//
	// This member is required.
	StatementId *string

	noSmithyDocumentSerde
}

type RemoveProfilePermissionOutput struct {

	// An identifier for the current revision of the profile permissions.
	RevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveProfilePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRemoveProfilePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveProfilePermission{}, middleware.After)
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
	if err = addOpRemoveProfilePermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveProfilePermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveProfilePermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "signer",
		OperationName: "RemoveProfilePermission",
	}
}
