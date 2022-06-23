// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the asssociation of an Amazon Web Services Identity and Access
// Management (IAM) role from a DB cluster. For more information on Amazon Aurora
// DB clusters, see  What is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
// see  Multi-AZ deployments with two readable standby DB instances
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
// in the Amazon RDS User Guide.
func (c *Client) RemoveRoleFromDBCluster(ctx context.Context, params *RemoveRoleFromDBClusterInput, optFns ...func(*Options)) (*RemoveRoleFromDBClusterOutput, error) {
	if params == nil {
		params = &RemoveRoleFromDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveRoleFromDBCluster", params, optFns, c.addOperationRemoveRoleFromDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveRoleFromDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveRoleFromDBClusterInput struct {

	// The name of the DB cluster to disassociate the IAM role from.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The Amazon Resource Name (ARN) of the IAM role to disassociate from the Aurora
	// DB cluster, for example arn:aws:iam::123456789012:role/AuroraAccessRole.
	//
	// This member is required.
	RoleArn *string

	// The name of the feature for the DB cluster that the IAM role is to be
	// disassociated from. For information about supported feature names, see
	// DBEngineVersion.
	FeatureName *string

	noSmithyDocumentSerde
}

type RemoveRoleFromDBClusterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveRoleFromDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRemoveRoleFromDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveRoleFromDBCluster{}, middleware.After)
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
	if err = addOpRemoveRoleFromDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveRoleFromDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveRoleFromDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RemoveRoleFromDBCluster",
	}
}
