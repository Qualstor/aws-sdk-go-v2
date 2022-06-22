// Code generated by smithy-go-codegen DO NOT EDIT.

package ecrpublic

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create or updates the catalog data for a public registry.
func (c *Client) PutRegistryCatalogData(ctx context.Context, params *PutRegistryCatalogDataInput, optFns ...func(*Options)) (*PutRegistryCatalogDataOutput, error) {
	if params == nil {
		params = &PutRegistryCatalogDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRegistryCatalogData", params, optFns, c.addOperationPutRegistryCatalogDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRegistryCatalogDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRegistryCatalogDataInput struct {

	// The display name for a public registry. The display name is shown as the
	// repository author in the Amazon ECR Public Gallery. The registry display name is
	// only publicly visible in the Amazon ECR Public Gallery for verified accounts.
	DisplayName *string

	noSmithyDocumentSerde
}

type PutRegistryCatalogDataOutput struct {

	// The catalog data for the public registry.
	//
	// This member is required.
	RegistryCatalogData *types.RegistryCatalogData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRegistryCatalogDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRegistryCatalogData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRegistryCatalogData{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRegistryCatalogData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRegistryCatalogData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr-public",
		OperationName: "PutRegistryCatalogData",
	}
}
