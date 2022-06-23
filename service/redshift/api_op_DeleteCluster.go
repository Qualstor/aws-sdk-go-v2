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

// Deletes a previously provisioned cluster without its final snapshot being
// created. A successful response from the web service indicates that the request
// was received correctly. Use DescribeClusters to monitor the status of the
// deletion. The delete operation cannot be canceled or reverted once submitted.
// For more information about managing clusters, go to Amazon Redshift Clusters
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in
// the Amazon Redshift Cluster Management Guide. If you want to shut down the
// cluster and retain it for future use, set SkipFinalClusterSnapshot to false and
// specify a name for FinalClusterSnapshotIdentifier. You can later restore this
// snapshot to resume using the cluster. If a final cluster snapshot is requested,
// the status of the cluster will be "final-snapshot" while the snapshot is being
// taken, then it's "deleting" once Amazon Redshift begins deleting the cluster.
// For more information about managing clusters, go to Amazon Redshift Clusters
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in
// the Amazon Redshift Cluster Management Guide.
func (c *Client) DeleteCluster(ctx context.Context, params *DeleteClusterInput, optFns ...func(*Options)) (*DeleteClusterOutput, error) {
	if params == nil {
		params = &DeleteClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCluster", params, optFns, c.addOperationDeleteClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteClusterInput struct {

	// The identifier of the cluster to be deleted. Constraints:
	//
	// * Must contain
	// lowercase characters.
	//
	// * Must contain from 1 to 63 alphanumeric characters or
	// hyphens.
	//
	// * First character must be a letter.
	//
	// * Cannot end with a hyphen or
	// contain two consecutive hyphens.
	//
	// This member is required.
	ClusterIdentifier *string

	// The identifier of the final snapshot that is to be created immediately before
	// deleting the cluster. If this parameter is provided, SkipFinalClusterSnapshot
	// must be false. Constraints:
	//
	// * Must be 1 to 255 alphanumeric characters.
	//
	// *
	// First character must be a letter.
	//
	// * Cannot end with a hyphen or contain two
	// consecutive hyphens.
	FinalClusterSnapshotIdentifier *string

	// The number of days that a manual snapshot is retained. If the value is -1, the
	// manual snapshot is retained indefinitely. The value must be either -1 or an
	// integer between 1 and 3,653. The default value is -1.
	FinalClusterSnapshotRetentionPeriod *int32

	// Determines whether a final snapshot of the cluster is created before Amazon
	// Redshift deletes the cluster. If true, a final cluster snapshot is not created.
	// If false, a final cluster snapshot is created before the cluster is deleted. The
	// FinalClusterSnapshotIdentifier parameter must be specified if
	// SkipFinalClusterSnapshot is false. Default: false
	SkipFinalClusterSnapshot bool

	noSmithyDocumentSerde
}

type DeleteClusterOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteCluster{}, middleware.After)
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
	if err = addOpDeleteClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DeleteCluster",
	}
}
