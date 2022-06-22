// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// In an environment account, get the detailed data for an environment account
// connection. For more information, see Environment account connections
// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-env-account-connections.html)
// in the Proton Administrator guide.
func (c *Client) GetEnvironmentAccountConnection(ctx context.Context, params *GetEnvironmentAccountConnectionInput, optFns ...func(*Options)) (*GetEnvironmentAccountConnectionOutput, error) {
	if params == nil {
		params = &GetEnvironmentAccountConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnvironmentAccountConnection", params, optFns, c.addOperationGetEnvironmentAccountConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnvironmentAccountConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnvironmentAccountConnectionInput struct {

	// The ID of the environment account connection that you want to get the detailed
	// data for.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetEnvironmentAccountConnectionOutput struct {

	// The detailed data of the requested environment account connection.
	//
	// This member is required.
	EnvironmentAccountConnection *types.EnvironmentAccountConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEnvironmentAccountConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEnvironmentAccountConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEnvironmentAccountConnection{}, middleware.After)
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
	if err = addOpGetEnvironmentAccountConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnvironmentAccountConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEnvironmentAccountConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "GetEnvironmentAccountConnection",
	}
}
