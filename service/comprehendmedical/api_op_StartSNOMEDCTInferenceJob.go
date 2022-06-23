// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous job to detect medical concepts and link them to the
// SNOMED-CT ontology. Use the DescribeSNOMEDCTInferenceJob operation to track the
// status of a job.
func (c *Client) StartSNOMEDCTInferenceJob(ctx context.Context, params *StartSNOMEDCTInferenceJobInput, optFns ...func(*Options)) (*StartSNOMEDCTInferenceJobOutput, error) {
	if params == nil {
		params = &StartSNOMEDCTInferenceJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSNOMEDCTInferenceJob", params, optFns, c.addOperationStartSNOMEDCTInferenceJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSNOMEDCTInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSNOMEDCTInferenceJobInput struct {

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that grants Amazon Comprehend Medical read access to your input data.
	//
	// This member is required.
	DataAccessRoleArn *string

	// The input properties for an entities detection job. This includes the name of
	// the S3 bucket and the path to the files to be analyzed.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// The language of the input documents. All documents must be in the same language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The output properties for a detection job.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// A unique identifier for the request. If you don't set the client request token,
	// Amazon Comprehend Medical generates one.
	ClientRequestToken *string

	// The user generated name the asynchronous InferSNOMEDCT job.
	JobName *string

	// An AWS Key Management Service key used to encrypt your output files. If you do
	// not specify a key, the files are written in plain text.
	KMSKey *string

	noSmithyDocumentSerde
}

type StartSNOMEDCTInferenceJobOutput struct {

	// The identifier generated for the job. To get the status of a job, use this
	// identifier with the StartSNOMEDCTInferenceJob operation.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSNOMEDCTInferenceJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartSNOMEDCTInferenceJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSNOMEDCTInferenceJob{}, middleware.After)
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
	if err = addIdempotencyToken_opStartSNOMEDCTInferenceJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartSNOMEDCTInferenceJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSNOMEDCTInferenceJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartSNOMEDCTInferenceJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartSNOMEDCTInferenceJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartSNOMEDCTInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartSNOMEDCTInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartSNOMEDCTInferenceJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartSNOMEDCTInferenceJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartSNOMEDCTInferenceJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartSNOMEDCTInferenceJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "StartSNOMEDCTInferenceJob",
	}
}
