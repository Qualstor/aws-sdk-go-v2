// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Revokes access to a cluster.
func (c *Client) RevokeEndpointAccess(ctx context.Context, params *RevokeEndpointAccessInput, optFns ...func(*Options)) (*RevokeEndpointAccessOutput, error) {
	if params == nil {
		params = &RevokeEndpointAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeEndpointAccess", params, optFns, c.addOperationRevokeEndpointAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeEndpointAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeEndpointAccessInput struct {

	// The Amazon Web Services account ID whose access is to be revoked.
	Account *string

	// The cluster to revoke access from.
	ClusterIdentifier *string

	// Indicates whether to force the revoke action. If true, the Redshift-managed VPC
	// endpoints associated with the endpoint authorization are also deleted.
	Force bool

	// The virtual private cloud (VPC) identifiers for which access is to be revoked.
	VpcIds []string

	noSmithyDocumentSerde
}

// Describes an endpoint authorization for authorizing Redshift-managed VPC
// endpoint access to a cluster across Amazon Web Services accounts.
type RevokeEndpointAccessOutput struct {

	// Indicates whether all VPCs in the grantee account are allowed access to the
	// cluster.
	AllowedAllVPCs bool

	// The VPCs allowed access to the cluster.
	AllowedVPCs []string

	// The time (UTC) when the authorization was created.
	AuthorizeTime *time.Time

	// The cluster identifier.
	ClusterIdentifier *string

	// The status of the cluster.
	ClusterStatus *string

	// The number of Redshift-managed VPC endpoints created for the authorization.
	EndpointCount int32

	// The Amazon Web Services account ID of the grantee of the cluster.
	Grantee *string

	// The Amazon Web Services account ID of the cluster owner.
	Grantor *string

	// The status of the authorization action.
	Status types.AuthorizationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRevokeEndpointAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRevokeEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRevokeEndpointAccess{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeEndpointAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokeEndpointAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "RevokeEndpointAccess",
	}
}
