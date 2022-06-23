// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Tests the auto-merging settings of your Identity Resolution Job without merging
// your data. It randomly selects a sample of matching groups from the existing
// matching results, and applies the automerging settings that you provided. You
// can then view the number of profiles in the sample, the number of matches, and
// the number of profiles identified to be merged. This enables you to evaluate the
// accuracy of the attributes in your matching list. You can't view which profiles
// are matched and would be merged. We strongly recommend you use this API to do a
// dry run of the automerging process before running the Identity Resolution Job.
// Include at least two matching attributes. If your matching list includes too few
// attributes (such as only FirstName or only LastName), there may be a large
// number of matches. This increases the chances of erroneous merges.
func (c *Client) GetAutoMergingPreview(ctx context.Context, params *GetAutoMergingPreviewInput, optFns ...func(*Options)) (*GetAutoMergingPreviewOutput, error) {
	if params == nil {
		params = &GetAutoMergingPreviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAutoMergingPreview", params, optFns, c.addOperationGetAutoMergingPreviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAutoMergingPreviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAutoMergingPreviewInput struct {

	// How the auto-merging process should resolve conflicts between different
	// profiles.
	//
	// This member is required.
	ConflictResolution *types.ConflictResolution

	// A list of matching attributes that represent matching criteria.
	//
	// This member is required.
	Consolidation *types.Consolidation

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

type GetAutoMergingPreviewOutput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The number of match groups in the domain that have been reviewed in this preview
	// dry run.
	NumberOfMatchesInSample int64

	// The number of profiles found in this preview dry run.
	NumberOfProfilesInSample int64

	// The number of profiles that would be merged if this wasn't a preview dry run.
	NumberOfProfilesWillBeMerged int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAutoMergingPreviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAutoMergingPreview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAutoMergingPreview{}, middleware.After)
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
	if err = addOpGetAutoMergingPreviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAutoMergingPreview(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAutoMergingPreview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "GetAutoMergingPreview",
	}
}
