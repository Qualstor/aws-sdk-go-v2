// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the Kubernetes version or AMI version of an Amazon EKS managed node
// group. You can update a node group using a launch template only if the node
// group was originally deployed with a launch template. If you need to update a
// custom AMI in a node group that was deployed with a launch template, then update
// your custom AMI, specify the new ID in a new version of the launch template, and
// then update the node group to the new version of the launch template. If you
// update without a launch template, then you can update to the latest available
// AMI version of a node group's current Kubernetes version by not specifying a
// Kubernetes version in the request. You can update to the latest AMI version of
// your cluster's current Kubernetes version by specifying your cluster's
// Kubernetes version in the request. For more information, see Amazon EKS
// optimized Amazon Linux 2 AMI versions
// (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)
// in the Amazon EKS User Guide. You cannot roll back a node group to an earlier
// Kubernetes version or AMI version. When a node in a managed node group is
// terminated due to a scaling action or update, the pods in that node are drained
// first. Amazon EKS attempts to drain the nodes gracefully and will fail if it is
// unable to do so. You can force the update if Amazon EKS is unable to drain the
// nodes as a result of a pod disruption budget issue.
func (c *Client) UpdateNodegroupVersion(ctx context.Context, params *UpdateNodegroupVersionInput, optFns ...func(*Options)) (*UpdateNodegroupVersionOutput, error) {
	if params == nil {
		params = &UpdateNodegroupVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNodegroupVersion", params, optFns, c.addOperationUpdateNodegroupVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNodegroupVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNodegroupVersionInput struct {

	// The name of the Amazon EKS cluster that is associated with the managed node
	// group to update.
	//
	// This member is required.
	ClusterName *string

	// The name of the managed node group to update.
	//
	// This member is required.
	NodegroupName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// Force the update if the existing node group's pods are unable to be drained due
	// to a pod disruption budget issue. If an update fails because pods could not be
	// drained, you can force the update after it fails to terminate the old node
	// whether or not any pods are running on the node.
	Force bool

	// An object representing a node group's launch template specification. You can
	// only update a node group using a launch template if the node group was
	// originally deployed with a launch template.
	LaunchTemplate *types.LaunchTemplateSpecification

	// The AMI version of the Amazon EKS optimized AMI to use for the update. By
	// default, the latest available AMI version for the node group's Kubernetes
	// version is used. For more information, see Amazon EKS optimized Amazon Linux 2
	// AMI versions
	// (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)
	// in the Amazon EKS User Guide. If you specify launchTemplate, and your launch
	// template uses a custom AMI, then don't specify releaseVersion, or the node group
	// update will fail. For more information about using launch templates with Amazon
	// EKS, see Launch template support
	// (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the
	// Amazon EKS User Guide.
	ReleaseVersion *string

	// The Kubernetes version to update to. If no version is specified, then the
	// Kubernetes version of the node group does not change. You can specify the
	// Kubernetes version of the cluster to update the node group to the latest AMI
	// version of the cluster's Kubernetes version. If you specify launchTemplate, and
	// your launch template uses a custom AMI, then don't specify version, or the node
	// group update will fail. For more information about using launch templates with
	// Amazon EKS, see Launch template support
	// (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the
	// Amazon EKS User Guide.
	Version *string

	noSmithyDocumentSerde
}

type UpdateNodegroupVersionOutput struct {

	// An object representing an asynchronous update.
	Update *types.Update

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNodegroupVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateNodegroupVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateNodegroupVersion{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateNodegroupVersionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateNodegroupVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNodegroupVersion(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateNodegroupVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateNodegroupVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateNodegroupVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateNodegroupVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateNodegroupVersionInput ")
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
func addIdempotencyToken_opUpdateNodegroupVersionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateNodegroupVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateNodegroupVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "UpdateNodegroupVersion",
	}
}
