// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deploys an application revision through the specified deployment group.
func (c *Client) CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) {
	if params == nil {
		params = &CreateDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeployment", params, optFns, c.addOperationCreateDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateDeployment operation.
type CreateDeploymentInput struct {

	// The name of an AWS CodeDeploy application associated with the IAM user or AWS
	// account.
	//
	// This member is required.
	ApplicationName *string

	// Configuration information for an automatic rollback that is added when a
	// deployment is created.
	AutoRollbackConfiguration *types.AutoRollbackConfiguration

	// The name of a deployment configuration associated with the IAM user or AWS
	// account. If not specified, the value configured in the deployment group is used
	// as the default. If the deployment group does not have a deployment configuration
	// associated with it, CodeDeployDefault.OneAtATime is used by default.
	DeploymentConfigName *string

	// The name of the deployment group.
	DeploymentGroupName *string

	// A comment about the deployment.
	Description *string

	// Information about how AWS CodeDeploy handles files that already exist in a
	// deployment target location but weren't part of the previous successful
	// deployment. The fileExistsBehavior parameter takes any of the following
	// values:
	//
	// * DISALLOW: The deployment fails. This is also the default behavior if
	// no option is specified.
	//
	// * OVERWRITE: The version of the file from the
	// application revision currently being deployed replaces the version already on
	// the instance.
	//
	// * RETAIN: The version of the file already on the instance is kept
	// and used as part of the new deployment.
	FileExistsBehavior types.FileExistsBehavior

	// If true, then if an ApplicationStop, BeforeBlockTraffic, or AfterBlockTraffic
	// deployment lifecycle event to an instance fails, then the deployment continues
	// to the next deployment lifecycle event. For example, if ApplicationStop fails,
	// the deployment continues with DownloadBundle. If BeforeBlockTraffic fails, the
	// deployment continues with BlockTraffic. If AfterBlockTraffic fails, the
	// deployment continues with ApplicationStop. If false or not specified, then if a
	// lifecycle event fails during a deployment to an instance, that deployment fails.
	// If deployment to that instance is part of an overall deployment and the number
	// of healthy hosts is not less than the minimum number of healthy hosts, then a
	// deployment to the next instance is attempted. During a deployment, the AWS
	// CodeDeploy agent runs the scripts specified for ApplicationStop,
	// BeforeBlockTraffic, and AfterBlockTraffic in the AppSpec file from the previous
	// successful deployment. (All other scripts are run from the AppSpec file in the
	// current deployment.) If one of these scripts contains an error and does not run
	// successfully, the deployment can fail. If the cause of the failure is a script
	// from the last successful deployment that will never run successfully, create a
	// new deployment and use ignoreApplicationStopFailures to specify that the
	// ApplicationStop, BeforeBlockTraffic, and AfterBlockTraffic failures should be
	// ignored.
	IgnoreApplicationStopFailures bool

	// The type and location of the revision to deploy.
	Revision *types.RevisionLocation

	// Information about the instances that belong to the replacement environment in a
	// blue/green deployment.
	TargetInstances *types.TargetInstances

	// Indicates whether to deploy to all instances or only to instances that are not
	// running the latest application revision.
	UpdateOutdatedInstancesOnly bool

	noSmithyDocumentSerde
}

// Represents the output of a CreateDeployment operation.
type CreateDeploymentOutput struct {

	// The unique ID of a deployment.
	DeploymentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDeployment{}, middleware.After)
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
	if err = addOpCreateDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "CreateDeployment",
	}
}
