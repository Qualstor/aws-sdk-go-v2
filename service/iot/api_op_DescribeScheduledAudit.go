// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a scheduled audit. Requires permission to access the
// DescribeScheduledAudit
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) DescribeScheduledAudit(ctx context.Context, params *DescribeScheduledAuditInput, optFns ...func(*Options)) (*DescribeScheduledAuditOutput, error) {
	if params == nil {
		params = &DescribeScheduledAuditInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScheduledAudit", params, optFns, c.addOperationDescribeScheduledAuditMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScheduledAuditOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScheduledAuditInput struct {

	// The name of the scheduled audit whose information you want to get.
	//
	// This member is required.
	ScheduledAuditName *string

	noSmithyDocumentSerde
}

type DescribeScheduledAuditOutput struct {

	// The day of the month on which the scheduled audit takes place. This is will be 1
	// through 31 or LAST. If days 29-31 are specified, and the month does not have
	// that many days, the audit takes place on the LAST day of the month.
	DayOfMonth *string

	// The day of the week on which the scheduled audit takes place, either one of SUN,
	// MON, TUE, WED, THU, FRI, or SAT.
	DayOfWeek types.DayOfWeek

	// How often the scheduled audit takes place, either one of DAILY, WEEKLY,
	// BIWEEKLY, or MONTHLY. The start time of each audit is determined by the system.
	Frequency types.AuditFrequency

	// The ARN of the scheduled audit.
	ScheduledAuditArn *string

	// The name of the scheduled audit.
	ScheduledAuditName *string

	// Which checks are performed during the scheduled audit. Checks must be enabled
	// for your account. (Use DescribeAccountAuditConfiguration to see the list of all
	// checks, including those that are enabled or use UpdateAccountAuditConfiguration
	// to select which checks are enabled.)
	TargetCheckNames []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScheduledAuditMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeScheduledAudit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeScheduledAudit{}, middleware.After)
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
	if err = addOpDescribeScheduledAuditValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScheduledAudit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeScheduledAudit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeScheduledAudit",
	}
}
