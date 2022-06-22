// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/migrationhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a created artifact of an AWS cloud resource, the target receiving the
// migration, with the migration task performed by a migration tool. This API has
// the following traits:
//
// * Migration tools can call the AssociateCreatedArtifact
// operation to indicate which AWS artifact is associated with a migration task.
//
// *
// The created artifact name must be provided in ARN (Amazon Resource Name) format
// which will contain information about type and region; for example:
// arn:aws:ec2:us-east-1:488216288981:image/ami-6d0ba87b.
//
// * Examples of the AWS
// resource behind the created artifact are, AMI's, EC2 instance, or DMS endpoint,
// etc.
func (c *Client) AssociateCreatedArtifact(ctx context.Context, params *AssociateCreatedArtifactInput, optFns ...func(*Options)) (*AssociateCreatedArtifactOutput, error) {
	if params == nil {
		params = &AssociateCreatedArtifactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateCreatedArtifact", params, optFns, c.addOperationAssociateCreatedArtifactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateCreatedArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateCreatedArtifactInput struct {

	// An ARN of the AWS resource related to the migration (e.g., AMI, EC2 instance,
	// RDS instance, etc.)
	//
	// This member is required.
	CreatedArtifact *types.CreatedArtifact

	// Unique identifier that references the migration task. Do not store personal data
	// in this field.
	//
	// This member is required.
	MigrationTaskName *string

	// The name of the ProgressUpdateStream.
	//
	// This member is required.
	ProgressUpdateStream *string

	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun bool

	noSmithyDocumentSerde
}

type AssociateCreatedArtifactOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateCreatedArtifactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateCreatedArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateCreatedArtifact{}, middleware.After)
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
	if err = addOpAssociateCreatedArtifactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateCreatedArtifact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateCreatedArtifact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "AssociateCreatedArtifact",
	}
}
