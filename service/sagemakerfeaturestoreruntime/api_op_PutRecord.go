// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakerfeaturestoreruntime

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used for data ingestion into the FeatureStore. The PutRecord API writes to both
// the OnlineStore and OfflineStore. If the record is the latest record for the
// recordIdentifier, the record is written to both the OnlineStore and
// OfflineStore. If the record is a historic record, it is written only to the
// OfflineStore.
func (c *Client) PutRecord(ctx context.Context, params *PutRecordInput, optFns ...func(*Options)) (*PutRecordOutput, error) {
	if params == nil {
		params = &PutRecordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRecord", params, optFns, c.addOperationPutRecordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRecordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRecordInput struct {

	// The name of the feature group that you want to insert the record into.
	//
	// This member is required.
	FeatureGroupName *string

	// List of FeatureValues to be inserted. This will be a full over-write. If you
	// only want to update few of the feature values, do the following:
	//
	// * Use
	// GetRecord to retrieve the latest record.
	//
	// * Update the record returned from
	// GetRecord.
	//
	// * Use PutRecord to update feature values.
	//
	// This member is required.
	Record []types.FeatureValue

	noSmithyDocumentSerde
}

type PutRecordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRecordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutRecord{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutRecord{}, middleware.After)
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
	if err = addOpPutRecordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRecord(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRecord(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "PutRecord",
	}
}
