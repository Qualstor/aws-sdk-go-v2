// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the specified connection alias for use with cross-Region redirection.
// For more information, see  Cross-Region Redirection for Amazon WorkSpaces
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html).
func (c *Client) CreateConnectionAlias(ctx context.Context, params *CreateConnectionAliasInput, optFns ...func(*Options)) (*CreateConnectionAliasOutput, error) {
	if params == nil {
		params = &CreateConnectionAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConnectionAlias", params, optFns, c.addOperationCreateConnectionAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConnectionAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectionAliasInput struct {

	// A connection string in the form of a fully qualified domain name (FQDN), such as
	// www.example.com. After you create a connection string, it is always associated
	// to your Amazon Web Services account. You cannot recreate the same connection
	// string with a different account, even if you delete all instances of it from the
	// original account. The connection string is globally reserved for your account.
	//
	// This member is required.
	ConnectionString *string

	// The tags to associate with the connection alias.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateConnectionAliasOutput struct {

	// The identifier of the connection alias.
	AliasId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConnectionAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateConnectionAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateConnectionAlias{}, middleware.After)
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
	if err = addOpCreateConnectionAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConnectionAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConnectionAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "CreateConnectionAlias",
	}
}
