// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a usage plan of a given plan Id.
func (c *Client) UpdateUsagePlan(ctx context.Context, params *UpdateUsagePlanInput, optFns ...func(*Options)) (*UpdateUsagePlanOutput, error) {
	if params == nil {
		params = &UpdateUsagePlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUsagePlan", params, optFns, c.addOperationUpdateUsagePlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUsagePlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The PATCH request to update a usage plan of a given plan Id.
type UpdateUsagePlanInput struct {

	// The Id of the to-be-updated usage plan.
	//
	// This member is required.
	UsagePlanId *string

	// For more information about supported patch operations, see Patch Operations
	// (https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).
	PatchOperations []types.PatchOperation

	noSmithyDocumentSerde
}

// Represents a usage plan used to specify who can assess associated API stages.
// Optionally, target request rate and quota limits can be set. In some cases
// clients can exceed the targets that you set. Don’t rely on usage plans to
// control costs. Consider using Amazon Web Services Budgets
// (https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-managing-costs.html)
// to monitor costs and WAF
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) to
// manage API requests.
type UpdateUsagePlanOutput struct {

	// The associated API stages of a usage plan.
	ApiStages []types.ApiStage

	// The description of a usage plan.
	Description *string

	// The identifier of a UsagePlan resource.
	Id *string

	// The name of a usage plan.
	Name *string

	// The AWS Markeplace product identifier to associate with the usage plan as a SaaS
	// product on AWS Marketplace.
	ProductCode *string

	// The target maximum number of permitted requests per a given unit time interval.
	Quota *types.QuotaSettings

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	// A map containing method level throttling information for API stage in a usage
	// plan.
	Throttle *types.ThrottleSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUsagePlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUsagePlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUsagePlan{}, middleware.After)
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
	if err = addOpUpdateUsagePlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUsagePlan(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateUsagePlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateUsagePlan",
	}
}
