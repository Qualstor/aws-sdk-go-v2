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

// Renames the specified organizational unit (OU). The ID and ARN don't change. The
// child OUs and accounts remain in place, and any attached policies of the OU
// remain attached. This operation can be called only from the organization's
// management account.
func (c *Client) UpdateOrganizationalUnit(ctx context.Context, params *UpdateOrganizationalUnitInput, optFns ...func(*Options)) (*UpdateOrganizationalUnitOutput, error) {
	if params == nil {
		params = &UpdateOrganizationalUnitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateOrganizationalUnit", params, optFns, c.addOperationUpdateOrganizationalUnitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateOrganizationalUnitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateOrganizationalUnitInput struct {

	// The unique identifier (ID) of the OU that you want to rename. You can get the ID
	// from the ListOrganizationalUnitsForParent operation. The regex pattern
	// (http://wikipedia.org/wiki/regex) for an organizational unit ID string requires
	// "ou-" followed by from 4 to 32 lowercase letters or digits (the ID of the root
	// that contains the OU). This string is followed by a second "-" dash and from 8
	// to 32 additional lowercase letters or digits.
	//
	// This member is required.
	OrganizationalUnitId *string

	// The new name that you want to assign to the OU. The regex pattern
	// (http://wikipedia.org/wiki/regex) that is used to validate this parameter is a
	// string of any of the characters in the ASCII character range.
	Name *string

	noSmithyDocumentSerde
}

type UpdateOrganizationalUnitOutput struct {

	// A structure that contains the details about the specified OU, including its new
	// name.
	OrganizationalUnit *types.OrganizationalUnit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateOrganizationalUnitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateOrganizationalUnit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateOrganizationalUnit{}, middleware.After)
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
	if err = addOpUpdateOrganizationalUnitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateOrganizationalUnit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateOrganizationalUnit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "UpdateOrganizationalUnit",
	}
}
