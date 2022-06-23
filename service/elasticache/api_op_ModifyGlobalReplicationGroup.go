// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the settings for a Global datastore.
func (c *Client) ModifyGlobalReplicationGroup(ctx context.Context, params *ModifyGlobalReplicationGroupInput, optFns ...func(*Options)) (*ModifyGlobalReplicationGroupOutput, error) {
	if params == nil {
		params = &ModifyGlobalReplicationGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyGlobalReplicationGroup", params, optFns, c.addOperationModifyGlobalReplicationGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyGlobalReplicationGroupInput struct {

	// This parameter causes the modifications in this request and any pending
	// modifications to be applied, asynchronously and as soon as possible.
	// Modifications to Global Replication Groups cannot be requested to be applied in
	// PreferredMaintenceWindow.
	//
	// This member is required.
	ApplyImmediately bool

	// The name of the Global datastore
	//
	// This member is required.
	GlobalReplicationGroupId *string

	// Determines whether a read replica is automatically promoted to read/write
	// primary if the existing primary encounters a failure.
	AutomaticFailoverEnabled *bool

	// A valid cache node type that you want to scale this Global datastore to.
	CacheNodeType *string

	// The name of the cache parameter group to use with the Global datastore. It must
	// be compatible with the major engine version used by the Global datastore.
	CacheParameterGroupName *string

	// The upgraded version of the cache engine to be run on the clusters in the Global
	// datastore.
	EngineVersion *string

	// A description of the Global datastore
	GlobalReplicationGroupDescription *string

	noSmithyDocumentSerde
}

type ModifyGlobalReplicationGroupOutput struct {

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different Amazon region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.
	//
	// * The GlobalReplicationGroupIdSuffix represents the name of
	// the Global datastore, which is what you use to associate a secondary cluster.
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyGlobalReplicationGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyGlobalReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyGlobalReplicationGroup{}, middleware.After)
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
	if err = addOpModifyGlobalReplicationGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyGlobalReplicationGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyGlobalReplicationGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ModifyGlobalReplicationGroup",
	}
}
