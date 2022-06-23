// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Confirms a job worker has received the specified job. Used for partner actions
// only.
func (c *Client) AcknowledgeThirdPartyJob(ctx context.Context, params *AcknowledgeThirdPartyJobInput, optFns ...func(*Options)) (*AcknowledgeThirdPartyJobOutput, error) {
	if params == nil {
		params = &AcknowledgeThirdPartyJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcknowledgeThirdPartyJob", params, optFns, c.addOperationAcknowledgeThirdPartyJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcknowledgeThirdPartyJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an AcknowledgeThirdPartyJob action.
type AcknowledgeThirdPartyJobInput struct {

	// The clientToken portion of the clientId and clientToken pair used to verify that
	// the calling entity is allowed access to the job and its details.
	//
	// This member is required.
	ClientToken *string

	// The unique system-generated ID of the job.
	//
	// This member is required.
	JobId *string

	// A system-generated random number that AWS CodePipeline uses to ensure that the
	// job is being worked on by only one job worker. Get this number from the response
	// to a GetThirdPartyJobDetails request.
	//
	// This member is required.
	Nonce *string

	noSmithyDocumentSerde
}

// Represents the output of an AcknowledgeThirdPartyJob action.
type AcknowledgeThirdPartyJobOutput struct {

	// The status information for the third party job, if any.
	Status types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAcknowledgeThirdPartyJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAcknowledgeThirdPartyJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAcknowledgeThirdPartyJob{}, middleware.After)
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
	if err = addOpAcknowledgeThirdPartyJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAcknowledgeThirdPartyJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAcknowledgeThirdPartyJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "AcknowledgeThirdPartyJob",
	}
}
