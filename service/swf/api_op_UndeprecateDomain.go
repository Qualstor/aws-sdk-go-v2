// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Undeprecates a previously deprecated domain. After a domain has been
// undeprecated it can be used to create new workflow executions or register new
// types. This operation is eventually consistent. The results are best effort and
// may not exactly reflect recent updates and changes. Access Control You can use
// IAM policies to control this action's access to Amazon SWF resources as
// follows:
//
// * Use a Resource element with the domain name to limit the action to
// only specified domains.
//
// * Use an Action element to allow or deny permission to
// call this action.
//
// * You cannot use an IAM policy to constrain this action's
// parameters.
//
// If the caller doesn't have sufficient permissions to invoke the
// action, or the parameter values fall outside the specified constraints, the
// action fails. The associated event attribute's cause parameter is set to
// OPERATION_NOT_PERMITTED. For details and example IAM policies, see Using IAM to
// Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) UndeprecateDomain(ctx context.Context, params *UndeprecateDomainInput, optFns ...func(*Options)) (*UndeprecateDomainOutput, error) {
	if params == nil {
		params = &UndeprecateDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UndeprecateDomain", params, optFns, c.addOperationUndeprecateDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UndeprecateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UndeprecateDomainInput struct {

	// The name of the domain of the deprecated workflow type.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type UndeprecateDomainOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUndeprecateDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUndeprecateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUndeprecateDomain{}, middleware.After)
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
	if err = addOpUndeprecateDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUndeprecateDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUndeprecateDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "UndeprecateDomain",
	}
}
