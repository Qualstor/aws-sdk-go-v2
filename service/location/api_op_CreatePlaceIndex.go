// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a place index resource in your AWS account. Use a place index resource
// to geocode addresses and other text queries by using the SearchPlaceIndexForText
// operation, and reverse geocode coordinates by using the
// SearchPlaceIndexForPosition operation, and enable autosuggestions by using the
// SearchPlaceIndexForSuggestions operation. If your application is tracking or
// routing assets you use in your business, such as delivery vehicles or employees,
// you may only use HERE as your geolocation provider. See section 82 of the AWS
// service terms (http://aws.amazon.com/service-terms) for more details.
func (c *Client) CreatePlaceIndex(ctx context.Context, params *CreatePlaceIndexInput, optFns ...func(*Options)) (*CreatePlaceIndexOutput, error) {
	if params == nil {
		params = &CreatePlaceIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePlaceIndex", params, optFns, c.addOperationCreatePlaceIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePlaceIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePlaceIndexInput struct {

	// Specifies the geospatial data provider for the new place index. This field is
	// case-sensitive. Enter the valid values as shown. For example, entering HERE
	// returns an error. Valid values include:
	//
	// * Esri – For additional information
	// about Esri
	// (https://docs.aws.amazon.com/location/latest/developerguide/esri.html)'s
	// coverage in your region of interest, see Esri details on geocoding coverage
	// (https://developers.arcgis.com/rest/geocode/api-reference/geocode-coverage.htm).
	//
	// *
	// Here – For additional information about HERE Technologies
	// (https://docs.aws.amazon.com/location/latest/developerguide/HERE.html)' coverage
	// in your region of interest, see HERE details on goecoding coverage
	// (https://developer.here.com/documentation/geocoder/dev_guide/topics/coverage-geocoder.html).
	// If you specify HERE Technologies (Here) as the data provider, you may not store
	// results
	// (https://docs.aws.amazon.com/location-places/latest/APIReference/API_DataSourceConfiguration.html)
	// for locations in Japan. For more information, see the AWS Service Terms
	// (https://aws.amazon.com/service-terms/) for Amazon Location Service.
	//
	// For
	// additional information , see Data providers
	// (https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html)
	// on the Amazon Location Service Developer Guide.
	//
	// This member is required.
	DataSource *string

	// The name of the place index resource. Requirements:
	//
	// * Contain only alphanumeric
	// characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and underscores (_).
	//
	// *
	// Must be a unique place index resource name.
	//
	// * No spaces allowed. For example,
	// ExamplePlaceIndex.
	//
	// This member is required.
	IndexName *string

	// Specifies the data storage option requesting Places.
	DataSourceConfiguration *types.DataSourceConfiguration

	// The optional description for the place index resource.
	Description *string

	// No longer used. If included, the only allowed value is RequestBasedUsage.
	//
	// Deprecated: Deprecated. If included, the only allowed value is
	// RequestBasedUsage.
	PricingPlan types.PricingPlan

	// Applies one or more tags to the place index resource. A tag is a key-value pair
	// that helps you manage, identify, search, and filter your resources. Format:
	// "key" : "value" Restrictions:
	//
	// * Maximum 50 tags per resource.
	//
	// * Each tag key
	// must be unique and must have exactly one associated value.
	//
	// * Maximum key
	// length: 128 Unicode characters in UTF-8.
	//
	// * Maximum value length: 256 Unicode
	// characters in UTF-8.
	//
	// * Can use alphanumeric characters (A–Z, a–z, 0–9), and the
	// following characters: + - = . _ : / @
	//
	// * Cannot use "aws:" as a prefix for a
	// key.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreatePlaceIndexOutput struct {

	// The timestamp for when the place index resource was created in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	CreateTime *time.Time

	// The Amazon Resource Name (ARN) for the place index resource. Used to specify a
	// resource across AWS.
	//
	// * Format example:
	// arn:aws:geo:region:account-id:place-index/ExamplePlaceIndex
	//
	// This member is required.
	IndexArn *string

	// The name for the place index resource.
	//
	// This member is required.
	IndexName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePlaceIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePlaceIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePlaceIndex{}, middleware.After)
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
	if err = addEndpointPrefix_opCreatePlaceIndexMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreatePlaceIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlaceIndex(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreatePlaceIndexMiddleware struct {
}

func (*endpointPrefix_opCreatePlaceIndexMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreatePlaceIndexMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "places." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreatePlaceIndexMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreatePlaceIndexMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePlaceIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "CreatePlaceIndex",
	}
}
