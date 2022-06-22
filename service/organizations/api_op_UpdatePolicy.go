// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing policy with a new name, description, or content. If you
// don't supply any parameter, that value remains unchanged. You can't change a
// policy's type. This operation can be called only from the organization's
// management account.
func (c *Client) UpdatePolicy(ctx context.Context, params *UpdatePolicyInput, optFns ...func(*Options)) (*UpdatePolicyOutput, error) {
	if params == nil {
		params = &UpdatePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePolicy", params, optFns, c.addOperationUpdatePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePolicyInput struct {

	// The unique identifier (ID) of the policy that you want to update. The regex
	// pattern (http://wikipedia.org/wiki/regex) for a policy ID string requires "p-"
	// followed by from 8 to 128 lowercase or uppercase letters, digits, or the
	// underscore character (_).
	//
	// This member is required.
	PolicyId *string

	// If provided, the new content for the policy. The text must be correctly
	// formatted JSON that complies with the syntax for the policy's type. For more
	// information, see Service Control Policy Syntax
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_scp-syntax.html)
	// in the Organizations User Guide.
	Content *string

	// If provided, the new description for the policy.
	Description *string

	// If provided, the new name for the policy. The regex pattern
	// (http://wikipedia.org/wiki/regex) that is used to validate this parameter is a
	// string of any of the characters in the ASCII character range.
	Name *string

	noSmithyDocumentSerde
}

type UpdatePolicyOutput struct {

	// A structure that contains details about the updated policy, showing the
	// requested changes.
	Policy *types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePolicy{}, middleware.After)
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
	if err = addOpUpdatePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "UpdatePolicy",
	}
}
