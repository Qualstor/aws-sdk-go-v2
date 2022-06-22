// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a program in a multiplex.
func (c *Client) UpdateMultiplexProgram(ctx context.Context, params *UpdateMultiplexProgramInput, optFns ...func(*Options)) (*UpdateMultiplexProgramOutput, error) {
	if params == nil {
		params = &UpdateMultiplexProgramInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMultiplexProgram", params, optFns, c.addOperationUpdateMultiplexProgramMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMultiplexProgramOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to update a program in a multiplex.
type UpdateMultiplexProgramInput struct {

	// The ID of the multiplex of the program to update.
	//
	// This member is required.
	MultiplexId *string

	// The name of the program to update.
	//
	// This member is required.
	ProgramName *string

	// The new settings for a multiplex program.
	MultiplexProgramSettings *types.MultiplexProgramSettings

	noSmithyDocumentSerde
}

// Placeholder documentation for UpdateMultiplexProgramResponse
type UpdateMultiplexProgramOutput struct {

	// The updated multiplex program.
	MultiplexProgram *types.MultiplexProgram

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMultiplexProgramMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMultiplexProgram{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMultiplexProgram{}, middleware.After)
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
	if err = addOpUpdateMultiplexProgramValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMultiplexProgram(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMultiplexProgram(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "UpdateMultiplexProgram",
	}
}
