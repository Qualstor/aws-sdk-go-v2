// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/Qualstor/aws-sdk-go-v2/aws/middleware"
	"github.com/Qualstor/aws-sdk-go-v2/aws/signer/v4"
	"github.com/Qualstor/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes a recipe. A recipe contains three items:
//
// * An algorithm that trains a
// model.
//
// * Hyperparameters that govern the training.
//
// * Feature transformation
// information for modifying the input data before training.
//
// Amazon Personalize
// provides a set of predefined recipes. You specify a recipe when you create a
// solution with the CreateSolution
// (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html) API.
// CreateSolution trains a model by using the algorithm in the specified recipe and
// a training dataset. The solution, when deployed as a campaign, can provide
// recommendations using the GetRecommendations
// (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html)
// API.
func (c *Client) DescribeRecipe(ctx context.Context, params *DescribeRecipeInput, optFns ...func(*Options)) (*DescribeRecipeOutput, error) {
	if params == nil {
		params = &DescribeRecipeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRecipe", params, optFns, c.addOperationDescribeRecipeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRecipeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecipeInput struct {

	// The Amazon Resource Name (ARN) of the recipe to describe.
	//
	// This member is required.
	RecipeArn *string

	noSmithyDocumentSerde
}

type DescribeRecipeOutput struct {

	// An object that describes the recipe.
	Recipe *types.Recipe

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRecipeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRecipe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRecipe{}, middleware.After)
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
	if err = addOpDescribeRecipeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecipe(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRecipe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DescribeRecipe",
	}
}
