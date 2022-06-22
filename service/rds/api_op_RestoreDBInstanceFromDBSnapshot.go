// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new DB instance from a DB snapshot. The target database is created
// from the source database restore point with most of the source's original
// configuration, including the default security group and DB parameter group. By
// default, the new DB instance is created as a Single-AZ deployment, except when
// the instance is a SQL Server instance that has an option group associated with
// mirroring. In this case, the instance becomes a Multi-AZ deployment, not a
// Single-AZ deployment. If you want to replace your original DB instance with the
// new, restored DB instance, then rename your original DB instance before you call
// the RestoreDBInstanceFromDBSnapshot action. RDS doesn't allow two DB instances
// with the same name. After you have renamed your original DB instance with a
// different identifier, then you can pass the original name of the DB instance as
// the DBInstanceIdentifier in the call to the RestoreDBInstanceFromDBSnapshot
// action. The result is that you replace the original DB instance with the DB
// instance created from the snapshot. If you are restoring from a shared manual DB
// snapshot, the DBSnapshotIdentifier must be the ARN of the shared DB snapshot.
// This command doesn't apply to Aurora MySQL and Aurora PostgreSQL. For Aurora,
// use RestoreDBClusterFromSnapshot.
func (c *Client) RestoreDBInstanceFromDBSnapshot(ctx context.Context, params *RestoreDBInstanceFromDBSnapshotInput, optFns ...func(*Options)) (*RestoreDBInstanceFromDBSnapshotOutput, error) {
	if params == nil {
		params = &RestoreDBInstanceFromDBSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBInstanceFromDBSnapshot", params, optFns, c.addOperationRestoreDBInstanceFromDBSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBInstanceFromDBSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RestoreDBInstanceFromDBSnapshotInput struct {

	// Name of the DB instance to create from the DB snapshot. This parameter isn't
	// case-sensitive. Constraints:
	//
	// * Must contain from 1 to 63 numbers, letters, or
	// hyphens
	//
	// * First character must be a letter
	//
	// * Can't end with a hyphen or
	// contain two consecutive hyphens
	//
	// Example: my-snapshot-id
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The identifier for the DB snapshot to restore from. Constraints:
	//
	// * Must match
	// the identifier of an existing DBSnapshot.
	//
	// * If you are restoring from a shared
	// manual DB snapshot, the DBSnapshotIdentifier must be the ARN of the shared DB
	// snapshot.
	//
	// This member is required.
	DBSnapshotIdentifier *string

	// A value that indicates whether minor version upgrades are applied automatically
	// to the DB instance during the maintenance window. If you restore an RDS Custom
	// DB instance, you must disable this parameter.
	AutoMinorVersionUpgrade *bool

	// The Availability Zone (AZ) where the DB instance will be created. Default: A
	// random, system-chosen Availability Zone. Constraint: You can't specify the
	// AvailabilityZone parameter if the DB instance is a Multi-AZ deployment. Example:
	// us-east-1a
	AvailabilityZone *string

	// Specifies where automated backups and manual snapshots are stored for the
	// restored DB instance. Possible values are outposts (Amazon Web Services
	// Outposts) and region (Amazon Web Services Region). The default is region. For
	// more information, see Working with Amazon RDS on Amazon Web Services Outposts
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in
	// the Amazon RDS User Guide.
	BackupTarget *string

	// A value that indicates whether to copy all tags from the restored DB instance to
	// snapshots of the DB instance. In most cases, tags aren't copied by default.
	// However, when you restore a DB instance from a DB snapshot, RDS checks whether
	// you specify new tags. If yes, the new tags are added to the restored DB
	// instance. If there are no new tags, RDS looks for the tags from the source DB
	// instance for the DB snapshot, and then adds those tags to the restored DB
	// instance. For more information, see  Copying tags to DB instance snapshots
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html#USER_Tagging.CopyTags)
	// in the Amazon RDS User Guide.
	CopyTagsToSnapshot *bool

	// The instance profile associated with the underlying Amazon EC2 instance of an
	// RDS Custom DB instance. The instance profile must meet the following
	// requirements:
	//
	// * The profile must exist in your account.
	//
	// * The profile must
	// have an IAM role that Amazon EC2 has permissions to assume.
	//
	// * The instance
	// profile name and the associated IAM role name must start with the prefix
	// AWSRDSCustom.
	//
	// For the list of permissions required for the IAM role, see
	// Configure IAM and your VPC
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc)
	// in the Amazon RDS User Guide. This setting is required for RDS Custom.
	CustomIamInstanceProfile *string

	// The compute and memory capacity of the Amazon RDS DB instance, for example
	// db.m4.large. Not all DB instance classes are available in all Amazon Web
	// Services Regions, or for all database engines. For the full list of DB instance
	// classes, and availability for your engine, see DB Instance Class
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide. Default: The same DBInstanceClass as the original
	// DB instance.
	DBInstanceClass *string

	// The database name for the restored DB instance. This parameter doesn't apply to
	// the MySQL, PostgreSQL, or MariaDB engines. It also doesn't apply to RDS Custom
	// DB instances.
	DBName *string

	// The name of the DB parameter group to associate with this DB instance. If you
	// don't specify a value for DBParameterGroupName, then RDS uses the default
	// DBParameterGroup for the specified DB engine. This setting doesn't apply to RDS
	// Custom. Constraints:
	//
	// * If supplied, must match the name of an existing
	// DBParameterGroup.
	//
	// * Must be 1 to 255 letters, numbers, or hyphens.
	//
	// * First
	// character must be a letter.
	//
	// * Can't end with a hyphen or contain two
	// consecutive hyphens.
	DBParameterGroupName *string

	// The DB subnet group name to use for the new instance. Constraints: If supplied,
	// must match the name of an existing DBSubnetGroup. Example: mydbsubnetgroup
	DBSubnetGroupName *string

	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection isn't enabled. For more information, see  Deleting a DB
	// Instance
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	DeletionProtection *bool

	// Specify the Active Directory directory ID to restore the DB instance in. The
	// domain/ must be created prior to this operation. Currently, you can create only
	// MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances in an Active
	// Directory Domain. For more information, see  Kerberos Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html)
	// in the Amazon RDS User Guide. This setting doesn't apply to RDS Custom.
	Domain *string

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service. This setting doesn't apply to RDS Custom.
	DomainIAMRoleName *string

	// The list of logs that the restored DB instance is to export to CloudWatch Logs.
	// The values in the list depend on the DB engine being used. For more information,
	// see Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide. This setting doesn't apply to RDS Custom.
	EnableCloudwatchLogsExports []string

	// A value that indicates whether to enable a customer-owned IP address (CoIP) for
	// an RDS on Outposts DB instance. A CoIP provides local or external connectivity
	// to resources in your Outpost subnets through your on-premises network. For some
	// use cases, a CoIP can provide lower latency for connections to the DB instance
	// from outside of its virtual private cloud (VPC) on your local network. This
	// setting doesn't apply to RDS Custom. For more information about RDS on Outposts,
	// see Working with Amazon RDS on Amazon Web Services Outposts
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in
	// the Amazon RDS User Guide. For more information about CoIPs, see Customer-owned
	// IP addresses
	// (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
	// in the Amazon Web Services Outposts User Guide.
	EnableCustomerOwnedIp *bool

	// A value that indicates whether to enable mapping of Amazon Web Services Identity
	// and Access Management (IAM) accounts to database accounts. By default, mapping
	// is disabled. For more information about IAM database authentication, see  IAM
	// Database Authentication for MySQL and PostgreSQL
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon RDS User Guide. This setting doesn't apply to RDS Custom.
	EnableIAMDatabaseAuthentication *bool

	// The database engine to use for the new instance. This setting doesn't apply to
	// RDS Custom. Default: The same as source Constraint: Must be compatible with the
	// engine of the source. For example, you can restore a MariaDB 10.1 DB instance
	// from a MySQL 5.6 snapshot. Valid Values:
	//
	// * mariadb
	//
	// * mysql
	//
	// * oracle-ee
	//
	// *
	// oracle-ee-cdb
	//
	// * oracle-se2
	//
	// * oracle-se2-cdb
	//
	// * postgres
	//
	// * sqlserver-ee
	//
	// *
	// sqlserver-se
	//
	// * sqlserver-ex
	//
	// * sqlserver-web
	Engine *string

	// Specifies the amount of provisioned IOPS for the DB instance, expressed in I/O
	// operations per second. If this parameter isn't specified, the IOPS value is
	// taken from the backup. If this parameter is set to 0, the new instance is
	// converted to a non-PIOPS instance. The conversion takes additional time, though
	// your DB instance is available for connections before the conversion starts. The
	// provisioned IOPS value must follow the requirements for your database engine.
	// For more information, see Amazon RDS Provisioned IOPS Storage to Improve
	// Performance
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS)
	// in the Amazon RDS User Guide. Constraints: Must be an integer greater than 1000.
	Iops *int32

	// License model information for the restored DB instance. This setting doesn't
	// apply to RDS Custom. Default: Same as source. Valid values: license-included |
	// bring-your-own-license | general-public-license
	LicenseModel *string

	// A value that indicates whether the DB instance is a Multi-AZ deployment. This
	// setting doesn't apply to RDS Custom. Constraint: You can't specify the
	// AvailabilityZone parameter if the DB instance is a Multi-AZ deployment.
	MultiAZ *bool

	// The network type of the DB instance. Valid values:
	//
	// * IPV4
	//
	// * DUAL
	//
	// The network
	// type is determined by the DBSubnetGroup specified for the DB instance. A
	// DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the IPv6
	// protocols (DUAL). For more information, see  Working with a DB instance in a VPC
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html)
	// in the Amazon RDS User Guide.
	NetworkType *string

	// The name of the option group to be used for the restored DB instance. Permanent
	// options, such as the TDE option for Oracle Advanced Security TDE, can't be
	// removed from an option group, and that option group can't be removed from a DB
	// instance after it is associated with a DB instance. This setting doesn't apply
	// to RDS Custom.
	OptionGroupName *string

	// The port number on which the database accepts connections. Default: The same
	// port as the original DB instance Constraints: Value must be 1150-65535
	Port *int32

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance. This setting doesn't apply to RDS Custom.
	ProcessorFeatures []types.ProcessorFeature

	// A value that indicates whether the DB instance is publicly accessible. When the
	// DB instance is publicly accessible, its Domain Name System (DNS) endpoint
	// resolves to the private IP address from within the DB instance's virtual private
	// cloud (VPC). It resolves to the public IP address from outside of the DB
	// instance's VPC. Access to the DB instance is ultimately controlled by the
	// security group it uses. That public access is not permitted if the security
	// group assigned to the DB instance doesn't permit it. When the DB instance isn't
	// publicly accessible, it is an internal DB instance with a DNS name that resolves
	// to a private IP address. For more information, see CreateDBInstance.
	PubliclyAccessible *bool

	// Specifies the storage type to be associated with the DB instance. Valid values:
	// standard | gp2 | io1 If you specify io1, you must also include a value for the
	// Iops parameter. Default: io1 if the Iops parameter is specified, otherwise gp2
	StorageType *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []types.Tag

	// The ARN from the key store with which to associate the instance for TDE
	// encryption. This setting doesn't apply to RDS Custom.
	TdeCredentialArn *string

	// The password for the given ARN from the key store in order to access the device.
	// This setting doesn't apply to RDS Custom.
	TdeCredentialPassword *string

	// A value that indicates whether the DB instance class of the DB instance uses its
	// default processor features. This setting doesn't apply to RDS Custom.
	UseDefaultProcessorFeatures *bool

	// A list of EC2 VPC security groups to associate with this DB instance. Default:
	// The default EC2 VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreDBInstanceFromDBSnapshotOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the operations CreateDBInstance,
	// CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances,
	// ModifyDBInstance, PromoteReadReplica, RebootDBInstance,
	// RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3,
	// RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreDBInstanceFromDBSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBInstanceFromDBSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBInstanceFromDBSnapshot{}, middleware.After)
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
	if err = addOpRestoreDBInstanceFromDBSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBInstanceFromDBSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreDBInstanceFromDBSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RestoreDBInstanceFromDBSnapshot",
	}
}
