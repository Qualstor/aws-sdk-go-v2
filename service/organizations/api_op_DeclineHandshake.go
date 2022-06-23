// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Declines a handshake request. This sets the handshake state to DECLINED and
// effectively deactivates the request. This operation can be called only from the
// account that received the handshake. The originator of the handshake can use
// CancelHandshake instead. The originator can't reactivate a declined request, but
// can reinitiate the process with a new handshake request. After you decline a
// handshake, it continues to appear in the results of relevant APIs for only 30
// days. After that, it's deleted.
func (c *Client) DeclineHandshake(ctx context.Context, params *DeclineHandshakeInput, optFns ...func(*Options)) (*DeclineHandshakeOutput, error) {
	if params == nil {
		params = &DeclineHandshakeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeclineHandshake", params, optFns, c.addOperationDeclineHandshakeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeclineHandshakeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeclineHandshakeInput struct {

	// The unique identifier (ID) of the handshake that you want to decline. You can
	// get the ID from the ListHandshakesForAccount operation. The regex pattern
	// (http://wikipedia.org/wiki/regex) for handshake ID string requires "h-" followed
	// by from 8 to 32 lowercase letters or digits.
	//
	// This member is required.
	HandshakeId *string

	noSmithyDocumentSerde
}

type DeclineHandshakeOutput struct {

	// A structure that contains details about the declined handshake. The state is
	// updated to show the value DECLINED.
	Handshake *types.Handshake

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeclineHandshakeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeclineHandshake{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeclineHandshake{}, middleware.After)
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
	if err = addOpDeclineHandshakeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeclineHandshake(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeclineHandshake(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DeclineHandshake",
	}
}
