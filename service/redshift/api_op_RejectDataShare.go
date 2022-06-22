// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// From a datashare consumer account, rejects the specified datashare.
func (c *Client) RejectDataShare(ctx context.Context, params *RejectDataShareInput, optFns ...func(*Options)) (*RejectDataShareOutput, error) {
	if params == nil {
		params = &RejectDataShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RejectDataShare", params, optFns, c.addOperationRejectDataShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RejectDataShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RejectDataShareInput struct {

	// The Amazon Resource Name (ARN) of the datashare to reject.
	//
	// This member is required.
	DataShareArn *string

	noSmithyDocumentSerde
}

type RejectDataShareOutput struct {

	// A value that specifies whether the datashare can be shared to a publicly
	// accessible cluster.
	AllowPubliclyAccessibleConsumers bool

	// An Amazon Resource Name (ARN) that references the datashare that is owned by a
	// specific namespace of the producer cluster. A datashare ARN is in the
	// arn:aws:redshift:{region}:{account-id}:{datashare}:{namespace-guid}/{datashare-name}
	// format.
	DataShareArn *string

	// A value that specifies when the datashare has an association between producer
	// and data consumers.
	DataShareAssociations []types.DataShareAssociation

	// The identifier of a datashare to show its managing entity.
	ManagedBy *string

	// The Amazon Resource Name (ARN) of the producer.
	ProducerArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRejectDataShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRejectDataShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRejectDataShare{}, middleware.After)
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
	if err = addOpRejectDataShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRejectDataShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRejectDataShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "RejectDataShare",
	}
}
