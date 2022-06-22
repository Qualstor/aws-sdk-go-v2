// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deregisters an account in Audit Manager. When you deregister your account from
// Audit Manager, your data isn’t deleted. If you want to delete your resource
// data, you must perform that task separately before you deregister your account.
// Either, you can do this in the Audit Manager console. Or, you can use one of the
// delete API operations that are provided by Audit Manager. To delete your Audit
// Manager resource data, see the following instructions:
//
// * DeleteAssessment
// (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessment.html)
// (see also: Deleting an assessment
// (https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-assessment.html)
// in the Audit Manager User Guide)
//
// * DeleteAssessmentFramework
// (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentFramework.html)
// (see also: Deleting a custom framework
// (https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-custom-framework.html)
// in the Audit Manager User Guide)
//
// * DeleteAssessmentFrameworkShare
// (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentFrameworkShare.html)
// (see also: Deleting a share request
// (https://docs.aws.amazon.com/audit-manager/latest/userguide/deleting-shared-framework-requests.html)
// in the Audit Manager User Guide)
//
// * DeleteAssessmentReport
// (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentReport.html)
// (see also: Deleting an assessment report
// (https://docs.aws.amazon.com/audit-manager/latest/userguide/generate-assessment-report.html#delete-assessment-report-steps)
// in the Audit Manager User Guide)
//
// * DeleteControl
// (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteControl.html)
// (see also: Deleting a custom control
// (https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-controls.html)
// in the Audit Manager User Guide)
//
// At this time, Audit Manager doesn't provide an
// option to delete evidence. All available delete operations are listed above.
func (c *Client) DeregisterAccount(ctx context.Context, params *DeregisterAccountInput, optFns ...func(*Options)) (*DeregisterAccountOutput, error) {
	if params == nil {
		params = &DeregisterAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterAccount", params, optFns, c.addOperationDeregisterAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterAccountInput struct {
	noSmithyDocumentSerde
}

type DeregisterAccountOutput struct {

	// The registration status of the account.
	Status types.AccountStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeregisterAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeregisterAccount{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "DeregisterAccount",
	}
}
