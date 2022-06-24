// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the configurable settings for the specified
// directory.
func (c *Client) DescribeSettings(ctx context.Context, params *DescribeSettingsInput, optFns ...func(*Options)) (*DescribeSettingsOutput, error) {
	if params == nil {
		params = &DescribeSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSettings", params, optFns, c.addOperationDescribeSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSettingsInput struct {

	// The identifier of the directory for which to retrieve information.
	//
	// This member is required.
	DirectoryId *string

	// The DescribeSettingsResult.NextToken value from a previous call to
	// DescribeSettings. Pass null if this is the first call.
	NextToken *string

	// The status of the directory settings for which to retrieve information.
	Status types.DirectoryConfigurationStatus

	noSmithyDocumentSerde
}

type DescribeSettingsOutput struct {

	// The identifier of the directory.
	DirectoryId *string

	// If not null, token that indicates that more results are available. Pass this
	// value for the NextToken parameter in a subsequent call to DescribeSettings to
	// retrieve the next set of items.
	NextToken *string

	// The list of SettingEntry objects that were retrieved. It is possible that this
	// list contains less than the number of items specified in the Limit member of the
	// request. This occurs if there are less than the requested number of items left
	// to retrieve, or if the limitations of the operation have been exceeded.
	SettingEntries []types.SettingEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSettings{}, middleware.After)
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
	if err = addOpDescribeSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeSettings",
	}
}