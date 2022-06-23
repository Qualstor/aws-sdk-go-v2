// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns domain configuration information about the specified domains, including
// the domain ID, domain endpoint, and domain ARN.
func (c *Client) DescribeDomains(ctx context.Context, params *DescribeDomainsInput, optFns ...func(*Options)) (*DescribeDomainsOutput, error) {
	if params == nil {
		params = &DescribeDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDomains", params, optFns, c.addOperationDescribeDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeDomains operation. By default, the
// API returns the status of all domains.
type DescribeDomainsInput struct {

	// The domains for which you want information.
	//
	// This member is required.
	DomainNames []string

	noSmithyDocumentSerde
}

// The result of a DescribeDomains request. Contains the status of the specified
// domains or all domains owned by the account.
type DescribeDomainsOutput struct {

	// The status of the domains requested in the DescribeDomains request.
	//
	// This member is required.
	DomainStatusList []types.DomainStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDomains{}, middleware.After)
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
	if err = addOpDescribeDomainsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDomains(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeDomains",
	}
}
