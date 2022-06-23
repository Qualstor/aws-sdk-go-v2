// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns details about a specific version of a specific application.
func (c *Client) GetApplicationVersion(ctx context.Context, params *GetApplicationVersionInput, optFns ...func(*Options)) (*GetApplicationVersionOutput, error) {
	if params == nil {
		params = &GetApplicationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApplicationVersion", params, optFns, c.addOperationGetApplicationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApplicationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApplicationVersionInput struct {

	// The unique identifier of the application.
	//
	// This member is required.
	ApplicationId *string

	// The specific version of the application.
	//
	// This member is required.
	ApplicationVersion *int32

	noSmithyDocumentSerde
}

type GetApplicationVersionOutput struct {

	// The specific version of the application.
	//
	// This member is required.
	ApplicationVersion *int32

	// The timestamp when the application version was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The content of the application definition. This is a JSON object that contains
	// the resource configuration/definitions that identify an application.
	//
	// This member is required.
	DefinitionContent *string

	// The name of the application version.
	//
	// This member is required.
	Name *string

	// The status of the application version.
	//
	// This member is required.
	Status types.ApplicationVersionLifecycle

	// The application description.
	Description *string

	// The reason for the reported status.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetApplicationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetApplicationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApplicationVersion{}, middleware.After)
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
	if err = addOpGetApplicationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetApplicationVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetApplicationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "m2",
		OperationName: "GetApplicationVersion",
	}
}
