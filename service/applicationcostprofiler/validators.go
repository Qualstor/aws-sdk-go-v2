// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationcostprofiler

import (
	"context"
	"fmt"
	"github.com/Qualstor/aws-sdk-go-v2/service/applicationcostprofiler/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpDeleteReportDefinition struct {
}

func (*validateOpDeleteReportDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteReportDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteReportDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteReportDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetReportDefinition struct {
}

func (*validateOpGetReportDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetReportDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetReportDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetReportDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpImportApplicationUsage struct {
}

func (*validateOpImportApplicationUsage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpImportApplicationUsage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ImportApplicationUsageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpImportApplicationUsageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutReportDefinition struct {
}

func (*validateOpPutReportDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutReportDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutReportDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutReportDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateReportDefinition struct {
}

func (*validateOpUpdateReportDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateReportDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateReportDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateReportDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDeleteReportDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteReportDefinition{}, middleware.After)
}

func addOpGetReportDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetReportDefinition{}, middleware.After)
}

func addOpImportApplicationUsageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpImportApplicationUsage{}, middleware.After)
}

func addOpPutReportDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutReportDefinition{}, middleware.After)
}

func addOpUpdateReportDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateReportDefinition{}, middleware.After)
}

func validateS3Location(v *types.S3Location) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3Location"}
	if v.Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Bucket"))
	}
	if v.Prefix == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Prefix"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSourceS3Location(v *types.SourceS3Location) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SourceS3Location"}
	if v.Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Bucket"))
	}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteReportDefinitionInput(v *DeleteReportDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteReportDefinitionInput"}
	if v.ReportId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetReportDefinitionInput(v *GetReportDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetReportDefinitionInput"}
	if v.ReportId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpImportApplicationUsageInput(v *ImportApplicationUsageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportApplicationUsageInput"}
	if v.SourceS3Location == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceS3Location"))
	} else if v.SourceS3Location != nil {
		if err := validateSourceS3Location(v.SourceS3Location); err != nil {
			invalidParams.AddNested("SourceS3Location", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutReportDefinitionInput(v *PutReportDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutReportDefinitionInput"}
	if v.ReportId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportId"))
	}
	if v.ReportDescription == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportDescription"))
	}
	if len(v.ReportFrequency) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ReportFrequency"))
	}
	if len(v.Format) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Format"))
	}
	if v.DestinationS3Location == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationS3Location"))
	} else if v.DestinationS3Location != nil {
		if err := validateS3Location(v.DestinationS3Location); err != nil {
			invalidParams.AddNested("DestinationS3Location", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateReportDefinitionInput(v *UpdateReportDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateReportDefinitionInput"}
	if v.ReportId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportId"))
	}
	if v.ReportDescription == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReportDescription"))
	}
	if len(v.ReportFrequency) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ReportFrequency"))
	}
	if len(v.Format) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Format"))
	}
	if v.DestinationS3Location == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationS3Location"))
	} else if v.DestinationS3Location != nil {
		if err := validateS3Location(v.DestinationS3Location); err != nil {
			invalidParams.AddNested("DestinationS3Location", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
