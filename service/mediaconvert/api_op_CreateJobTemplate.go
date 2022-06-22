// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new job template. For information about job templates see the User
// Guide at http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
func (c *Client) CreateJobTemplate(ctx context.Context, params *CreateJobTemplateInput, optFns ...func(*Options)) (*CreateJobTemplateOutput, error) {
	if params == nil {
		params = &CreateJobTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateJobTemplate", params, optFns, c.addOperationCreateJobTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateJobTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateJobTemplateInput struct {

	// The name of the job template you are creating.
	//
	// This member is required.
	Name *string

	// JobTemplateSettings contains all the transcode settings saved in the template
	// that will be applied to jobs created from it.
	//
	// This member is required.
	Settings *types.JobTemplateSettings

	// Accelerated transcoding can significantly speed up jobs with long, visually
	// complex content. Outputs that use this feature incur pro-tier pricing. For
	// information about feature limitations, see the AWS Elemental MediaConvert User
	// Guide.
	AccelerationSettings *types.AccelerationSettings

	// Optional. A category for the job template you are creating
	Category *string

	// Optional. A description of the job template you are creating.
	Description *string

	// Optional. Use queue hopping to avoid overly long waits in the backlog of the
	// queue that you submit your job to. Specify an alternate queue and the maximum
	// time that your job will wait in the initial queue before hopping. For more
	// information about this feature, see the AWS Elemental MediaConvert User Guide.
	HopDestinations []types.HopDestination

	// Specify the relative priority for this job. In any given queue, the service
	// begins processing the job with the highest value first. When more than one job
	// has the same priority, the service begins processing the job that you submitted
	// first. If you don't specify a priority, the service uses the default value 0.
	Priority int32

	// Optional. The queue that jobs created from this template are assigned to. If you
	// don't specify this, jobs will go to the default queue.
	Queue *string

	// Specify how often MediaConvert sends STATUS_UPDATE events to Amazon CloudWatch
	// Events. Set the interval, in seconds, between status updates. MediaConvert sends
	// an update at this interval from the time the service begins processing your job
	// to the time it completes the transcode or encounters an error.
	StatusUpdateInterval types.StatusUpdateInterval

	// The tags that you want to add to the resource. You can tag resources with a
	// key-value pair or with only a key.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateJobTemplateOutput struct {

	// A job template is a pre-made set of encoding instructions that you can use to
	// quickly create a job.
	JobTemplate *types.JobTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateJobTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateJobTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateJobTemplate{}, middleware.After)
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
	if err = addOpCreateJobTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJobTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateJobTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "CreateJobTemplate",
	}
}
