// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the retention settings for an AppInstance.
func (c *Client) GetAppInstanceRetentionSettings(ctx context.Context, params *GetAppInstanceRetentionSettingsInput, optFns ...func(*Options)) (*GetAppInstanceRetentionSettingsOutput, error) {
	if params == nil {
		params = &GetAppInstanceRetentionSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAppInstanceRetentionSettings", params, optFns, c.addOperationGetAppInstanceRetentionSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAppInstanceRetentionSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAppInstanceRetentionSettingsInput struct {

	// The ARN of the AppInstance.
	//
	// This member is required.
	AppInstanceArn *string

	noSmithyDocumentSerde
}

type GetAppInstanceRetentionSettingsOutput struct {

	// The retention settings for the AppInstance.
	AppInstanceRetentionSettings *types.AppInstanceRetentionSettings

	// The timestamp representing the time at which the specified items are retained,
	// in Epoch Seconds.
	InitiateDeletionTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAppInstanceRetentionSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAppInstanceRetentionSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAppInstanceRetentionSettings{}, middleware.After)
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
	if err = addOpGetAppInstanceRetentionSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAppInstanceRetentionSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAppInstanceRetentionSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetAppInstanceRetentionSettings",
	}
}
