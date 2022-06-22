// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about an event data store specified as either an ARN or the
// ID portion of the ARN.
func (c *Client) GetEventDataStore(ctx context.Context, params *GetEventDataStoreInput, optFns ...func(*Options)) (*GetEventDataStoreOutput, error) {
	if params == nil {
		params = &GetEventDataStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEventDataStore", params, optFns, c.addOperationGetEventDataStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEventDataStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEventDataStoreInput struct {

	// The ARN (or ID suffix of the ARN) of the event data store about which you want
	// information.
	//
	// This member is required.
	EventDataStore *string

	noSmithyDocumentSerde
}

type GetEventDataStoreOutput struct {

	// The advanced event selectors used to select events for the data store.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// The timestamp of the event data store's creation.
	CreatedTimestamp *time.Time

	// The event data store Amazon Resource Number (ARN).
	EventDataStoreArn *string

	// Indicates whether the event data store includes events from all regions, or only
	// from the region in which it was created.
	MultiRegionEnabled *bool

	// The name of the event data store.
	Name *string

	// Indicates whether an event data store is collecting logged events for an
	// organization in Organizations.
	OrganizationEnabled *bool

	// The retention period of the event data store, in days.
	RetentionPeriod *int32

	// The status of an event data store. Values can be ENABLED and PENDING_DELETION.
	Status types.EventDataStoreStatus

	// Indicates that termination protection is enabled.
	TerminationProtectionEnabled *bool

	// Shows the time that an event data store was updated, if applicable.
	// UpdatedTimestamp is always either the same or newer than the time shown in
	// CreatedTimestamp.
	UpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEventDataStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetEventDataStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEventDataStore{}, middleware.After)
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
	if err = addOpGetEventDataStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEventDataStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEventDataStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "GetEventDataStore",
	}
}
