// Code generated by smithy-go-codegen DO NOT EDIT.

package appintegrations

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/appintegrations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the DataIntegration. You cannot create a
// DataIntegration association for a DataIntegration that has been previously
// associated. Use a different DataIntegration, or recreate the DataIntegration
// using the CreateDataIntegration
// (https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html)
// API.
func (c *Client) GetDataIntegration(ctx context.Context, params *GetDataIntegrationInput, optFns ...func(*Options)) (*GetDataIntegrationOutput, error) {
	if params == nil {
		params = &GetDataIntegrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataIntegration", params, optFns, c.addOperationGetDataIntegrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataIntegrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataIntegrationInput struct {

	// A unique identifier.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetDataIntegrationOutput struct {

	// The Amazon Resource Name (ARN) for the DataIntegration.
	Arn *string

	// The KMS key for the DataIntegration.
	Description *string

	// A unique identifier.
	Id *string

	// The KMS key for the DataIntegration.
	KmsKey *string

	// The name of the DataIntegration.
	Name *string

	// The name of the data and how often it should be pulled from the source.
	ScheduleConfiguration *types.ScheduleConfiguration

	// The URI of the data source.
	SourceURI *string

	// One or more tags.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataIntegrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataIntegration{}, middleware.After)
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
	if err = addOpGetDataIntegrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataIntegration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataIntegration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "app-integrations",
		OperationName: "GetDataIntegration",
	}
}
