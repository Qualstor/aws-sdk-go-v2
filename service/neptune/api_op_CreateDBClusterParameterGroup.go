// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new DB cluster parameter group. Parameters in a DB cluster parameter
// group apply to all of the instances in a DB cluster. A DB cluster parameter
// group is initially created with the default parameters for the database engine
// used by instances in the DB cluster. To provide custom values for any of the
// parameters, you must modify the group after creating it using
// ModifyDBClusterParameterGroup. Once you've created a DB cluster parameter group,
// you need to associate it with your DB cluster using ModifyDBCluster. When you
// associate a new DB cluster parameter group with a running DB cluster, you need
// to reboot the DB instances in the DB cluster without failover for the new DB
// cluster parameter group and associated settings to take effect. After you create
// a DB cluster parameter group, you should wait at least 5 minutes before creating
// your first DB cluster that uses that DB cluster parameter group as the default
// parameter group. This allows Amazon Neptune to fully complete the create action
// before the DB cluster parameter group is used as the default for a new DB
// cluster. This is especially important for parameters that are critical when
// creating the default database for a DB cluster, such as the character set for
// the default database defined by the character_set_database parameter. You can
// use the Parameter Groups option of the Amazon Neptune console
// (https://console.aws.amazon.com/rds/) or the DescribeDBClusterParameters command
// to verify that your DB cluster parameter group has been created or modified.
func (c *Client) CreateDBClusterParameterGroup(ctx context.Context, params *CreateDBClusterParameterGroupInput, optFns ...func(*Options)) (*CreateDBClusterParameterGroupOutput, error) {
	if params == nil {
		params = &CreateDBClusterParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBClusterParameterGroup", params, optFns, c.addOperationCreateDBClusterParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBClusterParameterGroupInput struct {

	// The name of the DB cluster parameter group. Constraints:
	//
	// * Must match the name
	// of an existing DBClusterParameterGroup.
	//
	// This value is stored as a lowercase
	// string.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// The DB cluster parameter group family name. A DB cluster parameter group can be
	// associated with one and only one DB cluster parameter group family, and can be
	// applied only to a DB cluster running a database engine and engine version
	// compatible with that DB cluster parameter group family.
	//
	// This member is required.
	DBParameterGroupFamily *string

	// The description for the DB cluster parameter group.
	//
	// This member is required.
	Description *string

	// The tags to be assigned to the new DB cluster parameter group.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDBClusterParameterGroupOutput struct {

	// Contains the details of an Amazon Neptune DB cluster parameter group. This data
	// type is used as a response element in the DescribeDBClusterParameterGroups
	// action.
	DBClusterParameterGroup *types.DBClusterParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBClusterParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBClusterParameterGroup{}, middleware.After)
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
	if err = addOpCreateDBClusterParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBClusterParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBClusterParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBClusterParameterGroup",
	}
}
