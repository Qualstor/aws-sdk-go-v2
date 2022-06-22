// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a cell to replace the list of nested cells with a new list of nested
// cells.
func (c *Client) UpdateCell(ctx context.Context, params *UpdateCellInput, optFns ...func(*Options)) (*UpdateCellOutput, error) {
	if params == nil {
		params = &UpdateCellInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCell", params, optFns, c.addOperationUpdateCellMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCellOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCellInput struct {

	// The name of the cell.
	//
	// This member is required.
	CellName *string

	// A list of cell Amazon Resource Names (ARNs), which completely replaces the
	// previous list.
	//
	// This member is required.
	Cells []string

	noSmithyDocumentSerde
}

type UpdateCellOutput struct {

	// The Amazon Resource Name (ARN) for the cell.
	CellArn *string

	// The name of the cell.
	CellName *string

	// A list of cell ARNs.
	Cells []string

	// The readiness scope for the cell, which can be a cell Amazon Resource Name (ARN)
	// or a recovery group ARN. This is a list but currently can have only one element.
	ParentReadinessScopes []string

	// Tags on the resources.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCellMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCell{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCell{}, middleware.After)
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
	if err = addOpUpdateCellValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCell(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCell(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-readiness",
		OperationName: "UpdateCell",
	}
}
