// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables Health to work with Organizations. You can use the organizational view
// feature to aggregate events from all Amazon Web Services accounts in your
// organization in a centralized location. This operation also creates a
// service-linked role for the management account in the organization. To call this
// operation, you must meet the following requirements:
//
// * You must have a
// Business, Enterprise On-Ramp, or Enterprise Support plan from Amazon Web
// Services Support (http://aws.amazon.com/premiumsupport/) to use the Health API.
// If you call the Health API from an Amazon Web Services account that doesn't have
// a Business, Enterprise On-Ramp, or Enterprise Support plan, you receive a
// SubscriptionRequiredException error.
//
// * You must have permission to call this
// operation from the organization's management account. For example IAM policies,
// see Health identity-based policy examples
// (https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html).
//
// If
// you don't have the required support plan, you can instead use the Health console
// to enable the organizational view feature. For more information, see Aggregating
// Health events
// (https://docs.aws.amazon.com/health/latest/ug/aggregate-events.html) in the
// Health User Guide.
func (c *Client) EnableHealthServiceAccessForOrganization(ctx context.Context, params *EnableHealthServiceAccessForOrganizationInput, optFns ...func(*Options)) (*EnableHealthServiceAccessForOrganizationOutput, error) {
	if params == nil {
		params = &EnableHealthServiceAccessForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableHealthServiceAccessForOrganization", params, optFns, c.addOperationEnableHealthServiceAccessForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableHealthServiceAccessForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableHealthServiceAccessForOrganizationInput struct {
	noSmithyDocumentSerde
}

type EnableHealthServiceAccessForOrganizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableHealthServiceAccessForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableHealthServiceAccessForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableHealthServiceAccessForOrganization{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableHealthServiceAccessForOrganization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableHealthServiceAccessForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "EnableHealthServiceAccessForOrganization",
	}
}
