// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration of the storage optimizers for a table.
func (c *Client) UpdateTableStorageOptimizer(ctx context.Context, params *UpdateTableStorageOptimizerInput, optFns ...func(*Options)) (*UpdateTableStorageOptimizerOutput, error) {
	if params == nil {
		params = &UpdateTableStorageOptimizerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTableStorageOptimizer", params, optFns, c.addOperationUpdateTableStorageOptimizerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTableStorageOptimizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTableStorageOptimizerInput struct {

	// Name of the database where the table is present.
	//
	// This member is required.
	DatabaseName *string

	// Name of the table for which to enable the storage optimizer.
	//
	// This member is required.
	StorageOptimizerConfig map[string]map[string]string

	// Name of the table for which to enable the storage optimizer.
	//
	// This member is required.
	TableName *string

	// The Catalog ID of the table.
	CatalogId *string

	noSmithyDocumentSerde
}

type UpdateTableStorageOptimizerOutput struct {

	// A response indicating the success of failure of the operation.
	Result *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTableStorageOptimizerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTableStorageOptimizer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTableStorageOptimizer{}, middleware.After)
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
	if err = addOpUpdateTableStorageOptimizerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTableStorageOptimizer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTableStorageOptimizer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "UpdateTableStorageOptimizer",
	}
}
