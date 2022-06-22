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

// Updates a custom framework in Audit Manager.
func (c *Client) UpdateAssessmentFramework(ctx context.Context, params *UpdateAssessmentFrameworkInput, optFns ...func(*Options)) (*UpdateAssessmentFrameworkOutput, error) {
	if params == nil {
		params = &UpdateAssessmentFrameworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAssessmentFramework", params, optFns, c.addOperationUpdateAssessmentFrameworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAssessmentFrameworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssessmentFrameworkInput struct {

	// The control sets that are associated with the framework.
	//
	// This member is required.
	ControlSets []types.UpdateAssessmentFrameworkControlSet

	// The unique identifier for the framework.
	//
	// This member is required.
	FrameworkId *string

	// The name of the framework to be updated.
	//
	// This member is required.
	Name *string

	// The compliance type that the new custom framework supports, such as CIS or
	// HIPAA.
	ComplianceType *string

	// The description of the updated framework.
	Description *string

	noSmithyDocumentSerde
}

type UpdateAssessmentFrameworkOutput struct {

	// The name of the framework.
	Framework *types.Framework

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAssessmentFrameworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAssessmentFramework{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAssessmentFramework{}, middleware.After)
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
	if err = addOpUpdateAssessmentFrameworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAssessmentFramework(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAssessmentFramework(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "UpdateAssessmentFramework",
	}
}
