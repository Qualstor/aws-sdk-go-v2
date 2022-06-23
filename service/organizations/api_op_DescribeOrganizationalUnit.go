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

// Retrieves information about an organizational unit (OU). This operation can be
// called only from the organization's management account or by a member account
// that is a delegated administrator for an Amazon Web Services service.
func (c *Client) DescribeOrganizationalUnit(ctx context.Context, params *DescribeOrganizationalUnitInput, optFns ...func(*Options)) (*DescribeOrganizationalUnitOutput, error) {
	if params == nil {
		params = &DescribeOrganizationalUnitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationalUnit", params, optFns, c.addOperationDescribeOrganizationalUnitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationalUnitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationalUnitInput struct {

	// The unique identifier (ID) of the organizational unit that you want details
	// about. You can get the ID from the ListOrganizationalUnitsForParent operation.
	// The regex pattern (http://wikipedia.org/wiki/regex) for an organizational unit
	// ID string requires "ou-" followed by from 4 to 32 lowercase letters or digits
	// (the ID of the root that contains the OU). This string is followed by a second
	// "-" dash and from 8 to 32 additional lowercase letters or digits.
	//
	// This member is required.
	OrganizationalUnitId *string

	noSmithyDocumentSerde
}

type DescribeOrganizationalUnitOutput struct {

	// A structure that contains details about the specified OU.
	OrganizationalUnit *types.OrganizationalUnit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationalUnitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganizationalUnit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganizationalUnit{}, middleware.After)
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
	if err = addOpDescribeOrganizationalUnitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationalUnit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeOrganizationalUnit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DescribeOrganizationalUnit",
	}
}
