// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Puts the metadata key value pair for a specified schema version ID. A maximum of
// 10 key value pairs will be allowed per schema version. They can be added over
// one or more calls.
func (c *Client) PutSchemaVersionMetadata(ctx context.Context, params *PutSchemaVersionMetadataInput, optFns ...func(*Options)) (*PutSchemaVersionMetadataOutput, error) {
	if params == nil {
		params = &PutSchemaVersionMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSchemaVersionMetadata", params, optFns, c.addOperationPutSchemaVersionMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSchemaVersionMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSchemaVersionMetadataInput struct {

	// The metadata key's corresponding value.
	//
	// This member is required.
	MetadataKeyValue *types.MetadataKeyValuePair

	// The unique ID for the schema.
	SchemaId *types.SchemaId

	// The unique version ID of the schema version.
	SchemaVersionId *string

	// The version number of the schema.
	SchemaVersionNumber *types.SchemaVersionNumber

	noSmithyDocumentSerde
}

type PutSchemaVersionMetadataOutput struct {

	// The latest version of the schema.
	LatestVersion bool

	// The metadata key.
	MetadataKey *string

	// The value of the metadata key.
	MetadataValue *string

	// The name for the registry.
	RegistryName *string

	// The Amazon Resource Name (ARN) for the schema.
	SchemaArn *string

	// The name for the schema.
	SchemaName *string

	// The unique version ID of the schema version.
	SchemaVersionId *string

	// The version number of the schema.
	VersionNumber int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSchemaVersionMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutSchemaVersionMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutSchemaVersionMetadata{}, middleware.After)
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
	if err = addOpPutSchemaVersionMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSchemaVersionMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSchemaVersionMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "PutSchemaVersionMetadata",
	}
}
