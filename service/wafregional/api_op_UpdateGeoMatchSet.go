// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Inserts or deletes GeoMatchConstraint objects in an GeoMatchSet. For
// each GeoMatchConstraint object, you specify the following values:
//
// * Whether to
// insert or delete the object from the array. If you want to change an
// GeoMatchConstraint object, you delete the existing object and add a new one.
//
// *
// The Type. The only valid value for Type is Country.
//
// * The Value, which is a two
// character code for the country to add to the GeoMatchConstraint object. Valid
// codes are listed in GeoMatchConstraint$Value.
//
// To create and configure an
// GeoMatchSet, perform the following steps:
//
// * Submit a CreateGeoMatchSet
// request.
//
// * Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateGeoMatchSet request.
//
// * Submit an
// UpdateGeoMatchSet request to specify the country that you want AWS WAF to watch
// for.
//
// When you update an GeoMatchSet, you specify the country that you want to
// add and/or the country that you want to delete. If you want to change a country,
// you delete the existing country and add the new one. For more information about
// how to use the AWS WAF API to allow or block HTTP requests, see the AWS WAF
// Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) UpdateGeoMatchSet(ctx context.Context, params *UpdateGeoMatchSetInput, optFns ...func(*Options)) (*UpdateGeoMatchSetOutput, error) {
	if params == nil {
		params = &UpdateGeoMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGeoMatchSet", params, optFns, c.addOperationUpdateGeoMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGeoMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGeoMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// The GeoMatchSetId of the GeoMatchSet that you want to update. GeoMatchSetId is
	// returned by CreateGeoMatchSet and by ListGeoMatchSets.
	//
	// This member is required.
	GeoMatchSetId *string

	// An array of GeoMatchSetUpdate objects that you want to insert into or delete
	// from an GeoMatchSet. For more information, see the applicable data types:
	//
	// *
	// GeoMatchSetUpdate: Contains Action and GeoMatchConstraint
	//
	// * GeoMatchConstraint:
	// Contains Type and Value You can have only one Type and Value per
	// GeoMatchConstraint. To add multiple countries, include multiple
	// GeoMatchSetUpdate objects in your request.
	//
	// This member is required.
	Updates []types.GeoMatchSetUpdate

	noSmithyDocumentSerde
}

type UpdateGeoMatchSetOutput struct {

	// The ChangeToken that you used to submit the UpdateGeoMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGeoMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateGeoMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateGeoMatchSet{}, middleware.After)
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
	if err = addOpUpdateGeoMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGeoMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateGeoMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "UpdateGeoMatchSet",
	}
}
