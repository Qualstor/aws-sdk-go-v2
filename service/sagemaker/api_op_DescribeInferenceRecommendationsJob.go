// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides the results of the Inference Recommender job. One or more
// recommendation jobs are returned.
func (c *Client) DescribeInferenceRecommendationsJob(ctx context.Context, params *DescribeInferenceRecommendationsJobInput, optFns ...func(*Options)) (*DescribeInferenceRecommendationsJobOutput, error) {
	if params == nil {
		params = &DescribeInferenceRecommendationsJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInferenceRecommendationsJob", params, optFns, c.addOperationDescribeInferenceRecommendationsJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInferenceRecommendationsJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInferenceRecommendationsJobInput struct {

	// The name of the job. The name must be unique within an Amazon Web Services
	// Region in the Amazon Web Services account.
	//
	// This member is required.
	JobName *string

	noSmithyDocumentSerde
}

type DescribeInferenceRecommendationsJobOutput struct {

	// A timestamp that shows when the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// Returns information about the versioned model package Amazon Resource Name
	// (ARN), the traffic pattern, and endpoint configurations you provided when you
	// initiated the job.
	//
	// This member is required.
	InputConfig *types.RecommendationJobInputConfig

	// The Amazon Resource Name (ARN) of the job.
	//
	// This member is required.
	JobArn *string

	// The name of the job. The name must be unique within an Amazon Web Services
	// Region in the Amazon Web Services account.
	//
	// This member is required.
	JobName *string

	// The job type that you provided when you initiated the job.
	//
	// This member is required.
	JobType types.RecommendationJobType

	// A timestamp that shows when the job was last modified.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The Amazon Resource Name (ARN) of the Amazon Web Services Identity and Access
	// Management (IAM) role you provided when you initiated the job.
	//
	// This member is required.
	RoleArn *string

	// The status of the job.
	//
	// This member is required.
	Status types.RecommendationJobStatus

	// A timestamp that shows when the job completed.
	CompletionTime *time.Time

	// If the job fails, provides information why the job failed.
	FailureReason *string

	// The recommendations made by Inference Recommender.
	InferenceRecommendations []types.InferenceRecommendation

	// The job description that you provided when you initiated the job.
	JobDescription *string

	// The stopping conditions that you provided when you initiated the job.
	StoppingConditions *types.RecommendationJobStoppingConditions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInferenceRecommendationsJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInferenceRecommendationsJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInferenceRecommendationsJob{}, middleware.After)
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
	if err = addOpDescribeInferenceRecommendationsJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInferenceRecommendationsJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInferenceRecommendationsJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeInferenceRecommendationsJob",
	}
}
