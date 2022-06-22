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

// Creates a new function definition in the Data Catalog.
func (c *Client) CreateUserDefinedFunction(ctx context.Context, params *CreateUserDefinedFunctionInput, optFns ...func(*Options)) (*CreateUserDefinedFunctionOutput, error) {
	if params == nil {
		params = &CreateUserDefinedFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUserDefinedFunction", params, optFns, c.addOperationCreateUserDefinedFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserDefinedFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserDefinedFunctionInput struct {

	// The name of the catalog database in which to create the function.
	//
	// This member is required.
	DatabaseName *string

	// A FunctionInput object that defines the function to create in the Data Catalog.
	//
	// This member is required.
	FunctionInput *types.UserDefinedFunctionInput

	// The ID of the Data Catalog in which to create the function. If none is provided,
	// the Amazon Web Services account ID is used by default.
	CatalogId *string

	noSmithyDocumentSerde
}

type CreateUserDefinedFunctionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUserDefinedFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUserDefinedFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUserDefinedFunction{}, middleware.After)
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
	if err = addOpCreateUserDefinedFunctionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserDefinedFunction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUserDefinedFunction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateUserDefinedFunction",
	}
}
